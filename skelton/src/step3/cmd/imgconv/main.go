package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func readFile(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("書き込み対象ファイルが見つかりません")
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	return string(b), nil
}

func writeText(text string) error {
	fp, err := os.Create("output.txt")
	if err != nil {
		return fmt.Errorf("ファイルの作成に失敗しました")
	}
	defer fp.Close()
	fp.WriteString(text)
	return nil
}

func main() {
	var fileText []string
	flag.Parse()
	for _, v := range flag.Args() {
		text, err := readFile(v)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}
		fileText = append(fileText, text)
	}
	writeText(strings.Join(fileText, "\n"))
}
