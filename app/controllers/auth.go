/*
* @Author: Piyush Agrawal
* @Date:   2017-10-22 14:24:25
* @Last Modified by:   Piyush Agrawal
* @Last Modified time: 2017-10-23 00:26:19
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
	Endpoint:     "",
	RedirectURL:  "http://localhost:9000/Auth/Callback",
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