package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//go:generate go run gen.go

func copy(src, dst string) error {
	f, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(dst, f, 0644); err != nil {
		return err
	}

	return nil
}

func main() {
	matches, err := filepath.Glob("./*.tmpl")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, src := range matches {
		ext := filepath.Ext(src)
		dst := "../" + strings.TrimSuffix(src, ext) + ".go"
		if err := copy(src, dst); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
