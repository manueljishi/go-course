package main

import "fmt"

type contactInfo struct{
	email string
	zipcode int
}

type person struct{
	firstName string
	lastName string
	contactInfo contactInfo
}

func main()  {
	newPerson := person{
		firstName: "Alex",
		lastName: "Garcia",
		contactInfo: contactInfo{
			email: "alexgarcia@mail.com",
			zipcode: 46020,
		},
	}
	newPerson.print()
	newPerson.updateName("Pepe")
	newPerson.print()
}

func (p *person) updateName(newName string){
	(*p).firstName = newName
}

func (p person) print(){
	fmt.Printf("%+v", p)
}