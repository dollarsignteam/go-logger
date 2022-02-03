# Logger for Golang

It is based on the [zap][1] package.

## Install

```shell
go get github.com/dollarsignteam/go-logger
```

## Usage

```go
package main
import (
  "github.com/dollarsignteam/go-logger"
)
func main() {
  opts := logger.LoggerOptions{
    Level: "debug",
    Name:  "DEMO",
  }
  log := logger.NewLogger(opts)
  log.Debug("I am a debug log")
  log.Info("I am a info log")
  log.Warn("I am a warn log")
  log.Error("I am a error log")
}
```

Output

```shell
2022-02-03 14:48:58.277 +07:00 [go] 🟪 DEBUG  [DEMO] [go-logger/main.go:13 main.main] I am a debug log
2022-02-03 14:48:58.277 +07:00 [go] ⬜️ INFO   [DEMO] [go-logger/main.go:14 main.main] I am a info log
2022-02-03 14:48:58.277 +07:00 [go] 🟧 WARN   [DEMO] [go-logger/main.go:15 main.main] I am a warn log
2022-02-03 14:48:58.277 +07:00 [go] 🟥 ERROR  [DEMO] [go-logger/main.go:16 main.main] I am a error log
```

## Author

Dollarsign

## License

Licensed under the MIT License - see the [LICENSE][2] file for details.

[1]: https://github.com/uber-go/zap
[2]: https://github.com/dollarsignteam/go-logger/blob/main/LICENSE
