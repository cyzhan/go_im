package snowflake

import (
	"os"
	"strconv"

	"github.com/zlabwork/snowflake"
)

var Node *snowflake.Node

func NewIDGenerator() {
	nodeEnv, err := strconv.ParseInt(os.Getenv("SNOWFLAKE_NODE"), 10, 64)
	if err != nil {
		panic(err)
	}

	node, err := snowflake.NewNode(nodeEnv)
	if err != nil {
		panic(err)
	}

	Node = node
}

func GenerateInt64() int64 {
	return Node.Generate().Int64()
}
