package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"net"
	"os"
	"os/exec"
	"time"
)

func SearchIMEI(number string, wd selenium.WebDriver) TableItem {
	wd.Refresh()

	time.Sleep(4 * time.Second)
	elem, _ := wd.FindElement(selenium.ByID, "inputidfrag-inner")

	elem.SendKeys(number)
	elem.SendKeys(selenium.EnterKey)

	elem, _ = wd.FindElement(selenium.ByID, "__text9-__xmlview1--tabid-0")
	elem.Click()

	fmt.Println("IMEI : " + number)

	elem, _ = wd.FindElement(selenium.ByID, "__xmlview2--sitedsc-inner")
	dist, _ := elem.GetAttribute("value")
	fmt.Println("Distributor Name: " + dist)

	elem, _ = wd.FindElement(selenium.ByID, "__xmlview2--custmr-inner")
	retail, _ := elem.GetAttribute("value")
	fmt.Println("Retailer Name: " + retail)

	elem, _ = wd.FindElement(selenium.ByID, "__xmlview2--dte-inner")
	date, _ := elem.GetAttribute("value")
	fmt.Println("Date: " + date)

	return TableItem{number, dist, retail, "", date}
}

func SetupSelenium() (*selenium.Service, selenium.WebDriver) {

	jio_url := "https://partnercentral.jioconnect.com/group/guest/home"
	browserPath := GetBrowserPath("chromium")
	port, err := pickUnusedPort()

	var opts []selenium.ServiceOption
	service, err := selenium.NewChromeDriverService(GetChromeDriverPath(),
		port, opts...)

	if err != nil {
		fmt.Printf("Error starting the ChromeDriver server: %v", err)
	}

	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	chrCaps := chrome.Capabilities{
		Path: browserPath,
		Args: []string{
			"--no-sandbox",
		},
	}
	caps.AddChrome(chrCaps)

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://127.0.0.1:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}

	err = wd.Get(jio_url)

	if err != nil {
		panic(err)
	}
	wd.SetAsyncScriptTimeout(10 * time.Second)
	wd.SetImplicitWaitTimeout(10 * time.Second)
	wd.SetPageLoadTimeout(10 * time.Second)

	elem, err := (wd.FindElement(selenium.ByName, "username"))
	if err != nil {
		panic(err)
	}

	elem.SendKeys(username)
	elem, err = (wd.FindElement(selenium.ByName, "password"))
	if err != nil {
		panic(err)
	}
	elem.SendKeys(password)

	elem, _ = wd.FindElement(selenium.ByXPATH, "//input[@type='submit']")

	elem.Click()

	elem, err = wd.FindElement(selenium.ByXPATH, "//a[@title='"+username+"']")

	if err != nil {
		fmt.Println("Error loggin in")
		panic(err)
	}

	fmt.Println("Successfully Logged in")

	wd.Get("https://fiori.jioconnect.com/sap/bc/ui5_ui5/sap/zehys_dashboard/index.html")

	time.Sleep(3 * time.Second)

	elem, _ = wd.FindElement(selenium.ByID, "__shell1-header-hdr-begin")

	elem.Click()

	time.Sleep(2 * time.Second)
	elem, _ = wd.FindElement(selenium.ByID, "__panel0-__shell1-5-CollapsedImg")

	elem.Click()

	time.Sleep(2 * time.Second)
	elem, err = wd.FindElement(selenium.ByID, "__item1-__list0-__shell1-5-6")

	elem.Click()
	time.Sleep(3 * time.Second)

	elem, err = wd.FindElement(selenium.ByID, "inputidfrag-inner")

	if err != nil {
		panic(err)
	}
	fmt.Println("imei search box found")

	return service, wd

}

func pickUnusedPort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	port := l.Addr().(*net.TCPAddr).Port
	if err := l.Close(); err != nil {
		return 0, err
	}
	return port, nil
}

func GetBrowserPath(browser string) string {
	if _, err := os.Stat(browser); err != nil {
		path, err := exec.LookPath(browser)
		if err != nil {
			panic("Browser binary path not found")
		}
		return path
	}
	return browser
}
