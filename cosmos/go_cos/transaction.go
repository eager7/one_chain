package go_cos

import (
	"errors"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	draw "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/ibc"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/eager7/elog"
	"github.com/tendermint/tendermint/rpc/core/types"
	tt "github.com/tendermint/tendermint/rpc/core/types"
	"time"
)

var log = elog.NewLogger("cosmos", elog.InfoLevel)

type Transaction struct {
	Tx      auth.StdTx
	Receipt sdk.TxResponse
}

func (c *CosCli) ParseTransaction(block *core_types.ResultBlock) (txs []Transaction, err error) {
	for _, t := range block.Block.Txs {
		if tx, err := queryTx(c.cdc, c.ctx, t.Hash(), block); err != nil {
			return nil, errors.New("ParseTransaction err:" + err.Error())
		} else {
			txs = append(txs, tx)
		}
	}
	return txs, nil
}

func (c *CosCli) HandleTransactions(txs []Transaction) error {
	for _, tx := range txs {
		for _, m := range tx.Tx.Msgs {
			switch m.Route() {
			case "bank":
				switch m.Type() {
				case "send":
					msg, ok := m.(bank.MsgSend)
					if !ok {
						return errors.New("can't convert to send")
					}
					log.Debug(tx.Receipt.Height, "send-->from:", msg.FromAddress.String(), "to:", msg.ToAddress.String(), "amount:", msg.Amount, tx.Receipt.TxHash)
				case "multisend":
					msg, ok := m.(bank.MsgMultiSend)
					if !ok {
						return errors.New("can't convert to multisend")
					}
					log.Debug(tx.Receipt.Height, "multisend-->inputs:", msg.Inputs, "outputs:", msg.Outputs, tx.Receipt.TxHash)
				default:
					return errors.New("unknown msg type:" + m.Type())
				}
			case "staking":
				switch m.Type() {
				case "delegate":
					msg, ok := m.(staking.MsgDelegate)
					if !ok {
						return errors.New("can't convert to delegate")
					}
					log.Debug(tx.Receipt.Height, "delegate-->delegator_address:", msg.DelegatorAddress, "validator_address:", msg.ValidatorAddress, "amount:", msg.Amount, tx.Receipt.TxHash)
				case "edit_validator":
					msg, ok := m.(staking.MsgEditValidator)
					if !ok {
						return errors.New("can't convert to delegate")
					}
					log.Debug(tx.Receipt.Height, "edit_validator-->moniker:", msg.Moniker, "identity:", msg.Identity, "website:", msg.Website, "details:", msg.Details, "address:",
						msg.ValidatorAddress, "commission_rate:", msg.CommissionRate, "min_self_delegation:", msg.MinSelfDelegation, tx.Receipt.TxHash)
				case "begin_unbonding":
					msg, ok := m.(types.MsgUndelegate)
					if !ok {
						return errors.New("can't convert to begin_unbonding")
					}
					log.Debug(tx.Receipt.Height, "begin_unbonding-->delegate:", msg.DelegatorAddress.String(), "validator:", msg.ValidatorAddress.String(), "amount:", msg.Amount, tx.Receipt.TxHash)
				case "begin_redelegate":
					msg, ok := m.(staking.MsgBeginRedelegate)
					if !ok {
						return errors.New("can't convert to begin_redelegate")
					}
					log.Debug(tx.Receipt.Height, "begin_redelegate-->delegator_address:", msg.DelegatorAddress, "validator_src_address:", msg.ValidatorSrcAddress,
						"validator_dst_address:", msg.ValidatorDstAddress, "amount:", msg.Amount, tx.Receipt.TxHash)
				case "create_validator":
					msg, ok := m.(staking.MsgCreateValidator)
					if !ok {
						return errors.New("can't convert to create_validator")
					}
					log.Debug(tx.Receipt.Height, "create_validator-->description:", msg.Description, "commission:", msg.Commission, "min_self_delegation:", msg.MinSelfDelegation,
						"delegator_address:", msg.DelegatorAddress, "validator_address:", msg.ValidatorAddress, "pubkey:", msg.PubKey, "value:", msg.Value, tx.Receipt.TxHash)
				default:
					return errors.New("unknown msg type:" + m.Type())
				}
			case "distr":
				switch m.Type() {
				case "set_withdraw_address":
					msg, ok := m.(draw.MsgSetWithdrawAddress)
					if !ok {
						return errors.New("can't convert to set_withdraw_address")
					}
					log.Debug(tx.Receipt.Height, "set_withdraw_address-->delegator_address:", msg.DelegatorAddress, "withdraw_address:", msg.WithdrawAddress, tx.Receipt.TxHash)
				case "withdraw_delegator_reward":
					msg, ok := m.(draw.MsgWithdrawDelegatorReward)
					if !ok {
						return errors.New("can't convert to withdraw_delegator_reward")
					}
					log.Debug(tx.Receipt.Height, "withdraw_delegator_reward-->delegator_address:", msg.DelegatorAddress.String(), "validator_address:", msg.ValidatorAddress.String(), tx.Receipt.TxHash)
				case "withdraw_validator_rewards_all":
					msg, ok := m.(draw.MsgWithdrawValidatorCommission)
					if !ok {
						return errors.New("can't convert to withdraw_validator_rewards_all")
					}
					log.Debug(tx.Receipt.Height, "withdraw_validator_rewards_all-->validator_address:", msg.ValidatorAddress, tx.Receipt.TxHash)
				default:
					return errors.New("unknown msg type:" + m.Type())
				}
			case "gov":
				switch m.Type() {
				case "vote":
					msg, ok := m.(gov.MsgVote)
					if !ok {
						return errors.New("can't convert to vote")
					}
					log.Info(tx.Receipt.Height, "vote-->proposal_id:", msg.ProposalID, "voter:", msg.Voter, "option:", msg.Option.String(), tx.Receipt.TxHash)
				case "deposit":
					msg, ok := m.(gov.MsgDeposit)
					if !ok {
						return errors.New("can't convert to deposit")
					}
					log.Info(tx.Receipt.Height, "deposit-->proposal_id:", msg.ProposalID, "depositor:", msg.Depositor, "amount:", msg.Amount, tx.Receipt.TxHash)
				case "submit_proposal":
					msg, ok := m.(gov.MsgSubmitProposal)
					if !ok {
						return errors.New("can't convert to submit_proposal")
					}
					log.Debug(tx.Receipt.Height, "submit_proposal-->title:", msg.Title, "description:", msg.Description,
						"proposal_type:", msg.ProposalType.String(), "proposer:", msg.Proposer.String(), "initial_deposit:", msg.InitialDeposit.String(), tx.Receipt.TxHash)
				default:
					return errors.New("unknown msg type:" + m.Type())
				}
			case "ibc":
				switch m.Type() {
				case "receive":
					msg, ok := m.(ibc.MsgIBCReceive)
					if !ok {
						return errors.New("can't convert to receive")
					}
					log.Debug(tx.Receipt.Height, "receive-->src_addr:", msg.SrcAddr, "dest_addr:", msg.DestAddr, "coins:", msg.Coins, "src_chain:",
						msg.SrcChain, "dest_chain:", msg.DestChain, "address:", msg.Relayer, "sequence:", msg.Sequence, tx.Receipt.TxHash)
				case "transfer":
					msg, ok := m.(ibc.MsgIBCTransfer)
					if !ok {
						return errors.New("can't convert to receive")
					}
					log.Debug(tx.Receipt.Height, "receive-->src_addr:", msg.SrcAddr, "dest_addr:", msg.DestAddr, "coins:", msg.Coins, "src_chain:", msg.SrcChain, "dest_chain:", msg.DestChain, tx.Receipt.TxHash)
				default:
					return errors.New("unknown msg type:" + m.Type())
				}
			case "slashing":
				switch m.Type() {
				case "unjail":
					msg, ok := m.(slashing.MsgUnjail)
					if !ok {
						return errors.New("can't convert to unjail")
					}
					log.Debug(tx.Receipt.Height, "unjail-->address:", msg.ValidatorAddr, tx.Receipt.TxHash)
				default:
					return errors.New("unknown msg type:" + m.Type())
				}
			case "crisis":
				switch m.Type() {
				case "verify_invariant":
					msg, ok := m.(crisis.MsgVerifyInvariant)
					if !ok {
						return errors.New("can't convert to verify_invariant")
					}
					log.Debug(tx.Receipt.Height, "verify_invariant-->sender:", msg.Sender.String(), "invariant_module_name:", msg.InvariantModuleName, "invariant_route:", msg.InvariantRoute, tx.Receipt.TxHash)
				default:
					return errors.New("unknown msg type:" + m.Type())
				}
			default:
				return errors.New("unknown msg route:" + m.Route())
			}
		}
	}
	return nil
}

func queryTx(cdc *codec.Codec, cliCtx context.CLIContext, hash []byte, resBlock *tt.ResultBlock) (Transaction, error) {
	node, err := cliCtx.GetNode()
	if err != nil {
		return Transaction{}, err
	}
	resTx, err := node.Tx(hash, !cliCtx.TrustNode)
	if err != nil {
		return Transaction{}, err
	}
	out, tx, err := formatTxResult(cdc, resTx, resBlock)
	if err != nil {
		return Transaction{}, err
	}
	return Transaction{Tx: tx, Receipt: out}, nil
}

func formatTxResult(cdc *codec.Codec, resTx *tt.ResultTx, resBlock *tt.ResultBlock) (sdk.TxResponse, auth.StdTx, error) {
	tx, err := parseTx(cdc, resTx.Tx)
	if err != nil {
		return sdk.TxResponse{}, auth.StdTx{}, err
	}
	return sdk.NewResponseResultTx(resTx, tx, resBlock.Block.Time.Format(time.RFC3339)), tx, nil
}

func parseTx(cdc *codec.Codec, txBytes []byte) (auth.StdTx, error) {
	var tx auth.StdTx
	err := cdc.UnmarshalBinaryLengthPrefixed(txBytes, &tx)
	if err != nil {
		return auth.StdTx{}, err
	}
	return tx, nil
}
