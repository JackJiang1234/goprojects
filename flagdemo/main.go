package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

func main() {
	//testFlagParse()
	//subCommandParse()
	testCustomValueParse()
}

func testFlagParse() {
	var name string
	flag.StringVar(&name, "name", "go tour", "help")
	flag.StringVar(&name, "n", "go tour", "help")
	flag.Parse()

	log.Printf("name: %s", name)
}

var name string

func subCommandParse() {
	flag.Parse()
	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goCmd.StringVar(&name, "name", "go lanuage", "help")
	phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
	phpCmd.StringVar(&name, "n", "php language", "help")

	args := flag.Args()
	switch args[0] {
	case "go":
		_ = goCmd.Parse(args[1:])
	case "php":
		_ = phpCmd.Parse(args[1:])
	}
	log.Printf("name: %s", name)
}

func testCustomValueParse() {
	var name Name
	flag.Var(&name, "name", "help")
	flag.Parse()

	log.Printf("name %s", name)
}

type Name string

func (i *Name) String() string {
	return fmt.Sprint(*i)
}

func (i *Name) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("name flag already set")
	}
	*i = Name("jack:" + value)
	return nil
}
