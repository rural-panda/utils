package slist

import (
	"reflect"
	"testing"
)

func Test_safeList_PushBack(t *testing.T) {
	type args struct {
		param1 interface{}
		param2 interface{}
		typ    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "string",
			args: args{
				"china",
				"japan",
				"string",
			},
		},
		{
			name: "nil",
			args: args{
				nil,
				nil,
				"nil",
			},
		},
		{
			name: "slice",
			args: args{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				"slice",
			},
		},
		{
			name: "map",
			args: args{
				map[string]string{"china": "china"},
				map[string]string{"china": "japan"},
				"map",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			f := newSafeList()

			switch tt.args.typ {
			case "string":
				f.PushBack(tt.args.param1)
				f.PushBack(tt.args.param2)

				a := f.RemoveFront().(string)
				if !reflect.DeepEqual(a, tt.args.param1) {
					t.Errorf("case %s failed, want: %+v, res: %+v", tt.name, tt.args.param1, a)
				} else {
					t.Logf("case %s, want: %+v, res: %+v", tt.name, tt.args.param1, a)
				}
			case "slice":
				f.PushBack(tt.args.param1)
				f.PushBack(tt.args.param2)

				a := f.RemoveFront().([]int)
				if !reflect.DeepEqual(a, tt.args.param1) {
					t.Errorf("case %s failed, want: %+v, res: %+v", tt.name, tt.args.param1, a)
				} else {
					t.Logf("case %s, want: %+v, res: %+v", tt.name, tt.args.param1, a)
				}
			case "map":
				f.PushBack(tt.args.param1)
				f.PushBack(tt.args.param2)

				a := f.RemoveFront().(map[string]string)
				if !reflect.DeepEqual(a, tt.args.param1) {
					t.Errorf("case %s failed, want: %+v, res: %+v", tt.name, tt.args.param1, a)
				} else {
					t.Logf("case %s, want: %+v, res: %+v", tt.name, tt.args.param1, a)
				}
			case "nil":
				a := f.RemoveFront()
				if !reflect.DeepEqual(a, nil) {
					t.Errorf("case %s failed, want: %+v, res: %+v", tt.name, tt.args.param1, a)
				} else {
					t.Logf("case %s, want: %+v, res: %+v", tt.name, tt.args.param1, a)
				}
			}
		})
	}
}
