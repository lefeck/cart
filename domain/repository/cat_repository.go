package repository

import (
	"github.com/asveg/cart/domain/model"
	"github.com/jinzhu/gorm"
	"errors"
)

type ICartRepository interface {
	InitTable() error
	CreateCart(cart *model.Cart) (cartid int64, err error)
	FindCartByID(cartid int64) (cart *model.Cart,err error)
	DeleteCartByID(cartid int64)error
	UpdateCart(cart *model.Cart) error
	FindAll(UserID int64) (cartAll []model.Cart, err error)

	CleanCart(int64) error
	Incr(int64, int64) error
	Decr(int64, int64) error
}

type CartRepository struct {
	mysql *gorm.DB
}

func NewCartRepository(mysql *gorm.DB)ICartRepository  {
	return &CartRepository{mysql: mysql}
}

func (c *CartRepository)InitTable() error  {
	return c.mysql.CreateTable(&model.Cart{}).Error
}

//创建购物车
func (c *CartRepository) CreateCart(cart *model.Cart) (cartid int64, err error){
	db :=c.mysql.FirstOrCreate(cart,model.Cart{ProductID: cart.ProductID,SizeID: cart.SizeID,UserID: cart.UserID,})
	if db.Error !=nil {
		return 0,db.Error
	}
	if db.RowsAffected == 0 {
		return 0, errors.New("购物车添加失败")
	}
	return cart.ID,nil
}
//查找ID
func (c *CartRepository) FindCartByID(cartid int64) (cart *model.Cart, err error) {
	Cart := &model.Cart{}
	return cart,c.mysql.First(Cart,cartid).Error
}
//删除
func (c *CartRepository) DeleteCartByID(cartid int64) error {
	return  c.mysql.Where("id=?",cartid).Delete(&model.Cart{}).Error
}
//更新
func (c *CartRepository) UpdateCart(cart *model.Cart) error {
	return c.mysql.Model(cart).Update(cart).Error
}
//查找所有
func (c *CartRepository) FindAll(UserID int64) (cartAll []model.Cart, err error) {
	return cartAll,c.mysql.Where("user_id=?",UserID).Find(&cartAll).Error
}
// 清空购物车，删除购物车中user_id
func (c *CartRepository) CleanCart(userID int64) error {
	return c.mysql.Where("user_id=?",userID).Delete(&model.Cart{}).Error
}
//添加商品数量
func (c *CartRepository) Incr(cartID int64, num int64) error {
	cart:=&model.Cart{ID: cartID}
	return c.mysql.Model(cart).UpdateColumn("num",gorm.Expr("num + ?",num)).Error
}
//减少商品数量
func (c *CartRepository) Decr(cartID int64, num int64) error {
	cart:=&model.Cart{ID: cartID}
	db :=c.mysql.Model(cart).Where("num >= ?",num).UpdateColumn("num",gorm.Expr("num - ?",num))
	if db.Error !=nil {
		return db.Error
	}
	if db.RowsAffected ==0 {
		return errors.New("减少失败")
	}
	return nil
}

