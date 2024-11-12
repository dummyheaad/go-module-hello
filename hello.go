package hello

import "rsc.io/quote/v3"

// Print Hello world. string
func Hello() string {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}
