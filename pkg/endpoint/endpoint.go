package endpoint

// HashRequest specifies the request parameters for Hash method
type HashRequest struct {
	Password string `json:"password"`
}

// HashResponse specifies the response parameters for Hash method
type HashResponse struct {
	Hash string `json:"hash"`
	Err  string `json:"err,omitempty"`
}

// ValidateRequest specifies the request parameters for validate method
type ValidateRequest struct {
	Password string `json:"password"`
	Hash     string `json:"hash"`
}

// ValidateResponse specifies the response parameters for Validate method
type ValidateResponse struct {
	Valid bool   `json:"valid"`
	Err   string `json:"err,omitempty"`
}
