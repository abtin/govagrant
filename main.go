package main

import (
	"fmt"
	"govagrant/info"
	"os"
)

func main() {
	host, err := info.NewOsInfo()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(host)
	fmt.Println(host.Md5sum())
}
