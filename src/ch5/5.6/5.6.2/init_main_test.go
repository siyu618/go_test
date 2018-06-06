package main

import (
	"path/filepath"
	"os"
	"fmt"
)

var britishAmerican = "british-american.txt"
//
func init() {
	dir, _ := filepath.Split(os.Args[0])
	britishAmerican = filepath.Join(dir, britishAmerican)
}

func main()  {
	fmt.Println(britishAmerican)
}
