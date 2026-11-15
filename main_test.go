package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
	}{
		{name: "negative size", size: -10, wantLen: 0},
		{name: "zero size", size: 0, wantLen: 0},
		{name: "one element", size: 1, wantLen: 1},
		{name: "several elements", size: 50, wantLen: 50},
		{name: "large slice", size: 10000, wantLen: 10000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)
			if len(data) != tt.wantLen {
				t.Fatalf("generateRandomElements(%d): length = %d, want %d", tt.size, len(data), tt.wantLen)
			}
			for i, value := range data {
				if value < 1 || value > 1_000_000 {
					t.Fatalf("data[%d] = %d: expected value from 1 to 1000000", i, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "one element", data: []int{7}, want: 7},
		{name: "one zero", data: []int{0}, want: 0},
		{name: "max first", data: []int{9, 3, 1}, want: 9},
		{name: "max middle", data: []int{1, 9, 3}, want: 9},
		{name: "max last", data: []int{1, 3, 9}, want: 9},
		{name: "duplicate values", data: []int{3, 3, 3}, want: 3},
		{name: "negative elements", data: []int{-9, -3, -5}, want: -3},
		{name: "mixed elements", data: []int{-5, 0, 2, -1}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			if got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "one element", data: []int{7}, want: 7},
		{name: "less than eight", data: []int{1, 4, 2}, want: 4},
		{name: "exactly eight", data: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: 8},
		{name: "nine elements", data: []int{9, 1, 2, 3, 4, 5, 6, 7, 8}, want: 9},
		{name: "remainder has maximum", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 99}, want: 99},
		{name: "negative and short", data: []int{-9, -3, -5}, want: -3},
		{name: "negative and long", data: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1}, want: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := append([]int(nil), tt.data...)
			got := maxChunks(tt.data)
			if got != tt.want {
				t.Errorf("maxChunks(%v) = %d, want %d", tt.data, got, tt.want)
			}
			for i, value := range before {
				if tt.data[i] != value {
					t.Error("maxChunks modified the input slice")
					break
				}
			}
		})
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	data := generateRandomElements(100_003)
	if got, want := maxChunks(data), maximum(data); got != want {
		t.Errorf("maxChunks(data) = %d, maximum(data) = %d", got, want)
	}
}
