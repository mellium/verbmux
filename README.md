# verbmux

*verbmux* is an HTTP multiplexer that can be used to route HTTP requests in Go
programs based on their verb (GET, SET, HEAD, etc).
It handles OPTIONS requests automatically and is designed to be composable with
other HTTP routers.

Use it in your program with:

```go
import (
  "mellium.im/verbmux"
)
```
