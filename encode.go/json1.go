package main

import ( //packages required to import the package
	"encoding/json"
	"fmt"
	"log"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Password string   `json:"-"`
	Tags     []string `json:`
}

func main() {
	fmt.Println("Welcome to JSON video")
	//EncodeJson()
	Decodejson()
}
func EncodeJson() {
	lcocourse := []course{
		{"Ram", 56, "abc123", []string{"web-dev", "js"}},
		{"Shyam", 88, "iodfc123", []string{"full-stack", "js"}},
		{"Raju", 96, "adsd8923", []string{"google", "js"}},
	}
	//package this data as json data
	finaljson, err := json.MarshalIndent(lcocourse, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", finaljson)
}

func Decodejson() { //decoding json function
	jsondata := []byte(`
	{
		"coursename": "Ram",
                "Price": 56,
                "Tags": [
                        "web-dev",
                        "js"
                ]
	}
	`)
	var lcocourse course
	checkValid := json.Valid(jsondata)//checks validity of json

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsondata, &lcocourse)
		fmt.Printf("%#v\n", lcocourse)
	} else {
		fmt.Println("JSON was not valid")
	}
	//map
	var data map[string]interface{} //declaration of maps
	json.Unmarshal(jsondata, &data)
	fmt.Printf("%#v\n", data)

	for k, v := range data {
		fmt.Printf("Key is %v and value is %v", k, v)
	}

}
