package gildedrose

import "strings"

type Item struct {
	Name            string
	SellIn, Quality int
}

type agedBrieItem struct {
	Name            string
	SellIn, Quality int
}

type backstagePassItem struct {
	Name            string
	SellIn, Quality int
}

type sulfurasItem struct {
	Name            string
	SellIn, Quality int
}

type defaultItem struct {
	Name            string
	SellIn, Quality int
}

const (
	agedBrie      = "Aged Brie"
	sulfuras      = "Sulfuras"
	backstagePass = "Backstage passes"
)

func UpdateQualityOriginal(items []*Item) {
	for i := 0; i < len(items); i++ {

		if items[i].Name != "Aged Brie" && items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
			if items[i].Quality > 0 {
				if items[i].Name != "Sulfuras, Hand of Ragnaros" {
					items[i].Quality = items[i].Quality - 1
				}
			}
		} else {
			if items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 1
				if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].SellIn < 11 {
						if items[i].Quality < 50 {
							items[i].Quality = items[i].Quality + 1
						}
					}
					if items[i].SellIn < 6 {
						if items[i].Quality < 50 {
							items[i].Quality = items[i].Quality + 1
						}
					}
				}
			}
		}

		if items[i].Name != "Sulfuras, Hand of Ragnaros" {
			items[i].SellIn = items[i].SellIn - 1
		}

		if items[i].SellIn < 0 {
			if items[i].Name != "Aged Brie" {
				if items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].Quality > 0 {
						if items[i].Name != "Sulfuras, Hand of Ragnaros" {
							items[i].Quality = items[i].Quality - 1
						}
					}
				} else {
					items[i].Quality = items[i].Quality - items[i].Quality
				}
			} else {
				if items[i].Quality < 50 {
					items[i].Quality = items[i].Quality + 1
				}
			}
		}
	}
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		if strings.Contains(items[i].Name, agedBrie) {
			(*agedBrieItem)(items[i]).update()
			continue
		}
		if strings.Contains(items[i].Name, sulfuras) {
			(*sulfurasItem)(items[i]).update()
			continue
		}
		if strings.Contains(items[i].Name, backstagePass) {
			(*backstagePassItem)(items[i]).update()
			continue
		}
		(*defaultItem)(items[i]).update()
	}
}

func (i *Item) update() {
	if i.Quality < 0 {
		i.Quality = 0
	}
}

func (i *agedBrieItem) update() {
	i.Quality += 1
	if i.SellIn <= 0 {
		i.Quality += 1
	}
	if i.Quality > 50 {
		i.Quality = 50
	}
	i.SellIn -= 1
	(*Item)(i).update()
}

func (i *sulfurasItem) update() {
	(*Item)(i).update()
}

func (i *backstagePassItem) update() {
	i.Quality += 1
	if i.SellIn <= 0 {
		i.Quality = 0
	}
	if i.SellIn < 11 && i.SellIn > 0 {
		i.Quality += 1
	}
	if i.SellIn < 6 && i.SellIn > 0 {
		i.Quality += 1
	}
	if i.Quality > 50 {
		i.Quality = 50
	}
	i.SellIn -= 1
	(*Item)(i).update()
}

func (i *defaultItem) update() {
	i.Quality -= 1
	if i.SellIn <= 0 {
		i.Quality -= 1
	}
	if i.Quality > 50 {
		i.Quality = 50
	}
	i.SellIn -= 1
	(*Item)(i).update()
}
