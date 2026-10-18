package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExpandHome(t *testing.T) {
	expanded, err := ExpandHome("~")
	if err != nil {
		t.Fatal(err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}

	if expanded != filepath.Clean(home) {
		t.Errorf("expected %q, got %q", filepath.Clean(home), expanded)
	}
}
