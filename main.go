package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// SA1006: printf-style function with non-constant format string
	format := "Hello %s"
	fmt.Printf(format, "World")

	// SA1007: invalid URL in http.Get
	// http.Get("http://")

	// SA1015: using time.Tick in a way that will leak
	// ticker := time.Tick(time.Second)
	// _ = ticker

	// SA1019: using deprecated function
	// strings.Title is deprecated
	title := strings.Title("hello world")
	fmt.Println(title)

	// SA4003: comparing unsigned integer with negative value
	var x uint = 5
	if x < -1 {
		fmt.Println("This will never execute")
	}

	// SA4006: value assigned to variable is never read
	y := 10
	y = 20
	fmt.Println("y was assigned but first value never used")

	// SA5007: infinite recursive call
	// recursiveFunc()

	// SA6005: inefficient string concatenation in loop
	var result string
	items := []string{"a", "b", "c", "d", "e"}
	for _, item := range items {
		result += item // Should use strings.Builder
	}
	fmt.Println(result)

	// SA1002: invalid format string
	fmt.Printf("Hello %d", "world") // %d expects int, got string

	// SA1008: non-canonical key in http.Header map
	// header := make(http.Header)
	// header["content-type"] = []string{"application/json"} // should be "Content-Type"

	// SA1021: using regexp.MustCompile in a loop or hot path
	for i := 0; i < 3; i++ {
		re := regexp.MustCompile(`\d+`) // Should be compiled once outside loop
		fmt.Println(re.MatchString(fmt.Sprintf("test%d", i)))
	}

	// SA9003: empty branch
	if x > 0 {
		// Empty if block
	}

	// ST1003: poorly named function (should be camelCase)
	some_function()

	// U1000: unused function will be detected if function is not used
	unusedVariable := "this is unused"
	_ = unusedVariable // This line prevents the unused variable error

	// SA4009: argument is overwritten before first use
	processValue(getValue(), getValue()) // Second call overwrites first
}

// ST1003: function name should be camelCase
func some_function() {
	fmt.Println("This function has a poorly named identifier")
}

// SA5008: duplicate key in composite literal
type Config struct {
	Host string
	Port int
}

func createConfig() Config {
	return Config{
		Host: "localhost",
		Port: 8080,
		Port: 9090, // Duplicate key
	}
}

// U1000: this function is unused (will be detected by staticcheck)
func unusedFunction() {
	fmt.Println("This function is never called")
}

func getValue() string {
	return "value"
}

func processValue(a, b string) {
	fmt.Printf("Processing: %s, %s\n", a, b)
}

// SA9004: only the first constant has an explicit type
const (
	First  int = 1
	Second     = 2 // Should have explicit type
	Third      = 3 // Should have explicit type
)

// Recursive function that would cause SA5007 if uncommented
/*
func recursiveFunc() {
	recursiveFunc() // Infinite recursion
}
*/
