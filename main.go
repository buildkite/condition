package main

import (
	"fmt"
	"os"

	"github.com/buildkite/evaluate/repl"
)

func main() {
	fmt.Printf("Buildkite if condition evaluator\n")
	repl.Start(os.Stdin, os.Stdout)
}
