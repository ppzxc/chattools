package mongodb

import (
	"context"
	"database/sql"
	"github.com/ppzxc/chattools/storage/database"
	model2 "github.com/ppzxc/chattools/storage/database/model"
	"github.com/ppzxc/chattools/types"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func (m mongodb) registerAll(ctx context.Context, users []*model2.User) error {
	logrus.Debug("call func (m mongodb) registerAll(users []*model.User) error")
	var many []interface{}
	for i := 0; i < len(users); i++ {
		id, err := m.crudSequence.Next(ctx, database.MongoCollectionUser)
		if err != nil {
			return err
		}
		users[i].Id = id
		many = append(many, users[i])
	}
	return m.crudUser.InsertMany(ctx, many)
}

func (m mongodb) UserDeleteByUserId(ctx context.Context, userId int64) error {
	return m.crudUser.Delete(ctx, userId)
}

func (m mongodb) UserInsert(ctx context.Context, user model2.User) (int64, error) {
	id, err := m.crudSequence.Next(ctx, database.MongoCollectionUser)
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
			bson.D{{"$set", &model2.User{
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

func (m mongodb) UserFindAllByTopicId(ctx context.Context, topicId int64, paging model2.Paging) ([]model2.User, error) {
	cCtx, cancel := context.WithCancel(ctx)
	subs, err := m.crudSubs.FindManyByFilter(cCtx, bson.D{{"topic_id", topicId}})
	cancel()
	if err != nil {
		return nil, err
	}

	var users []model2.User
	for _, sub := range subs {
		var filter bson.D
		if paging != (model2.Paging{}) && paging.UpdatedAt != nil {
			filter = bson.D{{"_id", sub.UserId}, {"updated_at", bson.M{"$gt": paging.UpdatedAt}}}
		} else if paging != (model2.Paging{}) && paging.CreatedAt != nil {
			filter = bson.D{{"_id", sub.UserId}, {"created_at", bson.M{"$gt": paging.CreatedAt}}}
		} else {
			filter = bson.D{{"_id", sub.UserId}}
		}

		cCtx, cancel = context.WithCancel(ctx)
		user, err := m.crudUser.FindOneByFilter(cCtx, filter)
		//user, err := m.UserFindOneById(cCtx, sub.UserId)
		cancel()
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (m mongodb) UserFindAllByPaging(ctx context.Context, paging model2.Paging) ([]model2.User, error) {
	var filter bson.D
	//filter := bson.M{"updated_at": bson.M{"$gte": paging.UpdatedAt, "$lte": paging.CreatedAt}}

	if paging != (model2.Paging{}) && paging.UpdatedAt != nil {
		filter = bson.D{{"updated_at", bson.M{"$gt": paging.UpdatedAt}}}
	} else if paging != (model2.Paging{}) && paging.CreatedAt != nil {
		filter = bson.D{{"created_at", bson.M{"$gt": paging.CreatedAt}}}
	} else {
		filter = bson.D{}
	}

	return m.crudUser.FindManyByFilter(ctx, filter)
}

func (m mongodb) UserFindOneById(ctx context.Context, userId int64) (model2.User, error) {
	return m.crudUser.FindOneByFilter(ctx, bson.D{{"_id", userId}})
}

func (m mongodb) UserFindOneByEmail(ctx context.Context, email string) (model2.User, error) {
	user, err := m.crudUser.FindOneByFilter(ctx, bson.D{{database.MongoUsersFieldAuthentication + ".email", email}})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model2.User{}, sql.ErrNoRows
		}
		return model2.User{}, err
	}
	return user, nil
}

func (m mongodb) UserUpdate(ctx context.Context, user model2.User) error {
	return m.crudUser.FindOneAndUpdateByFilter(
		ctx,
		bson.D{{"_id", user.Id}},
		bson.D{{"$set", user}},
	)
}
