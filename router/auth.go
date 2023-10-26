package router

import (
	"github.com/DotsCore/go-web-framework/register"
	"github.com/DotsCore/go-web/app/http/validation"
)

var AuthRouter = register.HTTPRouter{
	Route: []register.Route{
		{
			Name:        "login",
			Path:        "/login",
			Action:      "AuthController@JWTAuthentication",
			Method:      "POST",
			Validation:  &validation.Credentials{},
			Description: "Perform login",
			Middleware:  []register.Middleware{},
		},
	},
}
