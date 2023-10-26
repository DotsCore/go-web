package module

import (
	"github.com/DotsCore/go-web-framework/register"
	"github.com/DotsCore/go-web/service"
)

var MainModule = register.DIModule{
	Name: "auth",
	Provides: []interface{}{
		service.ConnectDB,
		service.ConnectRedis,
		service.ConnectElastic,
		service.ConnectMongo,
	},
}
