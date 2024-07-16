package keycloak

import (
	"context"
	"fmt"

	"github.com/Nerzal/gocloak/v13"
)

const (
	KEYCLOAK_INSTANCE      = "http://10.10.0.3:8080"
	KEYCLOAK_REALM         = "ezt"
	KEYCLOAK_REALM_MASTER  = "master"
	KEYCLOAK_CLIENT_ID     = "ezt"
	KEYCLOAK_CLIENT_SECRET = "GRfeIc5kao4tCvKNhcmdpqpAREWSg32w"
	KEYCLOAK_ADMIN         = "admin"
	KEYCLOAK_ADMIN_PWD     = "keycloak"
)

// NewKeyCloakConnection create a new keycloak instance
func NewKeyCloakConnection() *gocloak.GoCloak {
	return gocloak.NewClient(KEYCLOAK_INSTANCE)
}

// NewAdminKeyCloakToken create a new connection to keycloak with admin account and return its access token
func NewAdminKeyCloakToken(keycloak *gocloak.GoCloak) string {
	token, err := keycloak.LoginAdmin(context.Background(), KEYCLOAK_ADMIN, KEYCLOAK_ADMIN_PWD, KEYCLOAK_REALM_MASTER)
	if err != nil {
		panic("Something wrong with the keycloak credentials or url")
	}
	return token.AccessToken
}

// Validate username and password of a user
func ValidateUser(username string, password string) (bool, string) {
	keycloak := gocloak.NewClient(KEYCLOAK_INSTANCE)
	client, err := keycloak.Login(context.Background(), KEYCLOAK_CLIENT_ID, KEYCLOAK_CLIENT_SECRET, KEYCLOAK_REALM, username, password)
	if err != nil {
		fmt.Println("Error: ", err)
		return false, ""
	}
	return true, client.AccessToken
}
