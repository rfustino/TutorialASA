package main

import (
	json "encoding/json"
	"fmt"

	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/mnemonic"
)

func main() {
	account1 := crypto.GenerateAccount()
	account2 := crypto.GenerateAccount()
	account3 := crypto.GenerateAccount()
	address1 := account1.Address.String()
	address2 := account2.Address.String()
	address3 := account3.Address.String()

	mnemonic1, err := mnemonic.FromPrivateKey(account1.PrivateKey)
	if err != nil {
		return
	}
	mnemonic2, err := mnemonic.FromPrivateKey(account2.PrivateKey)
	if err != nil {
		return
	}
	mnemonic3, err := mnemonic.FromPrivateKey(account3.PrivateKey)
	if err != nil {
		return
	}
	fmt.Printf("Account 1 = %s\n", address1)
	fmt.Printf("Account 2 = %s\n", address2)
	fmt.Printf("Account 3 = %s\n", address3)
	fmt.Printf("")
	fmt.Printf("Copy off accounts above and add TestNet Algo funds using the TestNet Dispenser at https://bank.testnet.algorand.network/\n")	
	fmt.Printf("Copy off the following mnemonic code for use in Step 1B\n")
	fmt.Printf("\n")	
	fmt.Printf("mnemonic1 := \"%s\"\n", mnemonic1)
	fmt.Printf("mnemonic2 := \"%s\"\n", mnemonic2)
	fmt.Printf("mnemonic3 := \"%s\"\n", mnemonic3)
}

// Terminal output should look similar to this...

// Account 1 = BQBRKJ7KSQO6WWX3QRSITPCZOEDOT6DX6MJVQQFQI72PDPTKWU5WKOOJ2I
// Account 2 = LBMZ2DNGB5N7AJNZQZPMTPDJRW3S6OB6YR4EHE5RYODRS67TILA5H23K4A
// Account 3 = 6AVEBUP35AHIKDA6UGSFM6LP433AUHVRZNW4KHAXOYCOYDYOLFJWZ7J3I4
// Copy off accounts above and add TestNet Algo funds using the TestNet Dispenser at https://bank.testnet.algorand.network/
// Copy off the following mnemonic code for use in Step 1B

// mnemonic1 := "fatigue laundry small early radar code supreme diary skin record slice distance bike skirt guard surround miss turtle horror frame train taxi column ability forest"
// mnemonic2 := "ill live coconut risk east flower snack develop boring enroll nest rice mistake pioneer rival account coffee depend bachelor orient rebel inform throw absent option"
// mnemonic3 := "broom bid found recall stick gas sample copy network mistake mind relief rely file disorder east asthma program filter hedgehog legal walnut wait about slogan"
