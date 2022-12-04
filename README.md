# FScanner

Scan filesystem for things.
```bash
fscanner
  --path "path-to-directory-to-start-scanning"
  [--dirsonly]"
```

At the moment, this is just a dummy project, which returns the found objects as output.

## Examples

Scan all objects in filesystem on given path:

`go run cmd/fscanner/main.go --path "path-to-directory-to-start-scanning"`

Scan for only directories in filesystem on given path:

`go run cmd/fscanner/main.go --path "path-to-directory-to-start-scanning" --dirsonly`
