package meta

import (
	"path/filepath"
	"strings"
)

var filter func(path string) bool

func init() {
	filter = func(path string) bool {
		ext := strings.ToLower(filepath.Ext(path))
		if ext == ".jpeg" {
			return true
		}
		if ext == ".jpg" {
			return true
		}
		if ext == ".png" {
			return true
		}
		if ext == ".webp" {
			return true
		}
		return false
	}
}
