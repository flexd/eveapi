package eveapi

import (
	"encoding/xml"
	"fmt"
	"github.com/whenhellfreezes/Stack"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const (
	AssetListUrl = "/char/AssetList.xml.aspx"
)

func (api API) GetAssets(charID string) (*Assets, error) {
	output := Assets{}
	arguments := url.Values{}
	arguments.Set("characterID", charID)
	arguments.Set("keyID", api.APIKey.ID)
	arguments.Set("vCode", api.APIKey.VCode)
	uri := api.Server + "/char/AssetList.xml.aspx"
	resp, err := http.PostForm(uri, arguments)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if api.Debug {
		io.Copy(os.Stdout, resp.Body)
	}
	err = parseAssets(xml.NewDecoder(resp.Body), &output)
	if err != nil {
		return nil, err
	}
	if output.Result.Error != nil {
		return nil, output.Result.Error
	}
	return &output, nil
}

type Item struct {
	itemID    int
	typeID    int
	quantity  int
	flag      int
	singleton bool
}

func (i Item) String() string {
	out := fmt.Sprintf("itemID: %d typeID: %d quantity: %d flag: %d singleton: %t \n", i.itemID, i.typeID, i.quantity, i.flag, i.singleton)
	return out
}

type Valueable interface {
	GetItem() *Item
	GetContents() []Valueable
}

type CompositeItem struct {
	Item     Item
	Contents []Valueable
}

func (c CompositeItem) GetItem() *Item {
	return &c.Item
}

func (c CompositeItem) GetContents() []Valueable {
	return c.Contents
}

func (c *CompositeItem) add(other Valueable) {
	c.Contents = append(c.Contents, other)
}

type Assets struct {
	Result APIResult
	Items  map[int]([]*CompositeItem) //Key refers to the locationID
}

func parseAssets(d *xml.Decoder, output *Assets) error {
	output.Items = make(map[int]([]*CompositeItem))
	result := APIResult{}
	containers := stack.NewStack(0) //Of type *CompositeItem
	var current_item Item
	var current_location int
	t, err := d.Token()
	if t == nil {
		fmt.Println("Empty Decoder")
	}
	for t != nil {
		switch se := t.(type) {
		case xml.StartElement:
			switch se.Name.Local {
			case "eveapi":
				for _, value := range se.Attr {
					if value.Name.Local == "version" {
						result.Version, _ = strconv.Atoi(value.Value)
					}
				}
			case "currentTime":
				d.DecodeElement(&result.CurrentTime, &se)
			case "cachedUntil":
				d.DecodeElement(&result.CachedUntil, &se)
			case "error":
				d.DecodeElement(&result.Error, &se)
			case "row":
				for _, value := range se.Attr {
					switch value.Name.Local {
					case "locationID":
						current_location, _ = strconv.Atoi(value.Value)
					case "itemID":
						current_item.itemID, _ = strconv.Atoi(value.Value)
					case "typeID":
						current_item.typeID, _ = strconv.Atoi(value.Value)
					case "quantity":
						current_item.quantity, _ = strconv.Atoi(value.Value)
					case "flag":
						current_item.flag, _ = strconv.Atoi(value.Value)
					case "singleton":
						current_item.singleton, _ = strconv.ParseBool(value.Value)
					default:
					}
				}
				var composite CompositeItem
				composite.Item = current_item
				if containers.Len() != 0 {
					top, err := containers.Peek()
					if err != nil {
						fmt.Println(err)
					}
					ptop := top.(*CompositeItem)
					ptop.add(&composite)
				} else {
					output.Items[current_location] = append(output.Items[current_location], &composite)
				}
				containers.Push(&composite)
			default:
			}
		case xml.EndElement:
			if se.Name.Local == "row" {
				containers.Pop()
			}
		default:
		}
		t, err = d.Token()
		if err != nil && err != io.EOF {
			fmt.Println(err)
		}
	}
	output.Result = result
	return nil
}
