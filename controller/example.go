package controller

import (
	"taego/dao"
	"taego/lib/merrors"
	"taego/lib/mlog"
	"taego/lib/morm"
	"taego/service/example"

	"github.com/gin-gonic/gin"
)

func Example(c *gin.Context) {

	// db
	us := []*dao.User{}
	err := morm.GetORM(dao.User{}).Query(GetSpan(c), "select name from user").Scan(&us)

	us2 := []dao.User{}
	err = morm.GetORM(dao.User{}).Query(GetSpan(c), "select name from user").Scan(&us2)

	us3 := dao.User{}
	err = morm.GetORM(dao.User{}).Query(GetSpan(c), "select name from user").Scan(&us3)

	us4 := []string{}
	err = morm.GetORM(dao.User{}).Query(GetSpan(c), "select name from user").Scan(&us4)

	us5 := ""
	err = morm.GetORM(dao.User{}).Query(GetSpan(c), "select name from user").Scan(&us5)
	mlog.Infof("us5: %s", us5)
	if err != nil {
		fail(c, merrors.New(err))
		return
	}

	// get date from the service
	req := &example.ReqExample{}
	res, err := example.GetExampleData(GetSpan(c), req)
	if err != nil {
		fail(c, merrors.New(err))
		return
	}

	// mlog.Info does not contain trace
	GetTrace(c).Log("some other things to do")

	// response
	success(c, string(res))
}
