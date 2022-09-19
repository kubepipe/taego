package controller

import (
	"context"

	"taego/lib/config"
	"taego/lib/merrors"
	"taego/mconst"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	token := c.Request.Header.Get("system-token")
	if token != config.Config.UString("system.token", "") {
		fail(c, merrors.Get(merrors.ERROR_UNAUTHORIZED))
		return
	}

	setUser(GetSpan(c), &mconst.Userinfo{
		Name: "admin",
		Erp:  "admin",
	})
}

const spanUserKey = "muser"

func setUser(ctx context.Context, userinfo *mconst.Userinfo) {
	ctx = context.WithValue(ctx, spanUserKey, userinfo)
}

func getUser(ctx context.Context) *mconst.Userinfo {
	if v := ctx.Value(spanUserKey); v != nil {
		if userinfo, ok := v.(*mconst.Userinfo); ok && userinfo != nil {
			return userinfo
		}
	}
	return &mconst.Userinfo{}
}
