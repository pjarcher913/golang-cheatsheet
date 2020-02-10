package JSON

import (
	"encoding/json"
	"fmt"
	"time"
)

type response1 struct {
	Page int
	Fruits []string
}

type response2 struct {
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}

type FruitBasket struct {
	Name    string
	Fruit   []string
	Id      int64  `json:"ref"`
	private string // An unexported field is not encoded.
	Created time.Time
}

func Exec() {
	//encodeData()
	//decodeData()
	data := encodeData()
	decodeData(data)
}

func encodeData() []uint8 {
	// Booleans
	testBool := true
	boolB, _ := json.Marshal(testBool)
	fmt.Println(string(boolB))

	// Ints
	testInt := 7
	intB, _ := json.Marshal(testInt)
	fmt.Println(string(intB))

	// Floats
	testFloat := 57.2
	fltB, _ := json.Marshal(testFloat)
	fmt.Println(string(fltB))

	// Strings
	testString := "gopher"
	strB, _ := json.Marshal(testString)
	fmt.Println(string(strB))

	// Arrays
	testArr := []string {"apple", "peach", "pear"}
	arrB, _ := json.Marshal(testArr)
	fmt.Println(string(arrB))

	// Maps
	testMap := map[string]int {"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(testMap)
	fmt.Println(string(mapB))

	// Custom structs
	res1D := &response1{
		Page:   1,
		Fruits: []string {"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string {"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))
	fmt.Printf("%T\n", res2B)

	res3D := FruitBasket{
		Name:    "Standard",
		Fruit:   []string {"Apple, Banana, Orange"},
		Id:      999,
		private: "Second-rate",
		Created: time.Now(),
	}
	res3B, _ := json.Marshal(res3D)
	fmt.Println(string(res3B))

	return res3B
}

func decodeData(dataJSON []uint8) {
	/* IF YOU KNOW THE STRUCT OF INCOMING DATA: */

	// Create destination var to store decoded data
	var storageStruct FruitBasket

	// Check for associated decoding errors
	// Syntax: json.Unmarshal(INCOMING_DATA, DESTINATION_FOR_DECODED_DATA)
	if err := json.Unmarshal(dataJSON, &storageStruct); err != nil {
		panic(err)
	}
	// Data can be accessed as normal for structs
	fmt.Println(storageStruct)
	for i := 0; i < len(storageStruct.Fruit); i++ {
		fmt.Println(storageStruct.Fruit[i])
	}
	fmt.Println(storageStruct.Name)

	/* IF YOU DO NOT KNOW THE STRUCT OF INCOMING DATA (or if ok using map since data not complex): */

	//// We need to provide a var where the JSON package can put the soon-to-be-decoded data.
	//// This map[string]interface{} will hold a map of strings to arbitrary data types.
	//var decodedData map[string]interface{}
	//
	//// Check for associated decoding errors
	//if err := json.Unmarshal(dataJSON, &decodedData); err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(decodedData)
	//fmt.Println(decodedData["fruits"])
	////fmt.Println(decodedData["fruits"][1])		// Can't do because the interface doesn't support indexing
	//fmt.Println(decodedData["page"])
	//fmt.Println(decodedData["TEST"])
}
