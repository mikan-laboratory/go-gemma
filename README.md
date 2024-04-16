# Go Gemma

## Description

Simple Go server that takes a token, command, and text and returns response from Gemma (2B parameter Google LLM). Uses Redis to cache responses. Easy deployment with Fly.io.

Developed for the `RapidRead` feature in [GhostRemix](https://ghostremix.com).

## Prerequisites

- Go 1.22
- Make
- Air
- Tilt

## Quickstart

1. Create `.env` file from `.env.example`.
2. Download Gemma zip [here](https://drive.google.com/file/d/1UexYG4stAjwyryQZxckJDxyi5A0w7WjN/view?usp=drive_link).
3. Create `build` folder and extract zip content there.
4. Run `tilt up` in project root.
5. Test with the command below.

```
curl -X POST -H "Content-Type: application/json" -d '{
  "command": "Summarize this post; Reply only with the summary;",
  "token": "your_token_here",
  "text": "Your input text goes here..."
}' http://localhost:8081/askGemma
```
