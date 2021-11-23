package mongodb

import (
	"context"
	"fmt"
	"github.com/ppzxc/chattools/storage/database"
	"github.com/ppzxc/chattools/storage/database/mongodb/repository"
	"github.com/ppzxc/chattools/storage/database/mongodb/repository/impl"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonoptions"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"reflect"
	"time"
)

type mongodb struct {
	connectTimeout time.Duration
	queryTimeout   time.Duration
	cli            *mongo.Client
	mdb            *mongo.Database
	crudUser       repository.User
	crudTopic      repository.Topic
	crudSeq        repository.Sequence
	crudNotify     repository.Notify
	crudMsg        repository.Message
	crudFile       repository.File
	crudSubs       repository.Subscription
}

func NewMongoDbInstance(closable context.Context, authMechanism string, authSource string, host string, port int, username string, dbName string, password string, connectTimeout time.Duration, queryTimeout time.Duration) (database.Service, error) {
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

	m.cli = client

	// mongodb conn ping
	pingCtx, cancel := context.WithTimeout(mongoClosable, connectTimeout)
	err = client.Ping(pingCtx, readpref.Primary())
	cancel()
	if err != nil {
		logrus.WithError(err).Error("mongodb ping fail")
		return nil, err
	}

	//defaultRegistry := bson.NewRegistryBuilder().Build()
	//defaultRegistry.LookupDecoder()
	var primitiveCodecs bson.PrimitiveCodecs
	rb := bsoncodec.NewRegistryBuilder()
	bsoncodec.DefaultValueEncoders{}.RegisterDefaultEncoders(rb)
	bsoncodec.DefaultValueDecoders{}.RegisterDefaultDecoders(rb)
	t := true
	rb.RegisterTypeDecoder(reflect.TypeOf(time.Time{}), bsoncodec.NewTimeCodec(&bsonoptions.TimeCodecOptions{UseLocalTimeZone: &t}))
	primitiveCodecs.RegisterPrimitiveCodecs(rb)

	registry := rb.Build()

	m.mdb = m.cli.Database(dbName, &options.DatabaseOptions{
		Registry: registry,
	})

	m.crudSeq = impl.NewSequenceRepository(m.mdb.Collection(database.MongoCollectionCounters), m.queryTimeout)
	m.crudUser = impl.NewUserRepository(m.mdb.Collection(database.MongoCollectionUser), m.queryTimeout)
	m.crudFile = impl.NewFileRepository(m.mdb.Collection(database.MongoCollectionFile), m.queryTimeout)
	m.crudTopic = impl.NewTopicRepository(m.mdb.Collection(database.MongoCollectionTopic), m.queryTimeout)
	m.crudNotify = impl.NewNotifyRepository(m.mdb.Collection(database.MongoCollectionNotify), m.queryTimeout)
	m.crudSubs = impl.NewSubscriptionRepository(m.mdb.Collection(database.MongoCollectionSubscriptions), m.queryTimeout)
	m.crudMsg = impl.NewMessageRepository(m.mdb.Collection(database.MongoCollectionMessage), m.queryTimeout)

	return m, nil
}

func (m mongodb) Close(ctx context.Context) error {
	closeCtx, cancel := context.WithTimeout(ctx, m.connectTimeout)
	err := m.cli.Disconnect(closeCtx)
	cancel()
	if err != nil {
		logrus.WithError(err).Error("disconnect error")
		return err
	}
	return nil
}
