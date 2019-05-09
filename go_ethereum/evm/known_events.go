package evm

import "strings"

var KnownEvents map[string]string

func init() {
	KnownEvents = make(map[string]string, 1)
	KnownEvents[`Transfer(address,address,uint256)`] = "ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	KnownEvents[`Approval(address,address,uint256)`] = "8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"
	KnownEvents["Pregnant(address,uint256,uint256,uint256)"] = `241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80`
	KnownEvents["Birth(address,uint256,uint256,uint256,uint256)"] = `0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5`
	KnownEvents[`ApprovalForAll(address,address,bool)`] = `17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31`
}

func KnownEvent(event string) string {
	if v, ok := KnownEvents[event]; ok {
		return strings.ToLower(v)
	}
	return ""
}
