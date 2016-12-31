package server

import (
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

type Entity struct {
	Email       string `datastore:"-"`
	Name        string
	Description string
	Quantity    string
	Price       string
	Currency    string
	Transaction string
	Status      string
	Time        time.Time
}

type Order struct {
	Meta    []int64
	Email   string
	List    []*Entity
	Context context.Context
}

func (r *Order) save() error {
	root := datastore.NewKey(r.Context, "user", r.Email, 0, nil)
	for _, e := range r.List {
		e.Time = time.Now()
		key := datastore.NewKey(r.Context, "order", "", 0, root)
		k, err := datastore.Put(r.Context, key, e)
		if err != nil {
			log.Infof(r.Context, "Error datastore: %v", err)
			return err
		}
		r.Meta = append(r.Meta, k.IntID())
	}
	return nil
}

func (r *Order) csv() (err error) {
	//root := datastore.NewKey(r.Context, "user", r.Email, 0, nil)
	//q := datastore.NewQuery("order").Ancestor(root).Order("__key__")
	q := datastore.NewQuery("order").Order("__key__")
	for c := q.Run(r.Context); ; {
		var e Entity
		k, err := c.Next(&e)
		if err == datastore.Done {
			err = nil
			break
		}
		if err != nil {
			log.Infof(r.Context, "Error datastore: %v", err)
			break
		}
		e.Email = k.Parent().StringID()
		r.List = append(r.List, &e)
	}
	return
}
