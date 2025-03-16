package tsdb

import (
	"database/sql"
	"log"

	"github.com/spf13/viper"

	_ "github.com/taosdata/driver-go/v3/taosRestful"
)

var TaosDB *sql.DB

//var TaosConn *af.Connector

func InitDatabaseTaos() {
	var taosDSN = viper.GetString("Taos.User") + ":" + viper.GetString("Taos.Password") + "@http(" + viper.GetString("Taos.Host") + ":" + viper.GetString("Taos.Port") + ")/"
	db, err := sql.Open("taosRestful", taosDSN)
	if err != nil {
		log.Panic("failed to connect TDengine, err:", err)
	}
	//defer TaosDB.Close()

	TaosDB = db

	log.Println("TDengine connected")
}

/**
func InitTaosConn() {
	conn, err := af.Open(viper.GetString("Taos.Host"), viper.GetString("Taos.User"), viper.GetString("Taos.Password"), "", viper.GetInt("Taos.Port"))
	if err != nil {
		log.Panic("fail to connect TDengine conn, err:", err)
	}
	//defer conn.Close()

	TaosConn = conn

	log.Println("TDengine conn connected")
}
*/
