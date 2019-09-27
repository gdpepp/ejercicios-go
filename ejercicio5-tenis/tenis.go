package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ball struct {
	accuracy int	//0 to 100%
	velocity int	//0 to 100%
	pass     bool
}

var game = make(chan ball)
var end = make(chan bool)

func main() {
	go djoko()
	go delpo()
	//saque
	game <- ball{accuracy: 100, velocity: 100, pass: true}
	<-end
	return
}

func delpo() {
	//wait for ball
	for true {
		if recv := (<-game).pass; !recv {
			return
		} else {
			fmt.Println("recibe Del Potro")
			bal := getBall()
			time.Sleep(1000 * time.Millisecond)
			game <- bal
		}
	}
}

func djoko() {
	//wait for ball
	for true {
		<-game
		fmt.Println("recibe Djokovich")
		bal := getBall()
		if !bal.pass {
			fmt.Println("La pelota queda en la red. Punto para Del Potro!")
			game <- ball{accuracy: 0, velocity: 0, pass: false}
			end <- true
			return
		}
		time.Sleep(1000 * time.Millisecond)
		game <- bal
	}
}

func getBall() ball {
	rand.Seed(time.Now().UnixNano()) //yields a constantly-changing number.
	vel := rand.Intn(100)
	acc := rand.Intn(100)
	pas := randBool() // function does not exists in rand
	fmt.Println("velocidad", vel, ", presicion", acc)
	return ball{accuracy: acc, velocity: vel, pass: pas}
}

func randBool() bool {
	// https://github.com/golang/go/issues/23804 (Proposal-Hold)
	return rand.Intn(2) == 0
}
