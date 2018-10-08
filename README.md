# dive-into-pact-go

Dive into [pact-go](https://github.com/pact-foundation/pact-go)

# Useful Documents and Information

* [docs.pact.io](https://docs.pact.io/)
* [What about Swagger Specification aka OpenAPI?](https://github.com/pact-foundation/pact-specification/issues/28)

# Build and Run

## Consumer

```sh
$ cd consumer
$ go build -o consumer cmd/recipes/main.go
$ ./consumer
```

Options

* `-h` ... Host where provider is running on (default: `localhost`)
* `-p` ... Port that provider is listening (default: `10000`)

## Provider

```sh
$ cd provider
$ go build -o provider cmd/recipes/main.go
$ ./provider
```

Options

* `-p` ... Port (default: `10000`)

## Provider

# Verification

## Consumer

```sh
$ cd consumer/client
$ go test
```

## Provider

```sh
$ cd provider/handler
$ go test
```
