# smm-go
API Wrapper for all popular smm panels

# Install
```
go get github.com/dropout1337/smm-go
```

# Example
```golang
package main

import (
	"fmt"
	smmgo "smm-go/smm"
	"strings"
)

const (
	BaseURL string = "https://smm.solar/api/v2"
	apiKey  string = "your apikey"
)

func main() {
	client := smmgo.New(BaseURL, apiKey)
	balance := client.GetBalance()
	services := client.GetServices()

	fmt.Printf("You have %s%s in your balance.\n", balance.Balance, balance.Currency)
	fmt.Printf("There are currently %d services avaliable!\n\n", len(*services))

	var option string
	fmt.Print("Would you like to print all the services? ")
	fmt.Scanln(&option)

	if strings.Contains(option, "y") {
		for _, service := range *services {
			fmt.Printf("%s\n - Type: %s\n - Rate: %s\n - Min: %s\n - Max: %s\n", service.Type, service.Name, service.Rate, service.Min, service.Max)
		}
	}
}
```
