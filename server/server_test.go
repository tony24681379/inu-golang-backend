package server

import (
	"testing"
)

func test(a, b int) int {
	return a + b
}

func TestServer(t *testing.T) {
	if test(1, 2) != 3 {
		t.Error("error")
	}
}
