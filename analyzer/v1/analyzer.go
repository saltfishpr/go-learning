package analyzer

import (
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"gormlinter/internal/util"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "gormlinter",
	Doc:      "Checks for unsafe chaining of gorm.DB methods",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	unsafeVars := make(map[types.Object]bool)

	nodeFilter := []ast.Node{
		(*ast.AssignStmt)(nil),
		(*ast.ValueSpec)(nil),
		(*ast.CallExpr)(nil),
	}

	insp.Preorder(nodeFilter, func(n ast.Node) {
		switch node := n.(type) {
		case *ast.AssignStmt:
			processAssignment(pass, node, unsafeVars)
		case *ast.ValueSpec:
			processValueSpec(pass, node, unsafeVars)
		case *ast.CallExpr:
			processCallExpr(pass, node, unsafeVars)
		}
	})

	return nil, nil
}

func processAssignment(pass *analysis.Pass, assign *ast.AssignStmt, unsafeVars map[types.Object]bool) {
	for i, rhs := range assign.Rhs {
		if isUnsafeChain(rhs, pass) {
			if len(assign.Lhs) > i {
				if ident, ok := assign.Lhs[i].(*ast.Ident); ok {
					obj := pass.TypesInfo.ObjectOf(ident)
					if obj != nil {
						unsafeVars[obj] = true
						reportUnsafeChain(pass, rhs.Pos())
					}
				}
			}
		}
	}
}

func processValueSpec(pass *analysis.Pass, spec *ast.ValueSpec, unsafeVars map[types.Object]bool) {
	for i, val := range spec.Values {
		if isUnsafeChain(val, pass) {
			if len(spec.Names) > i {
				name := spec.Names[i]
				obj := pass.TypesInfo.ObjectOf(name)
				if obj != nil {
					unsafeVars[obj] = true
					reportUnsafeChain(pass, val.Pos())
				}
			}
		}
	}
}

func processCallExpr(pass *analysis.Pass, call *ast.CallExpr, unsafeVars map[types.Object]bool) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return
	}

	receiver := sel.X
	receiverType := pass.TypesInfo.TypeOf(receiver)
	if !util.IsGormDBType(receiverType) {
		return
	}

	if ident, ok := receiver.(*ast.Ident); ok {
		obj := pass.TypesInfo.ObjectOf(ident)
		if obj != nil && unsafeVars[obj] {
			reportReuseRisk(pass, receiver.Pos(), ident.Name)
		}
	}
}

func isUnsafeChain(expr ast.Expr, pass *analysis.Pass) bool {
	chain := getMethodChain(expr)
	if len(chain) == 0 {
		return false
	}

	lastMethod := chain[len(chain)-1]
	if util.IsGormFinisherMethod(lastMethod) {
		return false
	}

	if util.IsGormNewSessionMethod(lastMethod) && isSessionCall(expr) {
		return false
	}

	return true
}

func getMethodChain(expr ast.Expr) []string {
	switch e := expr.(type) {
	case *ast.CallExpr:
		return getMethodChain(e.Fun)
	case *ast.SelectorExpr:
		return append(getMethodChain(e.X), e.Sel.Name)
	case *ast.Ident:
		return []string{}
	default:
		return []string{}
	}
}

func isSessionCall(expr ast.Expr) bool {
	call, ok := expr.(*ast.CallExpr)
	if !ok {
		return false
	}

	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	if !util.IsGormNewSessionMethod(sel.Sel.Name) {
		return false
	}

	if len(call.Args) != 1 {
		return false
	}

	arg := call.Args[0]
	if uarg, ok := arg.(*ast.UnaryExpr); ok && uarg.Op == token.AND {
		if cl, ok := uarg.X.(*ast.CompositeLit); ok {
			if sel, ok := cl.Type.(*ast.SelectorExpr); ok {
				return sel.Sel.Name == "Session" && strings.Contains(toString(sel.X), "gorm")
			}
		}
	}

	return false
}

func reportUnsafeChain(pass *analysis.Pass, pos token.Pos) {
	pass.Report(analysis.Diagnostic{
		Pos:     pos,
		Message: "gorm.DB chain must end with Finisher method or Session",
		SuggestedFixes: []analysis.SuggestedFix{
			{
				Message: "Add Session(&gorm.Session{})",
				TextEdits: []analysis.TextEdit{{
					Pos:     pos,
					NewText: []byte(".Session(&gorm.Session{})"),
				}},
			},
		},
	})
}

func reportReuseRisk(pass *analysis.Pass, pos token.Pos, varName string) {
	pass.Report(analysis.Diagnostic{
		Pos:     pos,
		Message: "Potential SQL condition pollution risk",
		SuggestedFixes: []analysis.SuggestedFix{
			{
				Message: "Initialize with Session(&gorm.Session{})",
				TextEdits: []analysis.TextEdit{{
					Pos:     pos,
					NewText: []byte(varName + ".Session(&gorm.Session{})"),
				}},
			},
		},
	})
}

func toString(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.SelectorExpr:
		return toString(e.X) + "." + e.Sel.Name
	}
	return ""
}
