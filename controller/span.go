package controller

import (
	"context"

	"taego/mconst"
)

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
