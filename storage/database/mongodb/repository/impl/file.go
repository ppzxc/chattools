package impl

import (
	"context"
	"errors"
	"fmt"
	"github.com/ppzxc/chattools/common/stats"
	"github.com/ppzxc/chattools/storage/database/model"
	"github.com/ppzxc/chattools/storage/database/mongodb/repository"
	"github.com/ppzxc/chattools/utils"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type file struct {
	//database     *mongo.Database
	collection   *mongo.Collection
	queryTimeout time.Duration
}

func NewFileRepository(collection *mongo.Collection, queryTimeout time.Duration) repository.File {
	return &file{
		//database:     db,
		collection:   collection,
		queryTimeout: queryTimeout,
	}
}

func (c file) DeleteOneByFilter(ctx context.Context, filter bson.D) (model.File, error) {
	start := time.Now()
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	var file model.File
	err := c.collection.FindOneAndDelete(cCtx, filter).Decode(&file)
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "c.collection.FindOne",
		"exec.time": time.Since(start).String(),
		"args":      fmt.Sprintf("%+#v", filter),
	})).Debug("sql execute")
	stats.QueryRecord(stats.SELECT, "file", "FindOneByFilter", start)

	if err != nil {
		return model.File{}, err
	}

	return file, nil
}

func (c file) FindOneByFilter(ctx context.Context, filter bson.D) (model.File, error) {
	start := time.Now()
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	var file model.File
	err := c.collection.FindOne(cCtx, filter).Decode(&file)
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "c.collection.FindOne",
		"exec.time": time.Since(start).String(),
		"args":      fmt.Sprintf("%+#v", filter),
	})).Debug("sql execute")
	stats.QueryRecord(stats.SELECT, "file", "FindOneByFilter", start)
	return file, err
}

func (c file) InsertOne(ctx context.Context, file model.File) error {
	start := time.Now()
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	now := time.Now()
	_, _ = c.collection.UpdateMany(cCtx,
		bson.D{{"type", "profile"}, {"from_user_id", file.FromUserId}},
		bson.D{{"$set", bson.M{
			"deleted_at": now,
		}}})
	cancel()

	cCtx, cancel = context.WithTimeout(ctx, c.queryTimeout)
	result, err := c.collection.InsertOne(cCtx, file)
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "c.collection.FindOne",
		"exec.time": time.Since(start).String(),
		"args":      fmt.Sprintf("%+#v", file),
	})).Debug("sql execute")
	stats.QueryRecord(stats.INSERT, "file", "InsertOne", start)

	if err != nil {
		return err
	}

	if result.InsertedID == nil {
		return errors.New("insert fail")
	}

	return nil
}
