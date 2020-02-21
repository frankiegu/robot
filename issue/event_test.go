package issue

import (
	"reflect"
	"testing"
)

func Test_decodeCommand(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *Command
	}{
		{
			"test command",
			args{"/promote 2 test"},
			&Command{"/promote", "2 test"},
		},
		{
			"test command",
			args{"   /promote 2 test"},
			&Command{"/promote", "2 test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeCommand(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}