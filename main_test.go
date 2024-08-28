package main

import "testing"

func TestRectangleArea(t *testing.T) {
	tests := []struct {
		length int
		width  int
		want   int
	}{
		{length: 2, width: 3, want: 6},   // 2 * 3 = 6
		{length: 5, width: 5, want: 25},  // 5 * 5 = 25
		{length: 10, width: 2, want: 20}, // 10 * 2 = 20
		{length: 0, width: 5, want: 0},   // 0 * 5 = 0
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := RectangleArea(tt.length, tt.width)
			if got != tt.want {
				t.Errorf("RectangleArea(%d, %d) = %d, want %d", tt.length, tt.width, got, tt.want)
			}
		})
	}
}

func TestRectanglePerimeter(t *testing.T) {
	tests := []struct {
		length int
		width  int
		want   int
	}{
		{length: 2, width: 3, want: 10},  // 2*(2 + 3) = 10
		{length: 5, width: 5, want: 20},  // 2*(5 + 5) = 20
		{length: 10, width: 2, want: 24}, // 2*(10 + 2) = 24
		{length: 0, width: 5, want: 10},  // 2*(0 + 5) = 10
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := RectanglePerimeter(tt.length, tt.width)
			if got != tt.want {
				t.Errorf("RectanglePerimeter(%d, %d) = %d, want %d", tt.length, tt.width, got, tt.want)
			}
		})
	}
}

func TestSquareArea(t *testing.T) {
	tests := []struct {
		side int
		want int
	}{
		{side: 2, want: 4},    // 2 * 2 = 4
		{side: 5, want: 25},   // 5 * 5 = 25
		{side: 10, want: 100}, // 10 * 10 = 100
		{side: 0, want: 0},    // 0 * 0 = 0
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := SquareArea(tt.side)
			if got != tt.want {
				t.Errorf("SquareArea(%d) = %d, want %d", tt.side, got, tt.want)
			}
		})
	}
}

func TestSquarePerimeter(t *testing.T) {
	tests := []struct {
		side int
		want int
	}{
		{side: 2, want: 8},   // 4 * 2 = 8
		{side: 5, want: 20},  // 4 * 5 = 20
		{side: 10, want: 40}, // 4 * 10 = 40
		{side: 0, want: 0},   // 4 * 0 = 0
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := SquarePerimeter(tt.side)
			if got != tt.want {
				t.Errorf("SquarePerimeter(%d) = %d, want %d", tt.side, got, tt.want)
			}
		})
	}
}
