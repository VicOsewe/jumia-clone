package dto

type User struct {
	UID         string `json:"uid,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	MiddleName  string `json:"middle_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Email       string `json:"email,omitempty"`
	Verified    bool   `json:"verified,omitempty"`
	PassWord    string `json:"pass_word,omitempty"`
}
