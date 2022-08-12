package fakedata

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
)

// SomeStructWithTags ...
type SomeStructWithTags struct {
	FirstName   string `faker:"first_name"`
	LastName    string `faker:"last_name"`
	Email       string `faker:"email"`
	Password    string `faker:"password"`
	PhoneNumber string `faker:"phone_number"`
	Date        string `faker:"date"`
	Time        string `faker:"time"`
	Timestamp   string `faker:"timestamp"`
}

func Example_withTags() {

	a := SomeStructWithTags{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", a)
}
