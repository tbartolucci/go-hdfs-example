package main

import (
	"fmt"
	"github.com/colinmarc/hdfs"
)

func main() {
	client, err := hdfs.New("localhost:8020")
	if err != nil {
		panic(err)
	}

	fsInfo, err := client.Stat(".")
	if err != nil {
		fmt.Errorf("%s", err)
	}

	fmt.Print(fsInfo)

	file, err := client.Open("/input/f1.txt")
	if err != nil {
		fmt.Errorf("%s", err)
	}


	buf := make([]byte, 59)
	file.ReadAt(buf, 48847)

	fmt.Println(string(buf))
}
