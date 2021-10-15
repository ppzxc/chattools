package impl

import (
	"context"
	"errors"
	"fmt"
	"github.com/ppzxc/chattools/database"
	"github.com/ppzxc/chattools/database/model"
	"github.com/ppzxc/chattools/database/mongodb/repository"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type notify struct {
	ctx          context.Context
	database     *mongo.Database
	collection   *mongo.Collection
	queryTimeout time.Duration
}

func NewNotifyRepository(db *mongo.Database, queryTimeout time.Duration) repository.Notify {
	return &notify{
		database:     db,
		collection:   db.Collection(database.MongoCollectionNotify),
		queryTimeout: queryTimeout,
	}
}

func (c notify) FindOneByFilter(ctx context.Context, filter bson.D) (*model.Notify, error) {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	var notify model.Notify
	start := time.Now()
	err := c.collection.FindOne(cCtx, filter).Decode(&notify)
	cancel()
	logrus.WithFields(logrus.Fields{
		"query":     "u.collection.FindOne",
		"exec.time": time.Since(start),
		"args":      fmt.Sprintf("%+#v", filter),
	}).Debug("sql execute")
	return &notify, err
}

func (c notify) FindManyFilter(ctx context.Context, filter bson.D) ([]*model.Notify, error) {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	cursor, err := c.collection.Find(cCtx, filter)
	cancel()
	logrus.WithFields(logrus.Fields{
		"query":     "u.collection.Find",
		"exec.time": time.Since(start),
		"args":      fmt.Sprintf("%+#v", filter),
	}).Debug("sql execute")
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
	ctx, cancel := context.WithTimeout(c.ctx, c.queryTimeout)
	start := time.Now()
	result, err := c.collection.InsertMany(ctx, many)
	cancel()
	logrus.WithFields(logrus.Fields{
		"query":     "u.collection.InsertMany",
		"exec.time": time.Since(start),
		"args":      fmt.Sprintf("%+#v", many),
	}).Debug("sql execute")
	if err != nil {
		return err
	}

	if result.InsertedIDs == nil || len(result.InsertedIDs) != len(many) {
		return errors.New("insert fail")
	}

	return nil
}

func (c notify) UpdateOne(ctx context.Context, notify *model.Notify) error {
	ctx, cancel := context.WithTimeout(c.ctx, c.queryTimeout)
	start := time.Now()
	result, err := c.collection.UpdateOne(ctx,
		bson.M{"_id": notify.Id}, bson.D{{"$set", notify}},
	)
	cancel()
	logrus.WithFields(logrus.Fields{
		"query":     "u.collection.UpdateOne",
		"exec.time": time.Since(start),
		"args":      fmt.Sprintf("%+#v", notify),
	}).Debug("sql execute")

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
