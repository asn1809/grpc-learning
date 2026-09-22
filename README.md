# grpc-learning

## Quick Start
1. Go to https://grpc.io/ 
1. Click on Go, will provide you steps.
1. Summary of the steps:
   1. Install Go -- [link](https://go.dev/doc/install)
   1. Install protoc (Protocol Buffer) -- [link](https://protobuf.dev/installation/)
1. Go plugins installation:
   1. Install protocol compiler plugins for Go
   ```sh
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

## Next steps
1. Create a go mod "go mod init <packageName>
1. Create directory structure:
    greet/
      server
      client
      proto
1. Create a dummy.protoc in greet/proto directory
1. Using terminal run the command "protoc -Igreet/proto --go_out=. --go_opt=module=github.com/asn1809/grpc-learning --go-grpc_out=. --go-grpc_opt=module=github.com/asn1809/grpc-learning greet/proto/dummy.proto"
1. make greet will generate the proto go files from the dummy.proto