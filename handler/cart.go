package handler

import (
	"github.com/wangjinh/cart/domain/model"
	"github.com/wangjinh/cart/domain/service"
	. "github.com/wangjinh/cart/proto/cart"
	"context"
	"github.com/wangjinh/cart/common"
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

func (c *Cart)AddCart(ctx context.Context, request *CartInfo, response *ResponseAdd) (err error ) {
	cart := &model.Cart{}
	common.SwapTo(request,cart)
	response.CartId,err = c.CartDataService.AddCart(cart)
	return err
}

func (c *Cart)CleanCart(ctx context.Context, request *Clean, response *Response) error  {
	if err := c.CartDataService.CleanCart(request.UserId); err !=nil {
		return err
	}
	response.Msg = "shopping cart clean success"
	return nil
}

func (c *Cart)Incr(ctx context.Context, request *Item,  response *Response) error  {
	if err := c.CartDataService.IncrNum(request.Id,request.ChangeNum); err !=nil {
		return err
	}
	response.Msg = "shopping cart increase success"
	return nil
}
func (c *Cart) Decr(ctx context.Context, request *Item, response *Response) error {
	if err := c.CartDataService.DecrNum(request.Id,request.ChangeNum); err !=nil {
		return err
	}
	response.Msg = "shopping cart decrease success"
	return nil
}

func (c *Cart) DeleteItemByID(ctx context.Context, request *CartID, response *Response) error  {
	if err := c.CartDataService.DeleteCart(request.Id); err !=nil {
		return err
	}
	response.Msg = "shopping cart delete successful"
	return nil
}

func (c *Cart) GetAll(ctx context.Context, request *CartFindAll, response *CartAll) error  {
	cartAll, err := c.CartDataService.FindAll(request.UserId)
	if err !=nil {
		return err
	}
	for _, v :=range cartAll {
		cart := &CartInfo{}
		if err := common.SwapTo(v,cart); err !=nil {
			return err
		}
		response.CartInfo = append(response.CartInfo,cart)
	}
	return nil
}