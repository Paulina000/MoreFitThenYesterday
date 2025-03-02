package main

import (
	"testing"
)

func Test_findAge(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAge(); got != tt.want {
				t.Errorf("findAge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findGender(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findGender(); got != tt.want {
				t.Errorf("findGender() = %v, want %v", got, tt.want)
			}
		})
	}
}
