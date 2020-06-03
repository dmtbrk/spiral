package spiral

import (
	"reflect"
	"testing"
)

func TestFlattenInt(t *testing.T) {
	type args struct {
		s [][]int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{"Empy", args{s: [][]int{}}, []int{}, false},
		{"3x3", args{s: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}, false},
		{"4x4", args{s: [][]int{{1, 2, 3, 1}, {4, 5, 6, 4}, {7, 8, 9, 7}, {7, 8, 9, 7}}}, []int{1, 2, 3, 1, 4, 7, 7, 9, 8, 7, 7, 4, 5, 6, 9, 8}, false},
		{"Nil", args{s: nil}, nil, true},
		{"Non-rectangular", args{s: [][]int{{1, 2, 3}, {4, 5, 6}}}, nil, true},
		{"Irregular shape", args{s: [][]int{{1, 2, 3}, {4, 5}}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FlattenInt(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("FlattenInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlattenInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
