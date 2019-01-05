# mailgunner

A super simplistic send only implementation of a [Mailgun](https://www.mailgun.com/) client in Go.

## Installation

Use as a module, import `github.com/shindakun/mailgunner/v2`, then run

```bash
go mod init module
go build
```

If you are new to modules checkout my post on [Dev.to for some more details](https://dev.to/shindakun/attempting-to-learn-go---lets-get-modular-390i).

```bash
go get github.com/shindakun/mailgunner
```

## Usage

See [example/main.go](./example/main.go) for a usage example.

## Change log

- Updated `New()` to take http.Client so we don't have to rely on `http.DefaultClient`