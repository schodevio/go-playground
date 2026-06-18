package main

import (
	"log"

	"github.com/Shopify/go-lua"
)

var script = `
bestBid = bestBid();
bestAsk = bestAsk();

spread = math.abs(bestBid - bestAsk);
print("Spread: " .. spread);
`

func main() {
	l := lua.NewState()
	lua.OpenLibraries(l)

	registerBestBid(l)
	registerBestAsk(l)

	if err := lua.DoString(l, script); err != nil {
		log.Fatal(err)
	}
}

func registerBestBid(l *lua.State) {
	l.Register("bestBid", func(l *lua.State) int {
		l.PushInteger(100) // Example best bid value
		return 1
	})
}

func registerBestAsk(l *lua.State) {
	l.Register("bestAsk", func(l *lua.State) int {
		l.PushInteger(105) // Example best ask value
		return 1
	})
}
