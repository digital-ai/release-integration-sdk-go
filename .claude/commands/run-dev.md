Run a development command or Go package. Pass the command in `$ARGUMENTS`.

```bash
if [ -z "$ARGUMENTS" ]; then echo "Usage: /run-dev <command or go run args>"; else $ARGUMENTS; fi
```
