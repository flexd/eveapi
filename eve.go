package eveapi

import (
	"errors"
	"net/url"
)

const (
	RefTypesURL     = "/Eve/RefTypes.xml.aspx"
	AllianceListURL = "/Eve/AllianceList.xml.aspx"
	CharacterIDURL  = "/eve/CharacterID.xml.aspx"
)

// CharacterName calls the /eve/CharacterName.xml.aspx endpoint with the parameter chars, which is a comma separate string of character names.
// It returns the results and any error encountered
func (api API) CharacterName(chars string) (*CharacterNameResult, error) {
	output := CharacterNameResult{}
	arguments := url.Values{}
	arguments.Add("names", chars)
	err := api.Call(CharacterIDURL, arguments, &output)
	if err != nil {
		return nil, err
	}
	if output.Error != nil {
		return nil, output.Error
	}
	return &output, nil
}

// Name2ID is a convenience wrapper around CharacterName.
// It takes a single parameter char string which is the name of the character you want the characterID of.
// It returns the characterID and any error encountered
func (api API) Name2ID(char string) (string, error) {
	names, err := api.CharacterName(char)
	if err != nil {
		return "", err
	}
	if len(names.Names) == 1 && names.Names[0].ID != "0" {
		return names.Names[0].ID, nil
	} else {
		return "", errors.New("Name2ID: No such character")
	}
}

// Namees2ID is a convenience wrapper around CharacterName.
// It takes a single parameter chars string which is a comma separated string of character names you want the details for.
// It returns the characters and any error encountered
func (api API) Names2ID(char string) ([]CharacterNameRow, error) {
	names, err := api.CharacterName(char)
	if err != nil {
		return nil, err
	}
	return names.Names, nil
}

// CharacterNameRow represents a single row in CharacterNameResult
type CharacterNameRow struct {
	ID   string `xml:"characterID,attr"`
	Name string `xml:"name,attr"`
}
type CharacterNameResult struct {
	APIResult
	Names []CharacterNameRow `xml:"result>rowset>row"`
}

func (api API) RefTypes() (*RefTypes, error) {
	output := RefTypes{}
	err := api.Call(RefTypesURL, nil, &output)
	if err != nil {
		return nil, err
	}
	if output.Error != nil {
		return nil, output.Error
	}
	return &output, nil
}

type RefTypes struct {
	APIResult
	RefTypes []struct {
		RefTypeID   string `xml:"refTypeID,attr"`
		RefTypeName string `xml:"refTypeName,attr"`
	} `xml:"result>rowset>row"`
}

func (api API) AllianceList() (*AllianceList, error) {
	output := AllianceList{}
	err := api.Call(AllianceListURL, nil, &output)
	if err != nil {
		return nil, err
	}
	if output.Error != nil {
		return nil, output.Error
	}
	return &output, nil
}

type AllianceList struct {
	APIResult
	Alliances []struct {
		Name           string  `xml:"name,attr"`
		Ticker         string  `xml:"shortName,attr"`
		AllianceID     string  `xml:"allianceID,attr"`
		ExecutorCorpID string  `xml:"executorCorpID,attr"`
		MemberCount    int     `xml:"memberCount,attr"`
		CorporationID  string  `xml:"corporationID,attr"`
		Created        eveTime `xml:"startDate,attr"`
	} `xml:"result>rowset>row"`
}
