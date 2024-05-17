package models

type Document struct {
	Id                int    `json:"id"`
	DocumentFile      string `json:"document_file"`
	DocumentPath      string `json:"document_path"`
	Status            string `json:"status"`
	NoCreditSalesForm string `json:"no_credit_sales_form"`
	DateOfLetter      string `json:"date_of_letter"`
	DateOfReceipt     string `json:"date_of_receipt"`
}
