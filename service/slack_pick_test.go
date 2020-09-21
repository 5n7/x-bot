package service

import "testing"

func Test_strJoin(t *testing.T) {
	type args struct {
		ss []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no element",
			args: args{
				ss: []string{""},
			},
			want: "",
		},
		{
			name: "1 element",
			args: args{
				ss: []string{"foo"},
			},
			want: "foo",
		},
		{
			name: "2 elements",
			args: args{
				ss: []string{"foo", "baz"},
			},
			want: "foo and baz",
		},
		{
			name: "3 elements",
			args: args{
				ss: []string{"foo", "bar", "baz"},
			},
			want: "foo, bar and baz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strJoin(tt.args.ss); got != tt.want {
				t.Errorf("strJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}
