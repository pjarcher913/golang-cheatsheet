package OOP

import (
	emp "OOP/support"
	"fmt"
)

func Exec() {
	newEmployee := emp.New("Joe", "Shmoe", "joeshmoe@gmail.com", 22, 10, 3)

	neObj := newEmployee.GetFullObject()
	fmt.Println(neObj)

	neFullName := newEmployee.GetName()
	neFirstName := neFullName[0]
	neLastName := neFullName[1]
	fmt.Printf("First Name: %s\nLast Name: %s\n", neFirstName, neLastName)

	neEmail := newEmployee.GetEmail()
	fmt.Printf("Email: %s\n", neEmail)

	newEmployee.SetEmail("newemail@outlook.com")

	neEmail = newEmployee.GetEmail()
	fmt.Printf("Email: %s\n", neEmail)

	neAge := newEmployee.GetAge()
	fmt.Printf("Age: %d\n", neAge)

	neTLD := newEmployee.GetTLD()
	fmt.Printf("TLD: %d\n", neTLD)

	neLDU := newEmployee.GetLDU()
	fmt.Printf("LDU: %d\n", neLDU)

	neRLD := newEmployee.GetRemainingLeaveDays()
	//neRLD := neTLD - neLDU
	fmt.Printf("RLD: %d\n", neRLD)

	newEmployee.PrintInfo()
}
