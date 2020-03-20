package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/antihax/optional"
	"golang.org/x/oauth2"

	api "github.com/adovenmuehle/mattermost-looker-plugin/openapi"
)

func main() {
	clientID := os.Getenv("LOOKER_API_KEY")
	clientSecret := os.Getenv("LOOKER_API_SECRET")
	client := api.NewAPIClient(api.NewConfiguration())

	accessToken, _, _ := client.ApiAuthApi.Login(context.Background(), &api.LoginOpts{ClientId: optional.NewString(clientID), ClientSecret: optional.NewString(clientSecret)})

	token := oauth2.Token{
		AccessToken: accessToken.AccessToken,
		Expiry:      time.Now().Add(time.Duration(accessToken.ExpiresIn)),
		TokenType:   accessToken.TokenType,
	}

	tokenSource := oauth2.StaticTokenSource(&token)

	auth := context.WithValue(context.Background(), api.ContextOAuth2, tokenSource)
	userID := optional.NewInt64(int64(47))

	r, _, _ := client.ContentApi.SearchContentFavorites(auth, &api.SearchContentFavoritesOpts{UserId: userID})

	prettyJSON, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))

	// r, _, _ := client.LookApi.RunLook(auth, int64(85), "md", &api.RunLookOpts{})

	// fmt.Println(r)
}
