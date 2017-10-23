/*
* @Author: Piyush Agrawal
* @Date:   2017-10-22 14:24:25
* @Last Modified by:   Piyush Agrawal
* @Last Modified time: 2017-10-23 22:05:15
*/

package controllers

import (
	"github.com/revel/revel"
	"os"
	"golang.org/x/oauth2"
)

type Auth struct {
	*revel.Controller
}

var Endpoint = oauth2.Endpoint{
    AuthURL:  "https://cloud.digitalocean.com/v1/oauth/authorize?scope=write",
    TokenURL: "https://cloud.digitalocean.com/v1/oauth/token",
}

var DO = &oauth2.Config{
	ClientID:     os.Getenv("clientid"),
	ClientSecret: os.Getenv("secret"),
	Endpoint: Endpoint,
	RedirectURL:  os.Getenv("redirect_url"),
}

// func (c Auth) Register() revel.Result {
// 	u := c.connected()
// 	me := map[string]interface{}{}
// 	if u != nil && u.AccessToken != "" {
// 		//....
// 	}

// 	authUrl := DO.AuthCodeURL("state", oauth2.AccessTypeOffline)
// 	return c.Render(me, authUrl)
// }

// func (c Auth) Callback() revel.Result {
// 	errr := c.Params.Query.Get("error")
// 	if len(errr) > 0 {
// 		c.ViewArgs["token"] = os.Getenv("clientid")
// 		c.ViewArgs["secret"] = os.Getenv("secret")
// 		c.ViewArgs["err"] = c.Params.Query.Get("error_description")
// 		return c.RenderTemplate("Auth/Register.html")
// 	}

// 	code := c.Params.Query.Get("code")
// 	return c.Render(code)
// }

// func (c Hotels) checkUser() revel.Result {
// 	if user := c.connected(); user == nil {
// 		c.Flash.Error("Please log in first")
// 		return c.Redirect(routes.Application.Index())
// 	}
// 	return nil
// }