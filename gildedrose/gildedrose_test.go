package gildedrose_test

import (
	"github.com/pascaldekloe/goe/verify"
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func TestUpdateQuality(t *testing.T) {
	tests := map[string]struct {
		given []*gildedrose.Item
		want  []*gildedrose.Item
	}{
		"-1 days": {
			given: []*gildedrose.Item{
				{Name: "Aged Brie", SellIn: -1, Quality: 10},
				{Name: "Aged Brie", SellIn: -1, Quality: 50},
				{Name: "Sulfuras, Hand of Ragnaros", SellIn: -1, Quality: 10},
				{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: -1, Quality: 10},
				{Name: "Conjured mana", SellIn: -1, Quality: 10},
				{Name: "default", SellIn: -1, Quality: 10},
			},
			want: []*gildedrose.Item{
				{Name: "Aged Brie", SellIn: -2, Quality: 12},
				{Name: "Aged Brie", SellIn: -2, Quality: 50},
				{Name: "Sulfuras, Hand of Ragnaros", SellIn: -1, Quality: 10},
				{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: -2, Quality: 0},
				{Name: "Conjured mana", SellIn: -2, Quality: 6},
				{Name: "default", SellIn: -2, Quality: 8},
			},
		},
		"0 days": {
			given: []*gildedrose.Item{
				{Name: "Aged Brie", SellIn: 0, Quality: 10},
				{Name: "Sulfuras, Hand of Ragnaros", SellIn: 0, Quality: 10},
				{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 0, Quality: 10},
				{Name: "Conjured mana", SellIn: 0, Quality: 10},
				{Name: "default", SellIn: 0, Quality: 10},
			},
			want: []*gildedrose.Item{
				{Name: "Aged Brie", SellIn: -1, Quality: 12},
				{Name: "Sulfuras, Hand of Ragnaros", SellIn: 0, Quality: 10},
				{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: -1, Quality: 0},
				{Name: "Conjured mana", SellIn: -1, Quality: 6},
				{Name: "default", SellIn: -1, Quality: 8},
			},
		},
		"1 day": {
			given: []*gildedrose.Item{
				{Name: "Aged Brie", SellIn: 1, Quality: 10},
				{Name: "Sulfuras, Hand of Ragnaros", SellIn: 1, Quality: 10},
				{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 1, Quality: 10},
				{Name: "Conjured mana", SellIn: 1, Quality: 10},
				{Name: "default", SellIn: 1, Quality: 10},
			},
			want: []*gildedrose.Item{
				{Name: "Aged Brie", SellIn: 0, Quality: 11},
				{Name: "Sulfuras, Hand of Ragnaros", SellIn: 1, Quality: 10},
				{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 0, Quality: 13},
				{Name: "Conjured mana", SellIn: 0, Quality: 8},
				{Name: "default", SellIn: 0, Quality: 9},
			},
		},
		"2 days": {
			given: []*gildedrose.Item{
				{Name: "Aged Brie", SellIn: 2, Quality: 10},
				{Name: "Sulfuras, Hand of Ragnaros", SellIn: 2, Quality: 10},
				{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 2, Quality: 10},
				{Name: "Conjured mana", SellIn: 2, Quality: 10},
				{Name: "default", SellIn: 2, Quality: 10},
			},
			want: []*gildedrose.Item{
				{Name: "Aged Brie", SellIn: 1, Quality: 11},
				{Name: "Sulfuras, Hand of Ragnaros", SellIn: 2, Quality: 10},
				{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 1, Quality: 13},
				{Name: "Conjured mana", SellIn: 1, Quality: 8},
				{Name: "default", SellIn: 1, Quality: 9},
			},
		},
		"6 days": {
			given: []*gildedrose.Item{
				{Name: "Aged Brie", SellIn: 6, Quality: 10},
				{Name: "Sulfuras, Hand of Ragnaros", SellIn: 6, Quality: 10},
				{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 6, Quality: 10},
				{Name: "Conjured mana", SellIn: 6, Quality: 10},
				{Name: "default", SellIn: 6, Quality: 10},
			},
			want: []*gildedrose.Item{
				{Name: "Aged Brie", SellIn: 5, Quality: 11},
				{Name: "Sulfuras, Hand of Ragnaros", SellIn: 6, Quality: 10},
				{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 5, Quality: 12},
				{Name: "Conjured mana", SellIn: 5, Quality: 8},
				{Name: "default", SellIn: 5, Quality: 9},
			},
		},
		"11 days": {
			given: []*gildedrose.Item{
				{Name: "Aged Brie", SellIn: 11, Quality: 10},
				{Name: "Sulfuras, Hand of Ragnaros", SellIn: 11, Quality: 10},
				{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 11, Quality: 10},
				{Name: "Conjured mana", SellIn: 11, Quality: 10},
				{Name: "default", SellIn: 11, Quality: 10},
			},
			want: []*gildedrose.Item{
				{Name: "Aged Brie", SellIn: 10, Quality: 11},
				{Name: "Sulfuras, Hand of Ragnaros", SellIn: 11, Quality: 10},
				{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 11},
				{Name: "Conjured mana", SellIn: 10, Quality: 8},
				{Name: "default", SellIn: 10, Quality: 9},
			},
		},
	}
	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			gildedrose.UpdateQuality(testCase.given)
			verify.Values(t, name, testCase.given, testCase.want)
		})
	}
}
