package time

import (
	"testing"
	"time"
)

func TestGetInterval(t *testing.T) {
	now := time.Now()
	type args struct {
		t1 time.Time
		t2 time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		// TODO: Add test cases.
		{name: "t1 < t2", args: args{
			t1: now,
			t2: now.Add(time.Second * 10),
		}, want: time.Second * 10},
		{name: "t1 > t2", args: args{
			t1: now.Add(time.Second * 10),
			t2: now,
		}, want: time.Second * 10},
		{name: "t1 == t2", args: args{
			t1: now,
			t2: now,
		}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInterval(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("GetInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
