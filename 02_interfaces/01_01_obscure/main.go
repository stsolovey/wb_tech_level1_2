package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type PaymentProcessor interface {
	ProcessPayment(amount float64) (bool, error)
	ProcessTax(amount float64) (bool, error)
}

type CreditCardProcessor struct {
	logger     *logrus.Logger
	CardNumber string
	CVV        string
	ExpiryDate string
}

func NewCreditCardProcessor(logger *logrus.Logger, cardNumber, cvv, expiryDate string) *CreditCardProcessor {
	return &CreditCardProcessor{
		logger:     logger,
		CardNumber: cardNumber,
		CVV:        cvv,
		ExpiryDate: expiryDate,
	}
}

func (ccp *CreditCardProcessor) ProcessPayment(amount float64) (bool, error) {
	ccp.logger.Infof("Processing credit card payment... %v", amount)

	return true, nil
}

func (ccp *CreditCardProcessor) ProcessTax(amount float64) (bool, error) {
	ccp.logger.Infof("Processing tax for credit card... %v", amount)

	return true, nil
}

type PayPalProcessor struct {
	logger *logrus.Logger
	Email  string
}

func NewPayPalProcessor(logger *logrus.Logger, email string) *PayPalProcessor {
	return &PayPalProcessor{
		logger: logger,
		Email:  email,
	}
}

func (pp *PayPalProcessor) ProcessPayment(amount float64) (bool, error) {
	pp.logger.Infof("Processing PayPal payment... %v", amount)

	return true, nil
}

func (pp *PayPalProcessor) ProcessTax(amount float64) (bool, error) {
	pp.logger.Infof("Processing tax for PayPal payment... %v", amount)

	return true, nil
}

func executePayment(p PaymentProcessor, amount float64, log *logrus.Logger) error {
	success, err := p.ProcessTax(amount)
	if err != nil {
		return fmt.Errorf("executePayment p.ProcessTax(amount): %w", err)
	}

	if success {
		log.Infof("Payment successful: %v", amount)
	}

	return nil
}

func main() {
	log := logrus.New()

	creditCardProcessor := NewCreditCardProcessor(log, "1234-5678-9012-3456", "123", "01/25")
	paypalProcessor := NewPayPalProcessor(log, "example@paypal.com")

	err := executePayment(creditCardProcessor, 100.0, log) //nolint:mnd
	if err != nil {
		log.WithError(err).Panic("Failed to execute creditCardProcessor payment")
	}

	err = executePayment(paypalProcessor, 150.0, log) //nolint:mnd
	if err != nil {
		log.WithError(err).Panic("Failed to execute paypalProcessor payment")
	}
}
