package models

import (
	"errors"
	"fmt"
	"log"

	"github.com/astaxie/beego/orm"
	. "github.com/micln/go-utils"
)

var TableNameDocument = `document`

type Document struct {
	Id       uint64
	Content  string `orm:"type(text)"`
	Keywords string `orm:"null"`
	Hash     string
}

func (d *Document) GetHash() string {
	if len(d.Hash) == 0 {
		d.Hash = Md5(d.Content)
	}
	return d.Hash
}

func (d *Document) IsDuplicate() bool {
	return orm.NewOrm().QueryTable(TableNameDocument).Filter(`hash__exact`, d.GetHash()).Exist()
}

func (d *Document) Save() (err error) {
	if len(d.Hash) == 0 {
		d.Hash = Md5(d.Content)
	}

	o := orm.NewOrm()
	if d.Id > 0 {
		_, err = o.Update(d)
	} else {
		if d.IsDuplicate() {
			err = errors.New(fmt.Sprintf("duplicate contents with hash(%s)", d.GetHash()))
		} else {
			_, err = o.Insert(d)
		}
	}

	if err != nil {
		log.Println("SavingError", *d)
	}

	return
}

func (d Document) Get() (docs []*Document) {
	o := orm.NewOrm()
	qs := o.QueryTable(TableNameDocument)

	if len(d.Content) > 0 {
		qs = qs.Filter(`content__contains`, d.Content)
	}

	if len(d.Hash) > 0 {
		qs = qs.Filter(`hash__contains`, d.Hash)
	}

	qs.All(&docs)

	return
}
