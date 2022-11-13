package application

import "context"

type applicationInterface interface {
	Startup(ctx context.Context)
	Close(ctx context.Context)
}

var binds []applicationInterface

func AddApp(app applicationInterface) {
	binds = append(binds, app)
}

func GetApps() []applicationInterface {
	return binds
}

func Startup(ctx context.Context) {
	for _, app := range binds {
		app.Startup(ctx)
	}
}

func Close(ctx context.Context) bool {
	for _, app := range binds {
		app.Close(ctx)
	}
	return true
}
