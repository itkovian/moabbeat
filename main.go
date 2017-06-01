package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/hpcugent/moabbeat/beater"
)

func main() {
	err := beat.Run("moabbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
