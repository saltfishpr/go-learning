// @file: others_test.go
// @date: 2021/3/9

package others

import (
	"fmt"
	"testing"
	"time"
)

var eggTests = []struct {
	in  [2]int
	out int
}{
	{in: [2]int{100, 100000}, out: 17},
	{in: [2]int{2, 100}, out: 14},
	{in: [2]int{3, 14}, out: 4},
}

func TestSuperEggDrop(t *testing.T) {
	for _, et := range eggTests {
		start := time.Now()
		result := superEggDrop(et.in[0], et.in[1])
		end := time.Now()
		delta := end.Sub(start)
		if result != et.out {
			t.Fatalf("get:%v, want false", result)
		}
		fmt.Printf("Spend time: %s\tResult is: %d\n", delta, result)
	}
}

var kmpTests = []struct {
	in  [2]string
	out int
}{
	{in: [2]string{"aaacaaab", "aaab"}, out: 4},
	{in: [2]string{"aaaaaaab", "aaab"}, out: 4},
	{in: [2]string{"CABAABABAC", "ABABC"}, out: -1},
}

func TestKMP(t *testing.T) {
	for _, kt := range kmpTests {
		result := kmp(kt.in[1], kt.in[0])
		if result != kt.out {
			t.Fatalf("get: %v, want: %v", result, kt.out)
		}
		fmt.Printf("Result is: %d\n", result)
	}
}
