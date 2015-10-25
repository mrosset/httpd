package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	if err := os.Chdir("testdata"); err != nil {
		t.Fatal(err)
	}
	main()
}
