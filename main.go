package main

import (
	"fmt"
	"math/rand"
	"time"
)

var hoursInADay = 24
var positions = 25

var sizes = [...]string{"E", "D", "C", "B", "A"}
var contentTypes = [...]string{"hour_image", "blank", "temperature"}

type imgDef struct {
	hour     int    // Hour of the day.
	size     string // One of five sizes.
	content  string // Blank, the img corresponding to the hour, or a temperature.
	position int    // Location on a 5*5 grid. Positions numbered 0-24.
}

func (i imgDef) String() string {
	return fmt.Sprintf("size=%s hour=%02d, grid_position=%02d, image_content=%s\n", i.size, i.hour, i.position, i.content)
}

func (i *imgDef) assign(h int, p *positionTable) {
	i.hour = rand.Intn(hoursInADay)
	i.size = sizes[rand.Intn(len(sizes))]
	i.content = contentTypes[rand.Intn(len(contentTypes))]
	i.position = p.choose(i.size)
}

type positionTable struct {
	mapping map[string][]int
}

func (p *positionTable) choose(size string) int {
	row := p.mapping[size]
	return row[rand.Intn(len(row))]
}

func (p *positionTable) generate(sizes [len(sizes)]string) {
	p.mapping = make(map[string][]int)
	for row, label := range sizes {
		slots := make([]int, 0)
		for cols := 1; cols <= row+1; cols++ {
			for span := 0; span < row+1; span++ {
				slots = append(slots, span+((cols-1)*5))
			}
		}
		p.mapping[label] = slots
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	pt := new(positionTable)
	pt.generate(sizes)
	imgTable := make(map[string][]imgDef)
	for i := 0; i < positions; i++ {
		var img imgDef
		img.assign(i, pt)
		imgTable[img.size] = append(imgTable[img.size], img)
	}
	for _, v := range sizes {
		for _, v2 := range imgTable[v] {
			fmt.Print(v2)
		}
	}
}
