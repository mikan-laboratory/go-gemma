package model

import (
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
	cmd.Stdin = strings.NewReader(sanitized)
	output, err := cmd.CombinedOutput()

	return output, err
}
