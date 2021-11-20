package common

import (
	"context"
	"errors"
	"fmt"
	"github.com/Nerzal/gocloak/v9"
)

var adminClient gocloak.GoCloak
var adminToken *gocloak.JWT
var authConfig clockConfig

func getAdminClient() gocloak.GoCloak {
	if adminClient == nil {
		fmt.Println("adminClient refreshed")
		adminClient = gocloak.NewClient(authConfig.clientUrl)
	}
	return adminClient
}

type clockConfig struct {
	clientUrl    string
	clientRealm  string
	clientId     string
	clientSecret string
	username     string
	password     string
}

func buildClockConfig(url, realm, clientId, clientSecret, username, password string) {
	authConfig = clockConfig{
		clientUrl:    url,
		clientRealm:  realm,
		clientId:     clientId,
		clientSecret: clientSecret,
		username:     username,
		password:     password,
	}
}

func ConnectKeyCloak(realm, basePath, username, password, clientId, clientSecret string) {
	buildClockConfig(basePath, realm, clientId, clientSecret, username, password)
	adminClient = gocloak.NewClient(basePath)
	ctx := context.Background()
	token, err := getAdminClient().Login(ctx, clientId, clientSecret, realm, username, password)
	if err != nil {
		fmt.Println(err)
	}
	adminToken = token
	if adminToken == nil {
		fmt.Println("Initialize failed: admin Token is nil")
	}

}

func RetrospectAndRefreshToken() {
	ctx := context.Background()

	if adminToken == nil {
		adminToken, _ = getAdminClient().Login(ctx, authConfig.clientId, authConfig.clientSecret, authConfig.clientRealm, authConfig.username, authConfig.password)
	}

	rptResult, err := getAdminClient().RetrospectToken(ctx, adminToken.AccessToken, authConfig.clientId, authConfig.clientSecret, authConfig.clientRealm)
	fmt.Println(rptResult)
	if !*rptResult.Active {
		ctx := context.Background()
		adminToken, err = getAdminClient().RefreshToken(ctx, adminToken.RefreshToken, authConfig.clientId, authConfig.clientSecret, authConfig.clientRealm)
		RetrospectAndRefreshToken()
		if err != nil {
			adminToken, _ = getAdminClient().Login(ctx, authConfig.clientId, authConfig.clientSecret, authConfig.clientRealm, authConfig.username, authConfig.password)
			fmt.Println("Refresh failed:" + err.Error())
		}
	}
}

func CreateUser(user *gocloak.User) (string, error) {

	if getAdminClient() == nil {
		return "", errors.New("key cloak server not connected")
	}
	if user == nil {
		return "", errors.New("new user is nil")
	}

	RetrospectAndRefreshToken()

	ctx := context.Background()
	res, err := getAdminClient().CreateUser(ctx, adminToken.AccessToken, authConfig.clientRealm, *user)
	if err != nil {
		fmt.Println("keycloak create user fail")
		fmt.Println(adminToken.AccessToken)
		fmt.Println(err)
	}
	return res, err
}

func SetPassword(userID string, password string) error {
	RetrospectAndRefreshToken()
	ctx := context.Background()
	err := getAdminClient().SetPassword(ctx, adminToken.AccessToken, userID, authConfig.clientRealm, password, false)
	return err
}

func ValidateToken(accessToken string) (*gocloak.UserInfo, error) {
	ctx := context.Background()
	return getAdminClient().GetUserInfo(ctx, accessToken, authConfig.clientRealm)
}
