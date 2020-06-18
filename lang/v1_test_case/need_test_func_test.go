package v1_test_case

import "testing"

// Generate test case: 1. select function, 2. cmd + N, 3. Test for selection
func TestFoo(t *testing.T) {
	type args struct {
		a string
		b string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "a param > b param",
			args: args{
				a: "y",
				b: "x",
			},
			want: "y",
		},
		{
			name: "a param < b param",
			args: args{
				a: "x",
				b: "y",
			},
			want: "y",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Foo(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Foo() = %v, want %v", got, tt.want)
			}
		})
	}
}