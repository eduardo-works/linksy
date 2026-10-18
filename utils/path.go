package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func ExpandHome(path string) (string, error) {
	if path != "~" && !strings.HasPrefix(path, "~/") {
		return path, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	home = filepath.Clean(home)

	if path == "~" {
		return home, nil
	}

	return filepath.Join(home, strings.TrimPrefix(path, "~/")), nil
}

func ExpandPath(path string) (string, error) {
	path, err := ExpandHome(path)
	if err != nil {
		return "", err
	}

	return filepath.Abs(path)
}
