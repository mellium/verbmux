# verbmux

[![Issue Tracker][badge]](https://github.com/mellium/xmpp/issues)
[![GoDoc](https://godoc.org/mellium.im/verbmux?status.svg)](https://pkg.go.dev/mellium.im/verbmux)
[![Chat](https://inverse.chat/badge.svg?room=mellium@conference.samwhited.com)](https://conversations.im/j/mellium@conference.samwhited.com)
[![License](https://img.shields.io/badge/license-FreeBSD-blue.svg)](https://opensource.org/licenses/BSD-2-Clause)

[![Buy Me A Coffee](https://www.buymeacoffee.com/assets/img/custom_images/purple_img.png)](https://www.buymeacoffee.com/samwhited)
[![Support Me](https://liberapay.com/assets/widgets/donate.svg)](https://liberapay.com/SamWhited/donate)

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
