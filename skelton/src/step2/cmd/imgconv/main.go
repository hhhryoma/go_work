package main

import (
	"fmt"
	"os"
)

func run() error {
	if len(os.Args) < 3 {
		return fmt.Errorf("引数が足りません")
	}
	for _, v := range os.Args {
		fmt.Println(v)
	}
	return nil
}

func main() {
	// flag.Parse()
	// for _, v := range flag.Args() {
	// 	fmt.Println(v)
	// }
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
