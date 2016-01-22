// generated by xgen -- DO NOT EDIT
package dql_test

import (
	"bytes"
	"fmt"

	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

var (
	USER        = schema.TABLE("sys_user", "USER")
	USER_ID     = USER.PRIMARY("id", "ID")
	USER_NAME   = USER.COLUMN("name", "NAME")
	USER_PASSWD = USER.COLUMN("passwd", "PASSWD")
	USER_EMAIL  = USER.COLUMN("email", "EMAIL")
)

func NewUser() *User {
	e := &User{}
	e.init()
	return e
}

func (me *User) Id() string {
	return me.id.Value()
}

func (me *User) SetId(v string) {
	me.id.SetValue(v)
}

func (me *User) Name() string {
	return me.name.Value()
}

func (me *User) SetName(v string) {
	me.name.SetValue(v)
}

func (me *User) Passwd() string {
	return me.passwd.Value()
}

func (me *User) SetPasswd(v string) {
	me.passwd.SetValue(v)
}

func (me *User) Email() string {
	return me.email.Value()
}

func (me *User) SetEmail(v string) {
	me.email.SetValue(v)
}

func (me *User) init() {
	me.table = USER
	me.initSetDict()
	me.initSetColumn()
	me.initSetDefault()
	me.initSetField()
}

func (me *User) initSetDict() {
}

func (me *User) initSetColumn() {
	
	me.id.SetColumn(USER_ID)
	me.name.SetColumn(USER_NAME)
	me.passwd.SetColumn(USER_PASSWD)
	me.email.SetColumn(USER_EMAIL)
}

func (me *User) initSetDefault() {
	
}

func (me *User) initSetField() {
	
	me.id.SetField(entity.DefaultField())
	me.name.SetField(entity.DefaultField())
	me.passwd.SetField(entity.DefaultField())
	me.email.SetField(entity.DefaultField())
}

func (me User) New() entity.Interface {
	return NewUser()
}

func (me *User) Get(column string) interface{} {
	switch column {
	case USER_ID.Name():
		return me.id.Value()
	case USER_NAME.Name():
		return me.name.Value()
	case USER_PASSWD.Name():
		return me.passwd.Value()
	case USER_EMAIL.Name():
		return me.email.Value()
	}
	return nil
}

func (me *User) GetPtr(column string) interface{} {
	switch column {
	case USER_ID.Name():
		return me.id.ValuePtr()
	case USER_NAME.Name():
		return me.name.ValuePtr()
	case USER_PASSWD.Name():
		return me.passwd.ValuePtr()
	case USER_EMAIL.Name():
		return me.email.ValuePtr()
	}
	return nil
}

func (me *User) Table() schema.Table {
	return me.table
}

func (me *User) Type(column string) (entity.Type, bool) {
	switch column {
	case USER_ID.Name():
		return &me.id, true
	case USER_NAME.Name():
		return &me.name, true
	case USER_PASSWD.Name():
		return &me.passwd, true
	case USER_EMAIL.Name():
		return &me.email, true
	}
	return nil, false
}

func (me *User) Column(field string) (schema.Column, bool) {
	switch strings.ToLowerFirst(field) {
	case "id":
		return USER_ID, true
	case "name":
		return USER_NAME, true
	case "passwd":
		return USER_PASSWD, true
	case "email":
		return USER_EMAIL, true
	}
	return nil, false
}

func (me *User) Columns() []schema.Column {
	return []schema.Column{
		USER_ID,
		USER_NAME,
		USER_PASSWD,
		USER_EMAIL,
	}
}

func (me *User) Names() []string {
	return []string{
		"id",
		"name",
		"passwd",
		"email",
	}
}

func (me *User) Value() *User {
	return me
}

func (me *User) SetString(field, value string) error {
	switch strings.ToLowerFirst(field) {
	case "id":
		return me.id.SetString(value)
	case "name":
		return me.name.SetString(value)
	case "passwd":
		return me.passwd.SetString(value)
	case "email":
		return me.email.SetString(value)
	}
	return nil
}

func (me *User) Validate() error {
	return nil
}

func (me *User) JSON() string {
	var b bytes.Buffer
	b.WriteString("{")
	b.WriteString(fmt.Sprintf(`,"id":%q`, me.id.String()))
	b.WriteString(fmt.Sprintf(`,"name":%q`, me.name.String()))
	b.WriteString(fmt.Sprintf(`,"passwd":%q`, me.passwd.String()))
	b.WriteString(fmt.Sprintf(`,"email":%q`, me.email.String()))
	b.WriteString("}")
	return b.String()
}
