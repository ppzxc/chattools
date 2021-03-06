package mongodb

import (
	"context"
	"github.com/ppzxc/chattools/common"
	"github.com/ppzxc/chattools/storage/database"
	"github.com/ppzxc/chattools/storage/database/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m mongodb) MessageFindByPaging(ctx context.Context, topicId int64, paging model.Paging) ([]model.Message, error) {
	//return m.crudMsg.FindManyByFilter(ctx, bson.D{{"topic_id", topicId}, {"sequence_id", bson.M{"$gte": paging.Offset, "$lt": paging.Offset + paging.Limit}}},
	var gt int64
	// in my case
	// limit 1~100
	// offset 1~999999999999
	// condition
	// offset 101~99999999 if
	// else offset == 1~99
	if paging.Offset > paging.Limit {
		gt = paging.Offset - paging.Limit
	} else {
		gt = 0
	}

	return m.crudMsg.FindManyByFilter(ctx, bson.D{{"topic_id", topicId}, {"sequence_id", bson.M{"$lte": paging.Offset, "$gt": gt}}},
		options.Find().SetSort(bson.D{{paging.By, paging.Order}}))
}

func (m mongodb) MessageMaxIdByTopicId(ctx context.Context, topicId int64) (maxId int64, err error) {
	findOptions := options.FindOptions{}
	findOptions.SetSort(bson.D{{"_id", -1}})
	findOptions.SetLimit(1)
	messages, err := m.crudMsg.FindManyByFilter(ctx, bson.D{{"topic_id", topicId}}, &findOptions)
	if err != nil {
		return 0, err
	}

	if len(messages) <= 0 {
		return 0, mongo.ErrNoDocuments
	}

	return messages[0].Id, nil
}

func (m mongodb) MessageInsert(ctx context.Context, message model.Message) (sequenceId int64, err error) {
	sequence, err := m.crudSeq.Next(ctx, m.crudSeq.TopicMaxSeq(database.MongoCollectionMessage, message.TopicId))
	if err != nil {
		return 0, err
	}
	message.SequenceId = sequence

	id, err := m.crudSeq.Next(ctx, database.MongoCollectionMessage)
	if err != nil {
		return 0, err
	}
	message.Id = id

	if message.Custom == nil {
		message.Custom = common.FromByteToMap([]byte("{}"))
	}

	err = m.crudMsg.InsertOne(ctx, message)
	if err != nil {
		return 0, err
	}

	return message.SequenceId, nil
}

func (m mongodb) MessageFindAllByTopicIdAndMoreThanSequenceId(ctx context.Context, topicId int64, sequenceId int64) ([]model.Message, error) {
	maxSequenceId, err := m.crudSeq.Current(ctx, database.MongoCollectionMessage, topicId)
	if err != nil {
		return nil, err
	}

	if maxSequenceId-sequenceId > common.FindCount {
		return m.crudMsg.FindManyByFilter(ctx, bson.D{{"topic_id", topicId}, {"sequence_id", bson.D{{"$gte", maxSequenceId - common.FindCount}}}})
	} else {
		return m.crudMsg.FindManyByFilter(ctx, bson.D{{"topic_id", topicId}, {"sequence_id", bson.D{{"$gte", sequenceId}}}})
	}
}

func (m mongodb) MessageFindMaxSequenceIdByTopicId(ctx context.Context, topicId int64) (int64, error) {
	sequence, err := m.crudSeq.Current(ctx, database.MongoCollectionMessage, topicId)
	if err != nil {
		return 0, err
	}
	return sequence, nil
}
