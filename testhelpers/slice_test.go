// +build unit

package testhelpers

import "testing"

func TestContains(t *testing.T) {
	tests := []struct {
		name string
		list []string
		s    string
		want bool
	}{
		{
			name: "should return true when list contains s",
			s:    "a",
			list: []string{"a", "b", "c", "d"},
			want: true,
		},
		{
			name: "should return true when list contains s",
			s:    "test3",
			list: []string{"test0", "test1", "test2", "test3", "test4"},
			want: true,
		},
		{
			name: "should return false when list does not contain s",
			s:    "abcd",
			list: []string{"efgh", "ijkl", "mnop", "qrst", "uvwx", "yz"},
			want: false,
		},
		{
			name: "should return false when list does not contain s",
			s:    "a",
			list: []string{"b", "c", "d"},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := StringSliceContains(tt.list, tt.s); got != tt.want {
				t.Fatalf("StringSliceContains() => got %v, want %v", got, tt.want)
			}
		})
	}
}
