# Introducing OpenAPI Specification

Introducing [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) for *Architecting the API first* approach. Gather requirements, design the API, define the schema and generate code.

In this project, using Go 1.24+. To manage the dependency, recommended tool is built in `go tool`:

```sh
go get -tool github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
```

Please visit the `oapi-codegen` documentation for more.

Server being used: `gin`

Example command:


```sh
# for medicines
oapi-codegen -config openapi/cfg-medicines.yaml medicines/medicines.yaml

# for users
oapi-codegen -config openapi/cfg-users.yaml users/users.yaml
```