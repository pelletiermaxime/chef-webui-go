package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"os"
	"time"
)

type Databags struct {
	*revel.Controller
}

func (c Databags) Show(id string) revel.Result {
	var databags map[string]string
	if err := cache.Get("databag_"+id, &databags); err != nil {
		ConnectChef()
		databag, ok, errgetdatabag := ChefConnection.GetDataByName(id)
		databags = databag
		if errgetdatabag != nil {
			fmt.Println("Error:", errgetdatabag)
			os.Exit(1)
		}
		if !ok {
			fmt.Println("Couldn't find that databag!")
		}
		cache.Set("databag_"+id, databags, 30*time.Minute)
	}
	// fmt.Println(databags)
	return c.Render(databags, id)
}

func (c Databags) ShowSub(id string, sub string) revel.Result {
	var databags map[string]string
	databagName := id + "/" + sub
	if err := cache.Get("databag_"+databagName, &databags); err != nil {
		ConnectChef()
		databag, ok, err := ChefConnection.GetDataByName(databagName)
		databags = databag
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		if !ok {
			fmt.Println("Couldn't find that databag!")
		}
		cache.Set("databag_"+databagName, databags, 30*time.Minute)
	}
	return c.Render(databags, id)
}
