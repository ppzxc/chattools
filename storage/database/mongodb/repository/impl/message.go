package impl

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ppzxc/chattools/common/stats"
	"github.com/ppzxc/chattools/storage/database/model"
	"github.com/ppzxc/chattools/storage/database/mongodb/repository"
	"github.com/ppzxc/chattools/utils"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type message struct {
	collection   *mongo.Collection
	queryTimeout time.Duration
}

func NewMessageRepository(collection *mongo.Collection, queryTimeout time.Duration) repository.Message {
	return &message{
		collection:   collection,
		queryTimeout: queryTimeout,
	}
}

func (c message) InsertOne(ctx context.Context, message model.Message) error {
	start := time.Now()
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	result, err := c.collection.InsertOne(cCtx, message)
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "c.collection.InsertOne",
		"exec.time": time.Since(start).String(),
		"args":      message,
	})).Debug("sql execute")
	stats.QueryRecord(stats.INSERT, "message", "InsertOne", start)

	if err != nil {
		return err
	}

	if result.InsertedID == nil {
		return errors.New("insert fail")
	}

	return nil
}

func (c message) InsertMany(ctx context.Context, messages []interface{}) error {
	start := time.Now()
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	result, err := c.collection.InsertMany(cCtx, messages)
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "c.collection.InsertMany",
		"exec.time": time.Since(start).String(),
		"args":      messages,
	})).Debug("sql execute")
	stats.QueryRecord(stats.INSERT, "message", "InsertMany", start)
	if err != nil {
		return err
	}

	if result.InsertedIDs == nil || len(result.InsertedIDs) != len(messages) {
		return errors.New("insert fail")
	}

	return nil
}

//func (c message) FindManyByFilterDescLimit(ctx context.Context, filter bson.D) ([]model.Message, error) {
//	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
//	start := time.Now()
//	cursor, err := c.collection.Find(cCtx, filter)
//	cancel()
//	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
//		"query":     "c.collection.Find",
//		"exec.time": time.Since(start).String(),
//		"args":      filter,
//	})).Debug("sql execute")
//	if err != nil {
//		return nil, err
//	}
//
//	var messages []model.Message
//	if err := cursor.All(ctx, &messages); err != nil {
//		return nil, err
//	}
//
//	if err := cursor.Err(); err != nil {
//		return nil, err
//	}
//
//	if err := cursor.Close(ctx); err != nil {
//		return nil, err
//	}
//
//	return messages, nil
//}

func (c message) FindManyByFilter(ctx context.Context, filter interface{}, options ...*options.FindOptions) ([]model.Message, error) {
	start := time.Now()
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	cursor, err := c.collection.Find(cCtx, filter, options...)
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "c.collection.Find",
		"exec.time": time.Since(start).String(),
		"args":      filter,
	})).Debug("sql execute")
	stats.QueryRecord(stats.SELECT, "message", "FindManyByFilter", start)
	if err != nil {
		return nil, err
	}

	var messages []model.Message
	if err := cursor.All(ctx, &messages); err != nil {
		return nil, err
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if err := cursor.Close(ctx); err != nil {
		return nil, err
	}

	return messages, nil
}

func (c message) Delete(ctx context.Context, filter bson.D) error {
	start := time.Now()
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	result, err := c.collection.DeleteMany(cCtx, filter)
	cancel()

	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "c.collection.DeleteMany",
		"exec.time": time.Since(start).String(),
		"args":      filter,
	})).Debug("sql execute")
	stats.QueryRecord(stats.DELETE, "message", "Delete", start)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return sql.ErrNoRows
	}
	return nil
}
