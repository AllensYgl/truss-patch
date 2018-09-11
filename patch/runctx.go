package patch

var RunContext = `
package server

    import (
        "log"
        "net"
        // 3d Party
        "google.golang.org/grpc"

        // This Service
        pb "{{.PBImportPath -}}"
        "{{.ImportPath -}} /svc"
        "{{.ImportPath -}} /handlers"
    )

    // Config contains the required fields for running a server
    type Config struct {
        GRPCAddr string
    }

    func NewEndpoints() svc.Endpoints {
        // Business domain.
        var service pb.{{.Service.Name}}Server
        {
            service = handlers.NewService()
            // Wrap Service with middlewares. See handlers/middlewares.go
            service = handlers.WrapService(service)
        }

        // Endpoint domain.
        var (
        {{range $i := .Service.Methods -}}
            {{ToLower $i.Name}}Endpoint = svc.Make{{$i.Name}}Endpoint(service)
        {{end}}
        )

        endpoints := svc.Endpoints{
        {{range $i := .Service.Methods -}}
            {{$i.Name}}Endpoint:    {{ToLower $i.Name}}Endpoint,
        {{end}}
        }

        // Wrap selected Endpoints with middlewares. See handlers/middlewares.go
        endpoints = handlers.WrapEndpoints(endpoints)

        return endpoints
    }

    // Run starts a new http server, gRPC server, and a debug server with the
    // passed config and logger
    func Run(cfg Config) {
        endpoints := NewEndpoints()

        // Mechanical domain.
        errc := make(chan error)

        // Interrupt handler.
        go handlers.InterruptHandler(errc)

        // gRPC transport.
        go func() {
            log.Println("transport", "gRPC","addr", cfg.GRPCAddr)
            ln, err := net.Listen("tcp", cfg.GRPCAddr)
            if err != nil {
                errc <- err
                return
            }

            srv := svc.MakeGRPCServer(endpoints)
            s := grpc.NewServer()
            pb.Register{{.Service.Name}}Server(s, srv)

            errc <- s.Serve(ln)
        }()

        // Run!
        log.Println("exit", <-errc)
	}
	`
