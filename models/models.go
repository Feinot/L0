package models

import "github.com/go-playground/validator/v10"

var (
	Cluster = "test-cluster"
	Prod    = "test-prod"
	Sub     = "test-su"
	Channel = "testcans"
)

type Order struct {
	OrderUID    string `json:"order_uid" validate:"required"`
	TrackNumber string `json:"track_number" `
	Entry       string `json:"entry" `

	Delivery Delivery `json:"delivery"`
	Payment  Payment  `json:"payment"`
	Items    []Item   `json:"items"`

	Locale            string `json:"locale" `
	InternalSignature string `json:"internal_signature" `
	CustomerID        string `json:"customer_id" `
	DeliveryService   string `json:"delivery_service" `
	ShardKey          string `json:"shardkey" `
	SmID              int    `json:"sm_id" validate:"numeric"`
	DateCreated       string `json:"date_created" `
	OofShard          string `json:"oof_shard" `
}

type Delivery struct {
	ID      int    `DB:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
	Order   string `db:"order_uid"`
}

type Payment struct {
	ID           int    `DB:"id"`
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount" validate:"numeric"`
	PaymentDt    int    `json:"payment_dt" validate:"numeric"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost" validate:"numeric"`
	GoodsTotal   int    `json:"goods_total" validate:"numeric"`
	CustomFee    int    `json:"custom_fee" validate:"numeric"`
	Order        string `db:"order_uid"`
}

type Item struct {
	ID          int    `DB:"id"`
	ChrtID      int    `json:"chrt_id" validate:"numeric"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price" validate:"numeric"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale" validate:"numeric"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price" validate:"numeric"`
	NmID        int    `json:"nm_id" validate:"numeric"`
	Brand       string `json:"brand"`
	Status      int    `json:"status" validate:"numeric"`
	Order       string `db:"order_uid"`
}

var (
	QOrder    = `INSERT INTO orders(uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shard_key, sm_id, date_created, oof_shard) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	QDelivery = `INSERT INTO delivery(name, phone, zip, city, address, region, email, order_uid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	QPayment  = `INSERT INTO payment(transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee, order_uid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	QItems = `INSERT INTO item(chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status, order_uid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
)
var (
	SOrder    = `select * from orders`
	SDelivery = `select * from delivery where order_uid = $1`
	SPayment  = `select * from payment where order_uid = $1`
	SItem     = `select * from item where order_uid = $1`
)

func (o *Order) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}
