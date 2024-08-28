package main

import "fmt"

type Identity struct {
	Id string
	Tags []string
}

func (i *Identity) SetId(id string) {
	i.Id = id
}

func (i *Identity) AddTags(tags []string) {
	i.Tags = append(i.Tags, tags...)
}

func (i *Identity) GetId() string {
	fmt.Println("ID: ", i.Id)
	return i.Id
}

func (i *Identity) GetTags() []string {
	fmt.Println("Tags: ", i.Tags)
	return i.Tags
}

