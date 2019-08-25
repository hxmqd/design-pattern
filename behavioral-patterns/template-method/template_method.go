package template_methodbn

import "fmt"

type IGame interface {
	initialize()
	startPlay()
	endPlay()
	Play()
}

type Game struct {
	IGame
}

func (g *Game)Play(){
	g.initialize()
	g.startPlay()
	g.endPlay()
}

type Cricket struct {
	Game
}

func (* Cricket) initialize(){
	fmt.Println("Cricket Game Initialized! Start playing")
}

func (* Cricket) endPlay(){
	fmt.Println("Cricket Game Finished!")
}

func (* Cricket) startPlay(){
	fmt.Println("Cricket Game Started. Enjoy the game!")
}

func NewCricket() IGame{
	cricket := &Cricket{}
	cricket.Game = Game{cricket}
	return cricket
}

type Football struct {
	Game
}

func (* Football) initialize(){
	fmt.Println("Football Game Initialized! Start playing")
}

func (* Football) endPlay(){
	fmt.Println("Football Game Finished!")
}

func (* Football) startPlay(){
	fmt.Println("Football Game Started. Enjoy the game!")
}

func NewFootball() IGame{
	football := &Football{}
	football.Game = Game{football}
	return football
}