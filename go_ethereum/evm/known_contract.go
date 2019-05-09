package evm

import "strings"

type Contract string

const (
	CryptoKitties = "KnownContracts"
	HyperDragons  = "HyperDragons"
	GodsUnchained = "GodsUnchained"
)

var KnownContracts map[Contract]string

func init() {
	KnownContracts = make(map[Contract]string, 1)
	KnownContracts[CryptoKitties] = "0x06012c8cf97BEaD5deAe237070F9587f8E7A266d"
	KnownContracts[HyperDragons] = "0x7fDcD2a1E52F10C28cB7732f46393e297eCaDDa1"
	KnownContracts[GodsUnchained] = "0x6EbeAf8e8E946F0716E6533A6f2cefc83f60e8Ab"

}

func KnownContract(name Contract) string {
	if v, ok := KnownContracts[name]; ok {
		return strings.ToLower(v)
	}
	return ""
}
