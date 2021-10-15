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
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type topic struct {
	database     *mongo.Database
	collection   *mongo.Collection
	queryTimeout time.Duration
}

func NewTopicRepository(db *mongo.Database, queryTimeout time.Duration) repository.Topic {
	return &topic{
		database:     db,
		collection:   db.Collection(database.MongoCollectionTopic),
		queryTimeout: queryTimeout,
	}
}

func (c topic) FindOneAndUpdateByFilter(ctx context.Context, filter bson.D, update bson.D) error {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	result := c.collection.FindOneAndUpdate(cCtx, filter, update, options.FindOneAndUpdate())
	cancel()
	logrus.WithFields(logrus.Fields{
		"query": "c.collection.FindOneAndUpdate",
		"exec.time": time.Since(start),
		"args1": filter,
		"args2": update,
	}).Debug("sql execute")
	return result.Err()
}

func (c topic) FindManyFilter(ctx context.Context, filter bson.D) ([]model.Topic, error) {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	cursor, err := c.collection.Find(cCtx, filter)
	cancel()
	logrus.WithFields(logrus.Fields{
		"query": "c.collection.Find",
		"exec.time": time.Since(start),
		"args": filter,
	}).Debug("sql execute")
	if err != nil {
		return nil, err
	}

	var topics []model.Topic
	if err := cursor.All(ctx, &topics); err != nil {
		return nil, err
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if err := cursor.Close(ctx); err != nil {
		return nil, err
	}

	return topics, nil
}

func (c topic) FindOneByFilter(ctx context.Context, filter bson.D) (model.Topic, error) {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	var topic model.Topic
	start := time.Now()
	err := c.collection.FindOne(cCtx, filter).Decode(&topic)
	cancel()
	logrus.WithFields(logrus.Fields{
		"query": "c.collection.FindOne",
		"exec.time": time.Since(start),
		"args": filter,
	}).Debug("sql execute")
	return topic, err
}

func (c topic) InsertOne(ctx context.Context, topic model.Topic) error {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	result, err := c.collection.InsertOne(cCtx, topic)
	cancel()
	logrus.WithFields(logrus.Fields{
		"query": "c.collection.InsertOne",
		"exec.time": time.Since(start),
		"args": topic,
	}).Debug("sql execute")
	if err != nil {
		return err
	}

	if result.InsertedID == nil {
		return errors.New("insert fail")
	}

	return nil
}

func (c topic) UpdateFilter(ctx context.Context, filter bson.D, update bson.D) error {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	result := c.collection.FindOneAndUpdate(cCtx, filter, update, options.FindOneAndUpdate())
	cancel()
	logrus.WithFields(logrus.Fields{
		"query": "c.collection.FindOneAndUpdate",
		"exec.time": time.Since(start),
		"args1": filter,
		"args2": update,
	}).Debug("sql execute")
	return result.Err()
}

func (c topic) Update(ctx context.Context, topic *model.Topic) error {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	result := c.collection.FindOneAndUpdate(cCtx,
		bson.D{{"_id", topic.Id}},
		bson.D{{"$set", topic}},
		options.FindOneAndUpdate())
	cancel()
	logrus.WithFields(logrus.Fields{
		"query": "c.collection.FindOneAndUpdate",
		"exec.time": time.Since(start),
		"args": topic,
	}).Debug("sql execute")
	return result.Err()
}

func (c topic) Delete(ctx context.Context, topicId int64) error {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	result, err := c.collection.DeleteOne(cCtx, bson.D{{Key: "_id", Value: topicId}})
	cancel()
	logrus.WithFields(logrus.Fields{
		"query": "c.collection.DeleteOne",
		"exec.time": time.Since(start),
		"args": topicId,
	}).Debug("sql execute")
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return sql.ErrNoRows
	}
	return nil
}
