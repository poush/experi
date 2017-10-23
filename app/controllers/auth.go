/*
* @Author: Piyush Agrawal
* @Date:   2017-10-22 14:24:25
* @Last Modified by:   Piyush Agrawal
* @Last Modified time: 2017-10-23 20:17:47
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

var DO = &oauth2.Config{
	ClientID:     os.Getenv("clientid"),
	ClientSecret: os.Getenv("secret"),
	Scopes:       ["write"],
	AuthURL:      "https://graph.facebook.com/oauth/authorize",
	TokenURL:     "https://graph.facebook.com/oauth/access_token",
	RedirectURL:  os.Getenv("redirecr_url"),
}

func (c Auth) Register() revel.Result {

}

func (c Auth) Callback() revel.Result {
	errr := c.Params.Query.Get("error")
	if len(errr) > 0 {
		c.ViewArgs["token"] = os.Getenv("clientid")
		c.ViewArgs["secret"] = os.Getenv("secret")
		c.ViewArgs["err"] = c.Params.Query.Get("error_description")
		return c.RenderTemplate("Auth/Register.html")
	}

	code := c.Params.Query.Get("code")
	return c.Render(code)
}