package models

type GeneralInformation struct {
	Id                    int    `json:"id"`
	BankName              string `json:"bank_name"`
	KCP                   string `json:"kcp"`
	SubProgram            string `json:"sub_program"`
	Analisis              string `json:"analisis"`
	CabangPencairan       string `json:"cabang_pencairan"`
	CabangAdmin           string `json:"cabang_admin"`
	TglAplikasi           string `json:"tgl_aplikasi"`
	TglPenerusan          string `json:"tgl_penerusan"`
	Segmen                string `json:"segmen"`
	NoAplikasi            int    `json:"no_aplikasi"`
	MarketInterestRate    int    `json:"market_interest_rate"`
	RequestedInterestRate int    `json:"requested_interest_rate"`
	DocumentFile          string `json:"document_file"`
	Status                string `json:"status"`
}
