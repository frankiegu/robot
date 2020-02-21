package drone_promote

import (
	"reflect"
	"testing"
)

func Test_decodeCmd(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *DronePromoteCmd
	}{
		{
			"test promote cmd",
			args{"32 test key=value"},
			&DronePromoteCmd{32,"test",map[string]string{"key":"value"}},
		},
		{
			"test promote cmd blanks",
			args{"  32 test   key=value"},
			&DronePromoteCmd{32,"test",map[string]string{"key":"value"}},
		},
		{
			"test promote cmd blanks",
			args{"  32 test   key=value   "},
			&DronePromoteCmd{32,"test",map[string]string{"key":"value"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeCmd(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitMultiBlank(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test split",
			args{"  43 test   key=value"},
			[]string{"43","test","key=value"},
		},
		{
			"test split 2",
			args{"  43 test   key=value  "},
			[]string{"43","test","key=value"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitMultiBlank(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitMultiBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}