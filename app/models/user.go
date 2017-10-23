/*
* @Author: Piyush Agrawal
* @Date:   2017-10-23 20:52:16
* @Last Modified by:   Piyush Agrawal
* @Last Modified time: 2017-10-23 22:50:24
*/
package models

import (
	"fmt"
)

type User struct {
	UserId             int
	Name               string
	Username           string
	Password           []byte
	DoToken			   string
	Github			   string
}

func (u *User) String() string {
	return fmt.Sprintf("User(%s)", u.Username)
}
