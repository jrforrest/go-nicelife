package gui

import "math/rand"

var cellColors = []uint32{
	0xff666699,
	0xff806699,
	0xff996699,
	0xff996680,
	0xff668099,
	0xff8B8BB1,
	0xffAFAFCA,
	0xff996666,
	0xff669999,
	0xffCACAAF,
	0xffB1B18B,
	0xff998066,
	0xff669980,
	0xff669966,
	0xff809966,
	0xff999966,
}

func randomCellColor() uint32 {
	i := rand.Int() % len(cellColors)
	return cellColors[i]
}
