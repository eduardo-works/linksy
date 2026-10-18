package configuration

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadOrCreate(t *testing.T) {
	original := "~/.linksy"
	fileName = tempFile(t)

	t.Cleanup(func() {
		fileName = original
	})

	config, err := LoadOrCreate()
	if err != nil {
		t.Fatal(err)
	}

	if config.HTTP.Port != 8080 {
		t.Errorf("expected port 8080, got %d", config.HTTP.Port)
	}

	if _, err := os.Stat(fileName); err != nil {
		t.Fatalf("file was not created: %v", err)
	}
}

func TestUpdate(t *testing.T) {
	original := "~/.linksy"
	fileName = tempFile(t)

	t.Cleanup(func() {
		fileName = original
	})

	config, err := LoadOrCreate()
	if err != nil {
		t.Fatal(err)
	}

	if config.HTTP.Port == 9234 {
		t.Errorf("expected port 9234, got %d", config.HTTP.Port)
	}

	config.HTTP.Port = 9234
	config.Update()
}

func tempFile(t *testing.T) string {
	return filepath.Join(t.TempDir(), ".linksy")
}
