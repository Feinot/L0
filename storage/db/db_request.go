package db

import (
	"log"
	"nats/models"
	"nats/storage/cashe"
)

func Recover() {

	cashe.Init()

	orderRows, err := db.Query(models.SOrder)
	if err != nil {
		log.Println(err)
	}

	defer orderRows.Close()

	for orderRows.Next() {
		order := models.Order{}

		orderRows.Scan(&order.OrderUID, &order.TrackNumber, &order.Entry, &order.Locale,
			&order.InternalSignature, &order.CustomerID, &order.DeliveryService,
			&order.ShardKey, &order.SmID, &order.DateCreated, &order.OofShard)

		delivery := models.Delivery{}

		deliveryRows, err := db.Query(models.SDelivery, order.OrderUID)
		if err != nil {
			log.Println(err)
		}
		defer deliveryRows.Close()

		for deliveryRows.Next() {
			deliveryRows.Scan(&delivery.ID, &delivery.Name, &delivery.Phone, &delivery.Zip,
				&delivery.City, &delivery.Address, &delivery.Region, &delivery.Email, &delivery.Order)
		}
		order.Delivery = delivery

		payment := models.Payment{}
		paymentRows, err := db.Query(models.SPayment, order.OrderUID)
		if err != nil {
			log.Println(err)
		}
		defer paymentRows.Close()

		for paymentRows.Next() {

			paymentRows.Scan(&payment.ID, &payment.Transaction, &payment.RequestID, &payment.Currency,
				&payment.Provider, &payment.Amount, &payment.PaymentDt, &payment.Bank,
				&payment.DeliveryCost, &payment.GoodsTotal, &payment.CustomFee, &payment.Order)
		}
		order.Payment = payment

		items := make([]models.Item, 0)

		itemRows, err := db.Query(models.SItem, order.OrderUID)
		if err != nil {
			log.Println(err)
		}
		defer itemRows.Close()

		for itemRows.Next() {
			i := models.Item{}
			itemRows.Scan(&i.ID, &i.ChrtID, &i.TrackNumber, &i.Price, &i.Rid, &i.Name,
				&i.Sale, &i.Size, &i.TotalPrice, &i.NmID, &i.Brand, &i.Status, &i.Order)
			items = append(items, i)
		}
		order.Items = items
		cashe.Set(order)

	}

}
