// +build windows

package main

import (
	"github.com/gobuffalo/packr"
	"io/ioutil"
	"os"
)

func GetChromeDriverPath() string {
	box := packr.NewBox("./driver")

	data, _ := box.MustBytes("chromedriver_windows")

	ioutil.WriteFile(os.TempDir()+"/chromedriver", data, 0744)

	return os.TempDir() + "/chromedriver"
}

func GetChromeBrowserPath() string {
	if _, err := os.Stat("chrome"); err != nil {
		path, err := exec.LookPath(browser)
		if err != nil {
			panic("Browser binary path not found")
		}
		return path
	}
	return browser
}
