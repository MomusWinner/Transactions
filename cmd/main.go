package main

import (
	"Transactions/database"
	"Transactions/internal/config"
	"Transactions/internal/dbconn"
	"context"
	"fmt"
	"log"
	"time"

	"database/sql"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	conf := config.LoadConfig(".")
	dbconn.Init(&conf)
	c := context.Background()
	customer_id, err := dbconn.DBQueries.CreateCustomer(
		c,
		database.CreateCustomerParams{
			Firstname: "Egor",
			Lastname:  "Vitalevich",
			Email:     "egor@mail.ru",
		},
	)

	if err == nil {
		log.Print("Couln't create customer")
	}

	// 1
	milk_id, err := dbconn.DBQueries.CreateProduct(c, database.CreateProductParams{
		Name:  "milk",
		Price: 100,
	})

	chees_id, err := dbconn.DBQueries.CreateProduct(c, database.CreateProductParams{
		Name:  "chees",
		Price: 100,
	})

	orders := []*OrderItem{
		{milk_id, 2, 200},
		{chees_id, 3, 300},
	}

	err = CreateOrders(c, dbconn.DB, dbconn.DBQueries, customer_id, orders)
	if err != nil {
		log.Print("Create order")
	} else {
		log.Print("Couln't create order")
	}

	// 2
	dbconn.DBQueries.UpdateComusmerEmail(c, database.UpdateComusmerEmailParams{
		Email: "newemail@mail.ru",
		ID:    customer_id,
	})

	// 3
	newproduct_id, err := dbconn.DBQueries.CreateProduct(c, database.CreateProductParams{
		Name:  "newproduct",
		Price: 100,
	})

	if err != nil {
		log.Print("Couln't create product")
	} else {
		log.Printf("Create product %v", newproduct_id)
	}
}

type OrderItem struct {
	productId int32
	quantity  int32
	subtotal  float64
}

func CreateOrders(
	ctx context.Context,
	db *sql.DB,
	queries *database.Queries,
	customerId int32,
	orders []*OrderItem,
) error {
	// Begin a transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Defer a rollback in case anything fails
	defer tx.Rollback()

	// Use the transaction to create queries
	qtx := queries.WithTx(tx)

	order_id, err := qtx.CreateOrder(ctx, database.CreateOrderParams{
		Customerid:  int32(customerId),
		Orderdate:   time.Now(),
		Totalamount: float64(len(orders)),
	})

	for _, orderItem := range orders {
		qtx.CreateOrderItem(ctx, database.CreateOrderItemParams{
			Orderid:   order_id,
			Productid: orderItem.productId,
			Quantity:  orderItem.quantity,
			Subtotal:  orderItem.subtotal,
		})
	}

	if err != nil {
		return fmt.Errorf("failed to debit account: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
