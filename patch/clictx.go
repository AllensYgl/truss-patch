package patch

var CliContext = `
package cli

    import (
        "{{.ImportPath -}} /svc/server"
    )

    var Config server.Config

    func init() {
        Config.GRPCAddr = "127.0.0.1:8888"
	}
	`
