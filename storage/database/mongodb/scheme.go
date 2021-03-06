package mongodb

import (
	"context"
	"github.com/ppzxc/chattools/storage/database"
	model2 "github.com/ppzxc/chattools/storage/database/model"
	"github.com/ppzxc/chattools/types"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

func (m *mongodb) dropTable(ctx context.Context, collectionName string) error {
	dropCtx, cancel := context.WithTimeout(ctx, m.queryTimeout)
	err := m.mdb.Collection(collectionName).Drop(dropCtx)
	cancel()
	if err != nil {
		return err
	}
	return nil
}

func (m *mongodb) InitializeTable(ctx context.Context, dropTableOnStart bool, createTableOnStart bool, testUserInsertOnStart bool) error {
	if dropTableOnStart {
		logrus.Debug("drop collections")
		_ = m.dropTable(ctx, database.MongoCollectionUser)
		_ = m.dropTable(ctx, database.MongoCollectionFile)
		_ = m.dropTable(ctx, database.MongoCollectionTopic)
		_ = m.dropTable(ctx, database.MongoCollectionNotify)
		_ = m.dropTable(ctx, database.MongoCollectionSubscriptions)
		_ = m.dropTable(ctx, database.MongoCollectionMessage)
		_ = m.dropTable(ctx, database.MongoCollectionCounters)
	}

	if createTableOnStart {
		logrus.Debug("create collections")
		_ = m.mdb.CreateCollection(ctx, database.MongoCollectionUser)
		_ = m.mdb.CreateCollection(ctx, database.MongoCollectionFile)
		_ = m.mdb.CreateCollection(ctx, database.MongoCollectionTopic)
		_ = m.mdb.CreateCollection(ctx, database.MongoCollectionNotify)
		_ = m.mdb.CreateCollection(ctx, database.MongoCollectionSubscriptions)
		_ = m.mdb.CreateCollection(ctx, database.MongoCollectionMessage)
		_ = m.mdb.CreateCollection(ctx, database.MongoCollectionCounters)

		//collTopic := m.mdb.Collection(mgdb.MONGO_COLLECTION_TOPIC)
		//_, _ = coll.Indexes().CreateOne(ctx, mongo.IndexModel{
		//	Keys: bson.D{{"user_id", 1}},
		//	Options: &options.IndexOptions{},
		//})

		fileCollection := m.mdb.Collection(database.MongoCollectionFile)
		_, _ = fileCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys:    bson.D{{"_id", 1}, {"deleted_at", 1}},
			Options: &options.IndexOptions{},
		})

		_, _ = fileCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys:    bson.D{{"type", 1}, {"from_user_id", 1}},
			Options: &options.IndexOptions{},
		})

		topicCollection := m.mdb.Collection(database.MongoCollectionTopic)
		_, _ = topicCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys:    bson.D{{"updated_at", 1}},
			Options: &options.IndexOptions{},
		})

		_, _ = topicCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys:    bson.D{{"created_at", 1}},
			Options: &options.IndexOptions{},
		})

		_, _ = topicCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys:    bson.D{{"_id", 1}, {"updated_at", 1}},
			Options: &options.IndexOptions{},
		})

		_, _ = topicCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys:    bson.D{{"_id", 1}, {"created_at", 1}},
			Options: &options.IndexOptions{},
		})

		collMessage := m.mdb.Collection(database.MongoCollectionMessage)
		_, _ = collMessage.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys:    bson.D{{"topic_id", 1}},
			Options: &options.IndexOptions{},
		})

		_, _ = collMessage.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys:    bson.D{{"topic_id", 1}, {"sequence_id", 1}},
			Options: &options.IndexOptions{},
		})

		coll := m.mdb.Collection(database.MongoCollectionSubscriptions)

		_, _ = coll.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys:    bson.D{{"user_id", 1}},
			Options: &options.IndexOptions{},
		})

		_, _ = coll.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys:    bson.D{{"topic_id", 1}},
			Options: &options.IndexOptions{},
		})

		_, _ = coll.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys:    bson.D{{"user_id", 1}, {"topic_id", 1}},
			Options: &options.IndexOptions{},
		})

		_, _ = coll.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys:    bson.D{{"user_id", 1}, {"topic_id", 1}, {"receive_sequence_id", 1}},
			Options: &options.IndexOptions{},
		})

		notifyCollection := m.mdb.Collection(database.MongoCollectionNotify)
		_, _ = notifyCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys:    bson.D{{"receive_user_id", 1}, {"created_at", 1}, {"deleted_at", 1}},
			Options: &options.IndexOptions{},
		})

		_, _ = notifyCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys:    bson.D{{"receive_user_id", 1}, {"updated_at", 1}, {"deleted_at", 1}},
			Options: &options.IndexOptions{},
		})

		_, _ = notifyCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys:    bson.D{{"receive_user_id", 1}, {"_id", 1}, {"deleted_at", 1}},
			Options: &options.IndexOptions{},
		})
	}

	if testUserInsertOnStart {
		letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

		logrus.Debug("create users")
		names := []string{
			"?????????A", "?????????B", "?????????C", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????", "?????????",
		}

		var insertUsers []*model2.User

		now := time.Now()
		hash, _ := bcrypt.GenerateFromPassword([]byte("asdfasdfasdf"), bcrypt.DefaultCost)
		expires := time.Now().Add(365 * 24 * time.Hour)

		for _, v := range names {
			b := make([]byte, 6)
			for i := range b {
				b[i] = letterBytes[rand.Intn(len(letterBytes))]
			}

			u := model2.User{
				State:     types.StateUserCreated,
				StatedAt:  &now,
				CreatedAt: &now,
				UpdatedAt: &now,
				DeletedAt: nil,
				Device: []*model2.Device{
					{
						DeviceId:        "-1",
						BrowserId:       "test",
						UserAgent:       "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
						OperationSystem: types.OsWindows,
						Platform:        types.PlatformDesktop,
						IsLogin:         false,
						CreatedAt:       &now,
						UpdatedAt:       &now,
					},
				},
				Authentication: &model2.Authentication{
					UserName:  v,
					Email:     string(b) + "@nanoit.kr",
					Password:  string(hash),
					AuthType:  types.UsingRotary,
					AuthLevel: types.StateAuthLevelUser,
					CreatedAt: &now,
					UpdatedAt: &now,
					Expires:   &expires,
				},
				Profile: &model2.Profile{
					Description: "rotary user",
					CreatedAt:   &now,
					UpdatedAt:   &now,
				},
			}

			logrus.WithField("auth.username", u.Authentication.UserName).Debug("create user")
			insertUsers = append(insertUsers, &u)
		}

		testUser := &model2.User{
			State:     types.StateUserCreated,
			StatedAt:  &now,
			CreatedAt: &now,
			UpdatedAt: &now,
			Device: []*model2.Device{
				{
					DeviceId:        "-1",
					BrowserId:       "test",
					UserAgent:       "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
					OperationSystem: types.OsWindows,
					Platform:        types.PlatformDesktop,
					IsLogin:         false,
					CreatedAt:       &now,
					UpdatedAt:       &now,
				},
			},
			Authentication: &model2.Authentication{
				UserName:  "asdf",
				Email:     "asdf@asdf.com",
				Password:  string(hash),
				AuthType:  types.StateAuthTypeId,
				AuthLevel: types.StateAuthLevelUser,
				Secret:    "testsecret",
				Expires:   &expires,
				CreatedAt: &now,
				UpdatedAt: &now,
				DeletedAt: &now,
			},
			Profile: &model2.Profile{
				Description: "test user",
				CreatedAt:   &now,
				UpdatedAt:   &now,
			},
		}

		insertUsers = append(insertUsers, testUser)

		return m.registerAll(ctx, insertUsers)
		//for _, u := range insertUsers {
		//	_ = m.UserInsert(u)
		//}
	}

	return nil
}
