package main


import "fmt"

func main() {
	fmt.Println("hello word")

	var mobil string = "BMW 2020"
	car,motor := "Honda" , "Beat"

	fmt.Println(car)
	fmt.Println(motor)

	fmt.Println(mobil)

	var newCars [4] string

	newCars[0] = "BMW"
	newCars[1] = "Toyota"
	newCars[2] = "Mitshubisi"
	newCars[3] = "Mercy"

	
	for i:=0; i < len(newCars); i++ {
		fmt.Println(newCars[i])
		i++		
	}

}