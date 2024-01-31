package client

import (
	"encoding/json"
	"io"
)

func GetBase[T interface{}](
	endpoint string, params Parameters, data map[string]T,
) (map[string]T, error) {
	client := NewClient()
	req, errReq := NewRequest(endpoint, params)
	if errReq != nil {
		return nil, errReq
	}
	response, errRsp := client.Do(req)
	if errRsp != nil {
		return nil, errRsp
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	errDecode := json.NewDecoder(response.Body).Decode(&data)
	return data, errDecode
}
