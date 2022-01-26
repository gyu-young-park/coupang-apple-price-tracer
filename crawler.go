package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func getExecCrawler(specifier string, e *colly.HTMLElement) (coupangItemInfo, error) {
	cardDiscount := e.ChildText(".ccid-txt")
	if cardDiscount == "" {
		cardDiscount = "none"
	}
	isOutOfStock := false
	if e.ChildText(".out-of-stock") != "" {
		isOutOfStock = true
	}
	data := coupangItemInfo{
		name:          e.ChildText(".name"),
		price:         e.ChildText(".price-value"),
		originalPrice: e.ChildText(".base-price"),
		cardDiscount:  cardDiscount,
		reward:        e.ChildText(".reward-cash-txt"),
		isOutOfStock:  isOutOfStock,
	}
	return data, nil
}

func (app *application) getIphone13info(specifier, url string) ([]coupangItemInfo, error) {
	dataList := []coupangItemInfo{}
	app.c.OnHTML(specifier, func(e *colly.HTMLElement) {
		data, err := getExecCrawler(specifier, e)
		if err != nil {
			fmt.Println("cannot get data")
		}
		dataList = append(dataList, data)
	})

	app.c.OnError(func(r *colly.Response, err error) {
		fmt.Println("getIphone13info: Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	app.c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	app.c.Visit(url)

	return dataList, nil
}

func (app *application) getMacBookProM1(specifier, url string) ([]coupangItemInfo, error) {
	dataList := []coupangItemInfo{}
	app.c.OnHTML(specifier, func(e *colly.HTMLElement) {
		data, err := getExecCrawler(specifier, e)
		if err != nil {
			fmt.Println("cannot get data")
		}
		dataList = append(dataList, data)
	})

	app.c.OnError(func(r *colly.Response, err error) {
		fmt.Println("getMacBookProM1: Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	app.c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	app.c.Visit(url)

	return dataList, nil
}
