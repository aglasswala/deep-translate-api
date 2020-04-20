package translate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type RequestWords struct {
	Text string
}

func TranslateObjects(words []string, targetLang string, uri string, subscriptionKey string) ([]byte, error) {
	u, _ := url.Parse(uri)
	q := u.Query()
	q.Add("to", "de")
	u.RawQuery = q.Encode()

	body := [10]RequestWords{}
	for i, val := range words {
		body[i] = RequestWords{
			Text: val,
		}
	}

	b, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(b))
	if err != nil {
		log.Fatal(err)
	}
	// Add required headers to the request
	req.Header.Add("Ocp-Apim-Subscription-Key", subscriptionKey)
	req.Header.Add("Content-Type", "application/json")

	// Call the Translator Text API
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var result interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Println("ERROR: Couldn't decode from result")
		return nil, err
	}

	// Format and print the response to terminal
	prettyJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Println("ERROR: Couldn't format and print the response to terminal")
		return nil, err
	}

	fmt.Printf("%s/n", prettyJSON)

	return prettyJSON, nil
}
