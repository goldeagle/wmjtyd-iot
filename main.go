// @title IoT平台API文档
// @version 1.0
// @description 物联网设备管理平台API文档
// @termsOfService http://swagger.io/terms/
// @contact.name API支持
// @contact.email support@wmjtyd.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

package main

import (
	"log"

	"wmjtyd-iot/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
