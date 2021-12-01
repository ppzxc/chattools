package impl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/ppzxc/chattools/common/stats"
	"github.com/ppzxc/chattools/storage/database/model"
	"github.com/ppzxc/chattools/storage/database/mongodb/repository"
	"github.com/ppzxc/chattools/types"
	"github.com/ppzxc/chattools/utils"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type subscription struct {
	//database     *mongo.Database
	collection   *mongo.Collection
	queryTimeout time.Duration
}

func NewSubscriptionRepository(collection *mongo.Collection, queryTimeout time.Duration) repository.Subscription {
	return &subscription{
		//database:     db,
		collection:   collection,
		queryTimeout: queryTimeout,
	}
}

func (c subscription) CountDocuments(ctx context.Context, filter interface{}) (count int64, err error) {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	count, err = c.collection.CountDocuments(cCtx, filter)
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "c.collection.CountDocuments",
		"exec.time": time.Since(start).String(),
		"args":      fmt.Sprintf("%+#v", filter),
	})).Debug("sql execute")
	stats.QueryRecord(stats.SELECT, "subscriptions", "CountDocuments", start)
	return
}

func (c subscription) FindOneByFilter(ctx context.Context, filter interface{}, options ...*options.FindOneOptions) (model.Subscription, error) {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	var subs model.Subscription
	start := time.Now()
	err := c.collection.FindOne(cCtx, filter, options...).Decode(&subs)
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "c.collection.FindOne",
		"exec.time": time.Since(start).String(),
		"args":      fmt.Sprintf("%+#v", filter),
	})).Debug("sql execute")
	stats.QueryRecord(stats.SELECT, "subscriptions", "FindOneByFilter", start)
	return subs, err
}

func (c subscription) FindManyByFilter(ctx context.Context, filter interface{}, options ...*options.FindOptions) ([]model.Subscription, error) {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	cursor, err := c.collection.Find(cCtx, filter, options...)
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "c.collection.Find",
		"exec.time": time.Since(start).String(),
		"args":      fmt.Sprintf("%+#v", filter),
	})).Debug("sql execute")
	stats.QueryRecord(stats.SELECT, "subscriptions", "FindManyByFilter", start)
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
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "c.collection.InsertOne",
		"exec.time": time.Since(start).String(),
		"args":      fmt.Sprintf("%+#v", subs),
	})).Debug("sql execute")
	stats.QueryRecord(stats.INSERT, "subscriptions", "InsertOne", start)
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
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "c.collection.DeleteMany",
		"exec.time": time.Since(start).String(),
		"args":      fmt.Sprintf("%+#v", filter),
	})).Debug("sql execute")
	stats.QueryRecord(stats.DELETE, "subscriptions", "DeleteAllByFilter", start)
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
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "c.collection.DeleteOne",
		"exec.time": time.Since(start).String(),
		"args":      fmt.Sprintf("%+#v", filter),
	})).Debug("DeleteOneByFilter")
	stats.QueryRecord(stats.DELETE, "subscriptions", "DeleteOneByFilter", start)
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
	result, err := c.collection.UpdateOne(cCtx, filter, update, &options.UpdateOptions{})
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "c.collection.UpdateOne",
		"exec.time": time.Since(start).String(),
		"args":      fmt.Sprintf("%+#v", filter),
	})).Debug("UpdateOneByFilter")
	stats.QueryRecord(stats.UPDATE, "subscriptions", "UpdateOneByFilter", start)

	if err != nil {
		return err
	}

	if result.MatchedCount <= 0 {
		return mongo.ErrNoDocuments
	}

	if result.ModifiedCount <= 0 {
		return types.ErrDataBaseUpdateRowAffectIsZero
	}

	return nil
}
