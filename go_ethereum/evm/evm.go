package evm

import (
	"encoding/hex"
	"golang.org/x/crypto/sha3"
	"strings"
)

const (
	None           = 0
	ErcStandard20  = 1
	ErcSimple20    = 2
	ErcStandard721 = 3
	ErcSimple721   = 4
)

var In165Interface = [4]byte{0x01, 0xff, 0xc9, 0xa7}
var InNonInterface = [4]byte{0xff, 0xff, 0xff, 0xff}
var In721Interface = [4]byte{0x80, 0xac, 0x58, 0xcd}
var In721Receiver = [4]byte{0x15, 0x0b, 0x7a, 0x02}
var In721Metadata = [4]byte{0x5b, 0x5e, 0x13, 0x9f}
var In721Enumerable = [4]byte{0x78, 0x0e, 0x9d, 0x63}

func CheckTokenInterface(code string) uint8 {
	typ := None
	if CheckStandardERC20Interface(code) {
		typ = ErcStandard20
	} else if CheckSimpleERC20Interface(code) {
		typ = ErcSimple20
	}

	if CheckStandardERC721Interface(code) {
		typ = ErcStandard721
	} else if CheckSimpleERC721Interface(code) {
		typ = ErcSimple721
	}
	return uint8(typ)
}

func CheckSimpleERC20Interface(code string) bool {
	if strings.Contains(code, KnowMethods["balanceOf(address)"]) &&
		strings.Contains(code, KnowMethods["transfer(address,uint256)"]) &&
		strings.Contains(code, KnowMethods["transferFrom(address,address,uint256)"]) &&
		strings.Contains(code, KnowMethods["approve(address,uint256)"]) &&
		strings.Contains(code, KnowMethods["allowance(address,address)"]) &&
		strings.Contains(code, KnownEvents["Transfer(address,address,uint256)"]) &&
		strings.Contains(code, KnownEvents["Approval(address,address,uint256)"]) {
		return true
	}
	return false
}

func CheckStandardERC20Interface(code string) bool {
	if strings.Contains(code, KnowMethods["totalSupply()"]) &&
		strings.Contains(code, KnowMethods["balanceOf(address)"]) &&
		strings.Contains(code, KnowMethods["transfer(address,uint256)"]) &&
		strings.Contains(code, KnowMethods["transferFrom(address,address,uint256)"]) &&
		strings.Contains(code, KnowMethods["approve(address,uint256)"]) &&
		strings.Contains(code, KnowMethods["allowance(address,address)"]) &&
		strings.Contains(code, KnownEvents["Transfer(address,address,uint256)"]) &&
		strings.Contains(code, KnownEvents["Approval(address,address,uint256)"]) {
		return true
	}
	return false
}

func CheckSimpleERC721Interface(code string) bool {
	if strings.Contains(code, KnownEvents["Transfer(address,address,uint256)"]) &&
		strings.Contains(code, KnownEvents["Approval(address,address,uint256)"]) &&

		strings.Contains(code, KnowMethods["balanceOf(address)"]) &&
		strings.Contains(code, KnowMethods["ownerOf(uint256)"]) &&
		strings.Contains(code, KnowMethods["transferFrom(address,address,uint256)"]) &&
		strings.Contains(code, KnowMethods["approve(address,uint256)"]) {
		return true
	}
	return false
}

func CheckStandardERC721Interface(code string) bool {
	if strings.Contains(code, KnownEvents["Transfer(address,address,uint256)"]) &&
		strings.Contains(code, KnownEvents["Approval(address,address,uint256)"]) &&
		strings.Contains(code, KnownEvents["ApprovalForAll(address,address,bool)"]) &&

		strings.Contains(code, KnowMethods["balanceOf(address)"]) &&
		strings.Contains(code, KnowMethods["ownerOf(uint256)"]) &&
		strings.Contains(code, KnowMethods["safeTransferFrom(address,address,uint256,bytes)"]) &&
		strings.Contains(code, KnowMethods["safeTransferFrom(address,address,uint256)"]) &&
		strings.Contains(code, KnowMethods["transferFrom(address,address,uint256)"]) &&
		strings.Contains(code, KnowMethods["approve(address,uint256)"]) &&
		strings.Contains(code, KnowMethods["setApprovalForAll(address,bool)"]) &&
		strings.Contains(code, KnowMethods["getApproved(uint256)"]) &&
		strings.Contains(code, KnowMethods["isApprovedForAll(address,address)"]) &&
		strings.Contains(code, KnowMethods["supportsInterface(bytes4)"]) {
		return true
	}
	return false
}

func EIP165Method(method string) string {
	hash := sha3.NewLegacyKeccak256()
	_, _ = hash.Write([]byte(method))
	return hex.EncodeToString(hash.Sum(nil)[0:4])
}

func EIP165Event(event string) string {
	hash := sha3.NewLegacyKeccak256()
	_, _ = hash.Write([]byte(event))
	return hex.EncodeToString(hash.Sum(nil))
}
