package demomodule

import (
	"time"

	mollie "github.com/gertcuykens/mollie-api-go"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

// Order struct
type Order struct {
	Email   string
	Method  string
	Issuer  string
	Server  string
	Product []Product
}

// Product struct
type Product struct {
	Email       string `json:"-"`
	Name        string
	Description string
	Quantity    float64
	Price       float64
	Currency    string
	Transaction string    `datastore:"-" json:"-"`
	Status      string    `datastore:"-" json:"-"`
	Time        time.Time `json:"-"`
}

func save(ctx context.Context, email string, transaction string, product []Product) (meta []int64, err error) {
	tx := datastore.NewKey(ctx, "transaction", transaction, 0, nil)
	for _, e := range product {
		e.Email = email
		e.Time = time.Now()
		key := datastore.NewKey(ctx, "order", "", 0, tx)
		k, err := datastore.Put(ctx, key, &e)
		if err != nil {
			log.Errorf(ctx, "Error datastore order.go save: %v", err)
			return nil, err
		}
		meta = append(meta, k.IntID())
	}
	return meta, nil
}

func csv2(ctx context.Context, email string) (product []Product, err error) {
	var q *datastore.Query
	if email == "" {
		q = datastore.NewQuery("order").Order("__key__")
	} else {
		q = datastore.NewQuery("order").Filter("Email =", email).Order("Email")
	}
	for c := q.Run(ctx); ; {
		var e Product
		k, err := c.Next(&e)
		if err == datastore.Done {
			err = nil
			break
		}
		if err != nil {
			break
		}
		// log.Debugf(ctx, "\n\n%v\n\n", e)
		e.Transaction = k.Parent().StringID()
		tx := datastore.NewKey(ctx, "transaction", e.Transaction, 0, nil)
		var t mollie.Transaction
		err = datastore.Get(ctx, tx, &t)
		if err == nil {
			e.Status = t.Status
			// e.Email = t.Description
			// e.Time = t.PaidDatetime
		}
		product = append(product, e)
	}
	return
}
