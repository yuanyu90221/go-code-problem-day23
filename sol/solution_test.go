package sol

import "testing"

func TestFindUnsortedArrayLen(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,6,3,8,10,9,10, 20]",
			args: args{
				nums: []int{1, 6, 3, 8, 10, 9, 10, 20},
			},
			want: 5,
		},
		{
			name: "[-3, -1,3, 4, 10]",
			args: args{
				nums: []int{-3, -1, 3, 4, 10},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindUnsortedArrayLen(tt.args.nums); got != tt.want {
				t.Errorf("FindUnsortedArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
