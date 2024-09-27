package goboom

func (node *GameObject) HasTag(tag string) bool {
	if node.Tags == nil {
		return false
	}
	for _, t := range node.Tags {
		if t == tag {
			return true
		}
	}

	return false
}

func (scene *GameObject) GetById(id string) *GameObject {
	for _, c := range scene.GetAll() {
		if c.Id == id {
			return c
		}
	}

	return nil
}


func (node *GameObject) GetAllTagged() []*GameObject {

	tagged := []*GameObject{}

	for _, c := range node.GetAll() {
		if len(c.Tags) > 0 {
			tagged = append(tagged, c)
		}
	}

	return tagged
}

func (node *GameObject) GetAll(tags ...string) []*GameObject {

	var objs = []*GameObject{}

	if node.GetParent() != nil {
		objs = append(objs, node)
	}

	for _, c := range node.GetChildren() {
		objs = append(objs, c.GetAll()...)
	}

	if len(tags) > 0 {
		return FilterByTags(objs, tags...)
	}

	return objs
}



func FilterByTags(objs []*GameObject, tags ...string) []*GameObject {

	withTags := []*GameObject{}

	for i, o := range objs {
		for _, t := range tags {
			if o.HasTag(t) {
				withTags = append(withTags, objs[i])
			}
		}
	}

	return withTags
}
