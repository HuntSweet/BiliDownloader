package main

import (
	"github.com/vbauerster/mpb/v5"
	"github.com/vbauerster/mpb/v5/decor"
)



func newBar(size int,title string) (*mpb.Bar) {

	bar := p.AddBar(int64(size),
		mpb.PrependDecorators(
			decor.Name(title + "  "),
			decor.Counters(decor.UnitKiB, "% .1f / % .1f"),

		),
		mpb.AppendDecorators(
			decor.Percentage(),

		),
		)

	return bar
}
