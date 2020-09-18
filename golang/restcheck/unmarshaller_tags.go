// Does a custom unmarshaller override tags?
// Let's find out!

package main

import (
	"encoding/json"
	"fmt"

	copier "github.com/jinzhu/copier"
)

// ForgotChocolate is a special error
type ForgotChocolate struct {
	ErrMsg          string
	AvailableBudget int
}

// Error implements the error interface on ForgotChocolate - now we can pass ForgotChocolate as a standard error
func (fc *ForgotChocolate) Error() string {
	return string(fc.ErrMsg)
}

// Grocery is our shopping cart
type Grocery struct {
	Milk      string `json:"mmilk"`
	Eggs      string `json:,omitempty` // doesn't even matter when unmarshalling
	Budget    int    `json:"budget"`
	Bread     string `json:"bread"`
	Chocolate string `json:"chocolate"`
}

// UnmarshalJSON will unmarshall grocery list and throw a special error, if we've forgotten the chocolate
func (g *Grocery) UnmarshalJSON(b []byte) error {

	// If we passed Grocery to json.Unmarshal directly, we would end up in endless recursion
	// Read into Alias instead, then copy into Grocery
	// TODO: can we optimize to read into g directly?
	type Alias Grocery

	var ag Alias

	err := json.Unmarshal(b, &ag)
	if err != nil {
		return err
	}

	// Wee need a deep copy here - check performance. TODO
	copier.Copy(g, &ag)

	// Custom checks
	if g.Chocolate == "" {

		// Unmarshal will fail with the following error
		// We want to catch and handle the error in main()
		myErr := ForgotChocolate{
			ErrMsg:          "FORGOT CHOCOLATE!!!",
			AvailableBudget: g.Budget,
		}

		return &myErr
	}

	// No errors
	return nil
}

func main() {
	// byteString received over REST
	jsonBytes := []byte(`{"mmilk":"Ljubljanske Mlekarne", "budget": 10}`)
	fmt.Printf("jsonBytes:\n%s\n", string(jsonBytes))

	// Read the byteString into our Grocery struct
	var a Grocery
	err := json.Unmarshal(jsonBytes, &a)
	if err != nil {
		// Catch possible errors and handle them
		fmt.Printf("Caught a nasty error: %s\n", err.Error())

		// Handle the error
		// We need to typecast here
		availableBudget := err.(*ForgotChocolate).AvailableBudget
		if availableBudget > 2 {
			fmt.Printf("Don't worry, we still have some money (%d €), add chocolate to the grocery list.\n", availableBudget)
		} else {
			panic(err)
		}
	}

	fmt.Printf("Unmarshalled struct:\n%+v\n", a)
}
