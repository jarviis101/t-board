package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseRepository struct {
}

func CreateBaseRepository() BaseRepository {
	return BaseRepository{}
}

func (br *BaseRepository) fromStringToObjectId(ids []string) []primitive.ObjectID {
	var objectIds []primitive.ObjectID

	for _, objectId := range ids {
		objectId, err := primitive.ObjectIDFromHex(objectId)
		if err != nil {
			continue
		}
		objectIds = append(objectIds, objectId)
	}

	return objectIds
}

//func (br *BaseRepository) removeElementFromFieldArray(
//	ctx context.Context,
//	collection *mongo.Collection,
//	filter bson.M,
//	fieldName string,
//	element primitive.ObjectID,
//) error {
//	update := bson.M{
//		"$pull": bson.M{fieldName: element},
//	}
//	_, err := collection.UpdateMany(ctx, filter, update, options.Update().SetUpsert(false))
//	if err != nil {
//		return err
//	}
//	return nil
//}
