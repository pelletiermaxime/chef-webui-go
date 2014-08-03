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

func (c Databags) Index() revel.Result {
	ConnectChef()
	databags, err := ChefConnection.GetData()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return c.Render(databags)
}

func (c Databags) Show(id string) revel.Result {
	databags := getDatabagCached(id)
	// fmt.Println(databags)
	return c.Render(databags, id)
}

func (c Databags) ShowSub(id string, sub string) revel.Result {
	databagName := id + "/" + sub
	databags := getDatabagCached(databagName)
	return c.Render(databags, id)
}

func getDatabagCached(databagName string) (databags map[string]string) {
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
	return databags
}
