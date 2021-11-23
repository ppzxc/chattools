package mongodb

import (
	"context"
	"github.com/ppzxc/chattools/common"
	"github.com/ppzxc/chattools/storage/database"
	"github.com/ppzxc/chattools/storage/database/model"
	"go.mongodb.org/mongo-driver/bson"
)

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
