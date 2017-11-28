package util

import (
	"os"
	"testing"
)

func TestWriteXlsFile(t *testing.T) {
	file := "test.xlsx"

	headers := []string{
		"field1", "field2", "field2",
	}

	items := [][]string{
		{"value for row 1 field 1", "value for row 1 field 2", "value for row 1 field 3"},
		{"value for row 2 field 1", "value for row 2 field 2", "value for row 2 field 3"},
		{"value for row 3 field 1", "value for row 3 field 2", "value for row 3 field 3"},
	}

	err := WriteXlsFile(file, headers, items)
	if err != nil {
		t.Fatalf("failed to create excel file: %v", err)
	}

	f, err := os.Open(file)
	if err != nil {
		t.Fatalf("failed to create excel file: %v", err)
	}

	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		t.Fatalf("failed to get file information: %v", err)
	}

	// 6117 is file size with headers only
	if fi.Size() <= 6117 {
		t.Fatal("failed to create excel file: probably file is empty")
	}

	err = os.Remove(file)
	if err != nil {
		t.Fatalf("failed to remove excel file: %v", err)
	}
}
