package trash

type Car struct {
	Model string
	Color string
}

// metodo
func (c Car) Start() {
	println(c.Model + " has been started!")
}

func (c Car) ChangeColorWithoutPointer(color string) {
	c.Color = color // duplicando o valor de c.Color na memoria - copia do color original
	println("Car color is " + c.Color + " now")
}

func (c *Car) ChangeColorWithPointer(color string) {
	c.Color = color // alterando o valor na memoria
	println("Car color is " + c.Color + " now")
}

// funcao
// func soma(x int, y int) int {
// 	return x + y
// }

func test() {
	a := 10
	// b := a // copiou o valor de a mas criou o seu proprio espaco na memoria
	// b = 20
	b := &a
	*b = 20 // altera o valor na memoria para 20, como o endereço de b na memoria é igual o endereço de a, acaba alterando o valor de a e de b

	println(&a) // endereço de a na memoria
	println(a)
	println(b)
	// println(a) // 10
	// println(b) // 10

	// car := Car{ // declarando e atribuindo a variável car
	// 	Model: "Ferrari",
	// 	Color: "Red",
	// }
	// car.Model = "Fiat" //alterando o valor do atributo Model
	// car.Start()
	// car.ChangeColorWithoutPointer("Blue")
	// car.ChangeColorWithPointer("Orange")
	// println(car.Color)
}
