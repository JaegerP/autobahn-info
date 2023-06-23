# Autobahn Info

This project is a simple demonstration on how to access a public API via Go. It utilizes the [resty library](https://github.com/go-resty/resty) for making requests and the [cobra cli library](https://github.com/spf13/cobra) to set up a simple CLI.

## Build

Simply run the go build tool:
```bash
go build
```

## Usage

So far, here are only two commands:

### Example 1: Get a list of all Highways in Germany
```bash
autobahn-info list
```

### Example 2: Get a list warnings for a given highway
```bash
autobahn-info warning --road-id A1
```
where A1 may be replaced by any valid road listed by Example 1

## Development

In order to extend this programm, make sure to have `cobra-cli` installed on your system. If you haven't, run the following command:
```bash
go get -u github.com/spf13/cobra@latest
```

You can add new subcommands by running:
```bash
cobra-cli add [-p parentCmd] <cmdName>
```