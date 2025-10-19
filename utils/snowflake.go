package utils

import (
	"time"

	"github.com/bwmarrin/snowflake"
)

var (
	node *snowflake.Node
)

// InitSnowflake 初始化雪花算法节点
func InitSnowflake(nodeID int64) error {
	var err error
	node, err = snowflake.NewNode(nodeID)
	return err
}

// GenerateID 生成雪花算法ID
func GenerateID() int64 {
	if node == nil {
		// 如果未初始化，使用默认节点ID 1
		InitSnowflake(1)
	}
	return node.Generate().Int64()
}

// GenerateStringID 生成字符串格式的雪花算法ID
func GenerateStringID() string {
	return node.Generate().String()
}

// ParseID 解析雪花算法ID，获取时间戳等信息
func ParseID(id int64) snowflake.ID {
	return snowflake.ID(id)
}

// GetTimestampFromID 从雪花算法ID中提取时间戳
func GetTimestampFromID(id int64) time.Time {
	timestamp := snowflake.ID(id).Time()
	return time.Unix(timestamp/1000, (timestamp%1000)*1000000)
}

// GetNodeIDFromID 从雪花算法ID中提取节点ID
func GetNodeIDFromID(id int64) int64 {
	return int64(snowflake.ID(id).Node())
}

// GetSequenceFromID 从雪花算法ID中提取序列号
func GetSequenceFromID(id int64) int64 {
	return int64(snowflake.ID(id).Step())
}
