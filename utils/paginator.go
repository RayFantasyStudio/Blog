package utils

import (
	"strconv"
	"fmt"
)

type Paginator  []*PaginatorItem

type PaginatorItem struct {
	Lable      string
	Link       string
	IsActive   bool
	IsDisabled bool
}

func NewPaginator(linkFormat string, curr, per int, total int64) *Paginator {
	var items []*PaginatorItem

	lastPage := int(total / int64(per))
	if (total % int64(per)) != 0 {
		lastPage++
	}

	if curr > lastPage {
		curr = lastPage
	}

	if curr > 1 {
		link := fmt.Sprintf(linkFormat, curr - 1)
		items = append(items, &PaginatorItem{Lable:"<", Link:link})
	}

	if curr - 2 > 1 {
		link := fmt.Sprintf(linkFormat, 1)
		items = append(items, &PaginatorItem{Lable:"1", Link:link})
		items = append(items, &PaginatorItem{Lable:"...", IsDisabled:true})
	}

	for i := -2; i <= 2; i++ {
		page := curr + i
		if page <= 0 || page > lastPage {
			continue
		}
		link := fmt.Sprintf(linkFormat, page)
		item := &PaginatorItem{Lable:strconv.Itoa(page), Link:link}
		if i == 0 {
			item.IsActive = true
		}
		items = append(items, item)
	}

	if curr + 2 < lastPage {
		items = append(items, &PaginatorItem{Lable:"...", IsDisabled:true})
		link := fmt.Sprintf(linkFormat, lastPage)
		items = append(items, &PaginatorItem{Lable:strconv.Itoa(lastPage), Link:link})
	}

	if curr < lastPage {
		link := fmt.Sprintf(linkFormat, curr + 1)
		items = append(items, &PaginatorItem{Lable:">", Link:link})
	}

	paginator := Paginator(items)
	return &paginator
}
