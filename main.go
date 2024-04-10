package main

import (
	"encoding/base64"
	"fmt"

	"github.com/yourusername/yourproject/node"
	"github.com/yourusername/yourproject/subscription"
)

func main() {
	subURL := "http://example.com/vmess/subscription"

	// 获取订阅
	encodedSub, err := subscription.FetchSubscription(subURL)
	if err != nil {
		fmt.Println("Error fetching subscription:", err)
		return
	}

	// 解码订阅内容
	decodedSub, err := subscription.DecodeBase64(encodedSub)
	if err != nil {
		fmt.Println("Error decoding subscription:", err)
		return
	}

	nodeInfos := []string{"node1Info", "node2Info"}

	var validNodes []string
	for _, nodeInfo := range nodeInfos {
		if node.TestNode(nodeInfo) && node.VerifyNode(nodeInfo) {
			validNodes = append(validNodes, nodeInfo)
		}
	}

	// 编码有效节点为新的订阅内容
	newSub := base64.StdEncoding.EncodeToString([]byte(fmt.Sprint(validNodes)))
	fmt.Println("New subscription:", newSub)
}
