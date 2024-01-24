package payment_api

import (
	"github.com/daqing/waterway/lib/repo"
)

type Payment struct {
	repo.Model

	UserId    repo.IdType
	UUID      string
	GoodsType string
	GoodsId   repo.IdType
	Cent      repo.PriceCent
	Action    string
	Note      string
	Status    PaymentStatus
}

const tableName = "payments"

func (m Payment) TableName() string { return tableName }

type PaymentStatus string

const FreshStatus PaymentStatus = "fresh"
const PaidStatus PaymentStatus = "paid"
