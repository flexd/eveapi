package eveapi

const (
	//AccountAPIKeyInfoURL is the endpoint for api key info
	AccountAPIKeyInfoURL = "/account/APIKeyInfo.xml.aspx"
)

//AccountAPIKeyInfo fetches info this key (such as characters attached)
func (api API) AccountAPIKeyInfo() (*APIKeyInfoResponse, error) {
	output := APIKeyInfoResponse{}
	err := api.Call(AccountAPIKeyInfoURL, nil, &output)
	if err != nil {
		return nil, err
	}

	if output.Error != nil {
		return nil, output.Error
	}

	return &output, nil
}

//APIKeyInfoResponse details the api key in use
type APIKeyInfoResponse struct {
	APIResult
	Key APIKey `xml:"result>key"`
}

//APIKey the api key being used
type APIKey struct {
	AccessMask int             `xml:"accessMask,attr"`
	Type       string          `xml:"type,attr"`
	Rows       []APIKeyInfoRow `xml:"rowset>row"`
}

//APIKeyInfoRow details the characters the api key is for
type APIKeyInfoRow struct {
	ID              int    `xml:"characterID,attr"`
	Name            string `xml:"characterName,attr"`
	CorporationID   int    `xml:"corporationID,attr"`
	CorporationName string `xml:"corporationName,attr"`
	AllianceID      int    `xml:"allianceID,attr"`
	AllianceName    string `xml:"allianceName,attr"`
	FactionID       int    `xml:"factionID,attr"`
	FactionName     string `xml:"factionName,attr"`
}
