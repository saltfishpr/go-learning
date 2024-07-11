## 331. 验证二叉树的前序序列化

序列化二叉树的一种方法是使用前序遍历。当我们遇到一个非空节点时，我们可以记录下这个节点的值。如果它是一个空节点，我们可以使用一个标记值记录，例如 `#`。

### Think

统计入度出度

- 每当多 1 个非空节点，可插入 null 的位置（槽）就多 1 个
- 每当多 1 个空节点，可插入 null 的位置（槽）就少 1 个

### Solution

```go
package main

func isValidSerialization(preorder string) bool {
	slots := 1
	for i := 0; i < len(preorder); {
		if slots == 0 {
			return false
		}
		switch preorder[i] {
		case ',':
			i++
		case '#':
			slots--
			i++
		default:
			for i < len(preorder) && preorder[i] != ',' {
				i++
			}
			slots++
		}
	}
	return slots == 0
}
```

### Tests

```go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValidSerialization(t *testing.T) {
	type args struct {
		preorder string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "示例 1",
			args: args{
				preorder: "9,3,4,#,#,1,#,#,2,#,6,#,#",
			},
			want: true,
		},
		{
			name: "示例 2",
			args: args{
				preorder: "1,#",
			},
			want: false,
		},
		{
			name: "示例 3",
			args: args{
				preorder: "9,#,#,1",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isValidSerialization(tt.args.preorder))
		})
	}
}

var result bool

func Benchmark_isValidSerialization(b *testing.B) {
	var res bool
	for n := 0; n < b.N; n++ {
		res = isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#")
	}
	result = res
}
```

```plaintext
=== RUN   Test_isValidSerialization
=== RUN   Test_isValidSerialization/示例_1
=== RUN   Test_isValidSerialization/示例_2
=== RUN   Test_isValidSerialization/示例_3
--- PASS: Test_isValidSerialization (0.00s)
    --- PASS: Test_isValidSerialization/示例_1 (0.00s)
    --- PASS: Test_isValidSerialization/示例_2 (0.00s)
    --- PASS: Test_isValidSerialization/示例_3 (0.00s)
PASS
coverage: 100.0% of statements
ok  	leetcode/331	0.003s	coverage: 100.0% of statements
```

```plaintext
goos: linux
goarch: amd64
pkg: leetcode/331
cpu: Intel(R) Xeon(R) Gold 6133 CPU @ 2.50GHz
Benchmark_isValidSerialization
Benchmark_isValidSerialization-4   	24098029	        49.87 ns/op	       0 B/op	       0 allocs/op
PASS
coverage: 91.7% of statements
ok  	leetcode/331	1.260s
```
