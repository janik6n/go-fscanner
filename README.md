# FScanner

Scan filesystem for things.

## API

```bash
fscanner
  --path "aws"
  --dirsonly"
```

## Examples

Scan all objects in filesystem on given path:
`go run cmd/fscanner/main.go --path "path-to-directory-to-start-scanning"`

Scan for only directories in filesystem on given path:
`go run cmd/fscanner/main.go --path "path-to-directory-to-start-scanning" --dirsonly`
