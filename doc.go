/*
Package ecobee is a Go client for the ecobee API.

For more information about the ecobee API, see the documentation:
https://www.ecobee.com/home/developer/api/introduction/index.shtml

Usage

You use the library by creating a Client object and calling methods on it. The
client can be created manually with NewClient.

For example, to authorize your app and retrieve a selection of thermostat data
for one or more thermostats.

	package main

	import (
		"bufio"
		"context"
		"fmt"
		"os"
		"time"

		"github.com/sfanous/go-ecobee"
		"github.com/sfanous/go-ecobee/objects"
	)

	func main() {
		fmt.Print("Please enter your ecobee Application Key: ")

		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		client := ecobee.NewClient(ecobee.WithApplicationKey(input.Text()))

		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		fmt.Println("Authorization in Progress...Step #1")

		authorizeResponse, err := client.PINAuthorization(ctx, ecobee.ScopeSmartWrite)
		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Printf("AuthorizeResponse %s\n", authorizeResponse)

		fmt.Printf("Please goto ecobee.com, login to the web portal and click on the settings tab. "+
			"Ensure the My Apps widget is enabled. If it is not click on the My Apps option in the menu on the left. "+
			"In the My Apps widget paste '%s' in the textbox labelled 'Enter your 4 digit pin to install your third party app' "+
			"and then click 'Install App'. The next screen will display any permissions the app requires and will ask you to click "+
			"'Authorize' to add the application.\n\nAfter completing this step please hit 'Enter' to continue", authorizeResponse.PIN())

		input.Scan()

		fmt.Println("Authorization in Progress...Step #2")

		ctx = context.Background()
		ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		tokensResponse, err := client.RequestTokens(ctx)
		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Printf("TokensResponse %s\n", tokensResponse)

		ctx = context.Background()
		ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		selection := objects.Selection{
			SelectionType:  ecobee.String("registered"),
			SelectionMatch: ecobee.String(""),
			IncludeEvents:  ecobee.Bool(true),
		}

		thermostatResponse, err := client.Thermostat(ctx, &selection, nil)
		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Printf("ThermostatResponse %s\n", thermostatResponse)
	}

*/
package ecobee
