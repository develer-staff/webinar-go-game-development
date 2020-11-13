package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/develersrl/webinar-go-game-development/shooter-game"
)

func main() {
	rand.Seed(time.Now().Unix())
	game := shooter.NewGame()
	if err := game.Run(); err != nil {
		log.Fatalf("Game error: %v", err)
	}
}
