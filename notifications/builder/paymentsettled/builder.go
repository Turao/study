package paymentsettled

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
	apiV1 "github.com/turao/topics/notifications/api/v1"
	"github.com/turao/topics/notifications/entity/notification"
)

type builder struct{}

// Newbuilder creates a new notification builder
func NewBuilder() *builder {
	return &builder{}
}

func (b *builder) NotificationType() string {
	return "payment_settled"
}

// SendNotification sends a notification
func (b *builder) BuildNotification(ctx context.Context, request apiV1.SendNotificationRequest) (*notification.Notification, error) {
	paymentDetails, err := b.fetchPaymentDetails(ctx, request.PaymentSettled.PaymentID)
	if err != nil {
		return nil, err
	}

	payerDetails, err := b.fetchPayerDetails(ctx, paymentDetails.PayerID)
	if err != nil {
		return nil, err
	}

	return &notification.Notification{
		ID:        uuid.Must(uuid.NewV4()).String(),
		Type:      b.NotificationType(),
		Recipient: request.Recipient,
		Subject:   "Payment settled!",
		Content:   b.buildContent(paymentDetails, payerDetails),
		Metadata:  b.buildMetadata(paymentDetails, request.Metadata),
		CreatedAt: time.Now(),
	}, nil
}

type paymentDetails struct {
	PaymentID                  string
	PayerID                    string
	PriceE5                    decimal.Decimal
	PriceCurrency              string
	PaymentMethodDisplayName   string
	PaymentMethodDisplayNumber string
	BillingDate                time.Time
}

func (b *builder) fetchPaymentDetails(ctx context.Context, paymentID string) (*paymentDetails, error) {
	return &paymentDetails{
		PaymentID:                  paymentID,
		PayerID:                    "payer-id",
		PriceE5:                    decimal.NewFromFloat(9.99),
		PriceCurrency:              "USD",
		PaymentMethodDisplayName:   "Apple Pay",
		PaymentMethodDisplayNumber: "**** **** *234",
		BillingDate:                time.Now().Add(time.Hour * -1),
	}, nil
}

type payerDetails struct {
	ID        string
	FirstName string
	LastName  string
}

func (b *builder) fetchPayerDetails(ctx context.Context, paymentID string) (*payerDetails, error) {
	return &payerDetails{
		ID:        "payer-id",
		FirstName: "John",
		LastName:  "Doe",
	}, nil
}

func (b *builder) buildContent(paymentDetails *paymentDetails, payerDetails *payerDetails) map[string]interface{} {
	return map[string]interface{}{
		"payer_first_name":              payerDetails.FirstName,
		"payer_last_name":               payerDetails.LastName,
		"price":                         paymentDetails.PriceE5.StringFixed(2),
		"price_currency":                paymentDetails.PriceCurrency,
		"payment_method_display_name":   paymentDetails.PaymentMethodDisplayName,
		"payment_method_display_number": paymentDetails.PaymentMethodDisplayNumber,
		"billing_date_yyyy_mm_dd":       paymentDetails.BillingDate.Format(time.DateOnly),
	}
}

func (b *builder) buildMetadata(paymentDetails *paymentDetails, extendedMetadata map[string]interface{}) map[string]interface{} {
	metadata := map[string]interface{}{
		"payer_id":   paymentDetails.PayerID,
		"payment_id": paymentDetails.PaymentID,
	}
	for key, value := range extendedMetadata {
		metadata[key] = value
	}
	return metadata
}
