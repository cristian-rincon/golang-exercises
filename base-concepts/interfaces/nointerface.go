package main

import "fmt"

type dog struct {}

type fish struct {}

type bird struct {}

func (dog) walk() string {
	return "i am a dog and i walk"
}

func (fish) swim() string {
	return "i am a fish and i swim"
}

func (bird) fly() string {
	return "i am a bird and i fly"
}




func main()  {
	p := dog{}
	f := fish{}
	b := bird{}

	fmt.Println(p.walk())
	fmt.Println(f.swim())
	fmt.Println(b.fly())

}