package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"net"
	"time"
)

func SearchIMEI(number string, wd selenium.WebDriver) TableItem {
	wd.Refresh()

	time.Sleep(6 * time.Second)
	elem, _ := wd.FindElement(selenium.ByID, "inputidfrag-inner")

	elem.SendKeys(number)
	//time.Sleep(2 * time.Second)
	elem.SendKeys(selenium.EnterKey)

	jcnum := ""

	elem, err := wd.FindElement(selenium.ByXPATH, "//span[text()='RRL Dispatch']")
	if err == nil {
		id, _ := elem.GetAttribute("id")
		id = "__text9-__xmlview1--tabid-" + id[len(id)-1:]
		elem, _ = wd.FindElement(selenium.ByID, id)
		jcnum, _ = elem.GetAttribute("innerHTML")
	}

	elem, err = wd.FindElement(selenium.ByXPATH, "//span[text()='Dispatch']")
	if err != nil {
		return TableItem{number, "", "", jcnum}
	}
	elem.Click()

	elem, _ = wd.FindElement(selenium.ByID, "__xmlview2--sitedsc-inner")
	dist, _ := elem.GetAttribute("value")

	elem, _ = wd.FindElement(selenium.ByID, "__xmlview2--custmr-inner")
	retail, _ := elem.GetAttribute("value")

	return TableItem{number, dist, retail, jcnum}
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
	wd.SetImplicitWaitTimeout(4 * time.Second)
	wd.SetPageLoadTimeout(15 * time.Second)

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
