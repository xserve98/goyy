// generated by xgen -- DO NOT EDIT
package role

import (
	"bytes"
	"fmt"

	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

var (
	POST_ENTITY           = schema.TABLE("sys_role_post", "ROLE POST")
	POST_ENTITY_ID        = POST_ENTITY.PRIMARY("id", "ID")
	POST_ENTITY_MEMO      = POST_ENTITY.COLUMN("memo", "MEMO")
	POST_ENTITY_CREATES   = POST_ENTITY.COLUMN("creates", "CREATES")
	POST_ENTITY_CREATER   = POST_ENTITY.CREATER("creater", "CREATER")
	POST_ENTITY_CREATED   = POST_ENTITY.CREATED("created", "CREATED")
	POST_ENTITY_MODIFIER  = POST_ENTITY.MODIFIER("modifier", "MODIFIER")
	POST_ENTITY_MODIFIED  = POST_ENTITY.MODIFIED("modified", "MODIFIED")
	POST_ENTITY_VERSION   = POST_ENTITY.VERSION("version", "VERSION")
	POST_ENTITY_DELETION  = POST_ENTITY.DELETION("deletion", "DELETION")
	POST_ENTITY_ARTIFICAL = POST_ENTITY.COLUMN("artifical", "ARTIFICAL")
	POST_ENTITY_HISTORY   = POST_ENTITY.COLUMN("history", "HISTORY")
	POST_ENTITY_ROLE_ID   = POST_ENTITY.COLUMN("role_id", "ROLE_ID")
	POST_ENTITY_POST_ID   = POST_ENTITY.COLUMN("post_id", "POST_ID")
)

func NewPostEntity() *PostEntity {
	e := &PostEntity{}
	e.init()
	return e
}

func (me *PostEntity) RoleId() string {
	return me.roleId.Value()
}

func (me *PostEntity) SetRoleId(v string) {
	me.roleId.SetValue(v)
}

func (me *PostEntity) PostId() string {
	return me.postId.Value()
}

func (me *PostEntity) SetPostId(v string) {
	me.postId.SetValue(v)
}

func (me *PostEntity) init() {
	me.table = POST_ENTITY

	if t, ok := me.Sys.Type("id"); ok {
		t.SetColumn(POST_ENTITY_ID)
	}
	if t, ok := me.Sys.Type("memo"); ok {
		t.SetColumn(POST_ENTITY_MEMO)
	}
	if t, ok := me.Sys.Type("creates"); ok {
		t.SetColumn(POST_ENTITY_CREATES)
	}
	if t, ok := me.Sys.Type("creater"); ok {
		t.SetColumn(POST_ENTITY_CREATER)
	}
	if t, ok := me.Sys.Type("created"); ok {
		t.SetColumn(POST_ENTITY_CREATED)
	}
	if t, ok := me.Sys.Type("modifier"); ok {
		t.SetColumn(POST_ENTITY_MODIFIER)
	}
	if t, ok := me.Sys.Type("modified"); ok {
		t.SetColumn(POST_ENTITY_MODIFIED)
	}
	if t, ok := me.Sys.Type("version"); ok {
		t.SetColumn(POST_ENTITY_VERSION)
	}
	if t, ok := me.Sys.Type("deletion"); ok {
		t.SetColumn(POST_ENTITY_DELETION)
	}
	if t, ok := me.Sys.Type("artifical"); ok {
		t.SetColumn(POST_ENTITY_ARTIFICAL)
	}
	if t, ok := me.Sys.Type("history"); ok {
		t.SetColumn(POST_ENTITY_HISTORY)
	}
	me.roleId.SetColumn(POST_ENTITY_ROLE_ID)
	me.postId.SetColumn(POST_ENTITY_POST_ID)

	if t, ok := me.Sys.Type("created"); ok {
		t.SetDefault("-62135596800")
	}
	if t, ok := me.Sys.Type("modified"); ok {
		t.SetDefault("-62135596800")
	}

	columns := []string{"id", "memo", "creates", "creater", "created", "modifier", "modified",
		"version", "deletion", "artifical", "history"}
	for _, c := range columns {
		if t, ok := me.Sys.Type(c); ok {
			t.SetField(entity.DefaultField())
		}
	}
	me.roleId.SetField(entity.DefaultField())
	me.postId.SetField(entity.DefaultField())
}

func (me PostEntity) New() entity.Interface {
	return NewPostEntity()
}

func (me *PostEntity) Get(column string) interface{} {
	switch column {
	case POST_ENTITY_ROLE_ID.Name():
		return me.roleId.Value()
	case POST_ENTITY_POST_ID.Name():
		return me.postId.Value()
	}
	return me.Sys.Get(column)
}

func (me *PostEntity) GetPtr(column string) interface{} {
	switch column {
	case POST_ENTITY_ROLE_ID.Name():
		return me.roleId.ValuePtr()
	case POST_ENTITY_POST_ID.Name():
		return me.postId.ValuePtr()
	}
	return me.Sys.GetPtr(column)
}

func (me *PostEntity) Table() schema.Table {
	return me.table
}

func (me *PostEntity) Type(column string) (entity.Type, bool) {
	switch column {
	case POST_ENTITY_ROLE_ID.Name():
		return &me.roleId, true
	case POST_ENTITY_POST_ID.Name():
		return &me.postId, true
	}
	return me.Sys.Type(column)
}

func (me *PostEntity) Column(field string) (schema.Column, bool) {
	switch strings.ToLowerFirst(field) {
	case "roleId":
		return POST_ENTITY_ROLE_ID, true
	case "postId":
		return POST_ENTITY_POST_ID, true
	}
	return me.Sys.Column(field)
}

func (me *PostEntity) Columns() []schema.Column {
	return []schema.Column{
		POST_ENTITY_ID,
		POST_ENTITY_MEMO,
		POST_ENTITY_CREATES,
		POST_ENTITY_CREATER,
		POST_ENTITY_CREATED,
		POST_ENTITY_MODIFIER,
		POST_ENTITY_MODIFIED,
		POST_ENTITY_VERSION,
		POST_ENTITY_DELETION,
		POST_ENTITY_ARTIFICAL,
		POST_ENTITY_HISTORY,
		POST_ENTITY_ROLE_ID,
		POST_ENTITY_POST_ID,
	}
}

func (me *PostEntity) Names() []string {
	return []string{
		"id",
		"memo",
		"creates",
		"creater",
		"created",
		"modifier",
		"modified",
		"version",
		"deletion",
		"artifical",
		"history",
		"roleId",
		"postId",
	}
}

func (me *PostEntity) Value() *PostEntity {
	return me
}

func (me *PostEntity) SetString(field, value string) error {
	switch strings.ToLowerFirst(field) {
	case "roleId":
		return me.roleId.SetString(value)
	case "postId":
		return me.postId.SetString(value)
	}
	return me.Sys.SetString(field, value)
}

func (me *PostEntity) Validate() error {
	return nil
}

func (me *PostEntity) JSON() string {
	var b bytes.Buffer
	b.WriteString("{")
	b.WriteString(fmt.Sprintf(`"id":%q`, me.Sys.Pk.Id()))
	b.WriteString(fmt.Sprintf(`,"memo":%q`, me.Sys.Memo()))
	b.WriteString(fmt.Sprintf(`,"creates":%q`, me.Sys.Creates()))
	b.WriteString(fmt.Sprintf(`,"creater":%q`, me.Sys.Creater()))
	b.WriteString(fmt.Sprintf(`,"created":%d`, me.Sys.Created()))
	b.WriteString(fmt.Sprintf(`,"modifier":%q`, me.Sys.Modifier()))
	b.WriteString(fmt.Sprintf(`,"modified":%d`, me.Sys.Modified()))
	b.WriteString(fmt.Sprintf(`,"version":%d`, me.Sys.Version()))
	b.WriteString(fmt.Sprintf(`,"deletion":%d`, me.Sys.Deletion()))
	b.WriteString(fmt.Sprintf(`,"artifical":%d`, me.Sys.Artifical()))
	b.WriteString(fmt.Sprintf(`,"history":%d`, me.Sys.History()))
	b.WriteString(fmt.Sprintf(`,"roleId":%q`, me.roleId.String()))
	b.WriteString(fmt.Sprintf(`,"postId":%q`, me.postId.String()))
	b.WriteString("}")
	return b.String()
}
