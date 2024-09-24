package response

type AccountResponse struct {
	AccountId      string `json:"account_id,omitempty"`
	DocumentNumber string `json:"document_number,omitempty"`
}
