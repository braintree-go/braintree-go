package braintree

import "testing"

func TestResponseHandlesBadXML(t *testing.T) {
	response := Response{Body: []byte("<xml></badxml/<>Ohdear>")}

	txResult, err := response.TransactionResult()
	if err == nil {
		t.Errorf("TransactionResult() did not return an error on bad XML")
	} else if txResult.Success() {
		t.Errorf("TransactionResult() returned a successful result on bad XML")
	}

	customerResult, err := response.CustomerResult()
	if err == nil {
		t.Errorf("CustomerResult() did not return an error on bad XML")
	} else if customerResult.Success() {
		t.Errorf("CustomerResult() returned a successful result on bad XML")
	}

	errorResult, err := response.ErrorResult()
	if err == nil {
		t.Errorf("ErrorResult() did not return an error on bad XML")
	} else if errorResult.Success() {
		t.Errorf("ErrorResult() returned a successful result on bad XML")
	}
}

func TestErrorResultPanicsForTransactions(t *testing.T) {
	result := ErrorResult{}

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Did not panic on ErrorResult.Transaction()")
		}
	}()

	result.Transaction()
}

func TestErrorResultPanicsForCustomer(t *testing.T) {
	result := ErrorResult{}

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Did not panic on ErrorResult.Customer()")
		}
	}()

	result.Customer()
}
