/*
NOTE: MAPS ARE KEY-VALUE PAIRS (DICTIONARIES)!
*/

package Maps

import (
	"fmt"
)

func Exec() {
	maps()
}

func maps() {
	// Define map
	// <VALUE> := make(map[<KEY>]<VALUE_DATA_TYPE>)
	emails := make(map[string]string)

	// Add/assign key-values ("KVs")
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"

	// Print
	fmt.Println("emails: ", emails)
	fmt.Println("emails[\"Bob\"]: ", emails["Bob"])
	fmt.Println("len(emails): ", len(emails))

	// Delete element from map
	delete(emails, "Bob")
	fmt.Println("emails (after \"Bob\" deletion): ", emails)

	// Declare map and add KVs on instantiation
	//emails_new := map[string]string {"Joe": "joe@gmail.com", "Beth": "beth@gmail.com"}
	emails_new := map[string]string {
		"Joe": "joe@gmail.com",
		"Beth": "beth@gmail.com",
	}
	fmt.Println("emails_new: ", emails_new)
}
