package main
import ("fmt")

type Animal interface{
	HappySound()
	AngrySound()
}

type Cat string

func (c Cat) HappySound(){
	fmt.Println("I m happy")
}

func (c Cat) AngrySound(){
	fmt.Println("I m so sad")
}
func (c Cat) Name()string{
	return string(c)
}

func main(){
	var kitty Animal
	kitty = Cat("Cat")
	kitty.AngrySound()
	kitty.HappySound()
	var kitty2 Cat = kitty.(Cat)
	fmt.Println(kitty2.Name())
}