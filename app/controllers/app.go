package controllers

import (
	"fmt"
	"github.com/marpaia/chef-golang"
	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"os"
	"time"
)

type App struct {
	*revel.Controller
}

var ChefConnection *chef.Chef

func (c App) Index() revel.Result {
	return c.Render()
}

func ConnectChef() {
	var err error
	if err = cache.Get("ChefConnection", &ChefConnection); err != nil {
		ChefConnection, err = chef.Connect()
		// ChefConnection.SSLNoVerify = true
		// ChefConnection, err = chef.ConnectBuilder("http://chef-server.libeo.com", "4000",
		// "", "mpelletier2", "/root/.chef/mpelletier2.pem", "")
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		cache.Set("ChefConnection", ChefConnection, 30*time.Minute)
	}
}

func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("Your name is required!")
	c.Validation.MinSize(myName, 3).Message("Your name is not long enough!")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	return c.Render(myName)
}
