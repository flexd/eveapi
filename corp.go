package eveapi

const (
	CorpContactListURL    = "/corp/ContactList.xml.aspx"
	CorpAccountBalanceURL = "/corp/AccountBalance.xml.aspx"
)

type Contact struct {
	ID       string `xml:"contactID,attr"`
	Name     string `xml:"contactName,attr"`
	Standing int    `xml:"standing,attr"`
}

type ContactSubList struct {
	Name     string    `xml:"name,attr"`
	Contacts []Contact `xml:"row"`
}

func (api API) CorpContactList() (*ContactList, error) {
	output := ContactList{}
	err := api.Call(CorpContactListURL, nil, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

type ContactList struct {
	APIResult
	ContactList []ContactSubList `xml:"result>rowset"`
}

func (c ContactList) Corporate() []Contact {
	for _, v := range c.ContactList {
		if v.Name == "corporateContactList" {
			return v.Contacts
		}
	}
	return nil
}
func (c ContactList) Alliance() []Contact {
	for _, v := range c.ContactList {
		if v.Name == "allianceContactList" {
			return v.Contacts
		}
	}
	return nil
}

type AccountBalance struct {
	APIResult
	Accounts []struct {
		ID      int     `xml:"accountID,attr"`
		Key     int     `xml:"accountKey,attr"`
		Balance float64 `xml:"balance,attr"`
	} `xml:"result>rowset>row"`
}

func (api API) CorpAccountBalances() (*AccountBalance, error) {
	output := AccountBalance{}
	err := api.Call(CorpAccountBalanceURL, nil, &output)
	if err != nil {
		return nil, err
	}
	if output.Error != nil {
		return nil, output.Error
	}
	return &output, nil
}
