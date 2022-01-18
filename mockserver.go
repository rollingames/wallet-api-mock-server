// Copyright (c) 2019-2022 The Rollin.Games developers
// All Rights Reserved.
// NOTICE: All information contained herein is, and remains
// the property of Rollin.Games and its suppliers,
// if any. The intellectual and technical concepts contained
// herein are proprietary to Rollin.Games
// Dissemination of this information or reproduction of this materia
// is strictly forbidden unless prior written permission is obtained
// from Rollin.Games.

package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/rollingames/wallet-api-mock-server/models"
	_ "github.com/rollingames/wallet-api-mock-server/preinit"
	_ "github.com/rollingames/wallet-api-mock-server/routers"
)

func main() {

	models.RegisterDataBase()
	models.RegisterModel()
	enableCORS()

	beego.Run()
}

func enableCORS() {
	v := beego.AppConfig.DefaultBool("enable_cors", true)
	if v {
		beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
			AllowAllOrigins: true,
			AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type", "X-LOGIN-TOKEN"},
			ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
		}))
	}
}
