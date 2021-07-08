package handler

import (
	"github.com/asveg/cart/domain/model"
	"github.com/asveg/cart/domain/service"
	cart "github.com/asveg/cart/proto/cart"
	"context"
	"github.com/asveg/cart/common"
)

type Cart struct{
	CartDataService service.ICartDataService
}

/*
type CartHandler interface {
	//添加购物车添加，添加信息，返回id和msg
	AddCart(context.Context, *CartInfo, *ResponseAdd) error
	//清除购物车，通过ID清除
	CleanCart(context.Context, *Clean, *Response) error
	//增加商品
	Incr(context.Context, *Item, *Response) error
	//减少商品
	Decr(context.Context, *Item, *Response) error
	//删除商品根据购物车id
	DeleteItemByID(context.Context, *CartID, *Response) error
	//获取所有购物车商品，通过用户id，返回所有商品
	GetAll(context.Context, *CartFindAll, *CartAll) error
}
 */

func (c *Cart)AddCart(ctx context.Context, request *cart.CartInfo, response *cart.ResponseAdd) (err error ) {
	cart := &model.Cart{}
	Carterr := common.SwapTo(request,cart)
	if Carterr !=nil {
		return err
	}
	response.CartId , err = c.CartDataService.AddCart(cart)
	return nil
}
//通过请求的userid，清空购物车。
func (c *Cart)CleanCart(ctx context.Context, request *cart.Clean, response *cart.Response) error  {
	if err :=c.CartDataService.CleanCart(request.UserId); err !=nil {
		return err
	}
	response.Msg="shopping cart clean successfully"
	return nil
}

func (c *Cart)Incr(ctx context.Context, request *cart.Item,  response *cart.Response) error {
	if err := c.CartDataService.IncrNum(request.Id, request.ChangeNum); err != nil{
		return err
    }
    response.Msg="shopping num increase success"
	return nil
}
//通过请求的cartid和product，减少product数量
func (c *Cart) Decr(ctx context.Context, request *cart.Item, response *cart.Response) error {
	if err := c.CartDataService.DecrNum(request.Id,request.ChangeNum); err !=nil {
		return err
	}
	response.Msg="shopping num decrease success"
	return nil
}
//删除product通过购物车id
func (c *Cart) DeleteItemByID(ctx context.Context, request *cart.CartID, response *cart.Response) error  {
	if err := c.CartDataService.DeleteCart(request.Id); err !=nil {
		return err
	}
	response.Msg="shopping delete product successfully"
	return nil
}
//获取所有购物车商品，通过用户id，返回所有商品
func (c *Cart) GetAll(ctx context.Context, request *cart.CartFindAll, response *cart.CartAll) error  {
	cartAll, err := c.CartDataService.FindAll(request.UserId)
	if err !=nil {
		return err
	}
	for _, v :=range cartAll {
		cart := &cart.CartInfo{}
		if err := common.SwapTo(v,cart); err !=nil {
			return err
		}
		response.CartInfo = append(response.CartInfo,cart)
	}
	return nil
}