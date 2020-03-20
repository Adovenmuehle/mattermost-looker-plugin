package main

import (
	"context"

	api "github.com/adovenmuehle/mattermost-looker-plugin/openapi"
)

func main() {
	auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{Key: "TEST123"})
}
