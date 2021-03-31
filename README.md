### gson

gson is a tool to format and compress json from a json file or string

### Installation

```shell
go get github.com/clearcodecn/gson
```

### Usage

```shell
$ gson -h 
gson is a tool to format and compress json from a json file or string
Usage:
        -b              beautiful output the json file
        -w              overwrite the file
        -f              filename, default is stdin
        -h              print usage
```

* read json from stdin

```shell
cat test.json | gson -b
```

* read json from file and overwrite it, format it.

```shell
gson -f test.json -w -b
```