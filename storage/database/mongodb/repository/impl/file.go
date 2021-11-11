package impl

import (
	"context"
	"errors"
	"fmt"
	"github.com/ppzxc/chattools/storage/database"
	"github.com/ppzxc/chattools/storage/database/model"
	"github.com/ppzxc/chattools/storage/database/mongodb/repository"
	"github.com/ppzxc/chattools/utils"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type file struct {
	database     *mongo.Database
	collection   *mongo.Collection
	queryTimeout time.Duration
}

func NewFileRepository(db *mongo.Database, queryTimeout time.Duration) repository.File {
	return &file{
		database:     db,
		collection:   db.Collection(database.MongoCollectionFile),
		queryTimeout: queryTimeout,
	}
}

func (c file) FindOneByFilter(ctx context.Context, filter bson.D) (model.File, error) {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	var file model.File
	start := time.Now()
	err := c.collection.FindOne(cCtx, filter).Decode(&file)
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "c.collection.FindOne",
		"exec.time": time.Since(start).String(),
		"args":      fmt.Sprintf("%+#v", filter),
	})).Debug("sql execute")
	return file, err
}

func (c file) InsertOne(ctx context.Context, file model.File) error {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	result, err := c.collection.InsertOne(cCtx, file)
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "c.collection.FindOne",
		"exec.time": time.Since(start).String(),
		"args":      fmt.Sprintf("%+#v", file),
	})).Debug("sql execute")
	if err != nil {
		return err
	}

	if result.InsertedID == nil {
		return errors.New("insert fail")
	}

	return nil
}
