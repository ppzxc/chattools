package mongodb

import (
	"context"
	"github.com/ppzxc/chattools/storage/database"
	model2 "github.com/ppzxc/chattools/storage/database/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (m mongodb) ProfileImageUpdate(ctx context.Context, file model2.File) (int64, error) {
	sequence, err := m.crudSeq.Next(ctx, database.MongoCollectionFile)
	if err != nil {
		return 0, err
	}
	file.Id = sequence

	if err := m.crudFile.InsertOne(ctx, file); err != nil {
		return 0, err
	}

	if err := m.crudUser.FindOneAndUpdateByFilter(
		ctx,
		bson.D{{"_id", file.FromUserId}},
		bson.D{{"$set", bson.M{
			"profile.file_id":    file.Id,
			"profile.updated_at": file.UpdatedAt,
		}}},
	); err != nil {
		return 0, err
	}

	return file.Id, nil
}

func (m mongodb) ProfileUpdateByUserId(ctx context.Context, profile model2.Profile) error {
	return m.crudUser.FindOneAndUpdateByFilter(
		ctx,
		bson.D{{"_id", profile.UserId}},
		bson.D{{"$set", bson.M{
			"profile.user_id":     profile.UserId,
			"profile.description": profile.Description,
			"profile.updated_at":  profile.UpdatedAt,
		}}},
	)
}
