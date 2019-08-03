package main

import (
	"context"
	"fmt"

	"github.com/coltiebaby/g-law/riot/v4/matches"
)

func main() {
	ctx := context.Background()
	m, _ := matches.GetTimeline(ctx, "3035593313")
	fmt.Println(m)
}
