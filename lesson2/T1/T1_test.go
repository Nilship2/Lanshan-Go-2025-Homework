package main

import (
	"reflect"
	"testing"
)

func TestTrans(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected map[int]int
	}{
		{
			name:     "空切片",
			input:    []int{},
			expected: map[int]int{},
		},
		{
			name:     "普通情况",
			input:    []int{1, 2, 2, 3, 3, 3},
			expected: map[int]int{1: 1, 2: 2, 3: 3},
		},
		{
			name:     "单个元素",
			input:    []int{5},
			expected: map[int]int{5: 1},
		},
		{
			name:     "所有元素相同",
			input:    []int{1, 1, 1, 1},
			expected: map[int]int{1: 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := trans(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("trans(%v) = %v, 期望 %v", tt.input, result, tt.expected)
			}
		})
	}
}
