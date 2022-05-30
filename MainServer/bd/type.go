package bd

import "encoding/json"

type Order struct {
	OrderUid string 		    `json:"order_uid" db:"order_uid"`
	TrackNumber string 		  `json:"track_number" db:"track_number"`
	Entry string 			      `json:"entry"`
	Delivery delivery 		  `json:"delivery"`
	Payment payment 		    `json:"payment"`
	Items []items 			    `json:"items"`
	Locale string 			    `json:"locale"`
	InternalSignature string`json:"internal_signature" db:"internal_signature"`
	CustomerId string		    `json:"customer_id" db:"customer_id"`
  DeliveryService string	`json:"delivery_service" db:"delivery_service"`
  Shardkey string			    `json:"shardkey"`
  SmId int64				      `json:"sm_id" db:"sm_id"`
	DateCreated string		  `json:"date_created" db:"date_created"`
	OofShard string			    `json:"oof_shard" db:"oof_shard"`


}

type delivery struct{
	Name string				`json:"name"`
  Phone string			`json:"phone"`
  Zip string				`json:"zip"`
  City string				`json:"city"`
  Address string		`json:"address"`
  Region string			`json:"region"`
  Email string			`json:"email"`
}
type payment struct{
	Transaction string	`json:"transaction"`
  RequestId string		`json:"request_id" db:"request_id"`
  Currency string			`json:"currency"`
  Provider string			`json:"provider"`
  Amount int64			  `json:"amount"`
  PaymentDt int64			`json:"payment_dt" db:"payment_dt"`
  Bank string				  `json:"bank"`
  DeliveryCost int64	`json:"delivery_cost" db:"delivery_cost"`
  GoodsTotal int64		`json:"goods_total" db:"goods_total"`
	CustomFee int64			`json:"custom_fee" db:"custom_fee"`
}
type items struct{
	ChrtId int64			`json:"chrt_id" db:"chrt_id"`
  TrackNumber string`json:"track_number" db:"track_number"`
  Price int64				`json:"price"`
  Rid string				`json:"rid"`
  Name string				`json:"name"`
  Sale int				  `json:"sale"`
  Size string				`json:"size"`
  TotalPrice int64	`json:"total_price" db:"total_price"`
  NmId int64				`json:"nm_id" db:"nm_id"`
  Brand string			`json:"brand"`
  Status int64			`json:"status"`
}

func (o *Order) String() string{
  data, _ := json.MarshalIndent(o, "", " ") 
  return string(data)
}