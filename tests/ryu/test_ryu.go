package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	ryu "github.com/specodec/specodec-runtime-golang/ryu"
)

func loadTests(filename string) []float64 {
	data, _ := os.ReadFile(filename)
	var result []float64
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || line[0] == '#' {
			continue
		}
		v, _ := strconv.ParseFloat(line, 64)
		result = append(result, v)
	}
	return result
}

func loadExpected(filename string) []string {
	data, _ := os.ReadFile(filename)
	var result []string
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			result = append(result, line)
		}
	}
	return result
}

func loadCoverage(filename string) []float64 {
	data, _ := os.ReadFile(filename)
	var result []float64
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || (line[0] < '0' || line[0] > '9') {
			continue
		}
		if idx := strings.IndexByte(line, '#'); idx >= 0 {
			line = strings.TrimSpace(line[:idx])
		}
		v, _ := strconv.ParseFloat(line, 64)
		result = append(result, v)
	}
	return result
}

func main() {
	base, _ := os.Getwd()
	passed, failed := 0, 0

	fmt.Println("=== Float32 Original (125 tests) ===")
	f32in := loadTests(filepath.Join(base, "test_cases_f32.txt"))
	f32exp := loadExpected(filepath.Join(base, "expected_f32.txt"))
	for i := 0; i < len(f32in) && i < len(f32exp); i++ {
		result := ryu.Float32ToString(float32(f32in[i]))
		if result == f32exp[i] {
			passed++
		} else {
			failed++
			if failed <= 5 {
				fmt.Printf("FAIL: %v => %s (expected %s)\n", f32in[i], result, f32exp[i])
			}
		}
	}
	fmt.Printf("%d/%d\n\n", len(f32in), len(f32in))

	fmt.Println("=== Float64 Original (102 tests) ===")
	f64in := loadTests(filepath.Join(base, "test_cases_f64.txt"))
	f64exp := loadExpected(filepath.Join(base, "expected_f64.txt"))
	for i := 0; i < len(f64in) && i < len(f64exp); i++ {
		result := ryu.Float64ToString(f64in[i])
		if result == f64exp[i] {
			passed++
		} else {
			failed++
			if failed <= 5 {
				fmt.Printf("FAIL: %v => %s (expected %s)\n", f64in[i], result, f64exp[i])
			}
		}
	}
	fmt.Printf("%d/%d\n\n", len(f64in), len(f64in))

	fmt.Println("=== Float32 Coverage (78 tests) ===")
	c32in := loadCoverage(filepath.Join(base, "test_cases_table_coverage.txt"))
	c32exp := loadExpected(filepath.Join(base, "expected_table_coverage.txt"))
	n := int(math.Min(float64(len(c32in)), float64(len(c32exp))))
	for i := 0; i < n; i++ {
		result := ryu.Float32ToString(float32(c32in[i]))
		if result == c32exp[i] {
			passed++
		} else {
			failed++
			if failed <= 5 {
				fmt.Printf("FAIL: %v => %s (expected %s)\n", c32in[i], result, c32exp[i])
			}
		}
	}
	fmt.Printf("%d/%d\n\n", n, n)

	fmt.Println("=== Float64 Coverage (616 tests) ===")
	c64in := loadCoverage(filepath.Join(base, "test_cases_f64_table_coverage.txt"))
	c64exp := loadExpected(filepath.Join(base, "expected_f64_table_coverage.txt"))
	n = int(math.Min(float64(len(c64in)), float64(len(c64exp))))
	for i := 0; i < n; i++ {
		result := ryu.Float64ToString(c64in[i])
		if result == c64exp[i] {
			passed++
		} else {
			failed++
			if failed <= 5 {
				fmt.Printf("FAIL: %v => %s (expected %s)\n", c64in[i], result, c64exp[i])
			}
		}
	}
	fmt.Printf("%d/%d\n\n", n, n)

	fmt.Printf("=== TOTAL: %d/%d ===\n", passed, passed+failed)
	if failed > 0 {
		os.Exit(1)
	}
}
