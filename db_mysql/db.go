package db_mysql



import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/go-sql-driver/mysql"
)
var DB *sql.DB
func ConnectDB()  {



		fmt.Println("链接数据库")
		//1.读取conf中的数据
		config := beego.AppConfig
		dbDriver := config.String("db_driverName")
		dbUser := config.String("db_user")
		dbPassword := config.String("db_password")
		dbIp := config.String("db_ip")
		dbName := config.String("db_name")

		connUrl := dbUser + ":" + dbPassword + "@tcp(" + dbIp + ")/" + dbName + "?charset=utf8"
		db, err := sql.Open(dbDriver, connUrl)
		if err != nil {
			panic("数据库连接失败，请重试")
		}
		DB = db


	}
	
