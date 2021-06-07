package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Quote struct {
	ID        int    `json:"id"`
	Author    string `json:"author"`
	Quote     string `json:"quote"`
	Category  string `json:"category"`
	CreatedAt int64  `json:"created_at"`
}

type Quotes []Quote

func (qs Quotes) GetLastID() (r int) {
	for _, q := range qs {
		if q.ID > r {
			r = q.ID
		}
	}
	return r
}

func (qs *Quotes) FillWithTestData(c int) {
	for i := 1; i <= c; i++ {
		q := Quote{
			ID:        qs.GetLastID() + 1,
			Author:    fmt.Sprintf("Author %d", i),
			Quote:     fmt.Sprintf("Quote %d", i),
			Category:  fmt.Sprintf("Category%d", i),
			CreatedAt: time.Now().Unix(),
		}
		*qs = append(*qs, q)
	}
}

func (qs *Quotes) Create(reqBody []byte) Quote {
	var q Quote
	q.ID = qs.GetLastID() + 1
	json.Unmarshal(reqBody, &q)
	*qs = append(*qs, q)
	return q
}

func (qs *Quotes) Update(reqBody []byte, i int) Quote {
	json.Unmarshal(reqBody, &(*qs)[i])
	return (*qs)[i]
}

func (qs *Quotes) DeleteByIndex(i int) {
	(*qs) = append((*qs)[:i], (*qs)[i+1:]...)
}

func (qs Quotes) FindIndexById(id int) (i int) {
	for i, q := range qs {
		if q.ID == id {
			return i
		}
	}
	panic("Not found")
}

func (qs Quotes) GetAllByCategory(name string) (fq Quotes) {
	for _, q := range qs {
		if q.Category == name {
			fq = append(fq, q)
		}
	}
	return
}
