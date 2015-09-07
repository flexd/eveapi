package eveapi

import (
	"fmt"
	"net/url"
	"strconv"
)

const (
	//CharAccountBalanceURL is the url for the account balance endpoint
	CharAccountBalanceURL = "/char/AccountBalance.xml.aspx"
	//CharSkillQueueURL is the url for the skill queue endpoint
	CharSkillQueueURL = "/char/SkillQueue.xml.aspx"
	//MarketOrdersURL is the url for the market orders endpoint
	MarketOrdersURL = "/char/MarketOrders.xml.aspx"
)

//AccountBalance is defined in corp.go

// CharAccountBalances calls /char/AccountBalance.xml.aspx
// Returns the account balance and any error if occured.
func (api API) CharAccountBalances(charID string) (*AccountBalance, error) {
	output := AccountBalance{}
	arguments := url.Values{}
	arguments.Add("characterID", charID)
	err := api.Call(CharAccountBalanceURL, arguments, &output)
	if err != nil {
		return nil, err
	}
	if output.Error != nil {
		return nil, output.Error
	}
	return &output, nil
}

//SkillQueueRow is an entry in a character's skill queue
type SkillQueueRow struct {
	Position  int     `xml:"queuePosition,attr"`
	TypeID    int     `xml:"typeID,attr"`
	Level     int     `xml:"level,attr"`
	StartSP   int     `xml:"startSP,attr"`
	EndSP     int     `xml:"endSP,attr"`
	StartTime eveTime `xml:"startTime,attr"`
	EndTime   eveTime `xml:"endTime,attr"`
}

func (s SkillQueueRow) String() string {
	return fmt.Sprintf("Position: %v, TypeID: %v, Level: %v, StartSP: %v, EndSP: %v, StartTime: %v, EndTime: %v", s.Position, s.TypeID, s.Level, s.StartSP, s.EndSP, s.StartTime, s.EndTime)
}

//SkillQueueResult is the result returned by the skill queue endpoint
type SkillQueueResult struct {
	APIResult
	SkillQueue []SkillQueueRow `xml:"result>rowset>row"`
}

// SkillQueue calls the API passing the parameter charID
// Returns a SkillQueueResult struct
func (api API) SkillQueue(charID string) (*SkillQueueResult, error) {
	output := SkillQueueResult{}
	arguments := url.Values{}
	arguments.Add("characterID", charID)
	err := api.Call(CharSkillQueueURL, arguments, &output)
	if err != nil {
		return nil, err
	}
	if output.Error != nil {
		return nil, output.Error
	}
	return &output, nil
}

//MarketOrdersResult is the result from calling the market orders endpoint
type MarketOrdersResult struct {
	APIResult
	Orders []MarketOrder `xml:"result>rowset>row"`
}

//MarketOrder is either a sell order or buy order
type MarketOrder struct {
	OrderID      int     `xml:"orderID,attr"`
	CharID       int     `xml:"charID,attr"`
	StationID    int     `xml:"stationID,attr"`
	VolEntered   int     `xml:"volEntered,attr"`
	VolRemaining int     `xml:"volRemaining,attr"`
	MinVolume    int     `xml:"minVolume,attr"`
	TypeID       int     `xml:"typeID,attr"`
	Range        int     `xml:"range,attr"`
	Division     int     `xml:"accountKey,attr"`
	Escrow       float64 `xml:"escrow,attr"`
	Price        float64 `xml:"price,attr"`
	IsBuyOrder   bool    `xml:"bid,attr"`
	Issued       eveTime `xml:"issued,attr"`
}

//MarketOrders returns the market orders for a given character
func (api API) MarketOrders(charID int) (*MarketOrdersResult, error) {
	output := MarketOrdersResult{}
	args := url.Values{}
	args.Add("characterID", strconv.Itoa(charID))
	err := api.Call(MarketOrdersURL, args, &output)
	if err != nil {
		return nil, err
	}
	if output.Error != nil {
		return nil, output.Error
	}
	return &output, nil
}
