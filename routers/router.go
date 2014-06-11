package routers

import (
	"exbook/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/page/:page:int", &controllers.MainController{})
	//登录
	beego.Router("/login", &controllers.LoginController{}, "get:Login;post:DoLogin")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/login/check_login", &controllers.LoginController{}, "get:CheckLogin")
	//注册
	beego.Router("/register", &controllers.RegisterController{}, "get:Reg;post:DoReg")
	beego.Router("/register/checkuname", &controllers.RegisterController{}, "post:Checkuname")

	beego.Router("/trade/:id:int", &controllers.TradeController{}, "get:Show")
	beego.Router("/addfav/trade", &controllers.AddFavController{}, "post:AddTrade")
	beego.Router("/addfav/book", &controllers.AddFavController{}, "post:AddBook")
	beego.Router("/about", &controllers.AboutController{})
	beego.Router("/bookshelf", &controllers.BookShelfController{})
	beego.Router("/mytrade", &controllers.MyTradeController{})
	beego.Router("/myoffer", &controllers.MyOfferController{})
	beego.Router("/mybookmark", &controllers.BookMarkController{})
	beego.Router("/addbook", &controllers.BookController{})
	beego.Router("/myprofile", &controllers.MyProfileController{})
	beego.Router("/myprofile/update", &controllers.MyProfileController{}, "post:UpdateProfile")
	beego.Router("/myprofile/changepwd", &controllers.MyProfileController{}, "post:ChangePwd")
	beego.Router("/addtrade", &controllers.AddController{}, "get:AddTrade")
	beego.Router("/search", &controllers.SearchController{})
	beego.Router("/trade/search", &controllers.SearchController{}, "get:SearchT")
	beego.Router("/trade/search/page/:page:int", &controllers.SearchController{}, "get:SearchT")
	beego.Router("/bookinv/rmbook/:invid:int", &controllers.AddFavController{}, "get:RmBook")
	beego.Router("/bookmark/rmbkt/:tid:int", &controllers.BookMarkController{}, "get:RemoveBT")
	beego.Router("/bookmark/rmbkb/:bid:int", &controllers.BookMarkController{}, "get:RemoveBK")
	beego.Router("/trade/remove/:tid:int", &controllers.MyTradeController{}, "get:RemoveTrade")
	beego.Router("/trade/popup", &controllers.MyTradeController{}, "post:PopTrade")
	beego.Router("/book/getbook", &controllers.GetBookController{}, "post:Getbinfo")
	beego.Router("/book/addinv", &controllers.GetBookController{}, "post:AddInv")
	beego.Router("/book/searchbook", &controllers.GetBookController{}, "post:Searchbinfo")
	beego.Router("/book/addtoleft", &controllers.AddController{}, "post:AddToLeft")
	beego.Router("/book/addtoright", &controllers.AddController{}, "post:AddToRight")
	beego.Router("/trade/addtrade", &controllers.MyTradeController{}, "post:PostTrade")
	beego.Router("/mytrade/:id:int", &controllers.MyTradeController{}, "get:Show")
	beego.Router("/inventory/read", &controllers.BookShelfController{}, "post:GetInv")
	beego.Router("/offer/addoffer", &controllers.AddController{}, "post:AddOffer")
	beego.Router("/offer/refuseoffer/:tid:int/:oid:int", &controllers.MyTradeController{}, "get:RefuseOffer")
	beego.Router("/offer/acceptoffer/:tid:int/:oid:int", &controllers.MyTradeController{}, "get:AcceptOffer")
	beego.Router("/offer/canceloffer/:tid:int/:oid:int", &controllers.MyTradeController{}, "get:CancelOffer")
	beego.Router("/offer/postreply", &controllers.MyTradeController{}, "post:PostReply")
	beego.Router("/offer/delreply/:tid:int/:rpid:int", &controllers.MyTradeController{}, "get:DelReply")
	beego.Router("/myoffer/:id:int", &controllers.MyOfferController{}, "get:Show")
	beego.Router("/offer/deleteoffer/:tid:int/:oid:int", &controllers.MyOfferController{}, "get:DeleteOffer")
	beego.Router("/offer/deletealloffer/:tid:int", &controllers.MyOfferController{}, "get:DeleteAllOffer")
	beego.Router("/offer/get_traderinfo", &controllers.MyOfferController{}, "post:GetTraderInfo")
	beego.Router("/offer/delpostreply/:tid:int/:rpid:int", &controllers.MyOfferController{}, "get:DeletePostReply")
}
