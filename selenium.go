package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"net"
	"strconv"
	"strings"
	"time"
)

func SearchIMEI(number string, wd selenium.WebDriver) TableItem {

	elem, _ := GetElement(wd, selenium.ByID, "idMenu")

	elem.Click()

	time.Sleep(2 * time.Second)
	elem, _ = GetElement(wd, selenium.ByID, "__panel0-__shell1-5-CollapsedImg")

	elem.Click()

	time.Sleep(2 * time.Second)
	elem, _ = GetElement(wd, selenium.ByID, "__item1-__list0-__shell1-5-6")

	elem.Click()

	time.Sleep(2 * time.Second)
	elem, _ = GetElement(wd, selenium.ByID, "inputidfrag-inner")

	time.Sleep(1 * time.Second)
	elem.SendKeys(number)
	time.Sleep(1 * time.Second)
	elem.SendKeys(selenium.EnterKey)

	jcnum := ""
	elem, err := GetElement(wd, selenium.ByXPATH, "//span[text()='RRL Dispatch']")
	if err == nil {
		id, _ := elem.GetAttribute("id")

		pos := strings.Index(id, "-")
		num, _ := strconv.Atoi(id[6:pos])
		num = num + 1
		id = id[:6] + strconv.Itoa(num) + id[pos:]

		elem, _ = GetElement(wd, selenium.ByID, id)
		jcnum, _ = elem.GetAttribute("innerHTML")
	}

	var dist string
	elem, err = GetElement(wd, selenium.ByXPATH, "//span[text()='Receipt']")
	if err == nil {
		id, _ := elem.GetAttribute("id")

		pos := strings.Index(id, "-")
		num, _ := strconv.Atoi(id[6:pos])
		num = num + 1
		id = id[:6] + strconv.Itoa(num) + id[pos:]

		elem, _ = wd.FindElement(selenium.ByID, id)
		fmt.Println("dist : " + dist)
	}

	elem, err = GetElement(wd, selenium.ByXPATH, "//span[text()='Dispatch']")
	if err != nil {
		return TableItem{number, "", "", jcnum}
	}
	elem.Click()

	elem, _ = GetElement(wd, selenium.ByXPATH, "//label[text()='Receiver']")
	temp, _ := elem.GetAttribute("for")
	elem, _ = GetElement(wd, selenium.ByID, temp)
	retail, _ := elem.GetAttribute("value")

	elem, _ = GetElement(wd, selenium.ByID, "idHome")
	elem.Click()

	time.Sleep(1 * time.Second)
	return TableItem{number, dist, retail, jcnum}
}

func GetElement(wd selenium.WebDriver, findBy, pattern string) (selenium.WebElement, error) {

	elementExists := func(wd selenium.WebDriver) (bool, error) {
		_, err := wd.FindElement(findBy, pattern)
		if err != nil {
			return false, nil
		}

		return true, nil
	}
	timeout, _ := time.ParseDuration("6s")
	interval, _ := time.ParseDuration("1s")
	wd.WaitWithTimeoutAndInterval(elementExists, timeout, interval)

	return wd.FindElement(findBy, pattern)

}

func SetupSelenium() (*selenium.Service, selenium.WebDriver) {

	jio_url := "https://partnercentral.jioconnect.com/group/guest/home"
	//browserPath := GetChromeBrowserPath()
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
	//chrCaps := chrome.Capabilities{
	//	Path: browserPath,
	//	Args: []string{
	//		"--no-sandbox",
	//	},
	//}
	//caps.AddChrome(chrCaps)

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://127.0.0.1:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}

	err = wd.Get(jio_url)

	if err != nil {
		panic(err)
	}
	wd.SetAsyncScriptTimeout(15 * time.Second)
	wd.SetPageLoadTimeout(15 * time.Second)

	elem, err := GetElement(wd, selenium.ByName, "username")
	if err != nil {
		panic(err)
	}

	elem.SendKeys(username)
	elem, err = GetElement(wd, selenium.ByName, "password")
	elem.SendKeys(password)

	elem, _ = GetElement(wd, selenium.ByXPATH, "//input[@type='submit']")

	elem.Click()

	elem, err = GetElement(wd, selenium.ByXPATH, "//a[@title='"+username+"']")

	if err != nil {
		fmt.Println("Error logging in")
		panic(err)
	}

	fmt.Println("Successfully Logged in")

	wd.Get("https://fiori.jioconnect.com/sap/bc/ui5_ui5/sap/zehys_dashboard/index.html")

	elem, _ = GetElement(wd, selenium.ByID, "idMenu")

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
