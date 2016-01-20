// generated by xgen -- DO NOT EDIT
package dict

import (
	"bytes"

	"gopkg.in/goyy/goyy.v0/data/entity"
)

func NewEntities(v int) *Entities {
	entities := &Entities{}
	entities.Make(v)
	return entities
}

type Entities struct {
	datas []*Entity
}

func (me *Entities) Make(cap int) {
	me.datas = make([]*Entity, 0, cap)
}

func (me *Entities) New() entity.Interface {
	return NewEntity()
}

func (me *Entities) Append(v entity.Interface) {
	me.datas = append(me.datas, v.(*Entity))
}

func (me *Entities) Len() int {
	return len(me.datas)
}

func (me *Entities) Cap() int {
	return cap(me.datas)
}

func (me *Entities) Index(v int) entity.Interface {
	return me.datas[v]
}

func (me *Entities) Slice() interface{} {
	return me.datas
}

func (me *Entities) Value(v int) *Entity {
	return me.datas[v]
}

func (me *Entities) Values() []*Entity {
	return me.datas
}

func (me *Entities) JSON() string {
	var b bytes.Buffer
	b.WriteString(`{"datas":[`)
	for i := 0; i < me.Len(); i++ {
		if i > 0 {
			b.WriteString(",")
		}
		b.WriteString(me.Index(i).JSON())
	}
	b.WriteString("]}")
	return b.String()
}
