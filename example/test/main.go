package main

import (
	"context"
	"fmt"

	dataloader "github.com/graph-gophers/dataloader/v7"
)

func main() {
	type User struct {
		ID        int
		Email     string
		FirstName string
		LastName  string
	}

	m := map[int]*User{
		5: {ID: 5, FirstName: "John", LastName: "Smith", Email: "john@example.com"},
	}

	batchFunc := func(_ context.Context, reqs []*dataloader.OneRequest[int, *User]) {
		var keys []int
		for _, req := range reqs {
			keys = append(keys, req.Key())
		}
		for i, k := range keys {
			reqs[i].OnDone(&dataloader.Result[*User]{Data: m[k]})
		}
	}

	// go-cache will automatically cleanup expired items on given duration
	cache := &dataloader.NoCache[int, *User]{}
	loader := dataloader.NewBatchedLoader(batchFunc, dataloader.WithCache[int, *User](cache))

	result, err := loader.Load(context.Background(), 5)()
	if err != nil {
		// handle error
		panic(err)
	}

	fmt.Printf("result: %+v", result)
	// Output: result: &{ID:5 Email:john@example.com FirstName:John LastName:Smith}
}
