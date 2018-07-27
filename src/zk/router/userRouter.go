package userHandler

import (
	"zk/util"
	"zk/mgosess"
	"zk/model"
	"gopkg.in/mgo.v2/bson"
	"zk/route"
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
//
//func GetOnlineUser(w http.ResponseWriter, r *http.Request) {
//	onlineUser:=httpsession.HttpSessions.SessionStart(w,r).Get("onlineUser")
//	util.Response{
//		Code:1,
//		Msg:"获取在线用户成功",
//		Data:map[string]interface{}{"user":onlineUser},
//	}.Response2(w)
//}
//func Logout(w http.ResponseWriter, r *http.Request) {
//	httpsession.HttpSessions.SessionDestroy(w,r)
//	util.Response{
//		Code:1,
//		Msg:"注销成功",
//	}.Response2(w)
//}
//
//func RegistHandler(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()
//	log.Println("====> path is : ",r.URL.Path," ====>")
//	log.Println(r.Form)
//
//	userMap := r.Form
//
//
//	session:=mgosess.OpenSession()
//	defer session.Close()
//	c := session.DB(mgosess.DB).C(model.UserCol)
//
//	//exits username ?
//	query := c.Find(bson.M{"username": userMap.Get("username")})
//	n, err := query.Count()
//	util.PrintErr(err)
//	if n!=0 {
//		util.Response{
//			Code:1,
//			Msg:"注册失败，姓名已存在",
//			Data:map[string]interface{}{"user":userMap},
//		}.Response2(w)
//		return
//	}
//
//
//	//insert new user
//	err = c.Insert(&model.User{
//		Username:userMap.Get("username"),
//		Password:userMap.Get("password"),
//	})
//	util.PrintErr(err)
//
//	//response
//
//	util.Response{
//		Code:1,
//		Msg:"注册成功",
//		Data:map[string]interface{}{"user":userMap},
//	}.Response2(w)
//
//}