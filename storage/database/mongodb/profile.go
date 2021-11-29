package mongodb

import (
	"context"
	"github.com/ppzxc/chattools/storage/database"
	"github.com/ppzxc/chattools/storage/database/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m mongodb) ProfileFind(ctx context.Context, userId int64) (model.Profile, error) {
	filter, err := m.crudUser.FindOneByFilter(ctx, bson.D{{"_id", userId}})
	if err != nil {
		return model.Profile{}, err
	}
	return *filter.Profile, nil
}

func (m mongodb) ProfileImageFind(ctx context.Context, userId int64) (model.File, error) {
	filter, err := m.crudUser.FindOneByFilter(ctx, bson.D{{"_id", userId}})
	if err != nil {
		return model.File{}, err
	}

	if filter.Profile.FileId > 0 {
		return m.crudFile.FindOneByFilter(ctx, bson.D{{"_id", filter.Profile.FileId}})
	} else {
		return model.File{}, mongo.ErrNoDocuments
	}
}

func (m mongodb) ProfileImageUpdate(ctx context.Context, file model.File) (int64, error) {
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

func (m mongodb) ProfileUpdateByUserId(ctx context.Context, profile model.Profile) error {
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
