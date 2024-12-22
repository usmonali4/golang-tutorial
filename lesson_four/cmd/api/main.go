package main

import (
	env "lesson4/internal"
	"log"
)

func main() {

	cfg := config{
		addr: env.GetString("ADDR"),
	}

	app := &application{
		config : cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}  