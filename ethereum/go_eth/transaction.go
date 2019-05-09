package go_eth

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/eager7/one_chain/common/utils"
	"github.com/eager7/one_chain/go_ethereum/evm"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"

	"math/big"
)

func (e *Eth) SendTransaction(private, to string, amount *big.Int, gasLimit uint64, data []byte) error {
	privateKey, err := crypto.HexToECDSA(private)
	if err != nil {
		return errors.New("send transaction private err:" + err.Error())
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	from := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := e.PendingNonceAt(from.Hex())
	if err != nil {
		return errors.New("send transaction nonce err:" + err.Error())
	}
	toAddr := common.HexToAddress(to)
	gasPrice, err := e.client.SuggestGasPrice(e.ctx)
	if err != nil {
		return errors.New("send transaction gas price err:" + err.Error())
	}
	gasBase, err := e.client.EstimateGas(e.ctx, ethereum.CallMsg{
		From:     from,
		To:       &toAddr,
		Gas:      0,
		GasPrice: gasPrice,
		Value:    amount,
		Data:     common.CopyBytes(data),
	})
	if err != nil {
		return err
	}
	tx := types.NewTransaction(nonce, common.HexToAddress(to), amount, gasBase+gasLimit, gasPrice, data)
	chainID, err := e.client.NetworkID(e.ctx)
	if err != nil {
		return errors.New("send transaction chain id err:" + err.Error())
	}
	sigTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return errors.New("send transaction sign err:" + err.Error())
	}
	w := utils.Writer{}
	_ = sigTx.EncodeRLP(&w)
	fmt.Println(hex.EncodeToString(w.Body.Bytes()))
	return e.client.SendTransaction(e.ctx, sigTx)
}

func (e *Eth) SendErcTransaction(private, contract, to string, amount *big.Int, gasLimit uint64, gasPrice *big.Int) error {
	var data []byte
	toAddr := common.HexToAddress(to)

	methodID := evm.EIP165Method("transfer(address,uint256)")
	data = append(data, methodID...)

	paddedAddress := common.LeftPadBytes(toAddr.Bytes(), 32)
	data = append(data, paddedAddress...)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	data = append(data, paddedAmount...)

	if gasLimit == 0 {
		var err error
		gasLimit, err = e.client.EstimateGas(e.ctx, ethereum.CallMsg{To: &toAddr, Data: data})
		if err != nil {
			return errors.New("send transaction estimate gas err:" + err.Error())
		}
	}

	return e.SendTransaction(private, contract, big.NewInt(0), gasLimit, data)
}

func (e *Eth) TransactionToRaw(tx *types.Transaction) string {
	return hex.EncodeToString(types.Transactions{tx}.GetRlp(0))
}

func (e *Eth) TransactionFromRaw(rawTx string) (*types.Transaction, error) {
	rawTxBytes, err := hex.DecodeString(rawTx)
	if err != nil {
		return nil, errors.New("transaction from raw err:" + err.Error())
	}
	tx := new(types.Transaction)
	if err := rlp.DecodeBytes(rawTxBytes, &tx); err != nil {
		return nil, errors.New("transaction from raw rlp err:" + err.Error())
	}
	return tx, nil
}
