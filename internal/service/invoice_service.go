package service

import (
	"github.com/lucasfontesgaspareto/imersao22/go-gateway/internal/domain"
	"github.com/lucasfontesgaspareto/imersao22/go-gateway/internal/dto"
)

type InvoiceService struct {
	repository     domain.InvoiceRepository
	accountService AccountService
}

func NewInvoiceService(repository domain.InvoiceRepository, accountService AccountService) *InvoiceService {
	return &InvoiceService{
		repository:     repository,
		accountService: accountService,
	}
}

func (s *InvoiceService) Create(input dto.CreateInvoiceInput) (*dto.InvoiceOutput, error) {
	accountOutput, err := s.accountService.FindByAPIKey(input.APIKey)
	if err != nil {
		return nil, err
	}

	invoice, err := dto.ToInvoice(input, accountOutput.ID)
	if err != nil {
		return nil, err
	}

	if err := invoice.Process(); err != nil {
		return nil, err
	}

	if invoice.Status == domain.StatusApproved {
		_, err = s.accountService.UpdateBalance(accountOutput.ID, invoice.Amount)
		if err != nil {
			return nil, err
		}
	}

	if err := s.repository.Save(invoice); err != nil {
		return nil, err
	}

	return dto.FromInvoice(invoice), nil
}

func (s *InvoiceService) GetByID(id, apiKey string) (*dto.InvoiceOutput, error) {
	invoice, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	accountOutput, err := s.accountService.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	if invoice.AccountID != accountOutput.ID {
		return nil, domain.ErrUnauthorizedAccess
	}

	return dto.FromInvoice(invoice), nil
}

func (s *InvoiceService) ListByAccount(accountID string) ([]*dto.InvoiceOutput, error) {
	invoices, err := s.repository.FindByAccountID(accountID)
	if err != nil {
		return nil, err
	}

	output := make([]*dto.InvoiceOutput, len(invoices))
	for i, invoice := range invoices {
		output[i] = dto.FromInvoice(invoice)
	}
	return output, nil
}

func (s *InvoiceService) ListByAccountAPIKey(apiKey string) ([]*dto.InvoiceOutput, error) {
	accountOutput, err := s.accountService.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	return s.ListByAccount(accountOutput.ID)
}
