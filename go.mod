module github.com/nebius/gosdk

go 1.22 // minimal supported version 1.22, tested all versions up to 1.24

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.31.0-20231030212536-12f9cba37c9d.2
	github.com/bufbuild/protovalidate-go v0.4.0
	github.com/cenkalti/backoff/v4 v4.3.0
	github.com/gofrs/flock v0.12.1
	github.com/golang-jwt/jwt/v4 v4.5.1
	github.com/google/uuid v1.6.0
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.1.0
	github.com/stretchr/testify v1.9.0
	go.uber.org/mock v0.2.0
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9
	golang.org/x/net v0.33.0
	golang.org/x/sync v0.10.0
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237
	google.golang.org/grpc v1.64.1
	google.golang.org/protobuf v1.33.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/antlr/antlr4/runtime/Go/antlr/v4 v4.0.0-20230512164433-5d1fd1a340c9 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/cel-go v0.18.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	github.com/stoewer/go-strcase v1.3.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240318140521-94a12d6c2237 // indirect
)
