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

type notify struct {
	//database     *mongo.Database
	collection   *mongo.Collection
	queryTimeout time.Duration
}

func NewNotifyRepository(collection *mongo.Collection, queryTimeout time.Duration) repository.Notify {
	return &notify{
		//database:     db,
		collection:   collection,
		queryTimeout: queryTimeout,
	}
}

func (c notify) FindOneByFilter(ctx context.Context, filter bson.D) (*model.Notify, error) {
	start := time.Now()
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	var notify model.Notify
	err := c.collection.FindOne(cCtx, filter).Decode(&notify)
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "u.collection.FindOne",
		"exec.time": time.Since(start).String(),
		"args":      fmt.Sprintf("%+#v", filter),
	})).Debug("sql execute")
	stats.QueryRecord(stats.SELECT, "notify", "FindOneByFilter", start)
	return &notify, err
}

func (c notify) FindManyFilter(ctx context.Context, filter bson.D) ([]*model.Notify, error) {
	start := time.Now()
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	cursor, err := c.collection.Find(cCtx, filter)
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "u.collection.Find",
		"exec.time": time.Since(start).String(),
		"args":      fmt.Sprintf("%+#v", filter),
	})).Debug("sql execute")
	stats.QueryRecord(stats.SELECT, "notify", "FindManyFilter", start)
	if err != nil {
		return nil, err
	}

	var notify []*model.Notify
	cCtx, cancel = context.WithTimeout(ctx, c.queryTimeout)
	if err := cursor.All(cCtx, &notify); err != nil {
		cancel()
		return nil, err
	}
	cancel()

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cCtx, cancel = context.WithTimeout(ctx, c.queryTimeout)
	if err := cursor.Close(cCtx); err != nil {
		cancel()
		return nil, err
	}
	cancel()

	return notify, nil
}

func (c notify) InsertMany(ctx context.Context, many []interface{}) error {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	result, err := c.collection.InsertMany(cCtx, many)
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "u.collection.InsertMany",
		"exec.time": time.Since(start).String(),
		"args":      fmt.Sprintf("%+#v", many),
	})).Debug("sql execute")
	stats.QueryRecord(stats.INSERT, "notify", "InsertMany", start)
	if err != nil {
		return err
	}

	if result.InsertedIDs == nil || len(result.InsertedIDs) != len(many) {
		return errors.New("insert fail")
	}

	return nil
}

func (c notify) InsertOne(ctx context.Context, one interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	result, err := c.collection.InsertOne(ctx, one)
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "u.collection.InsertOne",
		"exec.time": time.Since(start).String(),
		"args":      fmt.Sprintf("%+#v", one),
	})).Debug("sql execute")
	stats.QueryRecord(stats.INSERT, "notify", "InsertOne", start)
	if err != nil {
		return 0, err
	}

	return result.InsertedID.(int64), nil
}

func (c notify) UpdateOne(ctx context.Context, notify *model.Notify) error {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	result, err := c.collection.UpdateOne(cCtx,
		bson.M{"_id": notify.Id}, bson.D{{"$set", notify}},
	)
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "u.collection.UpdateOne",
		"exec.time": time.Since(start).String(),
		"args":      fmt.Sprintf("%+#v", notify),
	})).Debug("sql execute")
	stats.QueryRecord(stats.UPDATE, "notify", "UpdateOne", start)

	if err != nil {
		return err
	}

	if result.MatchedCount != 1 {
		return mongo.ErrNoDocuments
	}

	if result.ModifiedCount <= 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
