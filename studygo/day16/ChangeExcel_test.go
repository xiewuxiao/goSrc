package main

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func Test_step3(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			step3()
		})
	}
}
