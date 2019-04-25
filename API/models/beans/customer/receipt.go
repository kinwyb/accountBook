package customer

type ReceiptListReq struct {
	StartTime   string `description:"起始时间"`
	EndTime     string `description:"结束时间"`
	BankID      int64  `description:"银行ID"`
	ShopID      int64  `description:"店铺ID"`
	ReceiptType string `description:"收支类型ID"`
	Search      string `description:"搜索"`
	BillType    int64  `description:"单据类型"`
}

type ReceiptListResp struct {
	Data []*Receipt `description:"数据集合"`
}

type Receipt struct {
	Id          string  `description:"单据号"`
	Money       float64 `description:"交易金额"`
	Bank        string  `description:"银行"`
	Description string  `description:"备注"`
	Createtime  string  `description:"生成时间"`
	Operator    string  `description:"操作者"`
	Shop        string  `description:"店铺"`
	Type        string  `description:"单据类型"`
	MoneyType   string  `description:"货币类型"`
}

// 指定结束时间点金额
type ReceiptEndTimeMoneyCount struct {
	EndTime   string  `description:"结束时间"`
	BankID    int64   `description:"银行ID"`
	MoneyType string  `description:"货币类型"`
	Money     float64 `description:"金额"`
}
