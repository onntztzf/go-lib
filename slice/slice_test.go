package slice

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	type args struct {
		slice []interface{}
		size  int
	}
	tests := []struct {
		name string
		args args
		want [][]interface{}
	}{
		// TODO: Add test cases.
		{name: testing.CoverMode(), args: args{
			slice: []interface{}{1, 2, 3, 4, 5},
			size:  2,
		}, want: [][]interface{}{{1, 2}, {3, 4}, {5}}},
		{name: testing.CoverMode(), args: args{
			slice: []interface{}{1, 2, 3, 4, 5},
			size:  3,
		}, want: [][]interface{}{{1, 2, 3}, {4, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chunk(tt.args.slice, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContain(t *testing.T) {
	type args struct {
		slice  []interface{}
		target interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: testing.CoverMode(), args: args{
			slice:  []interface{}{int(1), int64(2)},
			target: int64(1),
		}, want: false},
		{name: testing.CoverMode(), args: args{
			slice:  append([]interface{}{}, int64(1)),
			target: int64(1),
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contain(tt.args.slice, tt.args.target); got != tt.want {
				t.Errorf("Contain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	type args struct {
		slice []interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}
