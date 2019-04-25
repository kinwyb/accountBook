package customer

type BankListCompateResp struct {
	BankID       int64   `description:"银行ID"`
	BankName     string  `description:"银行名称"`
	SCNY         float64 `description:"期初人民币"`
	ECNY         float64 `description:"期末人民币"`
	SUSD         float64 `description:"期初美金"`
	EUSD         float64 `description:"期末美金"`
	BankAccount  string  `description:"银行账户"`
	Contacts     string  `description:"联系人"`
	ContactPhone string  `description:"联系电话"`
}
