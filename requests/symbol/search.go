package requests_symbol

type SearchRequest struct {
	Symbol string `json:"symbol" validate:"required,alphanum,uppercase"`
}
