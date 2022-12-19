package request

import "encoding/json"

type Innova struct {
	InnovaOrderID            int     `json:"innova_order_id,omitempty"`
	CardNumber               string  `json:"card_number"`
	OrderNumber              string  `json:"order_number"`
	RequestNumber            string  `json:"request_number"`
	PaymentRef               string  `json:"payment_ref "`
	TotalAmountAfterDiscount float64 `json:"total_amount_after_discount"`
	RewardCode               string  `json:"reward_code"`
}

func (innova *Innova) Marshal() string {
	marshalIndent, _ := json.MarshalIndent(innova, "", "    ")
	return string(marshalIndent)
}
