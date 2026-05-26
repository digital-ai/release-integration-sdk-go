Run Go tests. Pass optional Go test arguments in `$ARGUMENTS`.

```bash
if [ -n "$ARGUMENTS" ]; then go test $ARGUMENTS; else go test -cover -v ./...; fi
```
