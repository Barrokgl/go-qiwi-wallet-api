package goqiwi

import "time"

type GetProfileParams struct {
	AuthInfoEnabled     bool `json:"authInfoEnabled" url:"authInfoEnabled"`
	ContractInfoEnabled bool `json:"contractInfoEnabled" url:"contractInfoEnabled"`
	UserInfoEnabled     bool `json:"userInfoEnabled" url:"userInfoEnabled"`
}

type GetProfileResult struct {
	AuthInfo     AuthInfo     `json:"authInfo"`
	ContractInfo ContractInfo `json:"contractInfo"`
	UserInfo     UserInfo     `json:"userInfo"`
}

type AuthInfo struct {
	PersonId         int64         `json:"personId"`
	RegistrationDate string        `json:"registrationDate"`
	BoundEmail       string        `json:"boundEmail"`
	IP               string        `json:"ip"`
	LastLoginDate    string        `json:"lastLoginDate"`
	MobilePinInfo    MobilePinInfo `json:"mobilePinInfo"`
	PassInfo         PassInfo      `json:"passInfo"`
}

type MobilePinInfo struct {
	LastMobilePinChange string `json:"lastMobilePinChange"`
	MobilePinUsed       bool   `json:"mobilePinUsed"`
	NextMobilePinChange string `json:"nextMobilePinChange"`
}

type PassInfo struct {
	LastPassChange string `json:"lastPassChange"`
	NextPassChange string `json:"nextPassChange"`
	PasswordUsed   bool   `json:"passwordUsed"`
}

type PinInfo struct {
	PinUsed bool `json:"pinUsed"`
}

type ContractInfo struct {
	Blocked            bool                 `json:"blocked"`
	ContractID         int64                `json:"contractId"`
	CreationDate       string               `json:"creationDate"`
	Features           []interface{}        `json:"features"`
	IdentificationInfo []IdentificationInfo `json:"identificationInfo"`
}

type IdentificationInfo struct {
	BankAlias           string `json:"bankAlias"`
	IdentificationLevel string `json:"identificationLevel"`
}

type UserInfo struct {
	DefaultPayCurrency int    `json:"defaultPayCurrency"`
	DefaultPaySource   int    `json:"defaultPaySource"`
	Email              string `json:"email"`
	FirstTxnID         int64  `json:"firstTxnId"`
	Language           string `json:"language"`
	Operator           string `json:"operator"`
	PhoneHash          string `json:"phoneHash"`
	PromoEnabled       string `json:"promoEnabled"`
}

type GetHistoryParams struct {
	Rows        int       `json:"rows" url:"rows"`
	Operation   string    `json:"operation" url:"operation"`
	Sources     []string  `json:"sources" url:"sources"`
	StartDate   time.Time `json:"startDate" url:"startDate"`
	EndDate     time.Time `json:"endDate" url:"endDate"`
	NextTxnDate time.Time `json:"nextTxnDate" url:"nextTxnDate"`
	NextTxnId   int64     `json:"nextTxnId" url:"nextTxnId"`
}

type GetHistoryResult struct {
	Data        []Transaction `json:"data"`
	NextTxnId   int64         `json:"nextTxnId"`
	NextTxnDate string        `json:"nextTxnDate"`
}

type Transaction struct {
	TxnId                  int64         `json:"txnId"`
	PersonId               int64         `json:"personId"`
	Date                   string        `json:"date"`
	ErrorCode              int           `json:"errorCode"`
	Error                  string        `json:"error"`
	Status                 string        `json:"status"`
	Type                   string        `json:"type"`
	StatusText             string        `json:"statusText"`
	TrmTxnId               int64         `json:"trmTxnId"`
	Account                string        `json:"account"`
	Sum                    Sum           `json:"sum"`
	Commission             Sum           `json:"commission"`
	Total                  Sum           `json:"total"`
	Provider               Provider      `json:"provider"`
	Comment                string        `json:"comment"`
	CurrencyRate           float64       `json:"currencyRate"`
	Extras                 []interface{} `json:"extras"`
	ChequeReady            bool          `json:"chequeReady"`
	BankDocumentAvailable  bool          `json:"bankDocumentAvailable"`
	BankDocumentReady      bool          `json:"bankDocumentReady"`
	RepeatPaymentEnabled   bool          `json:"repeatPaymentEnabled"`
	FavoritePaymentEnabled bool          `json:"favoritePaymentEnabled"`
	RegularPaymentEnabled  bool          `json:"regularPaymentEnabled"`
}

type Sum struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type Provider struct {
	ID          int64  `json:"id"`
	ShortName   string `json:"shortName"`
	LongName    string `json:"longName"`
	LogoUrl     string `json:"logoUrl"`
	Description string `json:"description"`
	Keys        string `json:"keys"`
	SiteUrl     string `json:"siteUrl"`
}
