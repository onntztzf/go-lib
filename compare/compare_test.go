package compare

import "testing"

func TestCompare(t *testing.T) {
	type A struct {
		Age int
	}
	type B struct {
		Age int
	}
	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: testing.CoverMode(), args: args{
			a: 1,
			b: 1,
		}, want: true},
		{name: testing.CoverMode(), args: args{
			a: 1,
			b: 2,
		}, want: false},
		{name: testing.CoverMode(), args: args{
			a: int(1),
			b: int64(1),
		}, want: false},
		{name: testing.CoverMode(), args: args{
			a: A{Age: 1},
			b: A{Age: 1},
		}, want: true},
		{name: testing.CoverMode(), args: args{
			a: A{Age: 1},
			b: B{Age: 1},
		}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
