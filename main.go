package main

import (
	"fmt"
	"log"

	"github.com/beerzezy/validate-pincode/validatepincode"
)

func main() {
	for {
		fmt.Println("Enter pincode :")
		var pinCode string
		_, err := fmt.Scan(&pinCode)
		if err != nil {
			log.Println(err)
		}

		isValid := validatepincode.IsValidPinCode(pinCode)
		fmt.Println("validate pincode :", isValid)
	}
}
