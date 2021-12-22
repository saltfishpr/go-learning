// @description: 雪花算法生成ID
// @file: snowflake.go
// @date: 2021/12/8

package utils

import (
	"sync"

	"learning/config"
	"learning/logger"

	"github.com/bwmarrin/snowflake"
)

var node struct {
	instance *snowflake.Node
	once     sync.Once
}

func NewNode() *snowflake.Node {
	node.once.Do(
		func() {
			n, err := snowflake.NewNode(config.GetInt64("node_id"))
			if err != nil {
				logger.Fatal(err)
			}
			node.instance = n
		},
	)
	return node.instance
}
