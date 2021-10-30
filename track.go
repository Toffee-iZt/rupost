package rupost

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const turl = "https://www.pochta.ru/tracking?p_p_id=trackingPortlet_WAR_portalportlet&p_p_lifecycle=2&p_p_resource_id=tracking.get-by-barcodes"

// Track returns tracking information for tracking codes.
func Track(barcodes ...string) ([]Response, *Error, error) {
	pf := url.Values{
		"barcodes": barcodes,
	}
	resp, err := http.PostForm(turl, pf)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	var response struct {
		Error    *Error     `json:"error"`
		Response []Response `json:"response"`
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, nil, err
	}

	return response.Response, response.Error, nil
}
