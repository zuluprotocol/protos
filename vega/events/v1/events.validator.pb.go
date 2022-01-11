// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vega/events/v1/events.proto

package v1

import (
	fmt "fmt"
	math "math"

	_ "code.vegaprotocol.io/protos/vega"
	_ "code.vegaprotocol.io/protos/vega/commands/v1"
	_ "code.vegaprotocol.io/protos/vega/oracles/v1"
	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *StakeLinking) Validate() error {
	return nil
}
func (this *CheckpointEvent) Validate() error {
	return nil
}
func (this *StreamStartEvent) Validate() error {
	return nil
}
func (this *RewardPayoutEvent) Validate() error {
	return nil
}
func (this *ValidatorScoreEvent) Validate() error {
	return nil
}
func (this *DelegationBalanceEvent) Validate() error {
	return nil
}
func (this *MarketEvent) Validate() error {
	return nil
}
func (this *TxErrorEvent) Validate() error {
	if oneOfNester, ok := this.GetTransaction().(*TxErrorEvent_OrderSubmission); ok {
		if oneOfNester.OrderSubmission != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.OrderSubmission); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("OrderSubmission", err)
			}
		}
	}
	if oneOfNester, ok := this.GetTransaction().(*TxErrorEvent_OrderAmendment); ok {
		if oneOfNester.OrderAmendment != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.OrderAmendment); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("OrderAmendment", err)
			}
		}
	}
	if oneOfNester, ok := this.GetTransaction().(*TxErrorEvent_OrderCancellation); ok {
		if oneOfNester.OrderCancellation != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.OrderCancellation); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("OrderCancellation", err)
			}
		}
	}
	if oneOfNester, ok := this.GetTransaction().(*TxErrorEvent_Proposal); ok {
		if oneOfNester.Proposal != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Proposal); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Proposal", err)
			}
		}
	}
	if oneOfNester, ok := this.GetTransaction().(*TxErrorEvent_VoteSubmission); ok {
		if oneOfNester.VoteSubmission != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.VoteSubmission); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("VoteSubmission", err)
			}
		}
	}
	if oneOfNester, ok := this.GetTransaction().(*TxErrorEvent_LiquidityProvisionSubmission); ok {
		if oneOfNester.LiquidityProvisionSubmission != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.LiquidityProvisionSubmission); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("LiquidityProvisionSubmission", err)
			}
		}
	}
	if oneOfNester, ok := this.GetTransaction().(*TxErrorEvent_WithdrawSubmission); ok {
		if oneOfNester.WithdrawSubmission != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.WithdrawSubmission); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("WithdrawSubmission", err)
			}
		}
	}
	if oneOfNester, ok := this.GetTransaction().(*TxErrorEvent_DelegateSubmission); ok {
		if oneOfNester.DelegateSubmission != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.DelegateSubmission); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("DelegateSubmission", err)
			}
		}
	}
	if oneOfNester, ok := this.GetTransaction().(*TxErrorEvent_UndelegateSubmission); ok {
		if oneOfNester.UndelegateSubmission != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.UndelegateSubmission); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("UndelegateSubmission", err)
			}
		}
	}
	if oneOfNester, ok := this.GetTransaction().(*TxErrorEvent_RestoreSnapshot); ok {
		if oneOfNester.RestoreSnapshot != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.RestoreSnapshot); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RestoreSnapshot", err)
			}
		}
	}
	if oneOfNester, ok := this.GetTransaction().(*TxErrorEvent_LiquidityProvisionCancellation); ok {
		if oneOfNester.LiquidityProvisionCancellation != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.LiquidityProvisionCancellation); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("LiquidityProvisionCancellation", err)
			}
		}
	}
	if oneOfNester, ok := this.GetTransaction().(*TxErrorEvent_LiquidityProvisionAmendment); ok {
		if oneOfNester.LiquidityProvisionAmendment != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.LiquidityProvisionAmendment); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("LiquidityProvisionAmendment", err)
			}
		}
	}
	return nil
}
func (this *TimeUpdate) Validate() error {
	return nil
}
func (this *EpochEvent) Validate() error {
	return nil
}
func (this *TransferResponses) Validate() error {
	for _, item := range this.Responses {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Responses", err)
			}
		}
	}
	return nil
}
func (this *PositionResolution) Validate() error {
	return nil
}
func (this *LossSocialization) Validate() error {
	return nil
}
func (this *TradeSettlement) Validate() error {
	return nil
}
func (this *SettlePosition) Validate() error {
	for _, item := range this.TradeSettlements {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("TradeSettlements", err)
			}
		}
	}
	return nil
}
func (this *SettleDistressed) Validate() error {
	return nil
}
func (this *MarketTick) Validate() error {
	return nil
}
func (this *AuctionEvent) Validate() error {
	return nil
}
func (this *ValidatorUpdate) Validate() error {
	return nil
}
func (this *KeyRotation) Validate() error {
	return nil
}
func (this *StateVar) Validate() error {
	return nil
}
func (this *BusEvent) Validate() error {
	if oneOfNester, ok := this.GetEvent().(*BusEvent_TimeUpdate); ok {
		if oneOfNester.TimeUpdate != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.TimeUpdate); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("TimeUpdate", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_TransferResponses); ok {
		if oneOfNester.TransferResponses != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.TransferResponses); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("TransferResponses", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_PositionResolution); ok {
		if oneOfNester.PositionResolution != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.PositionResolution); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PositionResolution", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_Order); ok {
		if oneOfNester.Order != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Order); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Order", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_Account); ok {
		if oneOfNester.Account != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Account); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Account", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_Party); ok {
		if oneOfNester.Party != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Party); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Party", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_Trade); ok {
		if oneOfNester.Trade != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Trade); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Trade", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_MarginLevels); ok {
		if oneOfNester.MarginLevels != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MarginLevels); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MarginLevels", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_Proposal); ok {
		if oneOfNester.Proposal != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Proposal); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Proposal", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_Vote); ok {
		if oneOfNester.Vote != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Vote); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Vote", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_MarketData); ok {
		if oneOfNester.MarketData != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MarketData); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MarketData", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_NodeSignature); ok {
		if oneOfNester.NodeSignature != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.NodeSignature); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("NodeSignature", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_LossSocialization); ok {
		if oneOfNester.LossSocialization != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.LossSocialization); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("LossSocialization", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_SettlePosition); ok {
		if oneOfNester.SettlePosition != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.SettlePosition); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("SettlePosition", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_SettleDistressed); ok {
		if oneOfNester.SettleDistressed != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.SettleDistressed); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("SettleDistressed", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_MarketCreated); ok {
		if oneOfNester.MarketCreated != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MarketCreated); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MarketCreated", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_Asset); ok {
		if oneOfNester.Asset != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Asset); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Asset", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_MarketTick); ok {
		if oneOfNester.MarketTick != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MarketTick); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MarketTick", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_Withdrawal); ok {
		if oneOfNester.Withdrawal != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Withdrawal); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Withdrawal", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_Deposit); ok {
		if oneOfNester.Deposit != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Deposit); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Deposit", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_Auction); ok {
		if oneOfNester.Auction != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Auction); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Auction", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_RiskFactor); ok {
		if oneOfNester.RiskFactor != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.RiskFactor); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RiskFactor", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_NetworkParameter); ok {
		if oneOfNester.NetworkParameter != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.NetworkParameter); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("NetworkParameter", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_LiquidityProvision); ok {
		if oneOfNester.LiquidityProvision != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.LiquidityProvision); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("LiquidityProvision", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_MarketUpdated); ok {
		if oneOfNester.MarketUpdated != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MarketUpdated); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MarketUpdated", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_OracleSpec); ok {
		if oneOfNester.OracleSpec != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.OracleSpec); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("OracleSpec", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_OracleData); ok {
		if oneOfNester.OracleData != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.OracleData); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("OracleData", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_DelegationBalance); ok {
		if oneOfNester.DelegationBalance != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.DelegationBalance); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("DelegationBalance", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_ValidatorScore); ok {
		if oneOfNester.ValidatorScore != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.ValidatorScore); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ValidatorScore", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_EpochEvent); ok {
		if oneOfNester.EpochEvent != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.EpochEvent); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("EpochEvent", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_ValidatorUpdate); ok {
		if oneOfNester.ValidatorUpdate != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.ValidatorUpdate); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ValidatorUpdate", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_StakeLinking); ok {
		if oneOfNester.StakeLinking != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.StakeLinking); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StakeLinking", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_RewardPayout); ok {
		if oneOfNester.RewardPayout != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.RewardPayout); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RewardPayout", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_Checkpoint); ok {
		if oneOfNester.Checkpoint != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Checkpoint); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Checkpoint", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_KeyRotation); ok {
		if oneOfNester.KeyRotation != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.KeyRotation); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("KeyRotation", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_StateVar); ok {
		if oneOfNester.StateVar != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.StateVar); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StateVar", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_NetworkLimits); ok {
		if oneOfNester.NetworkLimits != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.NetworkLimits); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("NetworkLimits", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_Market); ok {
		if oneOfNester.Market != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Market); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Market", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_TxErrEvent); ok {
		if oneOfNester.TxErrEvent != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.TxErrEvent); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("TxErrEvent", err)
			}
		}
	}
	return nil
}
