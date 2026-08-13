package orders

import "testing"

func TestCreateAndList(t *testing.T) {
	svc := NewService()

	created, err := svc.Create("acme", 12500)
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	if created.ID != "ord_000001" {
		t.Fatalf("unexpected id: %s", created.ID)
	}

	orders := svc.List()
	if len(orders) != 1 {
		t.Fatalf("expected 1 order, got %d", len(orders))
	}
}

func TestCreateRejectsInvalidInput(t *testing.T) {
	svc := NewService()

	if _, err := svc.Create("", 100); err != ErrInvalidCustomer {
		t.Fatalf("expected ErrInvalidCustomer, got %v", err)
	}
	if _, err := svc.Create("acme", 0); err != ErrInvalidAmount {
		t.Fatalf("expected ErrInvalidAmount, got %v", err)
	}
}
