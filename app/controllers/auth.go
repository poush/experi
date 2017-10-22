/*
* @Author: Piyush Agrawal
* @Date:   2017-10-22 14:24:25
* @Last Modified by:   Piyush Agrawal
* @Last Modified time: 2017-10-22 14:48:00
*/

package controllers

import (
	"github.com/revel/revel"
)

type Auth struct {
	*revel.Controller
}

func (c Auth) Register() revel.Result {
	return c.Render()
}
