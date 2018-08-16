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
