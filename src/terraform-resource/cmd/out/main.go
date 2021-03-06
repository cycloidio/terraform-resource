package main

import (
	"encoding/json"
	"log"
	"os"

	"terraform-resource/models"
	"terraform-resource/namer"
	"terraform-resource/out"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Expected path to sources as first arg")
	}
	sourceDir := os.Args[1]
	if err := os.Chdir(sourceDir); err != nil {
		log.Fatalf("Failed to access source dir '%s': %s", sourceDir, err)
	}

	req := models.OutRequest{}
	if err := json.NewDecoder(os.Stdin).Decode(&req); err != nil {
		log.Fatalf("Failed to read OutRequest: %s", err)
	}

	runner := out.Runner{
		SourceDir: sourceDir,
		LogWriter: os.Stderr,
		Namer:     namer.New(),
	}
	resp, err := runner.Run(req)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.NewEncoder(os.Stdout).Encode(resp); err != nil {
		log.Fatalf("Failed to write OutResponse: %s", err)
	}
}
