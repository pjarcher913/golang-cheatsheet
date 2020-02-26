package main

import (
	ArrSli "Arrays_Slices"
	"Closures"
	Conds "Conditionals"
	Funcs "Functions"
	"Interfaces"
	j "JSON"
	logger "Logs"
	"Loops"
	"Maps"
	"OOP"
	Packs "Packages"
	Ptrs "Pointers"
	"Ranges"
	Structs "Structures"
	"Vars"
	"fmt"
)

func main() {

	/* TODO: DYNAMIC EXERCISE EXECUTION */
	//exercises := []string {
	//	"Vars", "Packs", "Funcs", "ArrSli", "Conds", "Loops", "Maps", "Ranges", "Ptrs", "Closures",
	//	"Structs", "Interfaces", "Web", "OOP",
	//}
	//
	//for i := 1; i < len(exercises); i++ {
	//	exName := exercises[i]
	//	fmt.Printf("\n=== Starting App.%s.Exec() ===\n\n", exName)
	//	// TODO: execute function with concatenated name
	//}

	fmt.Print("\n=== Starting App.Vars.Exec() ===\n\n")
	Vars.Exec()
	fmt.Print("\n=== Starting App.Packages.Exec() ===\n\n")
	Packs.Exec()
	fmt.Print("\n=== Starting App.Functions.Exec() ===\n\n")
	Funcs.Exec()
	fmt.Print("\n=== Starting App.Arrays_Slices.Exec() ===\n\n")
	ArrSli.Exec()
	fmt.Print("\n=== Starting App.Conditionals.Exec() ===\n\n")
	Conds.Exec()
	fmt.Print("\n=== Starting App.Loops.Exec() ===\n\n")
	Loops.Exec()
	fmt.Print("\n=== Starting App.Maps.Exec() ===\n\n")
	Maps.Exec()
	fmt.Print("\n=== Starting App.Ranges.Exec() ===\n\n")
	Ranges.Exec()
	fmt.Print("\n=== Starting App.Pointers.Exec() ===\n\n")
	Ptrs.Exec()
	fmt.Print("\n=== Starting App.Closures.Exec() ===\n\n")
	Closures.Exec()
	fmt.Print("\n=== Starting App.Structures.Exec() ===\n\n")
	Structs.Exec()
	fmt.Print("\n=== Starting App.Interfaces.Exec() ===\n\n")
	Interfaces.Exec()
	fmt.Print("\n=== Starting App.WebServer.Exec() ===\n\n")
	//Web.Exec()	// disabled so web server doesn't actually initialize and start listening (blocks further exec)
	fmt.Print("\n=== Starting App.OOP.Exec() ===\n\n")
	OOP.Exec()
	fmt.Print("\n=== Starting App.JSON.Exec() ===\n\n")
	j.Exec()
	fmt.Print("\n=== Starting App.Logger.Exec() ===\n\n")
	logger.Exec()
}
