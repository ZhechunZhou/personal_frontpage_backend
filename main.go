package main

import (
	"git.wisecar.co/wisecar/web-api-golang/common"
	"git.wisecar.co/wisecar/web-api-golang/config"
)

func main() {
	var conf = Value()
	//connect database
	common.ConnectDb(conf.DBHost, conf.DBName, conf.User, conf.Password)
	//uncommented when the data model is changed
	//config.Migrate()

	//set up s3 bucket
	common.Builds3Config(conf.AWSAccessKeyID, conf.AWSSecretAccessKey, conf.AWSRegion, conf.AWSBucket)
	//set up keyClock
	common.ConnectKeyCloak(conf.CloakClientRealm, conf.CloakUrl, conf.CloakUser, conf.CloakPassword, conf.CloakAdminCliId, conf.ClockAdminCliSecret)

	//set up stripe
	common.SetUpStrip()
	mainRouter := config.SetupRouter()
	err := mainRouter.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		return
	}
}
