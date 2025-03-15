package util

import (
	"bytes"
	"fmt"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func GetCallReceiver(call *ssa.Call) ssa.Value {
	if call == nil {
		return nil
	}

	// 如果是普通函数调用而不是方法调用，返回nil
	if call.Call.IsInvoke() {
		return nil
	}

	// 获取调用的接收者
	recv := call.Call.Value
	if recv == nil {
		return nil
	}

	// 根据接收者的类型进行处理
	switch v := recv.(type) {
	case *ssa.Call:
		// 如果接收者是另一个方法调用，递归获取其接收者
		return GetCallReceiver(v)
	default:
		// 对于其他类型的接收者，直接返回
		return recv
	}
}

func PrintSSATree(dir string) {
	// 配置包加载
	cfg := &packages.Config{
		Mode: packages.LoadAllSyntax | packages.NeedModule,
	}

	// 加载包
	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		panic(err)
	}

	// 创建 SSA 程序
	prog, ssaPkgs := ssautil.Packages(pkgs, 0)

	// 构建 SSA
	prog.Build()

	// 打印 SSA 树
	for _, ssaPkg := range ssaPkgs {
		// 创建缓冲区
		var buf bytes.Buffer
		ssaPkg.WriteTo(&buf)

		for _, member := range ssaPkgs {
			member.WriteTo(&buf)
		}

		fmt.Println(buf.String())
	}
}
