package main

import "fmt"

type animal interface {
	move() string
}


type dog struct {}

type fish struct {}

type bird struct {}

func (dog) move() string {
	return "i am a dog and i walk"
}

func (fish) move() string {
	return "i am a fish and i swim"
}

func (bird) move() string {
	return "i am a bird and i fly"
}

func moveAnimal(a animal) {
	fmt.Println(a.move())
}


func main()  {
	p := dog{}
	f := fish{}
	b := bird{}

	moveAnimal(p)
	moveAnimal(f)
	moveAnimal(b)

}