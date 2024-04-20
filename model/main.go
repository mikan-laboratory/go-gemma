package model

import (
	"io"
	"os/exec"
	"strings"
)

func sanitizeText(text string) string {
	sanitized := strings.ReplaceAll(text, "'", `'\''`)
	return sanitized
}

func AskGemma(input string) ([]byte, error) {
	sanitized := sanitizeText(input)

	cmd := exec.Command("./build/_deps/gemma-build/gemma", "--tokenizer", "./build/tokenizer.spm", "--compressed_weights", "./build/2b-it-sfp.sbs", "--model", "2b-it", "--verbosity", "0")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	_, err = stdin.Write([]byte(sanitized + "\n"))
	if err != nil {
		return nil, err
	}
	stdin.Close()

	output, err := io.ReadAll(stdout)
	if err != nil {
		return nil, err
	}

	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	return output, nil
}
