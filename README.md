# GoGemma

## Description

Simple Go server that takes a token, command, and text and returns response from Gemma (2B parameter Google LLM). Uses Redis to cache responses. Easy deployment with Fly.io.

Leverages [Gemma CPP](https://github.com/google/gemma.cpp).

Developed for the `RapidRead` feature in [GhostRemix](https://ghostremix.com).

## Prerequisites

- Go 1.22
- Make
- Air
- Tilt

## Quickstart

1. Create `.env` file from `.env.example`.
2. Download Gemma from [Kaggle](https://www.kaggle.com/models/google/gemma/gemmaCpp/2b-it-sfp) or our Google drive link [here](https://drive.google.com/file/d/1Blx_O2FWV2-h71uGia0wtRb-5IaDwRX_/view?usp=sharing).
3. Create `libs` directory and unpack zip content there.
4. Run `tilt up` in project root.
5. Test with the command below.

```
curl -X POST -H "Content-Type: application/json" -d '{
  "command": "Summarize this post; Reply only with the summary;",
  "token": "your_token_here",
  "text": "Your input text goes here..."
}' http://localhost:8081/askGemma
```

## Test Docker Build

1. Build image and run container with `make all`.

2. Clean image and container with `make clean-all`.

## Deploy to Fly.io

### Prerequisites

1. Create [Fly.io](https://fly.io) account.

2. Authenticate with `flyctl auth login`.

3. Create app with `flyctl launch --no-deploy`.

### GitHub Actions

1. Navigate to the newly created application in the Fly.io dashboard and get a deploy token.

2. Set secrets in GitHub repository settings.

3. Manually trigger by going to Actions tab and selecting `Deploy`. Click `Run workflow` and enter the branch name to deploy.
   - You can update this action to trigger on push to `main` by changing the `on` section of the workflow file to `push: [main]`
