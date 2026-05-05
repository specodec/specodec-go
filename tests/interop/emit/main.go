package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func tryTest(name string, fn func()) (passed, failed int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("FAIL " + name + ": " + fmt.Sprint(r))
			failed = 1
		}
	}()
	fn()
	passed = 1
	return
}

func main() {
	vecDir := os.Getenv("VEC_DIR")
	outDir := os.Getenv("OUT_DIR")
	os.MkdirAll(filepath.Join(outDir, "scalars"), 0755)

	passed := 0
	failed := 0
	var p, f int

	p, f = runScalars(vecDir, outDir); passed += p; failed += f
	p, f = runAllTypes(vecDir, outDir); passed += p; failed += f

	fmt.Printf("emit-go: %d passed, %d failed\n", passed, failed)
	if failed > 0 { os.Exit(1) }
}
