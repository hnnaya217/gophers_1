package firsttask

import "errors"

var errInvalidAmount = errors.New("nilai harga barang dan uang pembeli harus lebih besar dari nol")

type transaction struct {
	itemPrice  int64
	paidAmount int64
}

type result struct {
	approved  bool
	change    int64
	shortfall int64
}

func newTransaction(itemPrice, paidAmount int64) (*transaction, error) {
	if itemPrice <= 0 || paidAmount <= 0 {
		return nil, errInvalidAmount
	}

	return &transaction{
		itemPrice:  itemPrice,
		paidAmount: paidAmount,
	}, nil
}

func (t *transaction) evaluate() result {
	if t.paidAmount < t.itemPrice {
		return result{
			approved:  false,
			shortfall: t.itemPrice - t.paidAmount,
		}
	}

	return result{
		approved: true,
		change:   t.paidAmount - t.itemPrice,
	}
}
