package zabbix

import "github.com/AlekSi/reflector"

// https://www.zabbix.com/documentation/2.0/manual/appendix/api/history/definitions
type History struct {
	ItemId      string    `json:"itemid,omitempty"`
	Clock       int       `json:"clock"`
	Value	    string    `json:"value"`
	Nanoseconds int       `json:"ns,omitempty"`
}

// Wrapper for item.get https://www.zabbix.com/documentation/2.0/manual/appendix/api/item/get
func (api *API) HistoriesGet(params Params) (res []History, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	response, err := api.CallWithError("history.get", params)
	if err != nil {
		return
	}

	reflector.MapsToStructs2(response.Result.([]interface{}), &res, reflector.Strconv, "json")
	return
}

