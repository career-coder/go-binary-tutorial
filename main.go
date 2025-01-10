package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

const _name = "yamlToJSON"

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {
	flag := flag.NewFlagSet(_name, flag.ExitOnError)
	input := flag.String("input", "", "input YAML document to convert")
	output := flag.String("output", "", "output JSON document (default to Stdout)")
	if err := flag.Parse(args); err != nil {
		return err
	}

	if input == nil || *input == "" {
		return errors.New("--input flag is required")
	}

	yamlBytes, err := os.ReadFile(*input)
	if err != nil {
		return fmt.Errorf("input YAML file failed to open: %w", err)
	}

	data := make(map[string]interface{})
	if err := yaml.Unmarshal(yamlBytes, data); err != nil {
		return err
	}

	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	if output == nil || *output == "" {
		return writeBytes(stdout, jsonBytes)
	}

	f, err := os.OpenFile(*output, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("output JSON file failed to open: %w", err)
	}
	defer f.Close()
	return writeBytes(f, jsonBytes)
}

func writeBytes(out io.Writer, bs []byte) error {
	_, err := out.Write(bs)
	return err
}
