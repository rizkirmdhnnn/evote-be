package middleware

import (
	"errors"
	"evote-be/app/models"
	"strconv"

	"github.com/goravel/framework/auth"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func Auth() http.Middleware {
	return func(ctx http.Context) {
		token := ctx.Request().Header("Authorization", "")
		if token == "" {
			ctx.Request().Abort(http.StatusUnauthorized)
			return
		}

		payload, err := facades.Auth(ctx).Parse(token)
		if err != nil {
			if errors.Is(err, auth.ErrorTokenExpired) {
				token, err = facades.Auth(ctx).Refresh()
				if err != nil {
					// Refresh time exceeded
					ctx.Request().Abort(http.StatusUnauthorized)
					return
				}

				token = "Bearer " + token
			} else {
				ctx.Request().Abort(http.StatusUnauthorized)
				return
			}
		}

		// You can get User in DB and set it to ctx
		var user models.User
		id, err := strconv.ParseUint(payload.Key, 10, 64)
		if err != nil {
			ctx.Request().Abort(http.StatusUnauthorized)
			return
		}
		user.ID = uint(id)
		// if err := facades.Auth(ctx).User(&user); err != nil {
		// 	ctx.Request().AbortWithStatus(http.StatusUnauthorized)
		// 	return
		// }
		// ctx.WithValue("user", user)
		ctx.WithValue("user", user)

		ctx.Response().Header("Authorization", token)
		ctx.Request().Next()
	}
}
