package userHandler

import (
	"zk/util"
	"zk/mgosess"
	"zk/model"
	"gopkg.in/mgo.v2/bson"
	"zk/route"
	"zk/httpsession"
)

func LoginRouter(response route.RouteResponse, request route.RouteRequest) {

	userMap:=request.Params

	session := mgosess.OpenSession()

	defer session.Close()

	query:=session.DB(mgosess.DB).C(model.UserCol).Find(bson.M{"username":userMap["username"],"password":userMap["password"]})
	count,err:=query.Count()
	util.PrintErr(err)
	if count<=0 {

		response.WriteJsonStr(util.Response{
			Code:-1,
			Msg:"登录失败，账户或者密码错误",
		})
	}else{
		user:=query.Limit(1)
		//httpsession.HttpSessions.SessionStart(w, r).Set("onlineUser",user)

		response.WriteJsonStr(util.Response{
			Code:1,
			Msg:"登录成功",
			Data:map[string]interface{}{"user":user},
		})

	}


}

func GetOnlineUserRouter(response route.RouteResponse, request route.RouteRequest) {
	onlineUser:=httpsession.HttpSessions.SessionStart(response.ResponseWriter,request.Request).Get("onlineUser")
	response.WriteJsonStr(util.Response{
		Code:1,
		Msg:"获取在线用户成功",
		Data:map[string]interface{}{"user":onlineUser},
	})
}
func LogoutRouter(response route.RouteResponse, request route.RouteRequest) {
	httpsession.HttpSessions.SessionDestroy(response.ResponseWriter,request.Request)
	response.WriteJsonStr(util.Response{
		Code:1,
		Msg:"注销成功",
	})
}

func RegistRouter(response route.RouteResponse, request route.RouteRequest) {

	userMap := request.Params


	session:=mgosess.OpenSession()
	defer session.Close()
	c := session.DB(mgosess.DB).C(model.UserCol)

	//exits username ?
	query := c.Find(bson.M{"username": userMap["username"]})
	n, err := query.Count()
	util.PrintErr(err)
	if n!=0 {
		response.WriteJsonStr(util.Response{
			Code:1,
			Msg:"注册失败，姓名已存在",
			Data:map[string]interface{}{"user":userMap},
		})
		return
	}


	//insert new user
	err = c.Insert(&model.User{
		Username:userMap["username"].(string),
		Password:userMap["password"].(string),
	})
	util.PrintErr(err)

	//response

	response.WriteJsonStr(util.Response{
		Code:1,
		Msg:"注册成功",
		Data:map[string]interface{}{"user":userMap},
	})

}