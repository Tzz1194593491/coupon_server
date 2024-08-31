package redis

import (
	"context"
	"strings"
)

func GetMasterCount() int {
	ctx := context.Background()
	nodes := RDB.ClusterNodes(ctx)
	split := strings.Split(nodes.Val(), "\n")
	master := 0
	for _, value := range split {
		i := strings.Index(value, "master")
		if i != -1 {
			master++
		}
	}
	return master
}
