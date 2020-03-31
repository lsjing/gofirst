package utils

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

var Config *ini.File
var RootPath string

func init()  {
	RootPath="/Users/danny0331/Documents/github/gofirst"
	var err error
	Config, err = ini.Load(RootPath + "/conf/config.ini")

	if err != nil {
		fmt.Printf("Failed to read file %v", err)
		os.Exit(1)
	}

}