package impl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/ppzxc/chattools/database"
	"github.com/ppzxc/chattools/database/model"
	"github.com/ppzxc/chattools/database/mongodb/repository"
	"github.com/ppzxc/chattools/types"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type subscription struct {
	database     *mongo.Database
	collection   *mongo.Collection
	queryTimeout time.Duration
}

func NewSubscriptionRepository(db *mongo.Database, queryTimeout time.Duration) repository.Subscription {
	return &subscription{
		database:     db,
		collection:   db.Collection(database.MongoCollectionSubscriptions),
		queryTimeout: queryTimeout,
	}
}

func (c subscription) FindOneByFilter(ctx context.Context, filter interface{}) (model.Subscription, error) {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	var subs model.Subscription
	start := time.Now()
	err := c.collection.FindOne(cCtx, filter).Decode(&subs)
	cancel()
	logrus.WithFields(logrus.Fields{
		"query":     "c.collection.FindOne",
		"exec.time": time.Since(start),
		"args":      fmt.Sprintf("%+#v", filter),
	})
	return subs, err
}

func (c subscription) FindManyByFilter(ctx context.Context, filter interface{}) ([]model.Subscription, error) {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	cursor, err := c.collection.Find(cCtx, filter)
	cancel()
	logrus.WithFields(logrus.Fields{
		"query":     "c.collection.Find",
		"exec.time": time.Since(start),
		"args":      fmt.Sprintf("%+#v", filter),
	})
	if err != nil {
		return nil, err
	}

	var subs []model.Subscription
	cCtx, cancel = context.WithTimeout(ctx, c.queryTimeout)
	if err := cursor.All(cCtx, &subs); err != nil {
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

	return subs, nil
}

func (c subscription) InsertOne(ctx context.Context, subs model.Subscription) error {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	result, err := c.collection.InsertOne(cCtx, subs)
	cancel()
	logrus.WithFields(logrus.Fields{
		"query":     "c.collection.InsertOne",
		"exec.time": time.Since(start),
		"args":      fmt.Sprintf("%+#v", subs),
	})
	if err != nil {
		return err
	}

	if result.InsertedID == nil {
		return errors.New("insert fail")
	}

	return nil
}

func (c subscription) DeleteAllByFilter(ctx context.Context, filter interface{}) error {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	result, err := c.collection.DeleteMany(cCtx, filter)
	cancel()
	logrus.WithFields(logrus.Fields{
		"query":     "c.collection.DeleteMany",
		"exec.time": time.Since(start),
		"args":      fmt.Sprintf("%+#v", filter),
	})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (c subscription) DeleteOneByFilter(ctx context.Context, filter interface{}) error {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	result, err := c.collection.DeleteOne(cCtx, filter)
	cancel()
	logrus.WithFields(logrus.Fields{
		"query":     "c.collection.DeleteOne",
		"exec.time": time.Since(start),
		"args":      fmt.Sprintf("%+#v", filter),
	})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (c subscription) UpdateOneByFilter(ctx context.Context, filter interface{}, update interface{}) error {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	t := true
	result, err := c.collection.UpdateOne(cCtx, filter, update, &options.UpdateOptions{Upsert: &t})
	cancel()
	logrus.WithFields(logrus.Fields{
		"query":     "c.collection.UpdateOne",
		"exec.time": time.Since(start),
		"args":      fmt.Sprintf("%+#v", filter),
	})

	if err != nil {
		return err
	}

	if result.MatchedCount != 1 {
		return mongo.ErrNoDocuments
	}

	if result.ModifiedCount <= 0 {
		return types.ErrDataBaseUpdateRowAffectIsZero
	}

	return nil
}
