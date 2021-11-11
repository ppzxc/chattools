package impl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/ppzxc/chattools/storage/database"
	"github.com/ppzxc/chattools/storage/database/model"
	"github.com/ppzxc/chattools/storage/database/mongodb/repository"
	"github.com/ppzxc/chattools/utils"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"time"
)

type user struct {
	database     *mongo.Database
	collection   *mongo.Collection
	logger       *zap.Logger
	queryTimeout time.Duration
}

func NewUserRepository(db *mongo.Database, queryTimeout time.Duration) repository.User {
	return &user{
		database:     db,
		collection:   db.Collection(database.MongoCollectionUser),
		queryTimeout: queryTimeout,
	}
}

func (c user) FindOneByFilter(ctx context.Context, filter bson.D) (model.User, error) {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	var findUser model.User
	start := time.Now()
	err := c.collection.FindOne(cCtx, filter).Decode(&findUser)
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "c.collection.FindOne",
		"exec.time": time.Since(start).String(),
		"args":      fmt.Sprintf("%+#v", filter),
	})).Debug("sql execute")
	return findUser, err
}

func (c user) FindOneAndDelete(ctx context.Context, userId int64) (*model.User, error) {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	var findUser model.User
	start := time.Now()
	err := c.collection.FindOneAndDelete(cCtx, bson.M{"_id": userId}).Decode(&findUser)
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "c.collection.FindOneAndDelete",
		"exec.time": time.Since(start).String(),
		"args":      fmt.Sprintf("%+#v", userId),
	})).Debug("sql execute")
	return &findUser, err
}

func (c user) FindManyByFilter(ctx context.Context, filter bson.D) ([]model.User, error) {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	cursor, err := c.collection.Find(cCtx, filter)
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "c.collection.Find",
		"exec.time": time.Since(start).String(),
		"args":      fmt.Sprintf("%+#v", filter),
	})).Debug("sql execute")
	if err != nil {
		return nil, err
	}

	var users []model.User
	cCtx, cancel = context.WithTimeout(ctx, c.queryTimeout)
	if err := cursor.All(cCtx, &users); err != nil {
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

	return users, nil
}

func (c user) InsertMany(ctx context.Context, many []interface{}) error {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	result, err := c.collection.InsertMany(cCtx, many)
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "c.collection.InsertMany",
		"exec.time": time.Since(start).String(),
		"args":      fmt.Sprintf("%+#v", many),
	})).Debug("sql execute")
	if err != nil {
		return err
	}

	if result.InsertedIDs == nil || len(result.InsertedIDs) != len(many) {
		return errors.New("insert fail")
	}

	return nil
}

func (c user) InsertOne(ctx context.Context, user model.User) error {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	result, err := c.collection.InsertOne(cCtx, user)
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "c.collection.InsertOne",
		"exec.time": time.Since(start).String(),
		"args":      fmt.Sprintf("%+#v", user),
	})).Debug("sql execute")
	if err != nil {
		return err
	}

	if result.InsertedID == nil {
		return errors.New("insert fail")
	}

	return nil
}

func (c user) FindOneAndUpdateByFilter(ctx context.Context, filter bson.D, update bson.D) error {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	result := c.collection.FindOneAndUpdate(cCtx, filter, update)
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "c.collection.FindOneAndUpdate",
		"exec.time": time.Since(start).String(),
		"args":      fmt.Sprintf("%+#v", filter),
		"args2":     fmt.Sprintf("%+#v", update),
	})).Debug("sql execute")
	return result.Err()
}

func (c user) Delete(ctx context.Context, userId int64) error {
	cCtx, cancel := context.WithTimeout(ctx, c.queryTimeout)
	start := time.Now()
	result, err := c.collection.DeleteOne(cCtx, bson.D{{Key: "_id", Value: userId}})
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "c.collection.DeleteOne",
		"exec.time": time.Since(start).String(),
		"args":      fmt.Sprintf("%+#v", userId),
	})).Debug("sql execute")
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return sql.ErrNoRows
	}
	return nil
}
