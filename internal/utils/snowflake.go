// @description: 雪花算法生成ID
// @file: snowflake.go
// @date: 2021/12/8

package utils

import (
	"sync"

	"learning/internal/config"

	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

var node struct {
	instance *snowflake.Node
	once     sync.Once
}

func NewNode() *snowflake.Node {
	node.once.Do(
		func() {
			cfg := config.GetConfig()
			n, err := snowflake.NewNode(cfg.NodeID)
			if err != nil {
				zap.S().Fatal(err)
			}
			node.instance = n
		},
	)
	return node.instance
}
