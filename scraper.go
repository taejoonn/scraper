package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/gocolly/colly"
)

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	URL := "https://medium.com/the-zap-project"

	collector := colly.NewCollector(
		colly.AllowURLRevisit(),
		colly.IgnoreRobotsTxt(),
	)

	var link string

	collector.OnHTML("div[class='u-lineHeightBase postItem u-marginRight3']", func(e *colly.HTMLElement) {
		link = e.ChildAttr("a[href]", "href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		collector.Visit(e.Request.AbsoluteURL(link))
	})

	collector.Visit(URL)

	openBrowser(link)
}
