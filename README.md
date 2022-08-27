# Aether
Aether is a microservice framework that lets you build robust, production-quality services with fast build times and minimal-configuration production infrastructure.

> Prefer to learn by example? **Check out the [HoroscopeService](https://github.com/nthnluu/aether/tree/main/examples/horoscope) example!**

## Concepts

### Services
A service is a set of one or more callable endpoints. A service describes the interface (in terms of RPCs) that a server provides.

### Server
A server is a set of one or more services can be run to serve RPC requests.

### Interceptor
An interceptor is a piece of code that extends the behavior of Aether RPC clients and servers. Interceptors are hooks that run before and after your actions and can be used to implement authentication, logging, instrumentation, etc.

### Module
A module is a self-contained, reusable piece of functionality that can be used by an Aether server. Modules are the fundamental building blocks of an Aether server: they are the API by which users configure and extend Aether servers: they can add interceptors and register services.

## Protobufs
Protobuf is an interface definition language (IDL) that allows us to define data types and services in a language neutral format. They are great because they provide a self-documented description of the service and its datatypes, plus shared types between languages, 

To add a new `.proto` file, add it to the `pb/` directory. Then, compile the Go proto library:
```shell
cd pb
./compile.sh
```
If `./compile.sh` returns a permission error, you can grant execute permissions by running:
```shell
chmod +x compile.sh # then rerun `./compile.sh`
```

## Creating a service

Create a service by writing a `<your service name>_service.proto` file and calling `aether create service pb/<your service name>_service.proto`

## Modules
A module is a self-contained, reuseable piece of functionality that servers built with Aether can use. It is the fundamental building block of an Aether server, and is the foundation in which users can:
- Set up a connection to another gRPC endpoint
- Register a gRPC service
- Add interceptors

This enables you to build features that are easy to share and reuse, such as:
- Database handles
- A high-level API for accessing a database
- Authentication
- Logging and metrics
- Glueing together multiple modules

To create a module, you must simply implement the `Module` interface, and register it by calling `b.RegisterModule(module)` in the server's `configure()` function.

## Best Practices

### Seperate datasource and business logic
Always keep your data logic (i.e. code that queries a database or makes an API call) seperate from your business logic. By decoupling these two things, your business logic code easier to read, and you make it easier to swap out databases for testing.

To do this, you should define a `Repository` interface inside `respository.go`. On this interface, you can define an abstract API for accessing the data so that your business logic is agnostic to the underlying data source(s).

See [examples/horoscope/service/repository.go](https://github.com/nthnluu/aether/blob/main/examples/horoscope/service/repository.go).
