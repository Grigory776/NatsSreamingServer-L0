package bd

import (
	"log"

	_"github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

func Connect(url string) (*sqlx.DB, error){
	conn, err := sqlx.Connect("postgres",url)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func AddEntry(conn *sqlx.DB, ord *Order){
	tmp := []interface{}{
		ord.Delivery.Name, ord.Delivery.Phone, ord.Delivery.Zip, ord.Delivery.City,
		ord.Delivery.Address, ord.Delivery.Region, ord.Delivery.Email,
	}
	req := `INSERT INTO delivery(name, phone, zip, city, address, region, email)
	VALUES ($1,$2,$3,$4,$5,$6,$7)`
	_,err := conn.Exec(req, tmp...)
	if err != nil{
		log.Println(err)
	}
	var delID int 
	req = `SELECT deliveryID FROM delivery ORDER BY deliveryID DESC LIMIT 1`
	err = conn.QueryRow(req).Scan(&delID) 
	if err != nil{
		log.Println(err)
	}

	tmp2 := []interface{}{
		ord.Payment.Transaction, ord.Payment.RequestId, ord.Payment.Currency,
		ord.Payment.Provider, ord.Payment.Amount, ord.Payment.PaymentDt, 
		ord.Payment.Bank, ord.Payment.DeliveryCost, ord.Payment.GoodsTotal,
		ord.Payment.CustomFee,
	}
	req = `INSERT INTO payment(transaction, request_id, currency, provider, amount, payment_dt,
		bank, delivery_cost, goods_total, custom_fee) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`
	_,err = conn.Exec(req, tmp2...)
	if err != nil{
		log.Println(err)
	}
	var paymID int
	req = `SELECT paymentID FROM payment ORDER BY paymentID DESC LIMIT 1`
	err = conn.QueryRow(req).Scan(&paymID) 
	if err != nil{
		log.Println(err)
	}

	tmp3 := [][]interface{}{}
	for _, val := range ord.Items{
		X := []interface{}{
			val.ChrtId, val.TrackNumber, val.Price, val.Rid,
			val.Name, val.Sale, val.Size, val.TotalPrice,
			val.NmId, val.Brand, val.Status,
		}
		tmp3 = append(tmp3, X)
	}
	req = `INSERT INTO items(chrt_id, track_number, price, rid, name, sale, size, total_price,
	nm_id, brand, status) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`
	for ind := range ord.Items{
		_,err = conn.Exec(req, tmp3[ind]...)
		if err != nil {
			log.Println(err)
		}
	}
	var itemID int // chrt_id должен быть один для элементов одного массива (для правильной выборки)
	req = `SELECT itemsID FROM items ORDER BY itemsID DESC LIMIT 1`
	err = conn.QueryRow(req).Scan(&itemID) 
	if err != nil{
		log.Println(err)
	}

	tmp4 := []interface{}{
		ord.OrderUid, ord.TrackNumber, ord.Entry, delID, paymID, itemID,
		ord.Locale, ord.InternalSignature, ord.CustomerId, ord.DeliveryService,
		ord.Shardkey, ord.SmId, ord.DateCreated, ord.OofShard,
	}
	req = `INSERT INTO orders(order_uid, track_number, entry, deliveryID, paymentID,
		itemsID, locale, internal_signature, customer_id, delivery_service, 
		shardkey, sm_id, date_created, oof_shard) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,
			$12,$13,$14)`
	_,err = conn.Exec(req, tmp4...)
	if err != nil{
		log.Println(err)
	}
}

func GetEntry(conn *sqlx.DB, uid string)(*Order,bool){
	var tmp string
	req := `SELECT order_uid FROM orders WHERE order_uid = $1`
	_ = conn.QueryRow(req, uid).Scan(&tmp) 
	if tmp == ""{
		return nil, false
	}

	var (
		delID, paymID, itemID int
		ord Order
	)

	req = `SELECT deliveryID, paymentID, itemsID FROM orders WHERE order_uid = $1`
	err := conn.QueryRow(req, uid).Scan(&delID,&paymID,&itemID) 
	if err != nil {
		log.Println(err)
	}

	req = `SELECT name, phone, zip, city, address, region, email FROM delivery WHERE deliveryID = $1`
	q, err := conn.Queryx(req, delID) 
	if err != nil {
		log.Println(err)
	}
	for q.Next(){
		err = q.StructScan(&ord.Delivery)
		if err != nil {
			log.Println(err)
		}	
	}

	req = `SELECT transaction, request_id, currency, provider, amount, payment_dt,
	bank, delivery_cost, goods_total, custom_fee FROM payment WHERE paymentID = $1`
	q, err = conn.Queryx(req, paymID) 
	if err != nil {
		log.Println(err)
	}
	for q.Next(){
		err = q.StructScan(&ord.Payment)
		if err != nil {
			log.Println(err)
		}	
	}

	req = `SELECT chrt_id, track_number, price, rid, name, sale, size, total_price,
		nm_id, brand, status 
		FROM items 
		WHERE chrt_id = (SELECT chrt_id FROM items WHERE itemsid=$1)`
	q, err = conn.Queryx(req, itemID) 
	if err != nil {
		log.Println(err)
	}
	var tmpItem items
	for q.Next(){
		err = q.StructScan(&tmpItem)
		if err != nil {
			log.Println(err)
		}
		ord.Items = append(ord.Items, tmpItem)	
	}

	req = `SELECT order_uid, track_number, entry,
	locale, internal_signature, customer_id, delivery_service, 
	shardkey, sm_id, date_created, oof_shard FROM orders WHERE order_uid = $1`
	q, err = conn.Queryx(req, uid) 
	if err != nil {
		log.Println(err)
	}
	for q.Next(){
		err = q.StructScan(&ord)
		if err != nil {
			log.Println(err)
		}	
	}
	
	return &ord,true
} 

func GetAllUid(conn *sqlx.DB) ([]string,error){
	var (
		tmp string
		res []string
	)

	req := `SELECT order_uid FROM orders`
	q, err := conn.Queryx(req) 
	if err != nil {
		return nil, err
	}
	for q.Next(){
		err = q.Scan(&tmp)
		if err != nil {
			log.Println(err)
		}	
		res = append(res, tmp)
	}

	return res, nil
}

func Close(conn *sqlx.DB){
	conn.Close()
}