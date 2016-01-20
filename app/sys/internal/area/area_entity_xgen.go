// generated by xgen -- DO NOT EDIT
package area

import (
	"bytes"
	"fmt"

	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

var (
	ENTITY              = schema.TABLE("sys_area", "AREA")
	ENTITY_ID           = ENTITY.PRIMARY("id", "ID")
	ENTITY_CODE         = ENTITY.COLUMN("code", "CODE")
	ENTITY_NAME         = ENTITY.COLUMN("name", "NAME")
	ENTITY_FULLNAME     = ENTITY.COLUMN("fullname", "FULLNAME")
	ENTITY_GENRE        = ENTITY.COLUMN("genre", "GENRE")
	ENTITY_LEAF         = ENTITY.COLUMN("leaf", "LEAF")
	ENTITY_GRADE        = ENTITY.COLUMN("grade", "GRADE")
	ENTITY_ORDINAL      = ENTITY.COLUMN("ordinal", "ORDINAL")
	ENTITY_PARENT_ID    = ENTITY.COLUMN("parent_id", "PARENT_ID")
	ENTITY_PARENT_IDS   = ENTITY.COLUMN("parent_ids", "PARENT_IDS")
	ENTITY_PARENT_CODES = ENTITY.COLUMN("parent_codes", "PARENT_CODES")
	ENTITY_PARENT_NAMES = ENTITY.COLUMN("parent_names", "PARENT_NAMES")
	ENTITY_MEMO         = ENTITY.COLUMN("memo", "MEMO")
	ENTITY_CREATES      = ENTITY.COLUMN("creates", "CREATES")
	ENTITY_CREATER      = ENTITY.CREATER("creater", "CREATER")
	ENTITY_CREATED      = ENTITY.CREATED("created", "CREATED")
	ENTITY_MODIFIER     = ENTITY.MODIFIER("modifier", "MODIFIER")
	ENTITY_MODIFIED     = ENTITY.MODIFIED("modified", "MODIFIED")
	ENTITY_VERSION      = ENTITY.VERSION("version", "VERSION")
	ENTITY_DELETION     = ENTITY.DELETION("deletion", "DELETION")
	ENTITY_ARTIFICAL    = ENTITY.COLUMN("artifical", "ARTIFICAL")
	ENTITY_HISTORY      = ENTITY.COLUMN("history", "HISTORY")
)

func NewEntity() *Entity {
	e := &Entity{}
	e.init()
	return e
}

func (me *Entity) init() {
	me.table = ENTITY

	if t, ok := me.Tree.Type("id"); ok {
		t.SetColumn(ENTITY_ID)
	}
	if t, ok := me.Tree.Type("code"); ok {
		t.SetColumn(ENTITY_CODE)
	}
	if t, ok := me.Tree.Type("name"); ok {
		t.SetColumn(ENTITY_NAME)
	}
	if t, ok := me.Tree.Type("fullname"); ok {
		t.SetColumn(ENTITY_FULLNAME)
	}
	if t, ok := me.Tree.Type("genre"); ok {
		t.SetColumn(ENTITY_GENRE)
	}
	if t, ok := me.Tree.Type("leaf"); ok {
		t.SetColumn(ENTITY_LEAF)
	}
	if t, ok := me.Tree.Type("grade"); ok {
		t.SetColumn(ENTITY_GRADE)
	}
	if t, ok := me.Tree.Type("ordinal"); ok {
		t.SetColumn(ENTITY_ORDINAL)
	}
	if t, ok := me.Tree.Type("parent_id"); ok {
		t.SetColumn(ENTITY_PARENT_ID)
	}
	if t, ok := me.Tree.Type("parent_ids"); ok {
		t.SetColumn(ENTITY_PARENT_IDS)
	}
	if t, ok := me.Tree.Type("parent_codes"); ok {
		t.SetColumn(ENTITY_PARENT_CODES)
	}
	if t, ok := me.Tree.Type("parent_names"); ok {
		t.SetColumn(ENTITY_PARENT_NAMES)
	}
	if t, ok := me.Tree.Type("memo"); ok {
		t.SetColumn(ENTITY_MEMO)
	}
	if t, ok := me.Tree.Type("creates"); ok {
		t.SetColumn(ENTITY_CREATES)
	}
	if t, ok := me.Tree.Type("creater"); ok {
		t.SetColumn(ENTITY_CREATER)
	}
	if t, ok := me.Tree.Type("created"); ok {
		t.SetColumn(ENTITY_CREATED)
	}
	if t, ok := me.Tree.Type("modifier"); ok {
		t.SetColumn(ENTITY_MODIFIER)
	}
	if t, ok := me.Tree.Type("modified"); ok {
		t.SetColumn(ENTITY_MODIFIED)
	}
	if t, ok := me.Tree.Type("version"); ok {
		t.SetColumn(ENTITY_VERSION)
	}
	if t, ok := me.Tree.Type("deletion"); ok {
		t.SetColumn(ENTITY_DELETION)
	}
	if t, ok := me.Tree.Type("artifical"); ok {
		t.SetColumn(ENTITY_ARTIFICAL)
	}
	if t, ok := me.Tree.Type("history"); ok {
		t.SetColumn(ENTITY_HISTORY)
	}

	if t, ok := me.Tree.Type("created"); ok {
		t.SetDefault("-62135596800")
	}
	if t, ok := me.Tree.Type("modified"); ok {
		t.SetDefault("-62135596800")
	}

	columns := []string{"id", "code", "name", "fullname", "genre", "leaf", "grade",
		"ordinal", "parent_id", "parent_ids", "parent_codes", "parent_names",
		"memo", "creates", "creater", "created", "modifier", "modified",
		"version", "deletion", "artifical", "history"}
	for _, c := range columns {
		if t, ok := me.Tree.Type(c); ok {
			t.SetField(entity.DefaultField())
		}
	}
}

func (me Entity) New() entity.Interface {
	return NewEntity()
}

func (me *Entity) Get(column string) interface{} {
	switch column {
	}
	return me.Tree.Get(column)
}

func (me *Entity) GetPtr(column string) interface{} {
	switch column {
	}
	return me.Tree.GetPtr(column)
}

func (me *Entity) Table() schema.Table {
	return me.table
}

func (me *Entity) Type(column string) (entity.Type, bool) {
	switch column {
	}
	return me.Tree.Type(column)
}

func (me *Entity) Column(field string) (schema.Column, bool) {
	switch strings.ToLowerFirst(field) {
	}
	return me.Tree.Column(field)
}

func (me *Entity) Columns() []schema.Column {
	return []schema.Column{
		ENTITY_ID,
		ENTITY_CODE,
		ENTITY_NAME,
		ENTITY_FULLNAME,
		ENTITY_GENRE,
		ENTITY_LEAF,
		ENTITY_GRADE,
		ENTITY_ORDINAL,
		ENTITY_PARENT_ID,
		ENTITY_PARENT_IDS,
		ENTITY_PARENT_CODES,
		ENTITY_PARENT_NAMES,
		ENTITY_MEMO,
		ENTITY_CREATES,
		ENTITY_CREATER,
		ENTITY_CREATED,
		ENTITY_MODIFIER,
		ENTITY_MODIFIED,
		ENTITY_VERSION,
		ENTITY_DELETION,
		ENTITY_ARTIFICAL,
		ENTITY_HISTORY,
	}
}

func (me *Entity) Names() []string {
	return []string{
		"id",
		"code",
		"name",
		"fullname",
		"genre",
		"leaf",
		"grade",
		"ordinal",
		"parentId",
		"parentIds",
		"parentCodes",
		"parentNames",
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
	}
}

func (me *Entity) Value() *Entity {
	return me
}

func (me *Entity) SetString(field, value string) error {
	switch strings.ToLowerFirst(field) {
	}
	return me.Tree.SetString(field, value)
}

func (me *Entity) Validate() error {
	return nil
}

func (me *Entity) JSON() string {
	var b bytes.Buffer
	b.WriteString("{")
	b.WriteString(fmt.Sprintf(`"id":%q`, me.Tree.Sys.Pk.Id()))
	b.WriteString(fmt.Sprintf(`,"memo":%q`, me.Tree.Sys.Memo()))
	b.WriteString(fmt.Sprintf(`,"creates":%q`, me.Tree.Sys.Creates()))
	b.WriteString(fmt.Sprintf(`,"creater":%q`, me.Tree.Sys.Creater()))
	b.WriteString(fmt.Sprintf(`,"created":%d`, me.Tree.Sys.Created()))
	b.WriteString(fmt.Sprintf(`,"modifier":%q`, me.Tree.Sys.Modifier()))
	b.WriteString(fmt.Sprintf(`,"modified":%d`, me.Tree.Sys.Modified()))
	b.WriteString(fmt.Sprintf(`,"version":%d`, me.Tree.Sys.Version()))
	b.WriteString(fmt.Sprintf(`,"deletion":%d`, me.Tree.Sys.Deletion()))
	b.WriteString(fmt.Sprintf(`,"artifical":%d`, me.Tree.Sys.Artifical()))
	b.WriteString(fmt.Sprintf(`,"history":%d`, me.Tree.Sys.History()))
	b.WriteString(fmt.Sprintf(`,"code":%q`, me.Tree.Code()))
	b.WriteString(fmt.Sprintf(`,"name":%q`, me.Tree.Name()))
	b.WriteString(fmt.Sprintf(`,"fullname":%q`, me.Tree.Fullname()))
	b.WriteString(fmt.Sprintf(`,"genre":%q`, me.Tree.Genre()))
	b.WriteString(fmt.Sprintf(`,"leaf":%d`, me.Tree.Leaf()))
	b.WriteString(fmt.Sprintf(`,"grade":%d`, me.Tree.Grade()))
	b.WriteString(fmt.Sprintf(`,"ordinal":%q`, me.Tree.Ordinal()))
	b.WriteString(fmt.Sprintf(`,"parentId":%q`, me.Tree.ParentId()))
	b.WriteString(fmt.Sprintf(`,"parentIds":%q`, me.Tree.ParentIds()))
	b.WriteString(fmt.Sprintf(`,"parentCodes":%q`, me.Tree.ParentCodes()))
	b.WriteString(fmt.Sprintf(`,"parentNames":%q`, me.Tree.ParentNames()))
	b.WriteString("}")
	return b.String()
}
