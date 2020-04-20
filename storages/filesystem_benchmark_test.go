package storages

import (
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
)

func BenchmarkCode(b *testing.B) {
	dir, _ := homedir.Dir()
	storage := &Filesystem{}

	err := storage.Init(filepath.Join(dir, "shawty"))
	if err != nil {
		b.Error(err)
	}

	for i := 0; i < b.N; i++ {
		storage.Code()
	}
}
