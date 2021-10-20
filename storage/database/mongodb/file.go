package mongodb

import (
	"context"
	"github.com/ppzxc/chattools/storage/database"
	"github.com/ppzxc/chattools/storage/database/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (m mongodb) FileInsert(ctx context.Context, file model.File) (int64, error) {
	sequence, err := m.crudSequence.Next(ctx, database.MongoCollectionFile)
	if err != nil {
		return 0, err
	}
	file.Id = sequence

	return file.Id, m.crudFile.InsertOne(ctx, file)
}

func (m mongodb) FileFindOneById(ctx context.Context, fileId int64) (model.File, error) {
	return m.crudFile.FindOneByFilter(ctx, bson.D{{"_id", fileId}})
}
