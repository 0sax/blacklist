package blacklist

import (
	"fmt"
	"github.com/0sax/err2"
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

	if r.Data != nil {
		blr = r.Data
		return
	}

	//if _, ok := r.Data.(bool); ok {
	//	err = err2.NewClientErr(nil, r.Message, 204)
	//	return
	//}
	//
	//if _, ok := r.Data.([]interface{}); ok {
	//	a := r.Data.([]interface{})
	//	for _, b := range a {
	//		bm := b.(map[string]interface{})
	//
	//		bl := BlacklistLoanRecord{
	//			CompanyName: bm["company_name"].(string),
	//			Name:        bm["name"].(string),
	//			Phone:       bm["phone"].(string),
	//			Email:       bm["email"].(string),
	//			Bvn:         bm["bvn"].(string),
	//			Gender:      bm["gender"].(string),
	//			LoanAmount:  bm["loan_amount"].(string),
	//			AmountPaid:  bm["amount_paid"].(string),
	//			Balance:     bm["balance"].(string),
	//			DueDate:     bm["due_date"].(string),
	//			Location:    bm["location"].(string),
	//			Date:        bm["Date"].(string),
	//		}
	//
	//		blr = append(blr, bl)
	//	}
	//	return
	//}

	err = err2.NewClientErr(nil,
		"internal error 5", 500)

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

	if r.Status == "error" {
		fmt.Println("error here 2") //debug delete
		err = err2.NewClientErr(nil, r.Message, 400)
		return
	}

	if a := r.Data.CRC[0]; a.Status == "error" {
		fmt.Println("error here 2") //debug delete
		err = err2.NewClientErr(nil, a.Message, 400)
		return
	} else {
		cd = a.Data
		return
	}

}
