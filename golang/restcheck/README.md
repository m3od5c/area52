# Checking REST data
When consuming REST requests, we want to check the validity of certain fields.   
We're going to override `UnmarshalJSON` methods next to type definitions in `models`.   

Some important questions:
* Does `UnmarshalJSON` oveeride json tags?

Todo:
* Pass a custom error using `models.ProblemDetails`

## reflect
Concerns about speed.  
Can't list const names - this is a game-stopper.   

## Implement Unmarshaller
Seems the most idiomatic approach, doesn't require additional function calls.   
Note that to we can only implement `Unmarshaller` in the original package (e.g. in models) - TODO: separate from `free5gc/`

See demo implementation for `NfType` in `./models/model_nf_type_unmarshaller.go`.   
```go
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
```

Example error output:
```bash
test: AMF
Unmarshalled a valid payload.
NfType: 
test: CorruptedNFType**!!--!!
panic: Invalid NfType: CorruptedNFType**!!--!!

goroutine 1 [running]:
main.main()
	/home/voje/git/stayinshape/golang/restcheck/main.go:45 +0x45e
exit status 2
```

## Some eamples
Apparently `omitempty` only works `go` -> `JSON` so we don't have to worry about unmarshalling.   
```go
type Grocery struct {
	Milk   string `json:"mmilk"`
	Eggs   string `json:,omitempty`
	Budget int    `json:"budget"`
	Bread  string `json:"bread"`
}
bs := []byte(`{"mmilk":"Ljubljanske Mlekarne", "budget": 10}`)
```
```bash
{
 "mmilk": "Ljubljanske Mlekarne",
 "Eggs": "",
 "budget": 10,
 "bread": ""
}
```

## Unmarshal and check
```


