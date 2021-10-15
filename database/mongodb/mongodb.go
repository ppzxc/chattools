package mongodb

import (
	"context"
	"fmt"
	"github.com/ppzxc/chattools/database"
	"github.com/ppzxc/chattools/database/mongodb/repository"
	"github.com/ppzxc/chattools/database/mongodb/repository/impl"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type mongodb struct {
	connectTimeout time.Duration
	queryTimeout   time.Duration
	mongoClient    *mongo.Client
	mongoDataBase  *mongo.Database

	crudUser     repository.User
	crudTopic    repository.Topic
	crudSequence repository.Sequence
	crudNotify   repository.Notify
	crudMessage  repository.Message
	crudFile     repository.File
	crudSubs     repository.Subscription
}

func NewMongoDbInstance(closable context.Context, authMechanism string, authSource string, host string, port int, username string, database string, password string, connectTimeout time.Duration, queryTimeout time.Duration) (database.Service, error) {
	// mongodb context
	mongoClosable, cancel := context.WithCancel(closable)

	// mongodb struct
	m := &mongodb{
		connectTimeout: connectTimeout,
		queryTimeout:   queryTimeout,
	}

	var passwordSet bool
	if password != "" {
		passwordSet = true
	}

	// mongodb conn
	connCtx, cancel := context.WithTimeout(mongoClosable, connectTimeout)
	opt := options.Client().ApplyURI(fmt.Sprintf("mongodb://%v:%v", host, port))
	//opt := options.Client()
	//opt.SetHosts([]string{conf.Config.Db.Host})
	opt.SetAuth(options.Credential{
		AuthMechanism: authMechanism,
		AuthSource:    authSource,
		Username:      username,
		Password:      password,
		PasswordSet:   passwordSet,
	})
	client, err := mongo.Connect(connCtx, opt)
	cancel()
	if err != nil {
		logrus.WithError(err).Error("mongodb connect fail")
		return nil, err
	}

	m.mongoClient = client

	// mongodb conn ping
	pingCtx, cancel := context.WithTimeout(mongoClosable, connectTimeout)
	err = client.Ping(pingCtx, readpref.Primary())
	cancel()
	if err != nil {
		logrus.WithError(err).Error("mongodb ping fail")
		return nil, err
	}

	m.mongoDataBase = m.mongoClient.Database(database)

	m.crudSequence = impl.NewSequenceRepository(m.mongoDataBase, m.queryTimeout)
	m.crudUser = impl.NewUserRepository(m.mongoDataBase, m.queryTimeout)
	m.crudFile = impl.NewFileRepository(m.mongoDataBase, m.queryTimeout)
	m.crudTopic = impl.NewTopicRepository(m.mongoDataBase, m.queryTimeout)
	m.crudNotify = impl.NewNotifyRepository(m.mongoDataBase, m.queryTimeout)
	m.crudSubs = impl.NewSubscriptionRepository(m.mongoDataBase, m.queryTimeout)
	m.crudMessage = impl.NewMessageRepository(m.mongoDataBase, m.queryTimeout)

	return m, nil
}

func (m mongodb) Close(ctx context.Context) error {
	closeCtx, cancel := context.WithTimeout(ctx, m.connectTimeout)
	err := m.mongoClient.Disconnect(closeCtx)
	cancel()
	if err != nil {
		logrus.WithError(err).Error("disconnect error")
		return err
	}
	return nil
}
