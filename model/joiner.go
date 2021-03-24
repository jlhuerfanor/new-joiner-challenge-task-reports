package model

type Joiner struct {
	IdNumber string `json:"idNumber"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Stack string `json:"stack"`
	Role string `json:"role"`
	EnglishLevel string `json:"englishLevel"`
	DomainExperience string `json:"domainExperience"`
}