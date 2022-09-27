# HoroscopeService Codelab

## Overview

### What you will build
TODO

### gRPC
TODO

### Protocol Buffers 
TODO

[Protocol buffer style guide](https://developers.google.com/protocol-buffers/docs/style)

## Define data structures
Let's get started by defining the data structures for our application. One of the benefits of protobufs is that it provides a language-neutral way to define our datatypes. This means we get auto-generated Typescript type definitions and Go structs from our protobufs.

Let's create a new file in the `pb/` directory called `horoscope_models.proto`. The definitions in a `.proto` file are simple: you add a message for each data structure you want to serialize, then specify a name and a type for each field in the message.
```proto=
syntax = "proto3";

package horoscope;

option go_package = "horoscope/pb";

enum ZodiacSign {
  ZODIAC_SIGN_UNSPECIFIED = 0;
  ZODIAC_SIGN_ARIES = 1;
  ZODIAC_SIGN_TAURUS = 2;
  ZODIAC_SIGN_GEMINI = 3;
  ZODIAC_SIGN_CANCER = 4;
  ZODIAC_SIGN_LEO = 5;
  ZODIAC_SIGN_VIRGO = 6;
  ZODIAC_SIGN_LIBRA = 7;
  ZODIAC_SIGN_SCORPIO = 8;
  ZODIAC_SIGN_SAGITTARIUS = 9;
  ZODIAC_SIGN_CAPRICORN = 10;
  ZODIAC_SIGN_AQUARIUS = 11;
  ZODIAC_SIGN_PISCES = 12;
}

message Fortune {
  string fortune = 1;
}
```

- The first line of the file specifies that you're using proto3 syntax. This must be the first non-empty, non-comment line of the file.
- The `package` defintion defines a namespace for the protobuf definitions to prevent naming conflicts. This is useful for grouping together definitions across multiple files into a common namespace, and should be a unique name based on the project name, and possibly based on the path of the file containing the protocol buffer type definitions.
- The `go_package` option defines the Go package import path for the generated code. This must be provided for any `.proto` files you intend to compile into a Go package. The Go package name will be the last path component of the import path. For example, our example will use a package name of "pb".
- `ZodiacSign` is an [enum](https://developers.google.com/protocol-buffers/docs/proto3#enum) that defines a pre-defined set of values: in this case, Zodiac signs.
- The `Fortune` [message](https://developers.google.com/protocol-buffers/docs/proto3#simple) definition specifies a data structure with one field (a field is a name/value pair) of type `string`.

## Define the service
Now that we have our data structure definitions, the next step is to define the RPC [service](https://developers.google.com/protocol-buffers/docs/proto3#services) and the method request and response types. To do so, create another file in `pb/` named `horoscope_service.proto`.

To define a service, you specify a named service in your `.proto` file:
```proto=
syntax = "proto3";

package horoscope;

import "horoscope_models.proto";

option go_package = "horoscope/pb";

service HoroscopeService {
   ...
}
```
**Note the `import`:** We can access our message definitions we defined in `horoscope_models.proto` by importing it.

Now you can define `rpc` methods inside your service definition, specifying their request and response types:
```proto=
service HoroscopeService {
  // Gets the daily horoscope for the given zodiac sign.
  rpc GetDailyHoroscope(GetDailyHoroscopeRequest) returns (GetDailyHoroscopeResponse);

  // Gets the horoscope for the given zodiac sign and date.
  rpc GetHoroscope(GetHoroscopeRequest) returns (GetHoroscopeResponse);

  // Suggest a fortune.
  rpc SuggestFortune(SuggestFortuneRequest) returns (SuggestFortuneResponse);
}
```

Our `horoscope_service.proto` file also contains protobuf message type definitions for all the request and response types used in our service methods:
```proto=
message GetDailyHoroscopeRequest {
   ZodiacSign zodiac_sign = 1;
}

message GetDailyHoroscopeResponse {
  Fortune fortune = 1;
}

message GetHoroscopeRequest {
  ZodiacSign zodiac_sign = 1;
  uint32 date = 2;
}

message GetHoroscopeResponse {
  Fortune fortune = 1;
}

message SuggestFortuneRequest {
  ZodiacSign zodiac_sign = 1;
  string suggestion = 2;
}

message SuggestFortuneResponse {
  bool success = 1;
}
```

## Compiling protobufs
Now that we defined our data structures and RPC service with protobuf, we can compile the `.proto` files into Go structs and interfaces. To do this, we need to use a command line tool called `protoc`. If you haven't installed the compiler, [download the package](https://developers.google.com/protocol-buffers/docs/downloads) and follow the instructions in the README.

Then, run the following command to install the Go protocol buffers plugin:
```bash=
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

Move into the `pb/` directory:
```bash=
cd pb
```

Create a directory to hold the generated `.go` files:
```bash=
mkdir out
```

Compile the protobufs:
```bash=
protoc --go_out=out --go_opt=paths=source_relative \
--go-grpc_out=out --go-grpc_opt=paths=source_relative \
*.proto
```

## Implement the service
Now that we've generated a Go package from our `.proto` files, we're ready to implement the business logic of our gRPC server.

In your Go project, create a new package called `service` and create a file named `server.go`.

In this file, define a struct that will implement the generated `pb.HoroscopeServer` interface:
```go=
package horoscope

import (
	pb "horoscope/pb/out"
)

// service is a struct that implements the interface generated the protocol buffer service definition.
type service struct {
	*pb.UnimplementedHoroscopeServiceServer
}

// GetDailyHoroscope gets the daily horoscope for the given zodiac sign.
func (s *service) GetDailyHoroscope(ctx context.Context, getDailyHoroscopeRequest *pb.GetDailyHoroscopeRequest) (*pb.GetDailyHoroscopeResponse, error) {
	return nil, aether.NotYetImplementedError("Implement me!")
}

// GetHoroscope gets the horoscope for the given zodiac sign and date.
func (s *service) GetHoroscope(ctx context.Context, createLinkRequest *pb.GetHoroscopeRequest) (*pb.GetHoroscopeResponse, error) {
	return nil, aether.NotYetImplementedError("Implement me!")
}

// SuggestFortune suggests a fortune.
func (s *service) SuggestFortune(ctx context.Context, suggestFortuneRequest *pb.SuggestFortuneRequest) (*pb.SuggestFortuneResponse, error) {
	return nil, aether.NotYetImplementedError("Implement me!")
}
```

Notice how we embed `*pb.UnimplementedHoroscopeServiceServer` in the struct. This is required by gRPC, and provides fallback functionality in case you don't implement an `rpc` method defined in `horoscope_service.proto`.

Now that we've implemented our service, we need to configure it as an Aether module. A module is a self-contained, reusable piece of functionality that can be used by an Aether server. Modules are the fundamental building blocks of an Aether server: they are the API by which users configure and extend Aether servers. 

An Aether module requires you to implement two methods: `Name() string` and `Configure(c *aether.ServerConfig) error`. The `Name()` method returns a string that provides a human-readable name for the module (useful for debugging), and the `Configure` method allows you to register your service with gRPC (and add interceptors and more). You can think of the `Configure` method as a container for the configuration you would typically do in `main.go`.

```go=
// module is a struct that implements Module. This module represents the Horoscope service.
type module struct {
	service *service
}

// Name is a method that returns a human-readable name for the module.
func (m *module) Name() string {
	return "Horoscope service"
}

// Configure is a function that is called with a `ServerConfig`. It can be used to install interceptors, register
// gRPC services, and more.
func (m *module) Configure(c *aether.ServerConfig) error {
	pb.RegisterHoroscopeServiceServer(c.GetGRPCServer(), m.service)
	return nil
}

// Module is a function that creates an instance of the `module` struct.
func Module(horoscopeRepository Repository) *module {
	return &module{
		service: &service{},
	}
}
```

## Install and run your service
One of the main advantages of Aether is a small `main.go` file. Typically, your `main.go` file would be fairly large, but Aether modules abstract away much of the configuration you typically would need to do yourself. Through Aether modules, configuration is coupled with the module that requires it. By installing modules on your server (via the `configure` function), Aether will configure your modules individually when your server is run.

```go=
package main

import (
	"flag"
	"horoscope/service"
	aether "github.com/nthnluu/aether/pkg"
)

var (
	port = flag.Int("port", 9999, "Port for serving gRPC requests")
)

func configure(c *aether.ServerConfig) error {
  // Install the Horoscope service module.
	c.InstallModule(service.Module())
	return nil
}

func main() {
	flag.Parse()

	// Run the server.
	aether.RunOrExit(configure, *port)
}
```
