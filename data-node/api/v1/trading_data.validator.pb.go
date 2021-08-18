// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: data-node/api/v1/trading_data.proto

package v1

import (
	fmt "fmt"
	math "math"

	_ "code.vegaprotocol.io/protos/vega"
	_ "code.vegaprotocol.io/protos/vega/commands/v1"
	_ "code.vegaprotocol.io/protos/vega/events/v1"
	_ "code.vegaprotocol.io/protos/vega/oracles/v1"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *DelegationsRequest) Validate() error {
	return nil
}
func (this *DelegationsResponse) Validate() error {
	for _, item := range this.Delegations {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Delegations", err)
			}
		}
	}
	return nil
}
func (this *StakeLinkingsRequest) Validate() error {
	return nil
}
func (this *StakeLinkingsResponse) Validate() error {
	for _, item := range this.StakeLinkings {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StakeLinkings", err)
			}
		}
	}
	return nil
}
func (this *GetNodeDataRequest) Validate() error {
	return nil
}
func (this *GetNodeDataResponse) Validate() error {
	if this.NodeData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.NodeData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("NodeData", err)
		}
	}
	return nil
}
func (this *GetNodesRequest) Validate() error {
	return nil
}
func (this *GetNodesResponse) Validate() error {
	for _, item := range this.Nodes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Nodes", err)
			}
		}
	}
	return nil
}
func (this *GetNodeByIDRequest) Validate() error {
	return nil
}
func (this *GetNodeByIDResponse) Validate() error {
	if this.Node != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Node); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Node", err)
		}
	}
	return nil
}
func (this *GetEpochRequest) Validate() error {
	return nil
}
func (this *GetEpochResponse) Validate() error {
	if this.Epoch != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Epoch); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Epoch", err)
		}
	}
	return nil
}
func (this *AssetsRequest) Validate() error {
	return nil
}
func (this *AssetsResponse) Validate() error {
	for _, item := range this.Assets {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Assets", err)
			}
		}
	}
	return nil
}
func (this *AssetByIDRequest) Validate() error {
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must not be an empty string`, this.Id))
	}
	return nil
}
func (this *AssetByIDResponse) Validate() error {
	if this.Asset != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Asset); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Asset", err)
		}
	}
	return nil
}
func (this *GetNodeSignaturesAggregateRequest) Validate() error {
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must not be an empty string`, this.Id))
	}
	return nil
}
func (this *GetNodeSignaturesAggregateResponse) Validate() error {
	for _, item := range this.Signatures {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Signatures", err)
			}
		}
	}
	return nil
}
func (this *OptionalProposalState) Validate() error {
	return nil
}
func (this *GetProposalsRequest) Validate() error {
	if this.SelectInState != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SelectInState); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SelectInState", err)
		}
	}
	return nil
}
func (this *GetProposalsResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetProposalsByPartyRequest) Validate() error {
	if this.PartyId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyId))
	}
	if this.SelectInState != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SelectInState); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SelectInState", err)
		}
	}
	return nil
}
func (this *GetProposalsByPartyResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetVotesByPartyRequest) Validate() error {
	if this.PartyId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyId))
	}
	return nil
}
func (this *GetVotesByPartyResponse) Validate() error {
	for _, item := range this.Votes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Votes", err)
			}
		}
	}
	return nil
}
func (this *GetNewMarketProposalsRequest) Validate() error {
	if this.SelectInState != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SelectInState); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SelectInState", err)
		}
	}
	return nil
}
func (this *GetNewMarketProposalsResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetUpdateMarketProposalsRequest) Validate() error {
	if this.MarketId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketId", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketId))
	}
	if this.SelectInState != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SelectInState); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SelectInState", err)
		}
	}
	return nil
}
func (this *GetUpdateMarketProposalsResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetNetworkParametersProposalsRequest) Validate() error {
	if this.SelectInState != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SelectInState); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SelectInState", err)
		}
	}
	return nil
}
func (this *GetNetworkParametersProposalsResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetNewAssetProposalsRequest) Validate() error {
	if this.SelectInState != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SelectInState); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SelectInState", err)
		}
	}
	return nil
}
func (this *GetNewAssetProposalsResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetProposalByIDRequest) Validate() error {
	if this.ProposalId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProposalId", fmt.Errorf(`value '%v' must not be an empty string`, this.ProposalId))
	}
	return nil
}
func (this *GetProposalByIDResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetProposalByReferenceRequest) Validate() error {
	if this.Reference == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Reference", fmt.Errorf(`value '%v' must not be an empty string`, this.Reference))
	}
	return nil
}
func (this *GetProposalByReferenceResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *ObserveGovernanceRequest) Validate() error {
	return nil
}
func (this *ObserveGovernanceResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *ObservePartyProposalsRequest) Validate() error {
	if this.PartyId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyId))
	}
	return nil
}
func (this *ObservePartyProposalsResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *ObserveProposalVotesRequest) Validate() error {
	if this.ProposalId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProposalId", fmt.Errorf(`value '%v' must not be an empty string`, this.ProposalId))
	}
	return nil
}
func (this *ObserveProposalVotesResponse) Validate() error {
	if this.Vote != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Vote); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Vote", err)
		}
	}
	return nil
}
func (this *ObservePartyVotesRequest) Validate() error {
	if this.PartyId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyId))
	}
	return nil
}
func (this *ObservePartyVotesResponse) Validate() error {
	if this.Vote != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Vote); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Vote", err)
		}
	}
	return nil
}
func (this *MarginLevelsSubscribeRequest) Validate() error {
	if this.PartyId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyId))
	}
	return nil
}
func (this *MarginLevelsSubscribeResponse) Validate() error {
	if this.MarginLevels != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.MarginLevels); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("MarginLevels", err)
		}
	}
	return nil
}
func (this *MarginLevelsRequest) Validate() error {
	if this.PartyId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyId))
	}
	return nil
}
func (this *MarginLevelsResponse) Validate() error {
	for _, item := range this.MarginLevels {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MarginLevels", err)
			}
		}
	}
	return nil
}
func (this *MarketsDataSubscribeRequest) Validate() error {
	return nil
}
func (this *MarketsDataSubscribeResponse) Validate() error {
	if this.MarketData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.MarketData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("MarketData", err)
		}
	}
	return nil
}
func (this *MarketDataByIDRequest) Validate() error {
	if this.MarketId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketId", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketId))
	}
	return nil
}
func (this *MarketDataByIDResponse) Validate() error {
	if this.MarketData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.MarketData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("MarketData", err)
		}
	}
	return nil
}
func (this *MarketsDataRequest) Validate() error {
	return nil
}
func (this *MarketsDataResponse) Validate() error {
	for _, item := range this.MarketsData {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MarketsData", err)
			}
		}
	}
	return nil
}
func (this *LastTradeRequest) Validate() error {
	if this.MarketId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketId", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketId))
	}
	return nil
}
func (this *LastTradeResponse) Validate() error {
	if this.Trade != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Trade); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Trade", err)
		}
	}
	return nil
}
func (this *MarketByIDRequest) Validate() error {
	if this.MarketId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketId", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketId))
	}
	return nil
}
func (this *MarketByIDResponse) Validate() error {
	if this.Market != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Market); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Market", err)
		}
	}
	return nil
}
func (this *PartyByIDRequest) Validate() error {
	if this.PartyId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyId))
	}
	return nil
}
func (this *PartyByIDResponse) Validate() error {
	if this.Party != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Party); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Party", err)
		}
	}
	return nil
}
func (this *PartiesRequest) Validate() error {
	return nil
}
func (this *PartiesResponse) Validate() error {
	for _, item := range this.Parties {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Parties", err)
			}
		}
	}
	return nil
}
func (this *TradesByPartyRequest) Validate() error {
	if this.Pagination != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pagination); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pagination", err)
		}
	}
	return nil
}
func (this *TradesByPartyResponse) Validate() error {
	for _, item := range this.Trades {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Trades", err)
			}
		}
	}
	return nil
}
func (this *TradesByOrderRequest) Validate() error {
	return nil
}
func (this *TradesByOrderResponse) Validate() error {
	for _, item := range this.Trades {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Trades", err)
			}
		}
	}
	return nil
}
func (this *AccountsSubscribeRequest) Validate() error {
	return nil
}
func (this *AccountsSubscribeResponse) Validate() error {
	if this.Account != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Account); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Account", err)
		}
	}
	return nil
}
func (this *OrdersSubscribeRequest) Validate() error {
	return nil
}
func (this *TradesSubscribeRequest) Validate() error {
	return nil
}
func (this *CandlesSubscribeRequest) Validate() error {
	if this.MarketId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketId", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketId))
	}
	return nil
}
func (this *CandlesSubscribeResponse) Validate() error {
	if this.Candle != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Candle); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Candle", err)
		}
	}
	return nil
}
func (this *MarketDepthSubscribeRequest) Validate() error {
	if this.MarketId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketId", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketId))
	}
	return nil
}
func (this *MarketDepthSubscribeResponse) Validate() error {
	if this.MarketDepth != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.MarketDepth); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("MarketDepth", err)
		}
	}
	return nil
}
func (this *MarketDepthUpdatesSubscribeRequest) Validate() error {
	if this.MarketId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketId", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketId))
	}
	return nil
}
func (this *MarketDepthUpdatesSubscribeResponse) Validate() error {
	if this.Update != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Update); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Update", err)
		}
	}
	return nil
}
func (this *PositionsSubscribeRequest) Validate() error {
	return nil
}
func (this *PositionsSubscribeResponse) Validate() error {
	if this.Position != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Position); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Position", err)
		}
	}
	return nil
}
func (this *OrdersByMarketRequest) Validate() error {
	if this.MarketId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketId", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketId))
	}
	if this.Pagination != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pagination); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pagination", err)
		}
	}
	return nil
}
func (this *OrdersByMarketResponse) Validate() error {
	for _, item := range this.Orders {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Orders", err)
			}
		}
	}
	return nil
}
func (this *OrdersByPartyRequest) Validate() error {
	if this.PartyId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyId))
	}
	if this.Pagination != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pagination); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pagination", err)
		}
	}
	return nil
}
func (this *OrdersByPartyResponse) Validate() error {
	for _, item := range this.Orders {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Orders", err)
			}
		}
	}
	return nil
}
func (this *OrderByMarketAndIDRequest) Validate() error {
	if this.MarketId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketId", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketId))
	}
	if this.OrderId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("OrderId", fmt.Errorf(`value '%v' must not be an empty string`, this.OrderId))
	}
	return nil
}
func (this *OrderByMarketAndIDResponse) Validate() error {
	if this.Order != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Order); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Order", err)
		}
	}
	return nil
}
func (this *OrderByReferenceRequest) Validate() error {
	if this.Reference == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Reference", fmt.Errorf(`value '%v' must not be an empty string`, this.Reference))
	}
	return nil
}
func (this *OrderByReferenceResponse) Validate() error {
	if this.Order != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Order); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Order", err)
		}
	}
	return nil
}
func (this *MarketsRequest) Validate() error {
	return nil
}
func (this *MarketsResponse) Validate() error {
	for _, item := range this.Markets {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Markets", err)
			}
		}
	}
	return nil
}
func (this *CandlesRequest) Validate() error {
	if this.MarketId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketId", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketId))
	}
	if !(this.SinceTimestamp > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("SinceTimestamp", fmt.Errorf(`value '%v' must be greater than '0'`, this.SinceTimestamp))
	}
	return nil
}
func (this *CandlesResponse) Validate() error {
	for _, item := range this.Candles {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Candles", err)
			}
		}
	}
	return nil
}
func (this *MarketDepthRequest) Validate() error {
	if this.MarketId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketId", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketId))
	}
	return nil
}
func (this *MarketDepthResponse) Validate() error {
	for _, item := range this.Buy {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Buy", err)
			}
		}
	}
	for _, item := range this.Sell {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Sell", err)
			}
		}
	}
	if this.LastTrade != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LastTrade); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LastTrade", err)
		}
	}
	return nil
}
func (this *TradesByMarketRequest) Validate() error {
	if this.MarketId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketId", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketId))
	}
	if this.Pagination != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pagination); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pagination", err)
		}
	}
	return nil
}
func (this *TradesByMarketResponse) Validate() error {
	for _, item := range this.Trades {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Trades", err)
			}
		}
	}
	return nil
}
func (this *PositionsByPartyRequest) Validate() error {
	if this.PartyId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyId))
	}
	return nil
}
func (this *PositionsByPartyResponse) Validate() error {
	for _, item := range this.Positions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Positions", err)
			}
		}
	}
	return nil
}
func (this *GetVegaTimeRequest) Validate() error {
	return nil
}
func (this *GetVegaTimeResponse) Validate() error {
	return nil
}
func (this *Pagination) Validate() error {
	return nil
}
func (this *OrdersSubscribeResponse) Validate() error {
	for _, item := range this.Orders {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Orders", err)
			}
		}
	}
	return nil
}
func (this *TradesSubscribeResponse) Validate() error {
	for _, item := range this.Trades {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Trades", err)
			}
		}
	}
	return nil
}
func (this *TransferResponsesSubscribeRequest) Validate() error {
	return nil
}
func (this *TransferResponsesSubscribeResponse) Validate() error {
	if this.Response != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Response); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Response", err)
		}
	}
	return nil
}
func (this *PartyAccountsRequest) Validate() error {
	return nil
}
func (this *PartyAccountsResponse) Validate() error {
	for _, item := range this.Accounts {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Accounts", err)
			}
		}
	}
	return nil
}
func (this *MarketAccountsRequest) Validate() error {
	return nil
}
func (this *MarketAccountsResponse) Validate() error {
	for _, item := range this.Accounts {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Accounts", err)
			}
		}
	}
	return nil
}
func (this *FeeInfrastructureAccountsRequest) Validate() error {
	return nil
}
func (this *FeeInfrastructureAccountsResponse) Validate() error {
	for _, item := range this.Accounts {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Accounts", err)
			}
		}
	}
	return nil
}
func (this *OrderByIDRequest) Validate() error {
	return nil
}
func (this *OrderByIDResponse) Validate() error {
	if this.Order != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Order); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Order", err)
		}
	}
	return nil
}
func (this *OrderVersionsByIDRequest) Validate() error {
	if this.Pagination != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pagination); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pagination", err)
		}
	}
	return nil
}
func (this *OrderVersionsByIDResponse) Validate() error {
	for _, item := range this.Orders {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Orders", err)
			}
		}
	}
	return nil
}
func (this *EstimateFeeRequest) Validate() error {
	if this.Order != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Order); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Order", err)
		}
	}
	return nil
}
func (this *EstimateFeeResponse) Validate() error {
	if this.Fee != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Fee); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Fee", err)
		}
	}
	return nil
}
func (this *EstimateMarginRequest) Validate() error {
	if this.Order != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Order); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Order", err)
		}
	}
	return nil
}
func (this *EstimateMarginResponse) Validate() error {
	if this.MarginLevels != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.MarginLevels); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("MarginLevels", err)
		}
	}
	return nil
}
func (this *ObserveEventBusRequest) Validate() error {
	return nil
}
func (this *ObserveEventBusResponse) Validate() error {
	for _, item := range this.Events {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Events", err)
			}
		}
	}
	return nil
}
func (this *WithdrawalsRequest) Validate() error {
	if this.PartyId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyId))
	}
	return nil
}
func (this *WithdrawalsResponse) Validate() error {
	for _, item := range this.Withdrawals {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Withdrawals", err)
			}
		}
	}
	return nil
}
func (this *WithdrawalRequest) Validate() error {
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must not be an empty string`, this.Id))
	}
	return nil
}
func (this *WithdrawalResponse) Validate() error {
	if this.Withdrawal != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Withdrawal); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Withdrawal", err)
		}
	}
	return nil
}
func (this *ERC20WithdrawalApprovalRequest) Validate() error {
	if this.WithdrawalId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("WithdrawalId", fmt.Errorf(`value '%v' must not be an empty string`, this.WithdrawalId))
	}
	return nil
}
func (this *ERC20WithdrawalApprovalResponse) Validate() error {
	return nil
}
func (this *DepositsRequest) Validate() error {
	if this.PartyId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyId))
	}
	return nil
}
func (this *DepositsResponse) Validate() error {
	for _, item := range this.Deposits {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Deposits", err)
			}
		}
	}
	return nil
}
func (this *DepositRequest) Validate() error {
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must not be an empty string`, this.Id))
	}
	return nil
}
func (this *DepositResponse) Validate() error {
	if this.Deposit != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Deposit); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Deposit", err)
		}
	}
	return nil
}
func (this *NetworkParametersRequest) Validate() error {
	return nil
}
func (this *NetworkParametersResponse) Validate() error {
	for _, item := range this.NetworkParameters {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("NetworkParameters", err)
			}
		}
	}
	return nil
}
func (this *LiquidityProvisionsRequest) Validate() error {
	return nil
}
func (this *LiquidityProvisionsResponse) Validate() error {
	for _, item := range this.LiquidityProvisions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("LiquidityProvisions", err)
			}
		}
	}
	return nil
}
func (this *OracleSpecRequest) Validate() error {
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must not be an empty string`, this.Id))
	}
	return nil
}
func (this *OracleSpecResponse) Validate() error {
	if this.OracleSpec != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.OracleSpec); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("OracleSpec", err)
		}
	}
	return nil
}
func (this *OracleSpecsRequest) Validate() error {
	return nil
}
func (this *OracleSpecsResponse) Validate() error {
	for _, item := range this.OracleSpecs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("OracleSpecs", err)
			}
		}
	}
	return nil
}
func (this *OracleDataBySpecRequest) Validate() error {
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must not be an empty string`, this.Id))
	}
	return nil
}
func (this *OracleDataBySpecResponse) Validate() error {
	for _, item := range this.OracleData {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("OracleData", err)
			}
		}
	}
	return nil
}
func (this *LastBlockHeightRequest) Validate() error {
	return nil
}
func (this *LastBlockHeightResponse) Validate() error {
	return nil
}
func (this *GetRewardDetailsRequest) Validate() error {
	if this.PartyId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyId", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyId))
	}
	return nil
}
func (this *GetRewardDetailsResponse) Validate() error {
	for _, item := range this.RewardDetails {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RewardDetails", err)
			}
		}
	}
	return nil
}
func (this *Checkpoint) Validate() error {
	return nil
}
func (this *CheckpointsRequest) Validate() error {
	return nil
}
func (this *CheckpointsResponse) Validate() error {
	for _, item := range this.Checkpoints {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Checkpoints", err)
			}
		}
	}
	return nil
}
