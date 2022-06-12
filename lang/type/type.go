package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
}

func (p *Person) Speak() string {
	return fmt.Sprintf("%s say hi", p.name)
}

// Human 是Person的别名，别名不是另一种类型
type Human = Person

// Dummy 被定义为Person类型，拥有和Person相同的数据结构，但不具备方法
type Dummy Person

// People 被定义为一种新的结构体，它内嵌了Person类型
type People struct {
	Person
}

func main() {
	person := Person{name: "person"}
	human := Human{"human"}
	dummy := Dummy{"dummy"}
	people := People{Person{"people"}}

	fmt.Println(person.Speak())                 // person say hi
	fmt.Println(human.Speak())                  // human say hi
	fmt.Printf("%v cannot speak\n", dummy.name) // dummy cannot speak
	fmt.Println(people.Speak())                 // people say hi

	// main.Person main.Person main.Dummy main.People
	fmt.Println(reflect.TypeOf(person), reflect.TypeOf(human), reflect.TypeOf(dummy), reflect.TypeOf(people))
}
