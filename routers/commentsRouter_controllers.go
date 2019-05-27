package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["beego-blog/controllers:ApiController"] = append(beego.GlobalControllerRouter["beego-blog/controllers:ApiController"],
        beego.ControllerComments{
            Method: "Upload",
            Router: `/api/upload`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-blog/controllers:ApiController"] = append(beego.GlobalControllerRouter["beego-blog/controllers:ApiController"],
        beego.ControllerComments{
            Method: "Down",
            Router: `/down`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
