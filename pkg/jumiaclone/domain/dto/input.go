package dto

// Recipient returns the details of a message recipient
type Recipient struct {
	Number    string `json:"number"`
	Cost      string `json:"cost"`
	Status    string `json:"status"`
	MessageID string `json:"messageID"`
}

// SMS returns the message details of a recipient
type SMS struct {
	Recipients []Recipient `json:"Recipients"`
}

// SendMessageResponse returns a message response with the recipient's details
type SendMessageResponse struct {
	SMSMessageData *SMS `json:"SMSMessageData"`
}
