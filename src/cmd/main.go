package main

import (
	"fmt"
	"github.com/colinmarc/hdfs"
	"os"
)

func main() {
	client, err := hdfs.New("localhost:9000")
	if err != nil {
		panic(err)
	}

	dir := "/example-data/"
	argsWithoutProg := os.Args[1:]
	srcFile := argsWithoutProg[0]

	err = client.Mkdir(dir, os.ModeDir)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}

	info, err := client.Stat(dir)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	fmt.Println(info)

	err = client.CopyToRemote(srcFile, fmt.Sprintf("%s%s", dir, "test.txt"))
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}

	fs, err := client.ReadDir(dir)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}

	for i := 0; i < len(fs); i++ {
		remoteFile := fmt.Sprintf("%s%s", dir, fs[i].Name())
		fmt.Println(remoteFile)
		if !fs[i].IsDir() {
			file, err := client.Open(remoteFile)
			if err != nil {
				fmt.Errorf("Error: %s", err)
			}
			buf := make([]byte, 59)
			file.ReadAt(buf, 48847)

			fmt.Println(string(buf))
		}
	}
}
