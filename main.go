package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36"
	c := colly.NewCollector(
		colly.UserAgent(userAgent),
	)

	app := application{
		c: c,
	}

	iphoneInfoList, err := app.getIphone13info(".baby-product-wrap > .descriptions", "https://www.coupang.com/np/promotion/85026")
	if err != nil {
		fmt.Println("failed get iphone info list")
	}
	fmt.Println(iphoneInfoList)

	macBookProM1List, err := app.getMacBookProM1(".baby-product-wrap > .descriptions", "https://www.coupang.com/np/promotion/51797")
	if err != nil {
		fmt.Println("failed get iphone info list")
	}
	fmt.Println(macBookProM1List)
}
