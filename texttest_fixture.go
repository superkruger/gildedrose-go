package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func main() {
	fmt.Println("OMGHAI!")

	days := 2
	var err error
	if len(os.Args) > 1 {
		days, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		days++
	}

	originalRes := original(days)
	refactoredRes := refactored(days)
	fmt.Printf("ORIGINAL:\n\n%s\n\n", originalRes)
	fmt.Printf("REFACTORED:\n\n%s\n\n", refactoredRes)

	fmt.Printf("Success: %v\n", originalRes == refactoredRes)
}

func original(days int) string {
	var items = []*gildedrose.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		{"Conjured Mana Cake", 3, 6}, // <-- :O
	}
	res := strings.Builder{}
	for day := 0; day < days; day++ {
		res.WriteString(fmt.Sprintf("-------- day %d --------\n", day))
		res.WriteString(fmt.Sprintln("Name, SellIn, Quality"))
		for i := 0; i < len(items); i++ {
			res.WriteString(fmt.Sprintln(items[i]))
		}
		res.WriteString(fmt.Sprintln(""))
		gildedrose.UpdateQualityOriginal(items)
	}
	return res.String()
}

func refactored(days int) string {
	var items = []*gildedrose.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		{"Conjured Mana Cake", 3, 6}, // <-- :O
	}
	res := strings.Builder{}
	for day := 0; day < days; day++ {
		res.WriteString(fmt.Sprintf("-------- day %d --------\n", day))
		res.WriteString(fmt.Sprintln("Name, SellIn, Quality"))
		for i := 0; i < len(items); i++ {
			res.WriteString(fmt.Sprintln(items[i]))
		}
		res.WriteString(fmt.Sprintln(""))
		gildedrose.UpdateQuality(items)
	}
	return res.String()
}
