package blacklist

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
)

func (c *Client) makeRequest(method, endpoint string, responseTarget interface{}) error {
	if reflect.TypeOf(responseTarget).Kind() != reflect.Ptr {
		return errors.New("blacklist sdk: responseTarget must be a pointer to a struct for JSON unmarshalling")
	}

	req, err := http.NewRequest(method, c.url+endpoint, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", c.apiKey)

	client := http.Client{
		Timeout: time.Second * 60,
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("\nResponse body: \n %v\n", string(b))

	if resp.StatusCode == 200 {
		err = json.Unmarshal(b, responseTarget)
		if err != nil {
			fmt.Println(err)
			fmt.Println("target sha", responseTarget)
			return err
		}
		return nil
	}
	//
	//err = Error{
	//	Code:     resp.StatusCode,
	//	Body:     string(b),
	//	Endpoint: req.URL.String(),
	//}
	return err
}
