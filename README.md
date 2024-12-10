
# Batchify

[![tag](https://img.shields.io/github/tag/samber/go-gladia.svg)](https://github.com/samber/go-gladia/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.18-%23007d9c)
[![GoDoc](https://godoc.org/github.com/samber/go-gladia?status.svg)](https://pkg.go.dev/github.com/samber/go-gladia)
![Build Status](https://github.com/samber/go-gladia/actions/workflows/test.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/samber/go-gladia)](https://goreportcard.com/report/github.com/samber/go-gladia)
[![Coverage](https://img.shields.io/codecov/c/github/samber/go-gladia)](https://codecov.io/gh/samber/go-gladia)
[![Contributors](https://img.shields.io/github/contributors/samber/go-gladia)](https://github.com/samber/go-gladia/graphs/contributors)
[![License](https://img.shields.io/github/license/samber/go-gladia)](./LICENSE)

A wrapper for [Gladia](https://gladia.io) API.

## 🚀 Install

```sh
go get github.com/samber/go-gladia
```

This library is v0 and follows SemVer strictly.

Some breaking changes might be made to exported APIs before v1.0.0.

## 🤠 Getting started

[GoDoc: https://godoc.org/github.com/samber/go-gladia](https://godoc.org/github.com/samber/go-gladia)

```go
import (
	"fmt"
	"os"
	"strings"

	gladia "github.com/samber/go-gladia/pkg/client"
)

client := gladia.NewGladiaClient(os.Getenv("GLADIA_API_KEY"))

transcription, err := client.GetTranscription(gladia.TranscriptionRequest{
    AudioURL:       "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
    Diarization:    true,
    DetectLanguage: true,
})
if err != nil {
    panic(err)
}

// Print the transcription
fmt.Printf("Languages: %s\n\n", strings.Join(transcription.Result.Transcription.Languages, ", "))
fmt.Printf("Transcript: %s\n", transcription.Result.Transcription.FullTranscript)
```

## 🤝 Contributing

- Ping me on Twitter [@samuelberthe](https://twitter.com/samuelberthe) (DMs, mentions, whatever :))
- Fork the [project](https://github.com/samber/go-gladia)
- Fix [open issues](https://github.com/samber/go-gladia/issues) or request new features

Don't hesitate ;)

```bash
# Install some dev dependencies
make tools

# Run tests
make test
# or
make watch-test
```

## 👤 Contributors

![Contributors](https://contrib.rocks/image?repo=samber/go-gladia)

## 💫 Show your support

Give a ⭐️ if this project helped you!

[![GitHub Sponsors](https://img.shields.io/github/sponsors/samber?style=for-the-badge)](https://github.com/sponsors/samber)

## 📝 License

Copyright © 2024 [Samuel Berthe](https://github.com/samber).

This project is [MIT](./LICENSE) licensed.
