package mapper

import "go.mongodb.org/mongo-driver/bson/primitive"

type BaseMapper struct {
}

func CreateBaseMapper() BaseMapper {
	return BaseMapper{}
}

func (bm *BaseMapper) fromObjectIdToString(objectIds []primitive.ObjectID) []string {
	var ids []string

	for _, objectId := range objectIds {
		ids = append(ids, objectId.Hex())
	}

	return ids
}
