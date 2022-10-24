package logic

import "bullbeckt/dao/mysql"

/*
 */
func CommunityList() (interface{}, error) {
	//查询数据库
	return mysql.CommunityList()
}
