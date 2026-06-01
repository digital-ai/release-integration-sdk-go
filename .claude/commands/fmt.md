Format Go files. Pass files or package paths in `$ARGUMENTS`; defaults to all non-generated Go files outside `.git`.

```bash
if [ -n "$ARGUMENTS" ]; then gofmt -w $ARGUMENTS; else gofmt -w $(git ls-files '*.go' ':!:api/release/openapi/*'); fi
```
