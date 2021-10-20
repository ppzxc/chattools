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
	dropCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	err := m.mongoDataBase.Collection(collectionName).Drop(dropCtx)
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
		_ = m.mongoDataBase.CreateCollection(ctx, database.MongoCollectionUser)
		_ = m.mongoDataBase.CreateCollection(ctx, database.MongoCollectionFile)
		_ = m.mongoDataBase.CreateCollection(ctx, database.MongoCollectionTopic)
		_ = m.mongoDataBase.CreateCollection(ctx, database.MongoCollectionNotify)
		_ = m.mongoDataBase.CreateCollection(ctx, database.MongoCollectionSubscriptions)
		_ = m.mongoDataBase.CreateCollection(ctx, database.MongoCollectionMessage)
		_ = m.mongoDataBase.CreateCollection(ctx, database.MongoCollectionCounters)

		//collTopic := m.mongoDataBase.Collection(mgdb.MONGO_COLLECTION_TOPIC)
		//_, _ = coll.Indexes().CreateOne(ctx, mongo.IndexModel{
		//	Keys: bson.D{{"user_id", 1}},
		//	Options: &options.IndexOptions{},
		//})

		collMessage := m.mongoDataBase.Collection(database.MongoCollectionMessage)
		_, _ = collMessage.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys:    bson.D{{"topic_id", 1}},
			Options: &options.IndexOptions{},
		})

		_, _ = collMessage.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys:    bson.D{{"topic_id", 1}, {"sequence_id", 1}},
			Options: &options.IndexOptions{},
		})

		coll := m.mongoDataBase.Collection(database.MongoCollectionSubscriptions)

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
	}

	if testUserInsertOnStart {
		letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

		logrus.Debug("create users")
		names := []string{
			"관리자A", "관리자B", "관리자C", "이영식", "조길조", "이제호", "김성태", "이장수", "강호권", "최호성", "이동철", "김태성", "천세근", "윤영길", "김백기", "김종철", "이상돈", "도재윤", "김종만", "김재진", "정두관", "정영석", "정종섭", "이대호", "이종석", "윤광호", "정삼국", "이승목", "김도협", "김태영", "김대연", "정용하", "박종철", "이상철", "최용철", "박경호", "최홍주", "임일중", "김진석", "이상문", "최정원", "송명곤", "최우섭", "안창영", "김영철", "김원주", "박원덕", "지방원", "한대훈", "최원석", "정영석", "김선열", "손병철", "최규식", "이언규", "박성준", "권동찬", "이상직", "조필석", "배진석", "손석만", "정희택", "김인하", "손정훈", "최정환", "김형기", "한종건", "윤명대", "권진영", "기노환", "이동우", "김보현", "신동일", "박기영", "정병태", "지인용", "안영진", "황병용", "박원규", "김주원", "오영만", "이희영", "이영철", "김동환", "최희정", "김성민",
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
