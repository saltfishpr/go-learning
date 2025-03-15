package analyzer

import (
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

var Analyzer = &analysis.Analyzer{
	Name:     "gormlinter",
	Doc:      "Checks for unsafe chaining of gorm.DB methods",
	Requires: []*analysis.Analyzer{buildssa.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	ssaInfo := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)

	// 用于存储 receiver -> finisher methods 调用的映射
	callMap := make(map[ssa.Value][]*ssa.Call)
	// 用于存储 receiver -> 赋值语句的映射
	assignmentMap := make(map[ssa.Value]*ssa.Call)

	// 遍历所有函数
	for _, fn := range ssaInfo.SrcFuncs {
		for _, block := range fn.Blocks {
			for _, instr := range block.Instrs {
				// 检查是否是方法调用
				if call, ok := instr.(*ssa.Call); ok {
					if call.Call.IsInvoke() {
						continue
					}

					// 获取调用的函数
					callee := call.Call.StaticCallee()
					if callee == nil {
						continue
					}

					// 检查是否是 gorm.DB 的方法调用
					recv := call.Call.Args[0]
					if !isGormDBType(recv.Type().(*types.Pointer)) {
						continue
					}

					// 如果是 finisher method,记录到 callMap
					if isFinisherMethod(callee.Name()) {
						callMap[recv] = append(callMap[recv], call)
					}

					// 记录赋值语句
					if assign, ok := instr.(*ssa.Call); ok {
						assignmentMap[assign] = call
					}
				}
			}
		}
	}

	// 分析每个 receiver 的调用
	for recv, calls := range callMap {
		if len(calls) > 1 {
			// 如果一个 receiver 有多个 finisher method 调用
			if !isSafeChaining(recv, assignmentMap) {
				pass.Reportf(calls[0].Pos(), "unsafe chaining of gorm.DB methods: multiple finisher methods called on same DB instance")
			}
		}
	}

	return nil, nil
}

// 检查类型是否是 *gorm.DB
func isGormDBType(t *types.Pointer) bool {
	named, ok := t.Elem().(*types.Named)
	if !ok {
		return false
	}

	return named.Obj().Pkg() != nil &&
		named.Obj().Pkg().Path() == "gorm.io/gorm" &&
		named.Obj().Name() == "DB"
}

// 检查是否是 finisher method
func isFinisherMethod(name string) bool {
	return finisherMethods[name]
}

// 检查是否是安全的链式调用
func isSafeChaining(recv ssa.Value, assignmentMap map[ssa.Value]*ssa.Call) bool {
	// 获取 receiver 的来源
	call, ok := assignmentMap[recv]
	if !ok {
		return true
	}

	// 检查是否调用了 new session methods
	if callee := call.Call.StaticCallee(); callee != nil {
		if isNewSessionMethod(callee.Name()) {
			return true
		}
	}

	// 递归检查上层 receiver
	return isSafeChaining(call.Call.Args[0], assignmentMap)
}

// 检查是否是 new session method
func isNewSessionMethod(name string) bool {
	return newSessionMethods[name]
}

var finisherMethods = map[string]bool{
	"Where":  true,
	"Find":   true,
	"First":  true,
	"Create": true,
	// ... 其他 finisher methods
}

var newSessionMethods = map[string]bool{
	"Session":     true,
	"WithContext": true,
	"Debug":       true,
}
