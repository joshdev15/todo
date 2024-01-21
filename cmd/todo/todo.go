package main

import (
	"todo/internal/flags"
	"todo/internal/memory"
)

func main() {
	memory.Open()
	defer memory.Close()
	flags.ReadFlags()
}
