package utils_test

import (
	"testing"

	"github.com/bopher/utils"
)

func TestFileExists(t *testing.T) {
	ok, err := utils.FileExists("./errors.go")
	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Fatal("failed!")
	}
}

func TestIsDirectory(t *testing.T) {
	ok, err := utils.IsDirectory(".")
	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Fatal("failed!")
	}
}

func TestFindFile(t *testing.T) {
	files := utils.FindFile(".", "ors_test.go$")
	if len(files) != 1 || files[0] != "errors_test.go" {
		t.Log(files)
		t.Fatal("failed!")
	}
}
