package client

type Client struct {
	ClientName    string   `json:"client_name"`
	Email         string   `json:"email"`
	Company       string   `json:"company"`
	ContactNumber string   `json:"contact_number"`
	Address       string   `json:"address"`
	Projects      []string `json:"projects"`
}
