package main

import (
	"github.com/wangjinh/cart/common"
	"github.com/wangjinh/cart/domain/repository"
	"github.com/wangjinh/cart/handler"
	"github.com/wangjinh/cart/proto/cart"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	service2 "github.com/wangjinh/cart/domain/service"
	"strconv"
)
var QPS = 100
func main() {
	//配置中心
	consulConfig, err := common.GetConsulConfig("192.168.10.168", 8500, "/micro/config")
	if err !=nil {
		log.Error(err)
	}
	//注册中心
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string {
			"192.168.10.168:8500",
		}
	})

	//链路追踪
	t,io,err := common.NewTracker("go.micro.service.cart","localhost:6343")
	if err !=nil {
		log.Error(err)
	}
	defer io.Close()
	//设置链路追踪
	opentracing.SetGlobalTracer(t)

	//数据库连接
	mysqlInfo := common.GetMysqlFromConsul(consulConfig,"mysql")
	mysqlInfoPort := strconv.FormatInt(mysqlInfo.Port,10)
	db,err := gorm.Open("mysql",mysqlInfo.User+":"+mysqlInfo.Pwd+"@("+mysqlInfo.Host+":"+mysqlInfoPort+")/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err !=nil {
		log.Error(err)
	}
	defer db.Close()
	db.SingularTable(true)

	//initialise database
	err = repository.NewCartRepository(db).InitTable()
	if err !=nil {
		log.Error(err)
	}

	//new 一个microservice 实例
	service := micro.NewService(
		micro.Name("go.micro.service.cart"),
		micro.Version("latest"),
		micro.Registry(consul),
		micro.Address("0.0.0.0:8708"),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
		)

	// Initialise service
	service.Init()

	//连接数据库
	cartDataService := service2.NewCartDateService(repository.NewCartRepository(db))

	// Register Handler， 将微服务的接口操作注册到处理器中
	go_micro_service_cart.RegisterCartHandler(service.Server(), &handler.Cart{
		cartDataService,
	})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
