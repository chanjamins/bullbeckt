package snowflake

import (
	"fmt"
	sf "github.com/bwmarrin/snowflake"
	"time"
)

var node *sf.Node

// 初始化
func Init(startTime string, machineId int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		fmt.Println("时间初始化失败！")
	}
	sf.Epoch = st.Unix() / 1000000
	node, err = sf.NewNode(machineId)
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}
