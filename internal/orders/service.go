package orders

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

var (
	ErrInvalidCustomer = errors.New("customer is required")
	ErrInvalidAmount   = errors.New("amount must be positive")
)

type Order struct {
	ID       string `json:"id"`
	Customer string `json:"customer"`
	Amount   int64  `json:"amount"`
}

type Service struct {
	mu     sync.Mutex
	nextID int64
	orders []Order
}

func NewService() *Service {
	return &Service{nextID: 1}
}

func (s *Service) Create(customer string, amount int64) (Order, error) {
	customer = strings.TrimSpace(customer)
	if customer == "" {
		return Order{}, ErrInvalidCustomer
	}
	if amount <= 0 {
		return Order{}, ErrInvalidAmount
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	order := Order{
		ID:       fmt.Sprintf("ord_%06d", s.nextID),
		Customer: customer,
		Amount:   amount,
	}
	s.nextID++
	s.orders = append(s.orders, order)
	return order, nil
}

func (s *Service) List() []Order {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]Order, len(s.orders))
	copy(result, s.orders)
	return result
}
