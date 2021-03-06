package impl

import (
	"context"
	"fmt"
	"github.com/ppzxc/chattools/common/stats"
	"github.com/ppzxc/chattools/storage/database/model"
	"github.com/ppzxc/chattools/storage/database/mongodb/repository"
	"github.com/ppzxc/chattools/utils"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type sequence struct {
	//database     *mongo.Database
	counter      *mongo.Collection
	queryTimeout time.Duration
}

func NewSequenceRepository(collection *mongo.Collection, queryTimeout time.Duration) repository.Sequence {
	return &sequence{
		//database:     db,
		counter:      collection,
		queryTimeout: queryTimeout,
	}
}

func (s sequence) TopicMaxSeq(collectionName string, topicId int64) string {
	return fmt.Sprintf("topic_%v_%v", topicId, collectionName)
}

func (s sequence) Current(ctx context.Context, collectionName string, topicId int64) (int64, error) {
	seq := model.Serial{}
	cCtx, cancel := context.WithCancel(ctx)
	start := time.Now()
	err := s.counter.FindOne(cCtx,
		bson.D{{Key: "_id", Value: s.TopicMaxSeq(collectionName, topicId)}},
	).Decode(&seq)
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "s.counter.FindOne",
		"exec.time": time.Since(start).String(),
	})).Debug("sql execute")
	stats.QueryRecord(stats.SELECT, "counters", "Current", start)
	if err != nil {
		return 0, err
	}
	return seq.Seq, nil
}

func (s sequence) Next(ctx context.Context, collectionName string) (int64, error) {
	seq := model.Serial{}
	cCtx, cancel := context.WithCancel(ctx)
	start := time.Now()
	err := s.counter.FindOneAndUpdate(cCtx,
		bson.D{{Key: "_id", Value: collectionName}},
		bson.D{{Key: "$inc", Value: bson.D{{Key: "seq", Value: 1}}}},
		options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After),
	).Decode(&seq)
	cancel()
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{
		"query":     "s.counter.FindOneAndUpdate",
		"exec.time": time.Since(start).String(),
	})).Debug("sql execute")
	stats.QueryRecord(stats.UPDATE, "counters", "Next", start)
	if err != nil {
		return 0, err
	}
	return seq.Seq, nil
}
