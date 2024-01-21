package main

import (
	"fmt"
)

func main() {
	cipher := "VJGTWF SMX VWE SHXWDTSMEW,\nVWJ KWZJ KMWKKW TAJFWF LJMY,\nZAFY VWK XJMWZDAFYK DWLRLW HXDSMEW\nMFV SF FMWKKWF FGUZ YWFMY."
	charCount := make(map[rune]int)
	for _, symbol := range cipher {
		charCount[symbol]++
	}
	for symbol, count := range charCount {
		fmt.Println(string(symbol), count)
	}
}
