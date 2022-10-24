package main

import (
	"bullbeckt/dao/mysql"
	"bullbeckt/setting"
	"fmt"
)

func main() {
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Println("mysql初始化失败...")
		return
	}

}
