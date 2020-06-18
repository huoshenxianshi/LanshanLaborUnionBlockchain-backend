package main

import (
	_ "RizhaoLanshanLabourUnion/docs"
	"RizhaoLanshanLabourUnion/routers"
	"RizhaoLanshanLabourUnion/security"
	"RizhaoLanshanLabourUnion/security/jwt"
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"os"
	"time"
)

// @title 岚山区劳动争议调解区块链平台
// @version 0.0.2
// @description 岚山区劳动争议调解区块链平台，后台采用golang开发，使用gin + gorm + gorbac + jwt开发
// @termsOfService http://hafrans.com
// @contact.name Chuuka Ro (Hafrans)
// @contact.url  http://hafrans.com/support
// @contact.email lvzh@hafrans.com

func main(){

	// init all components
	utils.InitSettings()
	dao.InitDB()
	dao.TryInitializeTables()
	security.InitRBAC()
	jwt.InitJwt()


	// initialize api services
	port, ok:= os.LookupEnv("PORT")
	if !ok{
		fmt.Println("Using Default Port 8088. You can set environment value \"PORT\" to change the port")
		port = "8088"
	}

	router := gin.New()

	router.MaxMultipartMemory = 20 << 20 // 20MiB

	gin.DisableConsoleColor()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// register routers
	routers.InitRouter(router)

	/*
	 * Index
	 */
	router.GET("/", func(context *gin.Context) {
		context.String(200,"BUPT HAFRANS SERVER")
	})


	/*
	 * server status ping
	 */
	router.GET("/ping", pingHandler)


	/*
	 * remain for test
	 */

	router.GET("/test", func(context *gin.Context) {
		pendingCaptcha := utils.CreateCaptcha("hello")
		context.JSON(200,pendingCaptcha)
	})


	/*
	 * inject swagger
	 */
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	/*
	 * run server
	 */
	router.Run(":"+port)

}



// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func pingHandler(ctx *gin.Context){
	ctx.JSON(200,gin.H{
		"status":0,
		"message":"pong",
		"data":map[string]interface{}{
			"timestamp":time.Now().Unix(),
			"datetime": time.Now().Format("2006-01-02 15:04:05"),
		},
	})
}
