package dto

type CreateAutoRequest struct{
	Marca string	 `json:"marca"`
	Modello string	 `json:"modello"`
	Targa string	 `json:"targa"`
}

type UpdateAutoRequest struct {
	Marca        string  `json:"marca"`
	Modello      string  `json:"modello"`
	Disponibile  bool    `json:"disponibile"`
}