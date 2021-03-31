package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	overWrite bool
	filename  string
	beauti    bool
	help      bool
)

func init() {
	flag.BoolVar(&overWrite, "w", false, "overwrite the current file")
	flag.StringVar(&filename, "f", "", "filename , default read from os.Stdin")
	flag.BoolVar(&beauti, "b", false, "output beautiful json")
	flag.BoolVar(&help, "h", false, "print usage")
}

func main() {
	flag.Parse()

	if help {
		printUsage()
		return
	}

	var reader io.Reader = os.Stdin
	if filename != "" {
		fi, err := os.OpenFile(filename, os.O_RDONLY, 0755)
		if err != nil {
			log.Fatal(err)
		}
		defer fi.Close()
		reader = fi
	}

	var object = make(map[string]interface{})
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = json.Unmarshal(data, &object)
	if err != nil {
		log.Fatal(err)
		return
	}

	var writer io.Writer = os.Stdout
	if overWrite && filename != "" {
		fi, err := os.OpenFile(filename, os.O_WRONLY, 0755)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer fi.Close()

		writer = fi
	}

	var output = bufio.NewWriter(writer)
	if beauti {
		data, err = json.MarshalIndent(object, "", "  ")
		if err != nil {
			log.Fatal(err)
			return
		}
	} else {
		data, err = json.Marshal(object)
		if err != nil {
			log.Fatal(err)
		}
	}
	_, err = output.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

var (
	usage = `gson: 
gson is a tool to format and compress json from a json file or string
Usage: 
	-b 		beautiful output the json file
	-w 		overwrite the file
	-f		filename, default is stdin
	-h		print usage
`
)

func printUsage() {
	print(usage)
}
