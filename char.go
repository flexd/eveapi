package eveapi

import (
	"fmt"
	"net/url"
)

const (
	CharAccountBalanceURL = "/char/AccountBalance.xml.aspx"
	CharSkillQueueURL     = "/char/SkillQueue.xml.aspx"
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
