// generated by xgen -- DO NOT EDIT
package user

import (
	"bytes"
	"fmt"

	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

var (
	ENTITY            = schema.TABLE("sys_user", "USER")
	ENTITY_ID         = ENTITY.PRIMARY("id", "ID")
	ENTITY_MEMO       = ENTITY.COLUMN("memo", "MEMO")
	ENTITY_CREATES    = ENTITY.COLUMN("creates", "CREATES")
	ENTITY_CREATER    = ENTITY.CREATER("creater", "CREATER")
	ENTITY_CREATED    = ENTITY.CREATED("created", "CREATED")
	ENTITY_MODIFIER   = ENTITY.MODIFIER("modifier", "MODIFIER")
	ENTITY_MODIFIED   = ENTITY.MODIFIED("modified", "MODIFIED")
	ENTITY_VERSION    = ENTITY.VERSION("version", "VERSION")
	ENTITY_DELETION   = ENTITY.DELETION("deletion", "DELETION")
	ENTITY_ARTIFICAL  = ENTITY.COLUMN("artifical", "ARTIFICAL")
	ENTITY_HISTORY    = ENTITY.COLUMN("history", "HISTORY")
	ENTITY_NAME       = ENTITY.COLUMN("name", "NAME")
	ENTITY_CODE       = ENTITY.COLUMN("code", "CODE")
	ENTITY_PASSWD     = ENTITY.COLUMN("passwd", "PASSWD")
	ENTITY_GENRE      = ENTITY.COLUMN("genre", "GENRE")
	ENTITY_EMAIL      = ENTITY.COLUMN("email", "EMAIL")
	ENTITY_TEL        = ENTITY.COLUMN("tel", "TEL")
	ENTITY_MOBILE     = ENTITY.COLUMN("mobile", "MOBILE")
	ENTITY_AREA_ID    = ENTITY.COLUMN("area_id", "AREA_ID")
	ENTITY_ORG_ID     = ENTITY.COLUMN("org_id", "ORG_ID")
	ENTITY_LOGIN_NAME = ENTITY.COLUMN("login_name", "LOGIN_NAME")
	ENTITY_LOGIN_IP   = ENTITY.COLUMN("login_ip", "LOGIN_IP")
	ENTITY_LOGIN_TIME = ENTITY.COLUMN("login_time", "LOGIN_TIME")
)

func NewEntity() *Entity {
	e := &Entity{}
	e.init()
	return e
}

func (me *Entity) Name() string {
	return me.name.Value()
}

func (me *Entity) SetName(v string) {
	me.name.SetValue(v)
}

func (me *Entity) Code() string {
	return me.code.Value()
}

func (me *Entity) SetCode(v string) {
	me.code.SetValue(v)
}

func (me *Entity) Passwd() string {
	return me.passwd.Value()
}

func (me *Entity) SetPasswd(v string) {
	me.passwd.SetValue(v)
}

func (me *Entity) Genre() string {
	return me.genre.Value()
}

func (me *Entity) SetGenre(v string) {
	me.genre.SetValue(v)
}

func (me *Entity) Email() string {
	return me.email.Value()
}

func (me *Entity) SetEmail(v string) {
	me.email.SetValue(v)
}

func (me *Entity) Tel() string {
	return me.tel.Value()
}

func (me *Entity) SetTel(v string) {
	me.tel.SetValue(v)
}

func (me *Entity) Mobile() string {
	return me.mobile.Value()
}

func (me *Entity) SetMobile(v string) {
	me.mobile.SetValue(v)
}

func (me *Entity) AreaId() string {
	return me.areaId.Value()
}

func (me *Entity) SetAreaId(v string) {
	me.areaId.SetValue(v)
}

func (me *Entity) OrgId() string {
	return me.orgId.Value()
}

func (me *Entity) SetOrgId(v string) {
	me.orgId.SetValue(v)
}

func (me *Entity) LoginName() string {
	return me.loginName.Value()
}

func (me *Entity) SetLoginName(v string) {
	me.loginName.SetValue(v)
}

func (me *Entity) LoginIp() string {
	return me.loginIp.Value()
}

func (me *Entity) SetLoginIp(v string) {
	me.loginIp.SetValue(v)
}

func (me *Entity) LoginTime() int64 {
	return me.loginTime.Value()
}

func (me *Entity) SetLoginTime(v int64) {
	me.loginTime.SetValue(v)
}

func (me *Entity) init() {
	me.table = ENTITY
	me.initSetDict()
	me.initSetColumn()
	me.initSetDefault()
	me.initSetField()
}

func (me *Entity) initSetDict() {
}

func (me *Entity) initSetColumn() {
	if t, ok := me.Sys.Type("id"); ok {
		t.SetColumn(ENTITY_ID)
	}
	if t, ok := me.Sys.Type("memo"); ok {
		t.SetColumn(ENTITY_MEMO)
	}
	if t, ok := me.Sys.Type("creates"); ok {
		t.SetColumn(ENTITY_CREATES)
	}
	if t, ok := me.Sys.Type("creater"); ok {
		t.SetColumn(ENTITY_CREATER)
	}
	if t, ok := me.Sys.Type("created"); ok {
		t.SetColumn(ENTITY_CREATED)
	}
	if t, ok := me.Sys.Type("modifier"); ok {
		t.SetColumn(ENTITY_MODIFIER)
	}
	if t, ok := me.Sys.Type("modified"); ok {
		t.SetColumn(ENTITY_MODIFIED)
	}
	if t, ok := me.Sys.Type("version"); ok {
		t.SetColumn(ENTITY_VERSION)
	}
	if t, ok := me.Sys.Type("deletion"); ok {
		t.SetColumn(ENTITY_DELETION)
	}
	if t, ok := me.Sys.Type("artifical"); ok {
		t.SetColumn(ENTITY_ARTIFICAL)
	}
	if t, ok := me.Sys.Type("history"); ok {
		t.SetColumn(ENTITY_HISTORY)
	}
	me.name.SetColumn(ENTITY_NAME)
	me.code.SetColumn(ENTITY_CODE)
	me.passwd.SetColumn(ENTITY_PASSWD)
	me.genre.SetColumn(ENTITY_GENRE)
	me.email.SetColumn(ENTITY_EMAIL)
	me.tel.SetColumn(ENTITY_TEL)
	me.mobile.SetColumn(ENTITY_MOBILE)
	me.areaId.SetColumn(ENTITY_AREA_ID)
	me.orgId.SetColumn(ENTITY_ORG_ID)
	me.loginName.SetColumn(ENTITY_LOGIN_NAME)
	me.loginIp.SetColumn(ENTITY_LOGIN_IP)
	me.loginTime.SetColumn(ENTITY_LOGIN_TIME)
}

func (me *Entity) initSetDefault() {
	if t, ok := me.Sys.Type("created"); ok {
		t.SetDefault("-62135596800")
	}
	if t, ok := me.Sys.Type("modified"); ok {
		t.SetDefault("-62135596800")
	}
}

func (me *Entity) initSetField() {
	for _, c := range entity.SysColumns {
		if t, ok := me.Sys.Type(c); ok {
			t.SetField(entity.DefaultField())
		}
	}
	me.name.SetField(entity.DefaultField())
	me.code.SetField(entity.DefaultField())
	me.passwd.SetField(entity.DefaultField())
	me.genre.SetField(entity.DefaultField())
	me.email.SetField(entity.DefaultField())
	me.tel.SetField(entity.DefaultField())
	me.mobile.SetField(entity.DefaultField())
	me.areaId.SetField(entity.DefaultField())
	me.orgId.SetField(entity.DefaultField())
	me.loginName.SetField(entity.DefaultField())
	me.loginIp.SetField(entity.DefaultField())
	me.loginTime.SetField(entity.DefaultField())
}

func (me Entity) New() entity.Interface {
	return NewEntity()
}

func (me *Entity) Get(column string) interface{} {
	switch column {
	case ENTITY_NAME.Name():
		return me.name.Value()
	case ENTITY_CODE.Name():
		return me.code.Value()
	case ENTITY_PASSWD.Name():
		return me.passwd.Value()
	case ENTITY_GENRE.Name():
		return me.genre.Value()
	case ENTITY_EMAIL.Name():
		return me.email.Value()
	case ENTITY_TEL.Name():
		return me.tel.Value()
	case ENTITY_MOBILE.Name():
		return me.mobile.Value()
	case ENTITY_AREA_ID.Name():
		return me.areaId.Value()
	case ENTITY_ORG_ID.Name():
		return me.orgId.Value()
	case ENTITY_LOGIN_NAME.Name():
		return me.loginName.Value()
	case ENTITY_LOGIN_IP.Name():
		return me.loginIp.Value()
	case ENTITY_LOGIN_TIME.Name():
		return me.loginTime.Value()
	}
	return me.Sys.Get(column)
}

func (me *Entity) GetPtr(column string) interface{} {
	switch column {
	case ENTITY_NAME.Name():
		return me.name.ValuePtr()
	case ENTITY_CODE.Name():
		return me.code.ValuePtr()
	case ENTITY_PASSWD.Name():
		return me.passwd.ValuePtr()
	case ENTITY_GENRE.Name():
		return me.genre.ValuePtr()
	case ENTITY_EMAIL.Name():
		return me.email.ValuePtr()
	case ENTITY_TEL.Name():
		return me.tel.ValuePtr()
	case ENTITY_MOBILE.Name():
		return me.mobile.ValuePtr()
	case ENTITY_AREA_ID.Name():
		return me.areaId.ValuePtr()
	case ENTITY_ORG_ID.Name():
		return me.orgId.ValuePtr()
	case ENTITY_LOGIN_NAME.Name():
		return me.loginName.ValuePtr()
	case ENTITY_LOGIN_IP.Name():
		return me.loginIp.ValuePtr()
	case ENTITY_LOGIN_TIME.Name():
		return me.loginTime.ValuePtr()
	}
	return me.Sys.GetPtr(column)
}

func (me *Entity) Table() schema.Table {
	return me.table
}

func (me *Entity) Type(column string) (entity.Type, bool) {
	switch column {
	case ENTITY_NAME.Name():
		return &me.name, true
	case ENTITY_CODE.Name():
		return &me.code, true
	case ENTITY_PASSWD.Name():
		return &me.passwd, true
	case ENTITY_GENRE.Name():
		return &me.genre, true
	case ENTITY_EMAIL.Name():
		return &me.email, true
	case ENTITY_TEL.Name():
		return &me.tel, true
	case ENTITY_MOBILE.Name():
		return &me.mobile, true
	case ENTITY_AREA_ID.Name():
		return &me.areaId, true
	case ENTITY_ORG_ID.Name():
		return &me.orgId, true
	case ENTITY_LOGIN_NAME.Name():
		return &me.loginName, true
	case ENTITY_LOGIN_IP.Name():
		return &me.loginIp, true
	case ENTITY_LOGIN_TIME.Name():
		return &me.loginTime, true
	}
	return me.Sys.Type(column)
}

func (me *Entity) Column(field string) (schema.Column, bool) {
	switch strings.ToLowerFirst(field) {
	case "name":
		return ENTITY_NAME, true
	case "code":
		return ENTITY_CODE, true
	case "passwd":
		return ENTITY_PASSWD, true
	case "genre":
		return ENTITY_GENRE, true
	case "email":
		return ENTITY_EMAIL, true
	case "tel":
		return ENTITY_TEL, true
	case "mobile":
		return ENTITY_MOBILE, true
	case "areaId":
		return ENTITY_AREA_ID, true
	case "orgId":
		return ENTITY_ORG_ID, true
	case "loginName":
		return ENTITY_LOGIN_NAME, true
	case "loginIp":
		return ENTITY_LOGIN_IP, true
	case "loginTime":
		return ENTITY_LOGIN_TIME, true
	}
	return me.Sys.Column(field)
}

func (me *Entity) Columns() []schema.Column {
	return []schema.Column{
		ENTITY_ID,
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
		ENTITY_NAME,
		ENTITY_CODE,
		ENTITY_PASSWD,
		ENTITY_GENRE,
		ENTITY_EMAIL,
		ENTITY_TEL,
		ENTITY_MOBILE,
		ENTITY_AREA_ID,
		ENTITY_ORG_ID,
		ENTITY_LOGIN_NAME,
		ENTITY_LOGIN_IP,
		ENTITY_LOGIN_TIME,
	}
}

func (me *Entity) Names() []string {
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
		"name",
		"code",
		"passwd",
		"genre",
		"email",
		"tel",
		"mobile",
		"areaId",
		"orgId",
		"loginName",
		"loginIp",
		"loginTime",
	}
}

func (me *Entity) Value() *Entity {
	return me
}

func (me *Entity) SetString(field, value string) error {
	switch strings.ToLowerFirst(field) {
	case "name":
		return me.name.SetString(value)
	case "code":
		return me.code.SetString(value)
	case "passwd":
		return me.passwd.SetString(value)
	case "genre":
		return me.genre.SetString(value)
	case "email":
		return me.email.SetString(value)
	case "tel":
		return me.tel.SetString(value)
	case "mobile":
		return me.mobile.SetString(value)
	case "areaId":
		return me.areaId.SetString(value)
	case "orgId":
		return me.orgId.SetString(value)
	case "loginName":
		return me.loginName.SetString(value)
	case "loginIp":
		return me.loginIp.SetString(value)
	case "loginTime":
		return me.loginTime.SetString(value)
	}
	return me.Sys.SetString(field, value)
}

func (me *Entity) Validate() error {
	return nil
}

func (me *Entity) JSON() string {
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
	b.WriteString(fmt.Sprintf(`,"name":%q`, me.name.String()))
	b.WriteString(fmt.Sprintf(`,"code":%q`, me.code.String()))
	b.WriteString(fmt.Sprintf(`,"passwd":%q`, me.passwd.String()))
	b.WriteString(fmt.Sprintf(`,"genre":%q`, me.genre.String()))
	b.WriteString(fmt.Sprintf(`,"email":%q`, me.email.String()))
	b.WriteString(fmt.Sprintf(`,"tel":%q`, me.tel.String()))
	b.WriteString(fmt.Sprintf(`,"mobile":%q`, me.mobile.String()))
	b.WriteString(fmt.Sprintf(`,"areaId":%q`, me.areaId.String()))
	b.WriteString(fmt.Sprintf(`,"orgId":%q`, me.orgId.String()))
	b.WriteString(fmt.Sprintf(`,"loginName":%q`, me.loginName.String()))
	b.WriteString(fmt.Sprintf(`,"loginIp":%q`, me.loginIp.String()))
	b.WriteString(fmt.Sprintf(`,"loginTime":%q`, me.loginTime.String()))
	b.WriteString("}")
	return b.String()
}
