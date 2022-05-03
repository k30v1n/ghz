# ghz

- ghz is a command line utility and Go package for load testing and benchmarking gRPC services
- intended to be used for testing and debugging services locally
- and in automated continous intergration environments for performance regression testing
- implemented as a Go library package 
- which can be used to programatically to implement performance tests as well

# Installation
- CLI version
- Golang package 1.11+ (https://jimkang.medium.com/install-go-on-mac-with-homebrew-5fa421fc55f5)    
    - go env -w GO111MODULE=auto
    - go mod init ghz
    - go get
    - go run unary_SayHello.go

# Links
- docs https://ghz.sh/
- source https://github.com/bojand/ghz
- go install https://go.dev/doc/install