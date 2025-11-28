package mwsclient

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Data    Data   `json:"data"`
	Message string `json:"message"`
}

type Data struct {
	Total    int      `json:"total"`
	Records  []Record `json:"records"`
	PageNum  int      `json:"pageNum"`
	PageSize int      `json:"pageSize"`
}

type Record struct {
	RecordID  string `json:"recordId"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	Fields    Fields `json:"fields"`
}

type Fields struct {
	AIAnalyzeDate                 int64    `json:"ai_analyze_date"`
	ID                            int      `json:"id"`
	URL                           URL      `json:"url"`
	Description                   string   `json:"description"`
	Date                          int64    `json:"date"`
	CommentsSummary               string   `json:"comments_summary"`
	PopularityPredictionReasoning string   `json:"popularity_prediction_reasoning"`
	PopularityConfidence          float64  `json:"popularity_confidence"`
	Likes                         int      `json:"likes"`
	Recomendation                 string   `json:"recomendation"`
	Views                         int      `json:"views"`
	Comments                      []string `json:"comments"`
	CommentsTone                  string   `json:"comments_tone"`
	Summary                       string   `json:"summary"`
	Topic                         string   `json:"topic"`
	PopularityPrediction          float64  `json:"popularity_prediction"`
	Author                        string   `json:"author"`
	Dislikes                      int      `json:"dislikes"`
}

type URL struct {
	Title   string `json:"title"`
	Text    string `json:"text"`
	Favicon string `json:"favicon"`
}

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
