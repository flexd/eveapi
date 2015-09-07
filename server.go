package eveapi

const (
	ServerStatusURL = "/server/ServerStatus.xml.aspx"
)

type ServerStatusResult struct {
	APIResult
	Open          bool `xml:"result>serverOpen"`
	OnlinePlayers int  `xml:"result>onlinePlayers"`
}

func (api API) ServerStatus() (*ServerStatusResult, error) {
	output := ServerStatusResult{}
	err := api.Call(ServerStatusURL, nil, &output)
	if err != nil {
		return nil, err
	}
	if output.Error != nil {
		return nil, output.Error
	}
	return &output, nil
}
