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
A service is simply a group of endpoints.

At a high level, to create a new service, you have to:
1. Create `<your service name>_models.proto` and `<your service name>_service.proto` to define your service's data types and RPC interface (including request/response types for each endpoint).
2. Create a folder in the `cmd` directory named after your service (all lowercase, one word, and singular).
3. 