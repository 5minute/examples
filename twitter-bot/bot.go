package main

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	consumerKey := "fill-me"
	consumerSecret := "fill-me"
	accessToken := "fill-me"
	accessTokenSecret := "fill-me"
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	user, _, err := client.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("Account: @%s (%s)\n", user.ScreenName, user.Name)
	_, _, err = client.Statuses.Update("New project is coming soon - stay tuned!", nil)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println("Twitted successfully")
}
