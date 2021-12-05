package utils_test

import (
	"testing"

	"github.com/bopher/utils"
)

func TestExtractNumbers(t *testing.T) {
	res := utils.ExtractNumbers("This is text with 123 and 456")
	if res != "123456" {
		t.Log(res)
		t.Fatal("failed!")
	}
}

func TestRandomStringFromCharset(t *testing.T) {
	res, err := utils.RandomStringFromCharset(5, "1234567")
	if err != nil {
		t.Fatal(err)
	}

	if len(res) != 5 {
		t.Log(res)
		t.Fatal("failed!")
	}
}

func TestRandomString(t *testing.T) {
	res, err := utils.RandomString(5)
	if err != nil {
		t.Fatal(err)
	}

	if len(res) != 5 {
		t.Log(res)
		t.Fatal("failed!")
	}
}

func TestSlugify(t *testing.T) {
	res := utils.Slugify("Hel", "    llo", "world")
	if res != "Hel-llo-world" {
		t.Log(res)
		t.Fatal("failed!")
	}
}

func TestConcatStr(t *testing.T) {
	res := utils.ConcatStr(" ", "John", "", "Doe")
	if res != "John Doe" {
		t.Log(res)
		t.Fatal("failed!")
	}
}

func TestFormatNumber(t *testing.T) {
	res := utils.FormatNumber("Total: %d$", 1230004)
	if res != "Total: 1,230,004$" {
		t.Log(res)
		t.Fatal("failed!")
	}
}
