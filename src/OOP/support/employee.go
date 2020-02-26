package support

import (
	"fmt"
)

type employee struct {
	firstName 		string
	lastName 		string
	email			string
	age 			int
	totalLeaveDays	int
	leaveDaysUsed	int
}

/* INITIALIZER */

func New(fn string, ln string, email string, age int, tld int, ldu int) employee {
	n := employee{fn, ln, email, age, tld, ldu}
	return n
}

/* GETTERS */

func (e *employee) GetFullObject() employee {
	return *e
}

func (e *employee) GetName() [2]string {
	eArr := [2]string {e.firstName, e.lastName}
	return eArr
}

func (e *employee) GetEmail() string {
	return e.email
}

func (e *employee) GetAge() int {
	return e.age
}

func (e *employee) GetTLD() int {
	return e.totalLeaveDays
}

func (e *employee) GetLDU() int {
	return e.leaveDaysUsed
}

func (e *employee) GetRemainingLeaveDays() int {
	return e.totalLeaveDays - e.leaveDaysUsed
}

/* SETTERS */

func (e *employee) SetName(newFN string, newLN string) {
	e.firstName = newFN
	e.lastName = newLN
}

func (e *employee) SetEmail(newEmail string) {
	e.email = newEmail
}

func (e *employee) SetAge(newAge int) {
	e.age = newAge
}

func (e *employee) SetTLD(newTLD int) {
	e.totalLeaveDays = newTLD
}

func (e *employee) SetLDU(newLDU int) {
	e.leaveDaysUsed = newLDU
}

/* PRINTERS */

func (e employee) PrintInfo() {
	fmt.Printf("First Name: %s\nLast Name: %s\nEmail: %s\nAge: %d\nLeave Days Remaining: %d\n",
		e.firstName, e.lastName, e.email, e.age, (e.totalLeaveDays - e.leaveDaysUsed))
}
