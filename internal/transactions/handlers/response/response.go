package response

type NewTransactionResponse struct {
	TransactionId   string  `json:"transaction_id,omitempty"`
	AccountId       string  `json:"account_id,omitempty"`
	OperationTypeId int     `json:"operation_type_id,omitempty"`
	Amount          float32 `json:"amount,omitempty"`
}
