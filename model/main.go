package model

import (
	"io"
	"log"
	"os/exec"
	"strings"
)

func sanitizeText(text string) string {
	sanitized := strings.ReplaceAll(text, "'", `'\''`)

	return sanitized
}

func AskGemma(input string) ([]byte, error) {
	sanitized := sanitizeText(input)

	cmd := exec.Command("./build/gemma", "--tokenizer", "./build/tokenizer.spm", "--compressed_weights", "./build/2b-it-sfp.sbs", "--model", "2b-it", "--verbosity", "0")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	defer stdin.Close()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	_, err = stdin.Write([]byte(sanitized + "\n\nPlease output double new line at the end.\n"))
	if err != nil {
		return nil, err
	}
	stdin.Close()

	var output []byte
	buffer := make([]byte, 4096)

	for {
		n, err := stdout.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		log.Printf("Model output: %s", string(buffer[:n]))

		output = append(output, buffer[:n]...)

		if strings.Contains(string(output), "\n\n") {
			return output, nil
		}
	}

	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	return output, nil
}
