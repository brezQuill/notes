package struct_

type Person struct {
	name string
	age  int
	sex  bool
}

func NewPerson(name string, age int, sex bool) Person {
	return Person{
		name: name,
		age:  age,
		sex:  sex,
	}
}
