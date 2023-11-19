package db

type User struct {
	UUID     string `gorm:"primaryKey" json:"uuid"`
	UserName string `gorm:"unique" json:"user_name"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "users"
}

type RateRecord struct {
	UUID         string  `gorm:"primaryKey" json:"uuid"`
	DepoRateLive float64 `json:"depo_rate_live"`
	DepoRate3    float64 `json:"depo_rate_3"`
	DepoRate6    float64 `json:"depo_rate_6"`
	DepoRate12   float64 `json:"depo_rate_12"`
	DepoRate24   float64 `json:"depo_rate_24"`
	DepoRate36   float64 `json:"depo_rate_36"`
	DepoRate60   float64 `json:"depo_rate_60"`
	LoanRate6    float64 `json:"loan_rate_6"`
	LoanRate12   float64 `json:"loan_rate_12"`
	LoanRate1236 float64 `json:"loan_rate_12_36"`
	LoanRate3660 float64 `json:"loan_rate_36_60"`
	LoanRate60   float64 `json:"loan_rate_60"`
}

func (RateRecord) TableName() string {
	return "rate_records"
}

type CalcHistory struct {
	HistId     string `gorm:"primaryKey" json:"id"`
	Time       int64  `json:"time"`
	Uuid       string `json:"uuid"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
}

func (CalcHistory) TableName() string {
	return "calc_histories"
}
