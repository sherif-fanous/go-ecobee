/*
Package ecobee provides a Go client for the ecobee API.

For more information about the ecobee API, see the documentation:
https://www.ecobee.com/home/developer/api/introduction/index.shtml

# Authentication

By design, the ecobee Client accepts any http.Client so OAuth2 requests can be
made by using the appropriate authenticated client.
Use the https://github.com/golang/oauth2 package to obtain an http.Client which
transparently authorizes requests.

# Usage

You use the library by creating a Client and invoking its methods. The client
can be created manually with NewClient.

The below example illustrates how to:

- Fetch a persisted OAuth2 token from a file in JSON format if one exists.

- Authorize an app and create the OAuth2 token using the ecobee PIN authorization method if no persisted OAuth2 token is found.

- Create an http.Client using the OAuth2 token. The Client will auto-refresh the token as necessary.

- Create an ecobee.Client.

- Retrieve a selection of thermostat data for one or more thermostats.

- Retrieve the OAuth2 token from the ecobee.Client (In case it was auto-refreshed).

- Persist the OAuth2 token to a file in JSON format

	package main

	import (
		"bufio"
		"context"
		"encoding/json"
		"flag"
		"fmt"
		"io/ioutil"
		"os"
		"time"

		"golang.org/x/oauth2"

		"github.com/sherif-fanous/go-ecobee"
		"github.com/sherif-fanous/go-ecobee/objects"
	)

	var applicationKey string
	var filename string
	var oauth2Endpoint = oauth2.Endpoint{
		AuthURL:  "https://api.ecobee.com/authorize",
		TokenURL: "https://api.ecobee.com/token",
	}

	func persistToken(oauth2Token *oauth2.Token) error {
		b, err := json.MarshalIndent(oauth2Token, "", " ")
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(filename, b, 0644); err != nil {
			return err
		}

		return nil
	}

	func fetchToken() (*oauth2.Token, error) {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			if os.IsNotExist(err) {
				return nil, nil
			}

			return nil, err
		}

		oauth2Token := oauth2.Token{}

		if err := json.Unmarshal(b, &oauth2Token); err != nil {
			return nil, err
		}

		return &oauth2Token, nil
	}

	func main() {
		flag.StringVar(&applicationKey, "key", "", "ecobee Application Key")
		flag.StringVar(&filename, "file", "", "token persistent store path")
		flag.Parse()

		oauth2Token, err := fetchToken()
		if err != nil {
			fmt.Println(err)

			return
		}

		if oauth2Token == nil {
			client := ecobee.NewClient()

			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()

			fmt.Println("PINAuthorization...Step #1")
			fmt.Println()

			authorizeResponse, err := client.PINAuthorization(ctx, applicationKey, ecobee.ScopeSmartWrite)
			if err != nil {
				fmt.Println(err)

				return
			}

			fmt.Printf("AuthorizeResponse %s\n", authorizeResponse)
			fmt.Println()

			fmt.Printf("Please goto ecobee.com, login to the web portal and click on the settings tab. "+
				"Ensure the My Apps widget is enabled. If it is not click on the My Apps option in the menu on the left. "+
				"In the My Apps widget paste '%s' in the textbox labelled 'Enter your 4 digit pin to install your third party app' "+
				"and then click 'Install App'. The next screen will display any permissions the app requires and will ask you to click "+
				"'Authorize' to add the application.\n\nAfter completing this step please hit 'Enter' to continue", authorizeResponse.PIN())

			input := bufio.NewScanner(os.Stdin)
			input.Scan()

			fmt.Println()
			fmt.Println("PINAuthorization...Step #2")
			fmt.Println()

			nowUTC := time.Now().UTC()

			ctx = context.Background()
			ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
			defer cancel()

			tokensResponse, err := client.RequestTokens(ctx, applicationKey, authorizeResponse.AuthorizationToken())
			if err != nil {
				fmt.Println(err)

				return
			}

			fmt.Printf("TokensResponse %s\n", tokensResponse)

			oauth2Token = &oauth2.Token{
				TokenType:    tokensResponse.TokenType(),
				AccessToken:  tokensResponse.AccessToken(),
				Expiry:       nowUTC.Add(time.Second * time.Duration(tokensResponse.ExpiresIn())),
				RefreshToken: tokensResponse.RefreshToken(),
			}
		}

		oauth2Config := &oauth2.Config{
			ClientID: applicationKey,
			Endpoint: oauth2Endpoint,
		}

		client := ecobee.NewClient(ecobee.WithHTTPClient(oauth2Config.Client(context.Background(), oauth2Token)))

		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		selection := objects.Selection{
			SelectionType:   ecobee.String("registered"),
			SelectionMatch:  ecobee.String(""),
			IncludeSettings: ecobee.Bool(true),
		}

		thermostatResponse, err := client.Thermostat(ctx, &selection, nil)
		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Printf("ThermostatResponse %s\n", thermostatResponse)

		// Get the updated token
		oauth2Token = client.OAuth2Token()

		if err := persistToken(oauth2Token); err != nil {
			fmt.Println(err)

			return
		}
	}
*/
package ecobee
