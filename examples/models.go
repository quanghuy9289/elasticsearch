package main

type Post struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Text      string   `json:"text"`
	Tags      []string `json:"tags"`
	CreatedAt int64    `json:"createdAt,omitempty"`
}

type BankAccount struct {
	AccountNumber int    `json:"account_number"`
	Balance       int64  `json:"balance"`
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	Age           int    `json:"age"`
	Gender        string `json:"gender"`
	Address       string `json:"address"`
	Employer      string `json:"employer"`
	Email         string `json:"email"`
	City          string `json:"city"`
	State         string `json:"state"`
}
