package goboom

import "fmt"

type Identity struct {
	Id string
	Tags []string
}

func (me *Identity) SetId(id string) {
	me.Id = id
}

func (me *Identity) AddTags(tags ...string) {
	me.Tags = append(me.Tags, tags...)
}

func (me *Identity) RemoveTag(tag string) {
	for i, t := range me.Tags {
		if t == tag {
			me.Tags = append(me.Tags[:i], me.Tags[i+1:]...)
			break
		}
	}
}

func (i *Identity) GetId() string {
	return i.Id
}

func (i *Identity) GetTags() []string {
	fmt.Println("Tags: ", i.Tags)
	return i.Tags
}

