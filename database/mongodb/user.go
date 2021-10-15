package mongodb

import (
	"context"
	"database/sql"
	"github.com/ppzxc/chattools/database"
	"github.com/ppzxc/chattools/database/model"
	"github.com/ppzxc/chattools/types"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func (m mongodb) registerAll(ctx context.Context, users []*model.User) error {
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

func (m mongodb) UserInsert(ctx context.Context, user model.User) (int64, error) {
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

func (m mongodb) UserFindAllByPaging(ctx context.Context, paging model.Paging) ([]model.User, error) {
	var filter bson.D
	//filter := bson.M{"updated_at": bson.M{"$gte": paging.UpdatedAt, "$lte": paging.CreatedAt}}

	if paging != (model.Paging{}) && paging.UpdatedAt != nil {
		filter = bson.D{{"updated_at", bson.M{"$gt": paging.UpdatedAt}}}
	} else if paging != (model.Paging{}) && paging.CreatedAt != nil {
		filter = bson.D{{"created_at", bson.M{"$gt": paging.CreatedAt}}}
	} else {
		filter = bson.D{}
	}

	return m.crudUser.FindManyByFilter(ctx, filter)
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
