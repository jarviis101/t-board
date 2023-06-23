package transformers

type BaseTransformer struct {
}

func CreateBaseTransformer() BaseTransformer {
	return BaseTransformer{}
}

func (bt *BaseTransformer) modifyIds(i []string) []*string {
	var ids []*string

	for _, id := range i {
		ids = append(ids, &id)
	}

	return ids
}
