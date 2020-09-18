package main

import (
	"encoding/json"
	"fmt"
	_ "reflect"

	"gotest.tools/assert"

	"github.com/voje/stayinshape/golang/restcheck/models"
)

func main() {
	// We'll be using models.NfProfile.NfType for demonstration

	// Empty payload
	payload := models.NfProfile{}
	payload.NfType = models.NfType_AMF // This is a valid NfType

	// Let's modify some fields.

	b, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	var validPayload models.NfProfile
	err = json.Unmarshal(b, &validPayload)
	if err != nil {
		panic(err)
	}
	fmt.Printf(`Unmarshalled a valid payload.
NfType: %s
`, validPayload.NfType)

	// Now let's create a corrupted payload
	payload.NfType = "CorruptedNFType**!!--!!"

	b, err = json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	var corruptedPayload models.NfProfile
	err = json.Unmarshal(b, &corruptedPayload)
	assert.EqualError(err)
}
