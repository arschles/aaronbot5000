package helix

import (
	"os"

	"github.com/nicklaw5/helix"
)

// GetHelixClient returns a pointer to the helix client
func GetHelixClient() (*helix.Client, error) {
	helixClient, err := helix.NewClient(&helix.Options{
		ClientID:     os.Getenv("TWITCH_CLIENT_ID"),
		ClientSecret: os.Getenv("TWITCH_CLIENT_SECRET"),
	})
	if err != nil {
		return nil, err
	}

	token, err := helixClient.GetAppAccessToken()
	if err != nil {
		return nil, err
	}

	helixClient.SetUserAccessToken(token.Data.AccessToken)
	return helixClient, nil
}
