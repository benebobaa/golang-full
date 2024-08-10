package handler

import "sync"

func NewUserHandler() *UserHandler {
	db := make(map[string]User)

	user1 := User{
		ID:            "USER-001",
		Username:      "user1",
		AccountBankID: "AC-001",
		Email:         "bene@gmail.com",
	}

	user2 := User{
		ID:            "USER-002",
		Username:      "user2",
		AccountBankID: "AC-002",
		Email:         "bene@beneboba.me",
	}

	db[user1.ID] = user1
	db[user2.ID] = user2

	return &UserHandler{
		db: db,
	}
}

func NewProductHandler() *ProductHandler {
	db := make(map[string]Product)

	product1 := Product{
		ID:    "P-001",
		Name:  "Product 1",
		Stock: 10,
		Price: 1000,
	}

	product2 := Product{
		ID:    "P-002",
		Name:  "Product 2",
		Stock: 5,
		Price: 2000,
	}

	db[product1.ID] = product1
	db[product2.ID] = product2

	return &ProductHandler{
		db:    db,
		mutex: &sync.RWMutex{},
	}
}

func NewPaymentHandler() *PaymentHandler {

	dbt := make(map[string]Transaction)
	dbb := make(map[string]Balance)

	balance1 := Balance{
		AccountID: "AC-001",
		Balance:   50000,
	}

	balance2 := Balance{
		AccountID: "AC-002",
		Balance:   100000,
	}

	dbb[balance1.AccountID] = balance1
	dbb[balance2.AccountID] = balance2

	return &PaymentHandler{
		dbT:   dbt,
		dbB:   dbb,
		mutex: &sync.RWMutex{},
	}
}
