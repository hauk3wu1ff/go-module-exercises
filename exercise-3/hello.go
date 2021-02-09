package hello

import (
	quoteV2 "rsc.io/quote/v2"
	"rsc.io/quote/v3"
)

func Hello() string {
	return quote.GlassV3()
}

func Opt() string {
	return quoteV2.OptV2()
}
