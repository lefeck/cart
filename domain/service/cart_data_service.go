package service

import (
	"github.com/wangjinh/cart/domain/model"
	"github.com/wangjinh/cart/domain/repository"

)

type ICartDataService interface {
	AddCart(cart *model.Cart)(cartid int64, err error)
	DeleteCart(cartid int64) error
	UpdateCart(cart *model.Cart) error
	FindCartByID(cartid int64)(cart *model.Cart, err error)
	FindAll(userid int64) (cart []model.Cart,err error)

	CleanCart(userid int64) error
	IncrNum(cartid int64, num int64) error
	DecrNum(cartid int64, num int64) error
}

type CartDateService struct {
	DataService repository.ICartRepository
}
func NewCartDateService(cartrepository repository.ICartRepository) ICartDataService {
	return &CartDateService{cartrepository}
}


func (c *CartDateService) AddCart(cart *model.Cart) (cartid int64, err error) {
	return c.DataService.CreateCart(cart)
}

func (c *CartDateService) DeleteCart(cartid int64) error {
	return c.DataService.DeleteCartByID(cartid)
}

func (c *CartDateService) UpdateCart(cart *model.Cart) error {
	return c.DataService.UpdateCart(cart)
}

func (c *CartDateService) FindCartByID(cartid int64) (cart *model.Cart, err error) {
	return c.DataService.FindCartByID(cartid)
}

func (c *CartDateService) FindAll(userid int64) (cart []model.Cart, err error) {
	return c.DataService.FindAll(userid)
}

func (c *CartDateService) CleanCart(userid int64) error {
	return c.DataService.CleanCart(userid)
}

func (c *CartDateService) IncrNum(cartid int64, num int64) error {
	return c.DataService.Incr(cartid,num)
}

func (c *CartDateService) DecrNum(cartid int64, num int64) error {
	return c.DataService.Decr(cartid, num)
}

