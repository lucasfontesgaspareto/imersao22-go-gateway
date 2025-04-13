package dto

import (
	"time"

	"github.com/lucasfontesgaspareto/imersao22/go-gateway/internal/domain"
)

const (
	StatusPending  = string(domain.StatusPending)
	StatusApproved = string(domain.StatusApproved)
	StatusRejected = string(domain.StatusRejected)
)

type CreateInvoiceInput struct {
	APIKey         string
	AccountID      string  `json:"account_id"`
	Amount         float64 `json:"amount"`
	Description    string  `json:"description"`
	PaymentType    string  `json:"payment_type"`
	CardNumber     string  `json:"card_number"`
	CVV            string  `json:"cvv"`
	ExpiryMonth    int     `json:"expiry_month"`
	ExpiryYear     int     `json:"expiry_year"`
	CardholderName string  `json:"cardholder_name"`
}

type InvoiceOutput struct {
	ID             string    `json:"id"`
	AccountID      string    `json:"account_id"`
	Amount         float64   `json:"amount"`
	Status         string    `json:"status"`
	Description    string    `json:"description"`
	PaymentType    string    `json:"payment_type"`
	CardLastDigits string    `json:"card_last_digits"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func ToInvoice(input CreateInvoiceInput, accountID string) (*domain.Invoice, error) {
	card := domain.CreditCard{
		Number:         input.CardNumber,
		CVV:            input.CVV,
		ExpiryMonth:    input.ExpiryMonth,
		ExpiryYear:     input.ExpiryYear,
		CardholderName: input.CardholderName,
	}

	return domain.NewInvoice(
		accountID,
		input.Amount,
		input.Description,
		input.PaymentType,
		card,
	)
}

func FromInvoice(Invoice *domain.Invoice) *InvoiceOutput {
	return &InvoiceOutput{
		ID:             Invoice.ID,
		AccountID:      Invoice.AccountID,
		Amount:         Invoice.Amount,
		Status:         string(Invoice.Status),
		Description:    Invoice.Description,
		PaymentType:    Invoice.PaymentType,
		CardLastDigits: Invoice.CardLastDigits,
		CreatedAt:      Invoice.CreatedAt,
		UpdatedAt:      Invoice.UpdatedAt,
	}
}
