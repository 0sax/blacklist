package blacklist

type BlacklistBVNResult struct {
	Status  string `json:"status"`
	Message string `json:"message"`

	Data interface{} `json:"data"`
}

type BlacklistLoanRecord struct {
	CompanyName string `json:"company_name"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Bvn         string `json:"bvn"`
	Gender      string `json:"gender"`
	LoanAmount  string `json:"loan_amount"`
	AmountPaid  string `json:"amount_paid"`
	Balance     string `json:"balance"`
	DueDate     string `json:"due_date"`
	Location    string `json:"location"`
	Date        string `json:"Date"`
}

type CRCBVNResult struct {
	//blacklist: null
	CRC []struct {
		Status  string   `json:"status"`
		Message string   `json:"message"`
		Data    *CRCData `json:"data"`
	} `json:"crc"`
}

type CRCData struct {
	MFBSummary          CRCSummary `json:"MFCREDIT_NANO_SUMMARY"`
	MortgageSummary     CRCSummary `json:"MGCREDIT_NANO_SUMMARY"`
	CreditNanoSummary   CRCSummary `json:"CREDIT_NANO_SUMMARY"`
	NanoConsumerProfile struct {
		Citizenship string `json:"CITIZENSHIP"`
		Dob         string `json:"DATE_OF_BIRTH"`
		FirstName   string `json:"FIRST_NAME"`
		LastName    string `json:"LAST_NAME"`
		Gender      string `json:"GENDER"`
	} `json:"NANO_CONSUMER_PROFILE"`
	ReportHeader struct {
		MailTo      string        `json:"MAILTO"`
		ProductName string        `json:"PRODUCTNAME"`
		Reason      []interface{} `json:"REASON"`
		ReportDate  string        `json:"REPORTDATE"`
	} `json:"REPORTHEADER"`
}

type CRCSummary struct {
	Summary struct {
		HasLoans         string `json:"HAS_CREDITFACILITIES"`
		LastReportUpdate string `json:"LAST_REPORTED_DATE"`
		DelinquentLoans  string `json:"NO_OF_DELINQCREDITFACILITIES"`
	} `json:"SUMMARY"`
}
