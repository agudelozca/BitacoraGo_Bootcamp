package service

import "app/internal"

// NewCustomersDefault creates new default service for customer entity.
func NewCustomersDefault(rp internal.RepositoryCustomer) *CustomersDefault {
	return &CustomersDefault{rp}
}

// CustomersDefault is the default service implementation for customer entity.
type CustomersDefault struct {
	// rp is the repository for customer entity.
	rp internal.RepositoryCustomer
}

// FindAll returns all customers.
func (s *CustomersDefault) FindAll() (c []internal.Customer, err error) {
	c, err = s.rp.FindAll()
	return
}

// FindTopActiveCustomersByAmountSpent returns the top active customers by amount spent.
func (s *CustomersDefault) FindTopActiveCustomersByAmountSpent(limit int) (c []internal.CustomerSpent, err error) {
	c, err = s.rp.FindTopActiveCustomersByAmountSpent(limit)
	return
}

// FindInvoicesByCondition returns the total invoices by customer condition.
func (s *CustomersDefault) FindInvoicesByCondition() (c []internal.CustomerInvoicesByCondition, err error) {
	c, err = s.rp.FindInvoicesByCondition()
	return
}

// Save saves the customer.
func (s *CustomersDefault) Save(c *internal.Customer) (err error) {
	err = s.rp.Save(c)
	return
}