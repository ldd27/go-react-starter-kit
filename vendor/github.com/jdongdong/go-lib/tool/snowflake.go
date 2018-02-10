package tool

import "github.com/bwmarrin/snowflake"

var node *snowflake.Node

func init() {
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
}

//func GenNewID() (int64, error) {
//iw, err := goSnowFlake.NewIdWorker(1)
//if err != nil {
//	return 0, err
//}
//if id, err := iw.NextId(); err != nil {
//	return 0, err
//} else {
//	return id, nil
//}
//}

func NewID() int64 {
	return node.Generate().Int64()
}

func NewStrID() string {
	return node.Generate().String()
}
