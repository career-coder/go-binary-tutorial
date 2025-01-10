package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	tmpDir := t.TempDir()
	inputFile := "./testdata/test.yaml"
	outputFile := filepath.Join(tmpDir, "output.json")
	args := []string{
		"--input",
		inputFile,
		"--output",
		outputFile,
	}
	require.NoError(t, run(args, io.Discard))

	got, err := os.ReadFile(outputFile)
	require.NoError(t, err)

	expect, err := os.ReadFile("./testdata/test.json")
	require.NoError(t, err)
	assert.Equal(t, string(expect), string(got))
}

func TestRunNoOutput(t *testing.T) {
	inputFile := "./testdata/test.yaml"
	args := []string{
		"--input",
		inputFile,
	}

	stdout := bytes.NewBuffer(nil)
	require.NoError(t, run(args, stdout))

	expect, err := os.ReadFile("./testdata/test.json")
	require.NoError(t, err)
	assert.Equal(t, string(expect), stdout.String())
}

func TestRunBadArgs(t *testing.T) {
	inputFile := "./testdata/not_exist.yaml"
	args := []string{
		"--input",
		inputFile,
	}

	stdout := bytes.NewBuffer(nil)
	assert.ErrorContains(t, run(args, stdout), "input YAML file")
}
