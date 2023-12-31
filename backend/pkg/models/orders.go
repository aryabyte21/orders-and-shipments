package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	BusinessID          primitive.ObjectID `bson:"businessId" json:"-"`
	OrderTotal          string             `bson:"total_price"`
	OrderNumber         int64              `bson:"order_number"`
	CustomerName        Customer           `bson:"customer"`
	CustomerEmail       string             `bson:"email"`
	CustomerOrderNumber string             `bson:"reference"`
	CustomerPhone       string             `bson:"customer.default_address.phone"`
	CustomerAddress     string             `bson:"shipping_address.address1"`
	ProviderOrderId     float64            `bson:"id"`
	CreatedAt           time.Time          `bson:"created_at"`
	ProviderCreatedAt   time.Time          `bson:"fulfillments.0.created_at"`
	Products            []Product          `bson:"line_items"`
	UniqueId            string             `bson:"token"`
	StripePaymentID     string             `bson:"stripe_payment_id,omitempty" json:"stripe_payment_id,omitempty"`
}

type OrderData struct {
	Business    Business     `json:"business"`
	Order       Order        `json:"order"`
	Shipments   []Shipment   `json:"shipments"`
	Checkpoints []Checkpoint `json:"checkpoints"`
}

type Product struct {
	Name  string `bson:"name"`
	Price string `bson:"price"`
}

type Customer struct {
	FirstName      string  `bson:"first_name"`
	LastName       string  `bson:"last_name"`
	Currency       string  `bson:"currency"`
	DefaultAddress Address `bson:"default_address"`
}

type Address struct {
	Phone    string `bson:"phone"`
	Address1 string `bson:"address1"`
	City     string `bson:"city"`
	Province string `bson:"province"`
	Zip      string `bson:"zip"`
	Country  string `bson:"country_name"`
}
