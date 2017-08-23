package goqiwi

type ProfileParams struct {
	AuthInfoEnabled     bool `json:"authInfoEnabled" url:"authInfoEnabled"`
	ContractInfoEnabled bool `json:"contractInfoEnabled" url:"contractInfoEnabled"`
	UserInfoEnabled     bool `json:"userInfoEnabled" url:"userInfoEnabled"`
}

type Profile struct {
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

// all dates should be formatted in RFC3339
type HistoryParams struct {
	Rows        int      `json:"rows" url:"rows"`
	Operation   string   `json:"operation" url:"operation"`
	Sources     []string `json:"sources" url:"sources"`
	StartDate   string   `json:"startDate" url:"startDate"`
	EndDate     string   `json:"endDate" url:"endDate"`
	NextTxnDate string   `json:"nextTxnDate" url:"nextTxnDate"`
	NextTxnId   int64    `json:"nextTxnId" url:"nextTxnId"`
}

type History struct {
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

// all dates should be formatted in RFC3339
type PaymentStatisticParams struct {
	StartDate string   `json:"startDate" url:"startDate"`
	EndDate   string   `json:"endDate" url:"endDate"`
	Operation string   `json:"operation" url:"operation"`
	Sources   []string `json:"sources" url:"sources"`
}

type PaymentStatistic struct {
	IncomingTotal []Sum `json:"incomingTotal"`
	OutgoingTotal []Sum `json:"outgoingTotal"`
}

type Balances struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	Alias      string `json:"alias"`
	FsAlias    string `json:"fsAlias"`
	Title      string `json:"title"`
	HasBalance bool   `json:"hasBalance"`
	Currency   string `json:"currency"`
	Type       Type   `json:"type"`
	Balance    Sum    `json:"balance"`
}

type Type struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type StandardRate struct {
	Content Content `json:"content"`
}

type Content struct {
	Terms Terms `json:"terms"`
}

type Terms struct {
	Cashbacks         []interface{}  `json:"cashbacks"`
	Commission        Commission     `json:"commission"`
	Description       string         `json:"description"`
	ID                string         `json:"id"`
	Identification    Identification `json:"identification"`
	Limits            []Limit        `json:"limits"`
	Overpayment       bool           `json:"overpayment"`
	RepeatablePayment bool           `json:"repeatablePayment"`
	Type              string         `json:"type"`
	Underpayment      bool           `json:"underpayment"`
}

type Identification struct {
	Required bool `json:"required"`
}

type Limit struct {
	Currency string  `json:"currency"`
	Max      float64 `json:"max"`
	Min      float64 `json:"min"`
}

type Commission struct {
	Ranges []Range `json:"ranges"`
}

type Range struct {
	Bound float32 `json:"bound"`
	Rate  float32 `json:"rate"`
	Min   float32 `json:"min"`
	Max   float32 `json:"max"`
	Fixed float32 `json:"fixed"`
}

type SpecialRateParams struct {
	Account       string        `json:"account"`
	PaymentMethod PaymentMethod `json:"paymentMethod"`
	PurchaseTotal PurchaseTotal `json:"purchaseTotal"`
}

type PurchaseTotal struct {
	Total Sum `json:"total"`
}

type PaymentMethod struct {
	Type      string `json:"type"`
	AccountId string `json:"accountId"`
}

type SpecialRate struct {
	ProviderId               string `json:"providerId"`
	WithdrawSum              Sum    `json:"withdrawSum"`
	EnrollmentSum            Sum    `json:"enrollmentSum"`
	QwCommission             Sum    `json:"qwCommission"`
	FundingSourceCommission  Sum    `json:"fundingSourceCommission"`
	WithdrawToEnrollmentRate int    `json:"withdrawToEnrollmentRate"`
}
