package main

type engine interface {
	start()
	stop()
}

type vehicle struct {
	engine
	speed float32
}

func main() {
	myCar := vehicle{speed: 0.0}

}
