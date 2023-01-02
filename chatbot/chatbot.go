package chatbot

import (
 "encoding/json"
 "fmt"
 "io/ioutil"
 "net/http"
 "net/url"
)

//------------------------------------------------------------------------------------

type ChatbotResponse struct {
	Cnt string `json:"cnt"`
}

//------------------------------------------------------------------------------------

func GetResponse(query string) string {


	EncodedQuery := url.QueryEscape(query)

    URL := fmt.Sprintf("http://api.brainshop.ai/get?bid=153868&key=rcKonOgrUFmn5usX&uid=1&msg=%s", EncodedQuery)
    
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	
	var responseObject ChatbotResponse
	json.Unmarshal(bodyBytes, &responseObject)
	return responseObject.Cnt
	
}
