package firsttask

import "testing"

func TestNewTransaction_InvalidAmount(t *testing.T) {
	cases := []struct {
		name       string
		itemPrice  int64
		paidAmount int64
	}{
		{"harga nol", 0, 10000},
		{"harga negatif", -5000, 10000},
		{"uang pembeli nol", 15000, 0},
		{"uang pembeli negatif", 15000, -1000},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := newTransaction(tc.itemPrice, tc.paidAmount); err != errInvalidAmount {
				t.Fatalf("expected errInvalidAmount, got %v", err)
			}
		})
	}
}

func TestEvaluate_UangKurang(t *testing.T) {
	tx, err := newTransaction(15000, 10000)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	r := tx.evaluate()
	if r.approved {
		t.Fatalf("expected transaction to be rejected")
	}
	if r.shortfall != 5000 {
		t.Fatalf("expected shortfall 5000, got %d", r.shortfall)
	}
}

func TestEvaluate_UangCukup(t *testing.T) {
	tx, err := newTransaction(15000, 20000)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	r := tx.evaluate()
	if !r.approved {
		t.Fatalf("expected transaction to be approved")
	}
	if r.change != 5000 {
		t.Fatalf("expected change 5000, got %d", r.change)
	}
}

func TestEvaluate_UangPas(t *testing.T) {
	tx, err := newTransaction(15000, 15000)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	r := tx.evaluate()
	if !r.approved {
		t.Fatalf("expected transaction to be approved when paid amount equals price")
	}
	if r.change != 0 {
		t.Fatalf("expected change 0, got %d", r.change)
	}
}
