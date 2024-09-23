package response

type UserAccountResponse struct {
	AccountId      string `json:"account_id,omitempty"`
	DocumentNumber string `json:"document_number,omitempty"`
}
