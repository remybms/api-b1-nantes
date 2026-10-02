package models

type Pokemon struct {
	Name            string `json:"name"`
	Image           string `json:"image"`
	Description     string `json:"description"`
	FirstType       string `json:"first_type"`
	SecondType      string `json:"second_type"`
	Generation      int    `json:"generation"`
	SignatureAttack string `json:"signature_attack"`
}
