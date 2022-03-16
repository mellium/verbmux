# verbmux

[![Issue Tracker][badge]](https://github.com/mellium/xmpp/issues)
[![Docs](https://pkg.go.dev/badge/mellium.im/verbmux)](https://pkg.go.dev/mellium.im/verbmux)
[![License](https://img.shields.io/badge/license-FreeBSD-blue.svg)](https://opensource.org/licenses/BSD-2-Clause)

<a href="https://opencollective.com/mellium" alt="Donate on Open Collective"><img src="https://opencollective.com/mellium/donate/button@2x.png?color=blue" width="200"/></a>

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

[badge]: https://img.shields.io/badge/style-mellium%2fxmpp-green.svg?longCache=true&style=popout-square&label=issues
