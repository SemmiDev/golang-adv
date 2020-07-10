package main
import (
	"fmt"
	"encoding/json"
)

type User struct{
	FullName string `json: "names"`
	Age int
}

func main() {
	var jsonString = `{
		"names": "Sam",
		"Age": 18
	}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user : ", data.FullName)
	fmt.Println("Age  : ", data.Age)

}