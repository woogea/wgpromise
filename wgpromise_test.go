package wgpromise

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_newPromise(t *testing.T) {
	type args struct {
		resolve func()
		reject  func()
	}
	tests := []struct {
		name string
		args args
		want *Promise
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newPromise(tt.args.resolve, tt.args.reject); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPromise() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createAndRun(t *testing.T) {
	p := newPromise(func() { fmt.Println("Success!") }, func() { fmt.Println("Failed!") })
	p.Run(func() bool {
		return true
	}).Run(func() bool {
		return false
	})
}

func Test_chain(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	p := newPromise(func() { fmt.Println(arr) }, func() { fmt.Println("failed") })
	p.Run(
		func() bool {
			//sorted := make([]int, len(arr))
			for i := 0; i < len(arr); i++ {
				minindex := findMinimum(arr[i:len(arr)])
				arr[i], arr[minindex+i] = swap(arr[i], arr[minindex+i])
			}
			return true
		})
	for p.stat == Pending {
	}
}

func swap(x int, y int) (int, int) {
	return y, x
}
func findMinimum(subarr []int) int {
	min := subarr[0]
	minindex := 0
	for i := 0; i < len(subarr); i++ {
		if subarr[i] < min {
			min = subarr[i]
			minindex = i
		}
	}
	return minindex
}
