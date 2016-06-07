package utils

import (
	"strconv"
	"net/http"
	"github.com/astaxie/beego"
)

type Paginator  []*PaginatorItem

type PaginatorItem struct {
	Lable      string
	Link       string
	IsActive   bool
	IsDisabled bool
}

func NewPaginator(req *http.Request, curr, per, total int) *Paginator {
	var items []*PaginatorItem
	baseLink := req.URL.RawPath
	beego.Debug(baseLink)

	lastPage := total / per
	if (total / per) == 0 {
		lastPage++
	}

	if curr > lastPage {
		curr = lastPage
	}

	if curr > 1 {
		pageStr := strconv.Itoa(curr - 1)
		// FIXME: 测试结束后改为路由分发形式，下同
		items = append(items, &PaginatorItem{Lable:"<", Link:baseLink + "?page=" + pageStr})
	}

	if curr - 2 > 1 {
		items = append(items, &PaginatorItem{Lable:"1", Link:baseLink + "?page=1"})
		items = append(items, &PaginatorItem{Lable:"...", IsDisabled:true})
	}

	for i := -2; i <= 2; i++ {
		page := curr + i
		if page <= 0 || page > lastPage {
			continue
		}
		pageStr := strconv.Itoa(page)
		item := &PaginatorItem{Lable:pageStr, Link:baseLink + "?page=" + pageStr}
		if i == 0 {
			item.IsActive = true
		}
		items = append(items, item)
	}

	if curr + 2 < lastPage {
		items = append(items, &PaginatorItem{Lable:"...", IsDisabled:true})
		lastPageStr := strconv.Itoa(lastPage)
		items = append(items, &PaginatorItem{Lable:lastPageStr, Link:baseLink + "?page=" + lastPageStr})
	}

	if curr < lastPage {
		pageStr := strconv.Itoa(curr + 1)
		items = append(items, &PaginatorItem{Lable:">", Link:baseLink + "?page=" + pageStr})
	}

	paginator := Paginator(items)
	return &paginator
}
