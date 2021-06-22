package ecrypto

import (
	"reflect"
	"testing"
)

func Test_BytesBlockSplit(t *testing.T) {
	tests := []struct {
		b    []byte
		k    int
		want [][]byte
	}{
		{
			b: []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			k: 3,
			want: [][]byte{
				{0, 1, 2},
				{3, 4, 5},
				{6, 7, 8},
				{9},
			},
		},
	}

	for _, tc := range tests {
		got := BytesBlockSplit(tc.b, tc.k)

		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func Test_Transpose(t *testing.T) {
	tests := []struct {
		bb   [][]byte
		want [][]byte
	}{
		{
			bb: [][]byte{
				{0, 1, 2},
				{3, 4, 5},
				{6, 7, 8},
				{9, 10, 11},
			},
			want: [][]byte{
				{0, 3, 6, 9},
				{1, 4, 7, 10},
				{2, 5, 8, 11},
			},
		},
		{
			bb: [][]byte{
				{0, 1, 2},
				{3, 4, 5},
				{6, 7, 8},
				{9, 10, 11},
				{12},
			},
			want: [][]byte{
				{0, 3, 6, 9, 12},
				{1, 4, 7, 10},
				{2, 5, 8, 11},
			},
		},
		{
			bb: [][]byte{
				{0, 1, 2},
				{3, 4, 5},
				{6, 7, 8},
				{9, 10, 11, 12},
			},
			want: nil,
		},
	}

	for _, tc := range tests {
		got := Transpose(tc.bb)

		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
