package controllers

import (
	"github.com/goravel/framework/contracts/http"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	_ "evote-be/docs"
)

/*********************************
1. Install swag
document: https://github.com/swaggo/http-swagger

go install github.com/swaggo/swag/cmd/swag@latest

2. Install http-swagger
go get -u github.com/swaggo/http-swagger

3. Optimize the document of endpoint: `app/http/controllers/swagger_controller.go`

4. Add route to `/route/web.go`

5. Init document
swag init

6. Run Server
air

7. Visit: http://localhost:3000/swagger/
 ********************************/

type SwaggerController struct {
	// Dependent services
}

func NewSwaggerController() *SwaggerController {
	return &SwaggerController{
		// Inject services
	}
}

func (r *SwaggerController) Index(ctx http.Context) http.Response {
	handler := httpSwagger.Handler()
	handler(ctx.Response().Writer(), ctx.Request().Origin())

	return nil
}
