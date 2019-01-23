package evttypes

type TRXJson struct {
	Expiration            string         `json:"expiration"`
	RefBlockNum           int            `json:"ref_block_num"`
	RefBlockPrefix        int            `json:"ref_block_prefix"`
	MaxCharge             int            `json:"max_charge"`
	Payer                 string         `json:"payer"`
	Actions               []SimpleAction `json:"actions"`
	TransactionExtensions []interface{}  `json:"transaction_extensions"`
}
