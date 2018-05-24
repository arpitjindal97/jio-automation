// +build linux

package main

import (
	"github.com/gobuffalo/packr"
	"io/ioutil"
	"os"
)

func GetChromeDriverPath() string {
	box := packr.NewBox("./driver")

	data, _ := box.MustBytes("chromedriver_linux")

	ioutil.WriteFile(os.TempDir()+"/chromedriver", data, 0744)

	return os.TempDir() + "/chromedriver"
}
