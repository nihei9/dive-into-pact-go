# dive-into-pact-go

Dive into pact-go

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
