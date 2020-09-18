package models

import (
	"encoding/json"
	"fmt"
)

// UnmarshalJSON unmarshals NfType and checks for correct fields.
func (nt *NfType) UnmarshalJSON(b []byte) (err error) {
	type Alias NfType

	var aux Alias

	err = json.Unmarshal(b, &aux)
	if err != nil {
		return err
	}
	fmt.Printf("test: %+v\n", aux)

	// Works so far...

	// switch aux.(NfType) {  // invalid type assertion: aux.(NfType) (non-interface type Alias on left)

	ntAux := NfType(aux)
	nt = &ntAux

	switch *nt {
	case NfType_NRF, NfType_UDM, NfType_AMF, NfType_SMF, NfType_AUSF, NfType_NEF, NfType_PCF, NfType_SMSF, NfType_NSSF, NfType_UDR, NfType_LMF, NfType_GMLC, NfType__5_G_EIR, NfType_SEPP, NfType_UPF, NfType_N3_IWF, NfType_AF, NfType_UDSF, NfType_BSF, NfType_CHF, NfType_NWDAF:
		return nil
	default:
		return fmt.Errorf("Invalid NfType: %s", *nt)
	}
	return nil
}
