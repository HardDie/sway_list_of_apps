package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sort"

	"github.com/joshuarubin/go-sway"
	_ "github.com/joshuarubin/go-sway"
)

type App struct {
	Name  string  `json:"name"`
	Class string  `json:"class"`
	AppId *string `json:"appId"`
}

type Workspace struct {
	Name string `json:"name"`
	Apps []App  `json:"apps"`
}

func TreeRecursion(nodes []*sway.Node) (appsInfo []App) {
	for _, app := range nodes {
		a := App{
			Name:  app.Name,
			AppId: app.AppID,
		}
		if app.WindowProperties != nil {
			a.Class = app.WindowProperties.Class
		}

		if len(a.Name) > 0 || a.AppId != nil || len(a.Class) > 0 {
			appsInfo = append(appsInfo, a)
		}

		if len(app.Nodes) > 0 {
			appsInfo = append(appsInfo, TreeRecursion(app.Nodes)...)
		}
	}
	return
}

func main() {
	ctx := context.Background()
	client, err := sway.New(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	tree, err := client.GetTree(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	var items []Workspace

	for _, display := range tree.Nodes {
		for _, workspace := range display.Nodes {
			items = append(items, Workspace{
				Name: workspace.Name,
				Apps: TreeRecursion(workspace.Nodes),
			})
		}
	}

	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Name < items[j].Name
	})

	data, _ := json.MarshalIndent(items, "", "	")
	fmt.Println(string(data))
}
