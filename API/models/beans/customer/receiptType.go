package customer

type ReceiptTypeTree struct {
	ID       int64              `description:"ID"`
	Name     string             `description:"名称"`
	Children []*ReceiptTypeTree `description:"子类" json:"children,omitempty"`
}

type ReceiptTypeAdd struct {
	ParentID int64  `description:"父级ID"`
	Name     string `description:"名称"`
}
