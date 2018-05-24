// +build windows

package main

import (
	"github.com/gobuffalo/packr"
	"io/ioutil"
	"os"
	"os/exec"
)

func GetChromeDriverPath() string {
	box := packr.NewBox("./driver")

	data, _ := box.MustBytes("chromedriver_windows")

	ioutil.WriteFile(os.TempDir()+"/chromedriver", data, 0744)

	return os.TempDir() + "/chromedriver"
}

func GetChromeBrowserPath() string {
	browser := "chrome"
	if _, err := os.Stat(browser); err != nil {
		path, err := exec.LookPath(browser)
		if err != nil {
			panic("Browser binary path not found")
		}
		return path
	}
	return browser
}
