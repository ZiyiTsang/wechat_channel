package interface_test_tool

import (
	"fmt"
	"testing"
)

type Animal interface {
	speak()
	walk()
}
type Duck struct {
	name  string
	class string //Duck or dog?
}

func (duck *Duck) speak() {
	fmt.Printf("%s says:I'm %s,Ga!Ga!Ga!\n", duck.name, duck.class)
}
func (duck *Duck) walk() {
	fmt.Printf("%s is walking with 2 legs\n", duck.name)
}

type Dog struct {
	name  string
	class string
}

func (dog *Dog) speak() {
	fmt.Printf("%s says:I'm %s,wang!wang!wang!\n", dog.name, dog.class)
}
func (dog *Dog) walk() {
	fmt.Printf("%s is walking with 4 legs\n", dog.name)
}

func TestInterface(t *testing.T) {
	//首先声明两个子类指针
	//指针1 --> 指向 Duck（子类）
	var duck_pointer *Duck = &Duck{name: "Xiao Ya", class: "Duck"}
	//指针2 --> 指向 Dog（子类）
	dog_pointer := new(Dog)
	dog_pointer.class = "Dog"
	dog_pointer.name = "Xiao Gou"
	//然后声明接口(没接触过GO的同学可以理解为父类)
	var animal Animal
	//其次将指针赋予接口
	fmt.Println("Now test:Duck")
	animal = duck_pointer
	animal.speak() //Xiao Ya says:I'm Duck,Ga!Ga!Ga!
	animal.walk()  //Xiao Ya is walking with 2 legs
	fmt.Println("Now twst:Dog")
	animal = dog_pointer
	animal.speak() //Xiao Gou says:I'm Dog,wang!wang!wang!
	animal.walk()  //Xiao Gou is walking with 4 legs
}
