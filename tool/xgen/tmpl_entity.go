// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplEntity = `// generated by xgen -- DO NOT EDIT
package {{.PackageName}}

import (
	"bytes"
	"fmt"{{if .IsTimeField}}
	"time"{{end}}

	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"{{if .IsValidationField}}
	"gopkg.in/goyy/goyy.v0/data/validate"{{end}}
	"gopkg.in/goyy/goyy.v0/util/strings"
){{range $e := .Entities}}

var (
	{{tname $e.Name $e.AllColumnMaxLen}} = schema.TABLE("{{lower $e.Table}}", "{{$e.GetComment}}"){{if eq $e.Extend "pk"}}
	{{uncamel $e.Name|upper}}_{{cname "id" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.PRIMARY("id", "` + i18N.Message("col.comment.id") + `"){{else if eq $e.Extend "sys"}}
	{{uncamel $e.Name|upper}}_{{cname "id" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.PRIMARY("id", "` + i18N.Message("col.comment.id") + `")
	{{uncamel $e.Name|upper}}_{{cname "memo" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.COLUMN("memo", "` + i18N.Message("col.comment.memo") + `")
	{{uncamel $e.Name|upper}}_{{cname "creates" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.COLUMN("creates", "` + i18N.Message("col.comment.creates") + `")
	{{uncamel $e.Name|upper}}_{{cname "creater" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.CREATER("creater", "` + i18N.Message("col.comment.creater") + `")
	{{uncamel $e.Name|upper}}_{{cname "created" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.CREATED("created", "` + i18N.Message("col.comment.created") + `")
	{{uncamel $e.Name|upper}}_{{cname "modifier" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.MODIFIER("modifier", "` + i18N.Message("col.comment.modifier") + `")
	{{uncamel $e.Name|upper}}_{{cname "modified" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.MODIFIED("modified", "` + i18N.Message("col.comment.modified") + `")
	{{uncamel $e.Name|upper}}_{{cname "version" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.VERSION("version", "` + i18N.Message("col.comment.version") + `")
	{{uncamel $e.Name|upper}}_{{cname "deletion" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.DELETION("deletion", "` + i18N.Message("col.comment.deletion") + `")
	{{uncamel $e.Name|upper}}_{{cname "artifical" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.COLUMN("artifical", "` + i18N.Message("col.comment.artifical") + `")
	{{uncamel $e.Name|upper}}_{{cname "history" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.COLUMN("history", "` + i18N.Message("col.comment.history") + `"){{else if eq $e.Extend "tree"}}
	{{uncamel $e.Name|upper}}_{{cname "id" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.PRIMARY("id", "` + i18N.Message("col.comment.id") + `")
	{{uncamel $e.Name|upper}}_{{cname "code" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.COLUMN("code", "` + i18N.Message("col.comment.code") + `")
	{{uncamel $e.Name|upper}}_{{cname "name" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.COLUMN("name", "` + i18N.Message("col.comment.name") + `")
	{{uncamel $e.Name|upper}}_{{cname "fullname" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.COLUMN("fullname", "` + i18N.Message("col.comment.fullname") + `")
	{{uncamel $e.Name|upper}}_{{cname "genre" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.COLUMN("genre", "` + i18N.Message("col.comment.genre") + `")
	{{uncamel $e.Name|upper}}_{{cname "leaf" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.COLUMN("leaf", "` + i18N.Message("col.comment.leaf") + `")
	{{uncamel $e.Name|upper}}_{{cname "grade" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.COLUMN("grade", "` + i18N.Message("col.comment.grade") + `")
	{{uncamel $e.Name|upper}}_{{cname "ordinal" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.COLUMN("ordinal", "` + i18N.Message("col.comment.ordinal") + `")
	{{uncamel $e.Name|upper}}_{{cname "parent_id" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.COLUMN("parent_id", "` + i18N.Message("col.comment.parent_id") + `")
	{{uncamel $e.Name|upper}}_{{cname "parent_ids" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.COLUMN("parent_ids", "` + i18N.Message("col.comment.parent_ids") + `")
	{{uncamel $e.Name|upper}}_{{cname "parent_codes" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.COLUMN("parent_codes", "` + i18N.Message("col.comment.parent_codes") + `")
	{{uncamel $e.Name|upper}}_{{cname "parent_names" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.COLUMN("parent_names", "` + i18N.Message("col.comment.parent_names") + `")
	{{uncamel $e.Name|upper}}_{{cname "memo" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.COLUMN("memo", "` + i18N.Message("col.comment.memo") + `")
	{{uncamel $e.Name|upper}}_{{cname "creates" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.COLUMN("creates", "` + i18N.Message("col.comment.creates") + `")
	{{uncamel $e.Name|upper}}_{{cname "creater" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.CREATER("creater", "` + i18N.Message("col.comment.creater") + `")
	{{uncamel $e.Name|upper}}_{{cname "created" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.CREATED("created", "` + i18N.Message("col.comment.created") + `")
	{{uncamel $e.Name|upper}}_{{cname "modifier" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.MODIFIER("modifier", "` + i18N.Message("col.comment.modifier") + `")
	{{uncamel $e.Name|upper}}_{{cname "modified" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.MODIFIED("modified", "` + i18N.Message("col.comment.modified") + `")
	{{uncamel $e.Name|upper}}_{{cname "version" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.VERSION("version", "` + i18N.Message("col.comment.version") + `")
	{{uncamel $e.Name|upper}}_{{cname "deletion" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.DELETION("deletion", "` + i18N.Message("col.comment.deletion") + `")
	{{uncamel $e.Name|upper}}_{{cname "artifical" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.COLUMN("artifical", "` + i18N.Message("col.comment.artifical") + `")
	{{uncamel $e.Name|upper}}_{{cname "history" $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.COLUMN("history", "` + i18N.Message("col.comment.history") + `"){{end}}{{range $f := $e.Fields}}{{if $f.IsPrimary}}
	{{uncamel $e.Name|upper}}_{{cname $f.Column $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.PRIMARY("{{lower $f.Column}}", "{{$f.GetComment}}"){{else if $f.IsVersion}}
	{{uncamel $e.Name|upper}}_{{cname $f.Column $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.VERSION("{{lower $f.Column}}", "{{$f.GetComment}}"){{else if $f.IsDeletion}}
	{{uncamel $e.Name|upper}}_{{cname $f.Column $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.DELETION("{{lower $f.Column}}", "{{$f.GetComment}}"){{else if $f.IsCreater}}
	{{uncamel $e.Name|upper}}_{{cname $f.Column $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.CREATER("{{lower $f.Column}}", "{{$f.GetComment}}"){{else if $f.IsCreated}}
	{{uncamel $e.Name|upper}}_{{cname $f.Column $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.CREATED("{{lower $f.Column}}", "{{$f.GetComment}}"){{else if $f.IsModifier}}
	{{uncamel $e.Name|upper}}_{{cname $f.Column $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.MODIFIER("{{lower $f.Column}}", "{{$f.GetComment}}"){{else if $f.IsModified}}
	{{uncamel $e.Name|upper}}_{{cname $f.Column $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.MODIFIED("{{lower $f.Column}}", "{{$f.GetComment}}"){{else if $f.IsTransient}}
	{{uncamel $e.Name|upper}}_{{cname $f.Column $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.TRANSIENT("{{lower $f.Column}}", "{{$f.GetComment}}"){{else}}
	{{uncamel $e.Name|upper}}_{{cname $f.Column $e.AllColumnMaxLen}} = {{uncamel $e.Name|upper}}.COLUMN("{{lower $f.Column}}", "{{$f.GetComment}}"){{end}}{{end}}
)

func New{{$e.Name}}() *{{$e.Name}} {
	e := &{{$e.Name}}{}
	e.init()
	return e
}
{{range $f := $e.Fields}}
func (me *{{$e.Name}}) {{upper1 $f.Name}}() {{$f.Type}} {
	return me.{{$f.Name}}.Value()
}

func (me *{{$e.Name}}) Set{{upper1 $f.Name}}(v {{$f.Type}}) {
	me.{{$f.Name}}.SetValue(v)
}
{{end}}
func (me *{{$e.Name}}) init() {
	me.table = {{uncamel $e.Name|upper}}
	me.initSetDict()
	me.initSetColumn()
	me.initSetDefault()
	me.initSetField()
}

func (me *{{$e.Name}}) initSetDict() {{"{"}}{{range $f := $e.Fields}}{{if notblank $f.Dict}}
	{{uncamel $e.Name|upper}}_{{upper $f.Column}}.SetDict("{{$f.Dict}}"){{end}}{{end}}
}

func (me *{{$e.Name}}) initSetColumn() {
	{{if eq $e.Extend "pk"}}if t, ok := me.Pk.Type("id"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "id"}})
	}{{else if eq $e.Extend "sys"}}if t, ok := me.Sys.Type("id"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "id"}})
	}
	if t, ok := me.Sys.Type("memo"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "memo"}})
	}
	if t, ok := me.Sys.Type("creates"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "creates"}})
	}
	if t, ok := me.Sys.Type("creater"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "creater"}})
	}
	if t, ok := me.Sys.Type("created"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "created"}})
	}
	if t, ok := me.Sys.Type("modifier"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "modifier"}})
	}
	if t, ok := me.Sys.Type("modified"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "modified"}})
	}
	if t, ok := me.Sys.Type("version"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "version"}})
	}
	if t, ok := me.Sys.Type("deletion"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "deletion"}})
	}
	if t, ok := me.Sys.Type("artifical"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "artifical"}})
	}
	if t, ok := me.Sys.Type("history"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "history"}})
	}{{else if eq $e.Extend "tree"}}if t, ok := me.Tree.Type("id"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "id"}})
	}
	if t, ok := me.Tree.Type("code"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "code"}})
	}
	if t, ok := me.Tree.Type("name"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "name"}})
	}
	if t, ok := me.Tree.Type("fullname"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "fullname"}})
	}
	if t, ok := me.Tree.Type("genre"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "genre"}})
	}
	if t, ok := me.Tree.Type("leaf"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "leaf"}})
	}
	if t, ok := me.Tree.Type("grade"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "grade"}})
	}
	if t, ok := me.Tree.Type("ordinal"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "ordinal"}})
	}
	if t, ok := me.Tree.Type("parent_id"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "parent_id"}})
	}
	if t, ok := me.Tree.Type("parent_ids"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "parent_ids"}})
	}
	if t, ok := me.Tree.Type("parent_codes"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "parent_codes"}})
	}
	if t, ok := me.Tree.Type("parent_names"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "parent_names"}})
	}
	if t, ok := me.Tree.Type("memo"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "memo"}})
	}
	if t, ok := me.Tree.Type("creates"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "creates"}})
	}
	if t, ok := me.Tree.Type("creater"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "creater"}})
	}
	if t, ok := me.Tree.Type("created"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "created"}})
	}
	if t, ok := me.Tree.Type("modifier"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "modifier"}})
	}
	if t, ok := me.Tree.Type("modified"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "modified"}})
	}
	if t, ok := me.Tree.Type("version"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "version"}})
	}
	if t, ok := me.Tree.Type("deletion"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "deletion"}})
	}
	if t, ok := me.Tree.Type("artifical"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "artifical"}})
	}
	if t, ok := me.Tree.Type("history"); ok {
		t.SetColumn({{uncamel $e.Name|upper}}_{{upper "history"}})
	}{{end}}{{range $f := $e.Fields}}
	me.{{$f.Name}}.SetColumn({{uncamel $e.Name|upper}}_{{upper $f.Column}}){{end}}
}

func (me *{{$e.Name}}) initSetDefault() {
	{{if eq $e.Extend "sys"}}if t, ok := me.Sys.Type("created"); ok {
		t.SetDefault("-62135596800")
	}
	if t, ok := me.Sys.Type("modified"); ok {
		t.SetDefault("-62135596800")
	}{{else if eq $e.Extend "tree"}}if t, ok := me.Tree.Type("created"); ok {
		t.SetDefault("-62135596800")
	}
	if t, ok := me.Tree.Type("modified"); ok {
		t.SetDefault("-62135596800")
	}{{end}}{{range $f := $e.Fields}}{{if notblank $f.Default}}
	me.{{$f.Name}}.SetDefault("{{$f.Default}}"){{end}}{{end}}
}

func (me *{{$e.Name}}) initSetField() {
	{{if eq $e.Extend "pk"}}if t, ok := me.Pk.Type("id"); ok {
		t.SetField(entity.DefaultField())
	}{{else if eq $e.Extend "sys"}}for _, c := range entity.SysColumns {
		if t, ok := me.Sys.Type(c); ok {
			t.SetField(entity.DefaultField())
		}
	}{{else if eq $e.Extend "tree"}}for _, c := range entity.TreeColumns {
		if t, ok := me.Tree.Type(c); ok {
			t.SetField(entity.DefaultField())
		}
	}{{end}}{{range $f := $e.Fields}}
	me.{{$f.Name}}.SetField(entity.DefaultField()){{end}}
}

func (me {{$e.Name}}) New() entity.Interface {
	return New{{$e.Name}}()
}

func (me *{{$e.Name}}) Get(column string) interface{} {
	switch column {{"{"}}{{range $f := $e.Fields}}
	case {{uncamel $e.Name|upper}}_{{uncamel $f.Column|upper}}.Name():
		return me.{{$f.Name}}.Value(){{end}}
	}{{if eq $e.Extend "pk"}}
	return me.Pk.Get(column){{else if eq $e.Extend "sys"}}
	return me.Sys.Get(column){{else if eq $e.Extend "tree"}}
	return me.Tree.Get(column){{else}}
	return nil{{end}}
}

func (me *{{$e.Name}}) GetPtr(column string) interface{} {
	switch column {{"{"}}{{range $f := $e.Fields}}
	case {{uncamel $e.Name|upper}}_{{uncamel $f.Column|upper}}.Name():
		return me.{{$f.Name}}.ValuePtr(){{end}}
	}{{if eq $e.Extend "pk"}}
	return me.Pk.GetPtr(column){{else if eq $e.Extend "sys"}}
	return me.Sys.GetPtr(column){{else if eq $e.Extend "tree"}}
	return me.Tree.GetPtr(column){{else}}
	return nil{{end}}
}

func (me *{{$e.Name}}) Table() schema.Table {
	return me.table
}

func (me *{{$e.Name}}) Type(column string) (entity.Type, bool) {
	switch column {{"{"}}{{range $f := $e.Fields}}
	case {{uncamel $e.Name|upper}}_{{uncamel $f.Column|upper}}.Name():
		return &me.{{$f.Name}}, true{{end}}
	}{{if eq $e.Extend "pk"}}
	return me.Pk.Type(column){{else if eq $e.Extend "sys"}}
	return me.Sys.Type(column){{else if eq $e.Extend "tree"}}
	return me.Tree.Type(column){{else}}
	return nil, false{{end}}
}

func (me *{{$e.Name}}) Column(field string) (schema.Column, bool) {
	switch strings.ToLowerFirst(field) {{"{"}}{{range $f := $e.Fields}}
	case "{{$f.Name}}":
		return {{uncamel $e.Name|upper}}_{{uncamel $f.Column|upper}}, true{{end}}
	}{{if eq $e.Extend "pk"}}
	return me.Pk.Column(field){{else if eq $e.Extend "sys"}}
	return me.Sys.Column(field){{else if eq $e.Extend "tree"}}
	return me.Tree.Column(field){{else}}
	return nil, false{{end}}
}

func (me *{{$e.Name}}) Columns() []schema.Column {
	return []schema.Column{{"{"}}{{if eq $e.Extend "pk"}}
		{{uncamel $e.Name|upper}}_ID,{{else if eq $e.Extend "sys"}}
		{{uncamel $e.Name|upper}}_ID,
		{{uncamel $e.Name|upper}}_MEMO,
		{{uncamel $e.Name|upper}}_CREATES,
		{{uncamel $e.Name|upper}}_CREATER,
		{{uncamel $e.Name|upper}}_CREATED,
		{{uncamel $e.Name|upper}}_MODIFIER,
		{{uncamel $e.Name|upper}}_MODIFIED,
		{{uncamel $e.Name|upper}}_VERSION,
		{{uncamel $e.Name|upper}}_DELETION,
		{{uncamel $e.Name|upper}}_ARTIFICAL,
		{{uncamel $e.Name|upper}}_HISTORY,{{else if eq $e.Extend "tree"}}
		{{uncamel $e.Name|upper}}_ID,
		{{uncamel $e.Name|upper}}_CODE,
		{{uncamel $e.Name|upper}}_NAME,
		{{uncamel $e.Name|upper}}_FULLNAME,
		{{uncamel $e.Name|upper}}_GENRE,
		{{uncamel $e.Name|upper}}_LEAF,
		{{uncamel $e.Name|upper}}_GRADE,
		{{uncamel $e.Name|upper}}_ORDINAL,
		{{uncamel $e.Name|upper}}_PARENT_ID,
		{{uncamel $e.Name|upper}}_PARENT_IDS,
		{{uncamel $e.Name|upper}}_PARENT_CODES,
		{{uncamel $e.Name|upper}}_PARENT_NAMES,
		{{uncamel $e.Name|upper}}_MEMO,
		{{uncamel $e.Name|upper}}_CREATES,
		{{uncamel $e.Name|upper}}_CREATER,
		{{uncamel $e.Name|upper}}_CREATED,
		{{uncamel $e.Name|upper}}_MODIFIER,
		{{uncamel $e.Name|upper}}_MODIFIED,
		{{uncamel $e.Name|upper}}_VERSION,
		{{uncamel $e.Name|upper}}_DELETION,
		{{uncamel $e.Name|upper}}_ARTIFICAL,
		{{uncamel $e.Name|upper}}_HISTORY,{{end}}{{range $f := $e.Fields}}
		{{uncamel $e.Name|upper}}_{{uncamel $f.Column|upper}},{{end}}
	}
}

func (me *{{$e.Name}}) Names() []string {
	return []string{{"{"}}{{if eq $e.Extend "pk"}}
		"id",{{else if eq $e.Extend "sys"}}
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
		"history",{{else if eq $e.Extend "tree"}}
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
		"history",{{end}}{{range $f := $e.Fields}}
		"{{$f.Name}}",{{end}}
	}
}

func (me *{{$e.Name}}) Value() *{{$e.Name}} {
	return me
}

func (me *{{$e.Name}}) SetString(field, value string) error {
	switch strings.ToLowerFirst(field) {{"{"}}{{range $f := $e.Fields}}
	case "{{$f.Name}}":
		return me.{{$f.Name}}.SetString(value){{end}}
	}{{if eq $e.Extend "pk"}}
	return me.Pk.SetString(field, value){{else if eq $e.Extend "sys"}}
	return me.Sys.SetString(field, value){{else if eq $e.Extend "tree"}}
	return me.Tree.SetString(field, value){{else}}
	return nil{{end}}
}

func (me *{{$e.Name}}) Validate() error {{"{"}}{{range $f := $e.Fields}}{{range $v := $f.Validations}}{{if eq $v.Name "required"}}
	if err := validate.Required(me.{{$f.Name}}.String()); err != nil {
		return err
	}{{end}}{{if eq $v.Name "min"}}
	if err := validate.Min(me.{{$f.Name}}.String(), {{$v.Value}}); err != nil {
		return err
	}{{end}}{{if eq $v.Name "max"}}
	if err := validate.Max(me.{{$f.Name}}.String(), {{$v.Value}}); err != nil {
		return err
	}{{end}}{{if eq $v.Name "range"}}
	if err := validate.Range(me.{{$f.Name}}.String(), {{$v.Value}}); err != nil {
		return err
	}{{end}}{{if eq $v.Name "minlen"}}
	if err := validate.Minlen(me.{{$f.Name}}.String(), {{$v.Value}}); err != nil {
		return err
	}{{end}}{{if eq $v.Name "maxlen"}}
	if err := validate.Maxlen(me.{{$f.Name}}.String(), {{$v.Value}}); err != nil {
		return err
	}{{end}}{{if eq $v.Name "rangelen"}}
	if err := validate.Rangelen(me.{{$f.Name}}.String(), {{$v.Value}}); err != nil {
		return err
	}{{end}}{{if eq $v.Name "email"}}
	if err := validate.Email(me.{{$f.Name}}.String()); err != nil {
		return err
	}{{end}}{{if eq $v.Name "url"}}
	if err := validate.URL(me.{{$f.Name}}.String()); err != nil {
		return err
	}{{end}}{{if eq $v.Name "ip"}}
	if err := validate.IP(me.{{$f.Name}}.String()); err != nil {
		return err
	}{{end}}{{if eq $v.Name "mobile"}}
	if err := validate.Mobile(me.{{$f.Name}}.String()); err != nil {
		return err
	}{{end}}{{if eq $v.Name "tel"}}
	if err := validate.Tel(me.{{$f.Name}}.String()); err != nil {
		return err
	}{{end}}{{if eq $v.Name "phone"}}
	if err := validate.Phone(me.{{$f.Name}}.String()); err != nil {
		return err
	}{{end}}{{if eq $v.Name "zipcode"}}
	if err := validate.Zipcode(me.{{$f.Name}}.String()); err != nil {
		return err
	}{{end}}{{if eq $v.Name "float"}}
	if err := validate.Float(me.{{$f.Name}}.String()); err != nil {
		return err
	}{{end}}{{if eq $v.Name "integer"}}
	if err := validate.Integer(me.{{$f.Name}}.String()); err != nil {
		return err
	}{{end}}{{if eq $v.Name "alpha"}}
	if err := validate.Alpha(me.{{$f.Name}}.String()); err != nil {
		return err
	}{{end}}{{if eq $v.Name "alrod"}}
	if err := validate.Alrod(me.{{$f.Name}}.String()); err != nil {
		return err
	}{{end}}{{if eq $v.Name "alnum"}}
	if err := validate.Alnum(me.{{$f.Name}}.String()); err != nil {
		return err
	}{{end}}{{if eq $v.Name "alnumrod"}}
	if err := validate.Alnumrod(me.{{$f.Name}}.String()); err != nil {
		return err
	}{{end}}{{if eq $v.Name "alnumhan"}}
	if err := validate.Alnumhan(me.{{$f.Name}}.String()); err != nil {
		return err
	}{{end}}{{if eq $v.Name "alnumhanrod"}}
	if err := validate.Alnumhanrod(me.{{$f.Name}}.String()); err != nil {
		return err
	}{{end}}{{if eq $v.Name "alhan"}}
	if err := validate.Alhan(me.{{$f.Name}}.String()); err != nil {
		return err
	}{{end}}{{if eq $v.Name "alhanrod"}}
	if err := validate.Alhanrod(me.{{$f.Name}}.String()); err != nil {
		return err
	}{{end}}{{if eq $v.Name "han"}}
	if err := validate.Han(me.{{$f.Name}}.String()); err != nil {
		return err
	}{{end}}{{if eq $v.Name "hanrod"}}
	if err := validate.Hanrod(me.{{$f.Name}}.String()); err != nil {
		return err
	}{{end}}{{end}}{{end}}
	return nil
}

func (me *{{$e.Name}}) JSON() string {
	var b bytes.Buffer
	b.WriteString("{"){{if eq $e.Extend "pk"}}
	b.WriteString(fmt.Sprintf(` + "`" + `"id":%q` + "`" + `, me.Sys.Pk.Id())){{else if eq $e.Extend "sys"}}
	b.WriteString(fmt.Sprintf(` + "`" + `"id":%q` + "`" + `, me.Sys.Pk.Id()))
	b.WriteString(fmt.Sprintf(` + "`," + `"memo":%q` + "`" + `, me.Sys.Memo()))
	b.WriteString(fmt.Sprintf(` + "`," + `"creates":%q` + "`" + `, me.Sys.Creates()))
	b.WriteString(fmt.Sprintf(` + "`," + `"creater":%q` + "`" + `, me.Sys.Creater()))
	b.WriteString(fmt.Sprintf(` + "`," + `"created":%d` + "`" + `, me.Sys.Created()))
	b.WriteString(fmt.Sprintf(` + "`," + `"modifier":%q` + "`" + `, me.Sys.Modifier()))
	b.WriteString(fmt.Sprintf(` + "`," + `"modified":%d` + "`" + `, me.Sys.Modified()))
	b.WriteString(fmt.Sprintf(` + "`," + `"version":%d` + "`" + `, me.Sys.Version()))
	b.WriteString(fmt.Sprintf(` + "`," + `"deletion":%d` + "`" + `, me.Sys.Deletion()))
	b.WriteString(fmt.Sprintf(` + "`," + `"artifical":%d` + "`" + `, me.Sys.Artifical()))
	b.WriteString(fmt.Sprintf(` + "`," + `"history":%d` + "`" + `, me.Sys.History())){{else if eq $e.Extend "tree"}}
	b.WriteString(fmt.Sprintf(` + "`" + `"id":%q` + "`" + `, me.Tree.Sys.Pk.Id()))
	b.WriteString(fmt.Sprintf(` + "`," + `"memo":%q` + "`" + `, me.Tree.Sys.Memo()))
	b.WriteString(fmt.Sprintf(` + "`," + `"creates":%q` + "`" + `, me.Tree.Sys.Creates()))
	b.WriteString(fmt.Sprintf(` + "`," + `"creater":%q` + "`" + `, me.Tree.Sys.Creater()))
	b.WriteString(fmt.Sprintf(` + "`," + `"created":%d` + "`" + `, me.Tree.Sys.Created()))
	b.WriteString(fmt.Sprintf(` + "`," + `"modifier":%q` + "`" + `, me.Tree.Sys.Modifier()))
	b.WriteString(fmt.Sprintf(` + "`," + `"modified":%d` + "`" + `, me.Tree.Sys.Modified()))
	b.WriteString(fmt.Sprintf(` + "`," + `"version":%d` + "`" + `, me.Tree.Sys.Version()))
	b.WriteString(fmt.Sprintf(` + "`," + `"deletion":%d` + "`" + `, me.Tree.Sys.Deletion()))
	b.WriteString(fmt.Sprintf(` + "`," + `"artifical":%d` + "`" + `, me.Tree.Sys.Artifical()))
	b.WriteString(fmt.Sprintf(` + "`," + `"history":%d` + "`" + `, me.Tree.Sys.History()))
	b.WriteString(fmt.Sprintf(` + "`," + `"code":%q` + "`" + `, me.Tree.Code()))
	b.WriteString(fmt.Sprintf(` + "`," + `"name":%q` + "`" + `, me.Tree.Name()))
	b.WriteString(fmt.Sprintf(` + "`," + `"fullname":%q` + "`" + `, me.Tree.Fullname()))
	b.WriteString(fmt.Sprintf(` + "`," + `"genre":%q` + "`" + `, me.Tree.Genre()))
	b.WriteString(fmt.Sprintf(` + "`," + `"leaf":%d` + "`" + `, me.Tree.Leaf()))
	b.WriteString(fmt.Sprintf(` + "`," + `"grade":%d` + "`" + `, me.Tree.Grade()))
	b.WriteString(fmt.Sprintf(` + "`," + `"ordinal":%q` + "`" + `, me.Tree.Ordinal()))
	b.WriteString(fmt.Sprintf(` + "`," + `"parentId":%q` + "`" + `, me.Tree.ParentId()))
	b.WriteString(fmt.Sprintf(` + "`," + `"parentIds":%q` + "`" + `, me.Tree.ParentIds()))
	b.WriteString(fmt.Sprintf(` + "`," + `"parentCodes":%q` + "`" + `, me.Tree.ParentCodes()))
	b.WriteString(fmt.Sprintf(` + "`," + `"parentNames":%q` + "`" + `, me.Tree.ParentNames())){{end}}{{range $f := $e.Fields}}
	b.WriteString(fmt.Sprintf(` + "`," + `"{{$f.Name}}":%q` + "`" + `, me.{{$f.Name}}.String())){{end}}
	b.WriteString("}")
	return b.String()
}{{end}}
`
