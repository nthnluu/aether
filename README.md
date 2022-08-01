# Aether
Aether is a microservice framework that lets you build robust, production-quality services with fast build times and minimal-configuration production infrastructure.

Aether provides libraries, automation, and tooling for production details/

## Concepts
### Services
A service is a collection of one or more callable endpoints. A service describes an interface (in terms of RPCs) that a server provides.

### Aether nodes
Aether nodes are the fundamental building block of an Aether server. An Aether node is a package of code and configuration that can be combined with other nodes to form an Aether server.

An Aether server is a binary that runs on Kubernetes and includes several types of nodes.

### Node types
The node type depends on the purpose of the node:
- **Executable nodes** represent callable services, and often contain most of the code and business logic for an application. They can listen for incoming requests.
- **Shared nodes** represent libraries or backends. They are a mechanism for centralizing reuseable bits of logic or functionality. Shared nodes can depend on other shared nodes if needed.
- **Server nodes** represent Aether servers. A server is an executeable that bundles one or more executable nodes along with the union of their dependencies into a single deployable entity.
  
## Protobufs
Protobuf is an interface definition language (IDL) that allows us to define data types and services in a language neutral format.

To add a new `.proto` file, add it to the `proto/` directory. Then, compile the Go proto library:
```shell
cd pb
./compile.sh
```
If `./compile.sh` returns a permission error, you can grant execute permissions by running:
```shell
chmod +x compile.sh # then rerun `./compile.sh`
```

## Creating a service

TLDR: Create a service by writing a `<your service name>_service.proto` file and calling `Aether create service pb/<your service name>_service.proto

A service is simply a group of endpoints.

At a high level, to create a new service, you have to:
1. Create `<your service name>_models.proto` and `<your service name>_service.proto` to define your service's data types and RPC interface (including request/response types for each endpoint).
2. Create a folder in the `cmd` directory named after your service (all lowercase, one word, and singular).
3. 