package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	a, err := app.NewApp(ctx)

	if err != nil {
		log.Fatal("failed to imput app %s", err)
	}
	err = a.Run()
}
