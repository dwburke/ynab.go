package payee

import (
	"fmt"

	"go.bmvs.io/ynab/api"
)

// NewService facilitates the creation of a new payee service instance
func NewService(c api.ClientReader) *Service {
	return &Service{c}
}

// Service wraps YNAB payee API endpoints
type Service struct {
	c api.ClientReader
}

// GetPayees fetches the list of payees from a budget
// https://api.youneedabudget.com/v1#/Payees/getPayees
func (s *Service) GetPayees(budgetID string) ([]*Payee, error) {
	resModel := struct {
		Data struct {
			Payees []*Payee `json:"payees"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/payees", budgetID)
	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.Payees, nil
}

// GetPayee fetches a specific payee from a budget
// https://api.youneedabudget.com/v1#/Payees/getPayeeById
func (s *Service) GetPayee(budgetID, payeeID string) (*Payee, error) {
	resModel := struct {
		Data struct {
			Payee *Payee `json:"payee"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/payees/%s", budgetID, payeeID)
	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.Payee, nil
}

// GetPayeeLocations fetches the list of payee locations from a budget
// https://api.youneedabudget.com/v1#/Payee_Locations/getPayeeLocations
func (s *Service) GetPayeeLocations(budgetID string) ([]*Location, error) {
	resModel := struct {
		Data struct {
			PayeeLocations []*Location `json:"payee_locations"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/payee_locations", budgetID)
	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.PayeeLocations, nil
}

// GetPayeeLocation fetches a specific payee location from a budget
// https://api.youneedabudget.com/v1#/Payee_Locations/getPayeeLocationById
func (s *Service) GetPayeeLocation(budgetID, payeeLocationID string) (*Location, error) {
	resModel := struct {
		Data struct {
			PayeeLocation *Location `json:"payee_location"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/payee_locations/%s", budgetID, payeeLocationID)
	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.PayeeLocation, nil
}

// GetPayeeLocationsByPayee fetches the list of locations of a specific payee from a budget
// https://api.youneedabudget.com/v1#/Payee_Locations/getPayeeLocationsByPayee
func (s *Service) GetPayeeLocationsByPayee(budgetID, payeeID string) ([]*Location, error) {
	resModel := struct {
		Data struct {
			PayeeLocations []*Location `json:"payee_locations"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/payees/%s/payee_locations", budgetID, payeeID)
	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.PayeeLocations, nil
}
