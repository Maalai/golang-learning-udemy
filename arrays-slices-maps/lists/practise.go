package lists

import (
	"fmt"

	"example.com/collections/product"
)

func mainPractise() {

	hobbies := [3]string{"spend time with family", "Playing Cricket", "watching youtbe"}

	// Task 1
	fmt.Println(hobbies)

	// Task 2
	fmt.Println(hobbies[0])

	filteredHobbies := hobbies[1:3]
	fmt.Println(filteredHobbies)

	// Task 3
	hobbiesSlice := []string{"spend time with family", "Playing Cricket"}

	hobbiesSlice1 := hobbies[0:2]

	fmt.Println(hobbiesSlice)
	fmt.Println(hobbiesSlice1)

	// Task 4
	reSliced := hobbiesSlice
	reSliced[0] = hobbies[1]
	reSliced[1] = hobbies[2]

	fmt.Println(reSliced)

	// Task 5
	courseGoals := []string{"Master Go", "Build services with Go"}

	courseGoals[1] = "Build Microservices with Go"

	courseGoals = append(courseGoals, "Create Go libraries and Open source it")

	fmt.Println(courseGoals)

	// Task 7
	products := []product.Product{product.New("Pen", 1, 99.99), product.New("Macbook", 2, 250000)}
	products = append(products, product.New("Notebook", 3, 50))
	fmt.Println(products)

}
