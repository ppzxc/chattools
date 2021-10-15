package impl

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ppzxc/chattools/database"
	"github.com/ppzxc/chattools/database/model"
	"github.com/ppzxc/chattools/database/mongodb/repository"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type message struct {
	database     *mongo.Database
	collection   *mongo.Collection
	queryTimeout time.Duration
}

func NewMessageRepository(db *mongo.Database, queryTimeout time.Duration) repository.Message {
	return &message{
		database:     db,
		collection:   db.Collection(database.MongoCollectionMessage),
		queryTimeout: queryTimeout,
	}
}

func (c message) InsertOne(ctx context.Context, message model.Message) error {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	result, err := c.collection.InsertOne(cCtx, message)
	cancel()
	logrus.WithFields(logrus.Fields{
		"query":     "c.collection.InsertOne",
		"exec.time": time.Since(start),
		"args":      message,
	})
	if err != nil {
		return err
	}

	if result.InsertedID == nil {
		return errors.New("insert fail")
	}

	return nil
}

func (c message) InsertMany(ctx context.Context, messages []interface{}) error {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	result, err := c.collection.InsertMany(cCtx, messages)
	cancel()
	logrus.WithFields(logrus.Fields{
		"query":     "c.collection.InsertMany",
		"exec.time": time.Since(start),
		"args":      messages,
	})
	if err != nil {
		return err
	}

	if result.InsertedIDs == nil || len(result.InsertedIDs) != len(messages) {
		return errors.New("insert fail")
	}

	return nil
}

func (c message) FindManyByFilter(ctx context.Context, filter bson.D) ([]model.Message, error) {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	cursor, err := c.collection.Find(cCtx, filter)
	cancel()
	logrus.WithFields(logrus.Fields{
		"query":     "c.collection.Find",
		"exec.time": time.Since(start),
		"args":      filter,
	})
	if err != nil {
		return nil, err
	}

	var messages []model.Message
	cCtx, cancel = context.WithTimeout(ctx, c.queryTimeout)
	if err := cursor.All(cCtx, &messages); err != nil {
		cancel()
		return nil, err
	}
	cancel()

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cCtx, cancel = context.WithTimeout(ctx, c.queryTimeout)
	if err := cursor.Close(ctx); err != nil {
		cancel()
		return nil, err
	}
	cancel()

	return messages, nil
}

func (c message) Delete(ctx context.Context, filter bson.D) error {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	result, err := c.collection.DeleteMany(cCtx, filter)
	cancel()

	logrus.WithFields(logrus.Fields{
		"query":     "c.collection.DeleteMany",
		"exec.time": time.Since(start),
		"args":      filter,
	})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return sql.ErrNoRows
	}
	return nil
}
