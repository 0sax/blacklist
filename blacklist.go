package blacklist

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/0sax/err2"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
)

type Client struct {
	url    string
	apiKey string
}

func NewBlackListClient(baseUrl, apiKey string) *Client {

	return &Client{
		url:    baseUrl,
		apiKey: apiKey,
	}
}

func (c *Client) SearchBlacklistFull(bvn string) (blr []BlacklistLoanRecord, err error) {

	var r BlacklistBVNResult

	err = c.makeRequest(
		"GET",
		fmt.Sprintf("bvn/%v", bvn),
		&r)
	if err != nil {
		return
	}

	//client errors
	if r.Status == "error" {
		err = err2.NewClientErr(nil, r.Message, 400)
		return
	}

	if _, ok := r.Data.(bool); ok {
		err = err2.NewClientErr(nil, r.Message, 204)
		return
	}

	if a, ok := r.Data.([]interface{}); ok {
		for _, b := range a {
			bm := b.(map[string]interface{})

			bl := BlacklistLoanRecord{
				CompanyName: bm["company_name"].(string),
				Name:        bm["name"].(string),
				Phone:       bm["phone"].(string),
				Email:       bm["email"].(string),
				Bvn:         bm["bvn"].(string),
				Gender:      bm["gender"].(string),
				LoanAmount:  bm["loan_amount"].(string),
				AmountPaid:  bm["amount_paid"].(string),
				Balance:     bm["balance"].(string),
				DueDate:     bm["due_date"].(string),
				Location:    bm["location"].(string),
				Date:        bm["Date"].(string),
			}

			blr = append(blr, bl)
		}
	}

	//if lr, ok := r.Data.([]BlacklistLoanRecord); ok {
	//
	//}
	//for _, res := range *r.Data {
	//	blr = append(blr, res)
	//}

	return

}

func (c *Client) SearchCRCFull(bvn string) (cd *CRCData, err error) {

	var r CRCBVNResult

	err = c.makeRequest(
		"GET",
		fmt.Sprintf("bvn-blacklist-crc-search/%v", bvn),
		&r)
	if err != nil {
		fmt.Println("error here 1") //debug delete
		fmt.Println(err)
		err = err2.NewServerErr(err)
		return
	}

	fmt.Printf("result %+v", r)

	if a := r.CRC[0]; a.Status == "error" {
		fmt.Println("error here 2") //debug delete
		err = err2.NewClientErr(nil, a.Message, 400)
		return
	}

	cd = r.CRC[0].Data

	return

}

func (c *Client) makeRequest(method, endpoint string, responseTarget interface{}) error {
	if reflect.TypeOf(responseTarget).Kind() != reflect.Ptr {
		return errors.New("blacklist sdk: responseTarget must be a pointer to a struct for JSON unmarshalling")
	}

	req, err := http.NewRequest(method, c.url+endpoint, nil)
	if err != nil {
		fmt.Println("error hereA") //debug delete
		return err
	}

	req.Header.Set("Authorization", c.apiKey)

	client := http.Client{
		Timeout: time.Second * 20,
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error hereC") //debug delete
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error hereD") //debug delete
		return err
	}

	fmt.Println("\nHeres the body \n %v\n", string(b)) //debug delete

	if resp.StatusCode == 200 {
		err = json.Unmarshal(b, responseTarget)
		if err != nil {
			fmt.Println("error hereE") //debug delete
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
