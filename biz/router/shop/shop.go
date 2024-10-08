// Code generated by hertz generator. DO NOT EDIT.

package shop

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	shop "xzdp/biz/handler/shop"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_shop := root.Group("/shop", _shopMw()...)
		_shop.GET("/:id", append(_shopinfoMw(), shop.ShopInfo)...)
		{
			_of := _shop.Group("/of", _ofMw()...)
			_of.GET("/type", append(_shopoftypeMw(), shop.ShopOfType)...)
			_type := _of.Group("/type", _typeMw()...)
			_type.GET("/geo", append(_shopoftypegeoMw(), shop.ShopOfTypeGeo)...)
		}
	}
	{
		_shop_type := root.Group("/shop-type", _shop_typeMw()...)
		_shop_type.GET("/list", append(_shoplistMw(), shop.ShopList)...)
	}
}
