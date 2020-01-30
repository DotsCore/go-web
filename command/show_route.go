package command

import (
	"github.com/olekukonko/tablewriter"
	"go.uber.org/dig"
	"ikdev/go-web/config"
	"os"
	"strings"
)

type ShowRoute struct {
	Signature   string
	Description string
}

func (c *ShowRoute) Register() {
	c.Signature = "show:route"
	c.Description = "Show active Go-Web routes"
}

// Show the current go-web routes
func (c *ShowRoute) Run(sc *dig.Container, args string) {
	var data [][]string
	routes := config.ConfigurationWeb()

	// Parse single route
	showSingleRoute(routes.Routes, &data)

	// Parse groups
	showGroupRoutes(routes.Groups, &data)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"METHOD", "ACTION", "PATH", "PREFIX", "MIDDLEWARE"})

	for _, v := range data {
		table.Append(v)
	}

	table.Render()
}

// Show single routes
func showSingleRoute(routes map[string]config.Route, data *[][]string) {
	for _, r := range routes {
		*data = append(*data, []string{r.Method, r.Action, r.Path, r.Prefix, strings.ToLower(strings.Join(r.Middleware, ", "))})
	}
}

// Show groups
func showGroupRoutes(routes map[string]config.Group, data *[][]string) {
	for _, g := range routes {
		var middleware []string
		middleware = append(middleware, g.Middleware...)
		for _, gr := range g.Routes {
			middleware = append(middleware, gr.Middleware...)
			*data = append(*data, []string{gr.Method, gr.Action, gr.Path, g.Prefix, strings.ToLower(strings.Join(middleware, ", "))})
		}
	}
}