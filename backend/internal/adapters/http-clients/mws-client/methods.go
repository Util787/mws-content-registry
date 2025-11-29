package mwsclient

import (
	"github.com/Util787/mws-content-registry/internal/models"
	"encoding/json"
	"fmt"
)

func (mwsClient *MWSClient) TakeAll() (Response, error) {
	res, err := mwsClient.client.R().Get(mwsClient.MWSUrl)
	if err != nil {
		return Response{}, err
	}
	var response Response

	err = json.Unmarshal(res.Body(), &response)
	if err != nil {
		return Response{}, err
	}
	fmt.Println("////////////////////////////////////////////////////////////////////////////////")
	fmt.Println("start")
	for i, rec := range response.Data.Records {
		fmt.Println(i)
		fmt.Println(rec.Fields)
		fmt.Println("////////////////////////////////////////////////////////////////////////////////")
	}
	fmt.Println("end")
	return response, err
}

func (mwsClient *MWSClient) TakeByID() {
	res, err := mwsClient.client.R().Get(mwsClient.MWSUrl)
	fmt.Println(err, res)
}
