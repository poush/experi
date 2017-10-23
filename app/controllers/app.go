package controllers

import (
	"github.com/revel/revel"
	// "golang.org/x/crypto/bcrypt"
	// "github.com/poush/sharkDeploy/app/models"
	// "github.com/poush/sharkDeploy/app/routes"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render();
}
// func (c Application) AddUser() revel.Result {
// 	if user := c.connected(); user != nil {
// 		c.ViewArgs["user"] = user
// 	}
// 	return nil
// }

// func (c Application) connected() *models.User {
// 	if c.ViewArgs["user"] != nil {
// 		return c.ViewArgs["user"].(*models.User)
// 	}
// 	if username, ok := c.Session["user"]; ok {
// 		return c.getUser(username)
// 	}
// 	return nil
// }

// func (c Application) getUser(username string) *models.User {
// 	users, err := c.Txn.Select(models.User{}, `select * from User where Username = ?`, username)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if len(users) == 0 {
// 		return nil
// 	}
// 	return users[0].(*models.User)
// }

// func (c Application) Index() revel.Result {
// 	if c.connected() != nil {
// 		return c.Redirect(routes.Auth.Index())
// 	}
// 	c.Flash.Error("Please log in first")
// 	return c.Render()
// }

// func (c Application) Register() revel.Result {
// 	return c.Render()
// }

// func (c Application) SaveUser(user models.User, verifyPassword string) revel.Result {
// 	c.Validation.Required(verifyPassword)
// 	c.Validation.Required(verifyPassword == user.Password).
// 		Message("Password does not match")
// 	user.Validate(c.Validation)

// 	if c.Validation.HasErrors() {
// 		c.Validation.Keep()
// 		c.FlashParams()
// 		return c.Redirect(routes.Auth.Register())
// 	}

// 	user.HashedPassword, _ = bcrypt.GenerateFromPassword(
// 		[]byte(user.Password), bcrypt.DefaultCost)
// 	err := c.Txn.Insert(&user)
// 	if err != nil {
// 		panic(err)
// 	}

// 	c.Session["user"] = user.Username
// 	c.Flash.Success("Welcome, " + user.Name)
// 	return c.Redirect(routes.Auth.Index())
// }

// func (c Application) Login(username, password string, remember bool) revel.Result {
// 	user := c.getUser(username)
// 	if user != nil {
// 		err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
// 		if err == nil {
// 			c.Session["user"] = username
// 			if remember {
// 				c.Session.SetDefaultExpiration()
// 			} else {
// 				c.Session.SetNoExpiration()
// 			}
// 			c.Flash.Success("Welcome, " + username)
// 			return c.Redirect(routes.Auth.Index())
// 		}
// 	}

// 	c.Flash.Out["username"] = username
// 	c.Flash.Error("Login failed")
// 	return c.Redirect(routes.Auth.Index())
// }

// func (c Application) Logout() revel.Result {
// 	for k := range c.Session {
// 		delete(c.Session, k)
// 	}
// 	return c.Redirect(routes.Auth.Index())
// }