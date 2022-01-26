package main

import "github.com/gocolly/colly"

type coupangItemInfo struct {
	name          string
	price         string
	originalPrice string
	cardDiscount  string
	reward        string
	isOutOfStock  bool
}

type application struct {
	c *colly.Collector
}
