package main

import (
	"todo/internal/flags"
	"todo/internal/memory"
)

func main() {
	memory.Open()
	flags.ReadFlags()
	memory.Close()
}
