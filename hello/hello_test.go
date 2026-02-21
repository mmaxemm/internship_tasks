package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	if Hello() != "Hello, world!" {
		t.Errorf("Expected 'Hello, world!', got '%s'", Hello())
	}
}

