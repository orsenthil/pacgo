package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// This only works with *Person, does not work with Person
// Only works with Test 2 and Test 3

func (p *Person) String() string {
	return fmt.Sprintf("%s is %d", p.Name, p.Age)
}

// This works for both *Person and Person, BUT you can't modify the value and it takes up more space.
// Works with Test 1, Test 2, Test 3 and Test 4

/* func (p Person) String() string {
	return fmt.Sprintf("%s is %d", p.Name, p.Age)
        } */

// Pass by Value

func test1() {
	p := Person{"Steve", 28}
	printPerson1(p)
	updatePerson1(p)
	printPerson1(p)
}

func updatePerson1(p Person) {
	p.Age = 32
	printPerson1(p)
}

func printPerson1(p Person) {
	fmt.Printf("String: %v | Name: %v | Age: %d\n",
		p,
		p.Name,
		p.Age)
}

// Pass by Refrence

func test2() {
	p := &Person{"Steve", 28}
	printPerson2(p)
	updatePerson2(p)
	printPerson2(p)
}

func updatePerson2(p *Person) {
	p.Age = 32
	printPerson2(p)
}

func printPerson2(p *Person) {
	fmt.Printf("String: %v | Name: %v | Age: %d\n",
		p,
		p.Name,
		p.Age)
}

// Pass by Reference (requires more typing)

func test3() {
	p := Person{"Steve", 28}
	printPerson3(&p)
	updatePerson3(&p)
	printPerson3(&p)
}

func updatePerson3(p *Person) {
	p.Age = 32
	printPerson3(p)
}

func printPerson3(p *Person) {
	fmt.Printf("String: %v | Name: %v | Age: %d\n",
		p,
		p.Name,
		p.Age)
}

func test4() {
	p := &Person{"Steve", 28}
	printPerson4(*p)
	updatePerson4(*p)
	printPerson4(*p)
}

func updatePerson4(p Person) {
	p.Age = 32
	printPerson4(p)
}

func printPerson4(p Person) {
	fmt.Printf("String: %v | Name: %v | Age: %d\n",
		p,
		p.Name,
		p.Age)
}

func main() {
	test1()
	test2()
	test3()
	test4()
}
