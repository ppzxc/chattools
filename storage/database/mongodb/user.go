package mongodb

import (
	"context"
	"database/sql"
	"github.com/ppzxc/chattools/storage/database"
	"github.com/ppzxc/chattools/storage/database/model"
	"github.com/ppzxc/chattools/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func (m mongodb) registerAll(ctx context.Context, users []*model.User) error {
	var many []interface{}
	for i := 0; i < len(users); i++ {
		id, err := m.crudSeq.Next(ctx, database.MongoCollectionUser)
		if err != nil {
			return err
		}
		users[i].Id = id
		many = append(many, users[i])
	}
	return m.crudUser.InsertMany(ctx, many)
}

func (m mongodb) UserMaxId(ctx context.Context) (int64, error) {
	option := options.FindOptions{}
	option.SetSort(bson.D{{"_id", -1}})
	option.SetLimit(1)

	users, err := m.crudUser.FindManyByFilter(ctx, bson.D{}, &option)
	if err != nil {
		return 0, err
	}
	if len(users) <= 0 {
		return 0, mongo.ErrNoDocuments
	}

	return users[0].Id, nil
}

func (m mongodb) UserCountDocuments(ctx context.Context) (int64, error) {
	return m.crudUser.CountDocuments(ctx, bson.D{})
}

func (m mongodb) UserDeleteByUserId(ctx context.Context, userId int64) error {
	return m.crudUser.Delete(ctx, userId)
}

func (m mongodb) UserInsert(ctx context.Context, user model.User) (int64, error) {
	id, err := m.crudSeq.Next(ctx, database.MongoCollectionUser)
	if err != nil {
		return 0, err
	}
	user.Id = id
	return id, m.crudUser.InsertOne(ctx, user)
}

func (m mongodb) UserLogout(ctx context.Context, userId int64) error {
	user, err := m.crudUser.FindOneByFilter(ctx, bson.D{{"_id", userId}})
	if err != nil {
		return err
	}

	now := time.Now()
	switch user.Authentication.AuthType {
	case types.StateAuthTypeId, types.StateAuthTypeToken, types.StateAuthTypeRotary:
		return m.crudUser.FindOneAndUpdateByFilter(
			ctx,
			bson.D{{"_id", user.Id}},
			bson.D{{"$set", &model.User{
				State:     types.StateUserInactive,
				StatedAt:  &now,
				UpdatedAt: &now,
			}}},
		)
	case types.StateAuthTypeAnonymous:
		return m.UserDeleteByUserId(ctx, userId)
	default:
		return nil
	}
}

func (m mongodb) UserFindAllByTopicIdAndPaging(ctx context.Context, topicId int64, paging model.Paging) ([]model.User, error) {
	var filter bson.D
	if paging != (model.Paging{}) && paging.UpdatedAt != nil {
		filter = bson.D{{"topic_id", topicId}, {"updated_at", bson.M{"$gt": paging.UpdatedAt}}}
	} else if paging != (model.Paging{}) && paging.CreatedAt != nil {
		filter = bson.D{{"topic_id", topicId}, {"created_at", bson.M{"$gt": paging.CreatedAt}}}
		//} else if paging != (model.Paging{}) && paging.Id > 0 {
		//	filter = bson.D{{"topic_id", topicId}, {"_id", bson.M{"$gt": paging.Id}}}
	} else {
		filter = bson.D{{"topic_id", topicId}}
	}

	cCtx, cancel := context.WithCancel(ctx)
	subs, err := m.crudSubs.FindManyByFilter(cCtx, filter)
	cancel()
	if err != nil {
		return nil, err
	}

	var users []model.User
	for _, sub := range subs {
		cCtx, cancel = context.WithCancel(ctx)
		user, err := m.crudUser.FindOneByFilter(cCtx, bson.D{{"_id", sub.UserId}})
		cancel()
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (m mongodb) UserFindAllByPaging(ctx context.Context, paging model.Paging) ([]model.User, error) {
	var filter bson.D
	var opt *options.FindOptions
	if paging != (model.Paging{}) && paging.Offset > 0 && paging.Limit > 0 {
		filter = bson.D{{"_id", bson.M{"$gte": paging.Offset, "$lt": paging.Offset + paging.Limit}}}
		opt = options.Find().SetSort(bson.D{{paging.By, paging.Order}})
	} else if paging != (model.Paging{}) && paging.UpdatedAt != nil {
		filter = bson.D{{"updated_at", bson.M{"$gt": paging.UpdatedAt}}}
		opt = options.Find().SetSort(bson.D{{"_id", -1}})
	} else if paging != (model.Paging{}) && paging.CreatedAt != nil {
		filter = bson.D{{"created_at", bson.M{"$gt": paging.CreatedAt}}}
		opt = options.Find().SetSort(bson.D{{"_id", -1}})
	} else if paging != (model.Paging{}) && paging.Id > 0 {
		filter = bson.D{{"_id", bson.M{"$lt": paging.Id}}}
		opt = options.Find().SetSort(bson.D{{"_id", -1}})
	} else {
		filter = bson.D{}
		opt = options.Find().SetSort(bson.D{{"_id", -1}})
	}

	return m.crudUser.FindManyByFilter(ctx, filter, opt)
}

func (m mongodb) UserFindOneById(ctx context.Context, userId int64) (model.User, error) {
	return m.crudUser.FindOneByFilter(ctx, bson.D{{"_id", userId}})
}

func (m mongodb) UserFindOneByEmail(ctx context.Context, email string) (model.User, error) {
	user, err := m.crudUser.FindOneByFilter(ctx, bson.D{{database.MongoUsersFieldAuthentication + ".email", email}})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, sql.ErrNoRows
		}
		return model.User{}, err
	}
	return user, nil
}

func (m mongodb) UserUpdate(ctx context.Context, user model.User) error {
	return m.crudUser.FindOneAndUpdateByFilter(
		ctx,
		bson.D{{"_id", user.Id}},
		bson.D{{"$set", user}},
	)
}
