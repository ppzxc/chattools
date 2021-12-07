package mongodb

import (
	"context"
	"database/sql"
	"github.com/ppzxc/chattools/storage/database"
	"github.com/ppzxc/chattools/storage/database/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TopicInsert is call to require new room
func (m mongodb) TopicInsert(ctx context.Context, topic model.Topic) (int64, error) {
	id, err := m.crudSeq.Next(ctx, database.MongoCollectionTopic)
	if err != nil {
		return 0, err
	}
	topic.Id = id

	err = m.crudTopic.InsertOne(ctx, topic)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m mongodb) TopicDelete(ctx context.Context, topicId int64) error {
	_ = m.crudMsg.Delete(ctx, bson.D{{"topic_id", topicId}})
	return m.crudTopic.Delete(ctx, topicId)
}

func (m mongodb) TopicDeleteByEmptySubs(ctx context.Context) error {
	all, err := m.crudTopic.FindManyFilter(ctx, bson.D{}, nil)
	if err != nil {
		return err
	}

	// topic empty
	if all == nil || len(all) <= 0 {
		return nil
	}

	for _, topic := range all {
		subs, err := m.crudSubs.FindManyByFilter(ctx, bson.M{"topic_id": topic.Id})
		if err != nil {
			return err
		}

		if err != nil {
			if err != mongo.ErrNoDocuments {
				err := m.crudTopic.Delete(ctx, topic.Id)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		} else if subs != nil || len(subs) <= 0 {
			err := m.crudTopic.Delete(ctx, topic.Id)
			if err != nil {
				return err
			}

			_ = m.crudMsg.Delete(ctx, bson.D{{"topic_id", topic.Id}})
		}
	}

	return nil
}

func (m mongodb) TopicFindAll(ctx context.Context, paging model.Paging) ([]model.Topic, error) {
	var filter bson.D
	//filter := bson.M{"updated_at": bson.M{"$gte": paging.UpdatedAt, "$lte": paging.CreatedAt}}

	if paging != (model.Paging{}) && paging.UpdatedAt != nil {
		filter = bson.D{{"updated_at", bson.M{"$gt": paging.UpdatedAt}}}
	} else if paging != (model.Paging{}) && paging.CreatedAt != nil {
		filter = bson.D{{"created_at", bson.M{"$gt": paging.CreatedAt}}}
	} else if paging != (model.Paging{}) && paging.Id > 0 {
		filter = bson.D{{"_id", bson.M{"$lt": paging.Id}}}
	} else {
		filter = bson.D{}
	}

	return m.crudTopic.FindManyFilter(ctx, filter, options.Find().
		SetSort(bson.D{{"_id", -1}}).
		SetLimit(100))
}

func (m mongodb) TopicFindIdsByUserId(ctx context.Context, userId int64) ([]int64, error) {
	subs, err := m.crudSubs.FindManyByFilter(ctx, bson.M{"user_id": userId})
	if err != nil {
		if err == sql.ErrNoRows {
			return []int64{}, nil
		}
		if err == mongo.ErrNoDocuments {
			return []int64{}, nil
		}
		return nil, err
	}

	var topicIds []int64
	for _, v := range subs {
		topicIds = append(topicIds, v.TopicId)
	}

	return topicIds, nil
}

func (m mongodb) TopicFindAllByUserId(ctx context.Context, userId int64, paging model.Paging) ([]model.Topic, error) {
	subs, err := m.crudSubs.FindManyByFilter(ctx, bson.M{"user_id": userId})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	var topicIds []int64
	for _, v := range subs {
		topicIds = append(topicIds, v.TopicId)
	}

	if len(topicIds) <= 0 {
		return nil, sql.ErrNoRows
	}

	var filter bson.D
	var opt *options.FindOptions
	if paging != (model.Paging{}) && paging.Offset > 0 && paging.Limit > 0 {
		filter = bson.D{{"_id", bson.M{"$in": topicIds}}, {"_id", bson.M{"$gte": paging.Offset, "$lt": paging.Offset + paging.Limit}}}
		opt = options.Find().SetSort(bson.D{{paging.By, paging.Order}})
	} else if paging != (model.Paging{}) && paging.UpdatedAt != nil {
		filter = bson.D{{"_id", bson.M{"$in": topicIds}}, {"updated_at", bson.M{"$gt": paging.UpdatedAt}}}
		opt = options.Find().
			SetSort(bson.D{{"_id", 1}})
	} else if paging != (model.Paging{}) && paging.CreatedAt != nil {
		filter = bson.D{{"_id", bson.M{"$in": topicIds}}, {"created_at", bson.M{"$gt": paging.CreatedAt}}}
		opt = options.Find().
			SetSort(bson.D{{"_id", 1}})
	} else if paging != (model.Paging{}) && paging.Id > 0 {
		filter = bson.D{{"_id", bson.M{"$in": topicIds}}, {"_id", bson.M{"$gt": paging.Id}}}
		opt = options.Find().
			SetSort(bson.D{{"_id", 1}})
	} else {
		filter = bson.D{{"_id", bson.M{"$in": topicIds}}}
		opt = options.Find().SetSort(bson.D{{"_id", 1}})
	}

	return m.crudTopic.FindManyFilter(ctx, filter, opt)
}

func (m mongodb) TopicFindOneById(ctx context.Context, topicId int64) (model.Topic, error) {
	return m.crudTopic.FindOneByFilter(ctx, bson.D{{"_id", topicId}})
}

func (m mongodb) TopicMaxIdByUserId(ctx context.Context, userId int64) (maxTopicId int64, err error) {
	option := options.FindOptions{}
	option.SetSort(bson.D{{"topic_id", -1}})
	option.SetLimit(1)

	subs, err := m.crudSubs.FindManyByFilter(ctx, bson.D{{"user_id", userId}}, &option)
	if err != nil {
		return 0, err
	}
	if len(subs) <= 0 {
		return 0, mongo.ErrNoDocuments
	}

	return subs[0].TopicId, nil
}

func (m mongodb) TopicCountDocumentsByUserId(ctx context.Context, userId int64) (count int64, err error) {
	count, err = m.crudSubs.CountDocuments(ctx, bson.D{{"user_id", userId}})
	if err == nil && count <= 0 {
		return 0, mongo.ErrNoDocuments
	}
	return
}
