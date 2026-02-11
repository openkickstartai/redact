# redact

Find leaked secrets in git repos. Scans entire history.

## Install

```bash
git clone https://github.com/openkickstartai/redact.git
cd redact && go build -o redact .
```

## Usage

```bash
redact scan .                    # scan current dir
redact scan . --format sarif     # SARIF for GitHub
redact hook install              # pre-commit hook
```

## Detects

AWS keys, GitHub tokens, private keys, Slack webhooks, DB connection strings, high-entropy strings.

## Testing

```bash
go test -v ./...
```
