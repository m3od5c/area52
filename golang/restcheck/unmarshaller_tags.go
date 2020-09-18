// Does a custom unmarshaller override tags?
// Let's find out!

package main

import (
	"encoding/json"
	"fmt"
)

type Grocery struct {
	Milk   string `json:"mmilk"`
	Eggs   string `json:,omitempty`
	Budget int    `json:"budget"`
	Bread  string `json:"bread"`
}

func (g *Grocery) UnmarshalJSON

func main() {

	// In this example, we didn't add Eggs and Bread to the byteString
	// While Eggs get ommitted (omitempty), Bread is added as a default value (empty string)
	bs := []byte(`{"mmilk":"Ljubljanske Mlekarne", "budget": 10}`)
	var a Grocery
	err := json.Unmarshal(bs, &a)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", a)

}
