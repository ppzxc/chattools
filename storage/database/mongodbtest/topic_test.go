package mongodbtest

import (
	"context"
	"github.com/ppzxc/chattools/common/stats"
	"github.com/ppzxc/chattools/storage/database/model"
	"github.com/ppzxc/chattools/types"
	"github.com/ppzxc/chattools/utils/mono"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
	"time"
)

func topicCreate(ctx context.Context, t *testing.T, count int, owner int, time time.Time) {
	for i := 0; i < count; i++ {
		insertTopic := model.Topic{
			State:     types.StateTopicCreated,
			StatedAt:  &time,
			Name:      mono.GetMONOID(),
			Owner:     int64(owner),
			Private:   false,
			CreatedAt: &time,
			UpdatedAt: &time,
			//DeletedAt: nil,
			//Message:   []*model.Message{},
		}

		insertId, err := service.TopicInsert(ctx, insertTopic)
		require.NoError(t, err)
		insertTopic.Id = insertId
	}
}

func TestTopicTimezone(t *testing.T) {
	stats.Writer = stats.Initialize()
	mono.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_ = service.InitializeTable(ctx, true, true, false)

	now := time.Now()
	insertTopic := model.Topic{
		State:     types.StateTopicCreated,
		StatedAt:  &now,
		Name:      mono.GetMONOID(),
		Owner:     5,
		Private:   false,
		CreatedAt: &now,
		UpdatedAt: &now,
		//DeletedAt: nil,
		//Message:   []*model.Message{},
	}

	queryCtx, cancel := context.WithCancel(ctx)
	insertTopicId, err := service.TopicInsert(queryCtx, insertTopic)
	cancel()
	require.NoError(t, err)
	assert.Greater(t, insertTopicId, int64(0))

	insertTopic.Id = insertTopicId

	queryCtx, cancel = context.WithCancel(ctx)
	findTopic, err := service.TopicFindOneById(queryCtx, insertTopic.Id)
	cancel()
	require.NoError(t, err)

	assert.Equal(t, insertTopic.Id, findTopic.Id)
	assert.Equal(t, insertTopic.State, findTopic.State)
	timeEquals(t, insertTopic.StatedAt, findTopic.StatedAt)
	assert.Equal(t, insertTopic.Name, findTopic.Name)
	assert.Equal(t, insertTopic.Owner, findTopic.Owner)
	assert.Equal(t, insertTopic.Private, findTopic.Private)
	timeEquals(t, insertTopic.CreatedAt, findTopic.CreatedAt)
	timeEquals(t, insertTopic.UpdatedAt, findTopic.UpdatedAt)
	assert.Equal(t, insertTopic.Message, findTopic.Message)

	//fmt.Printf("################################## %v : %v", insertTopic.CreatedAt.Local(), findTopic.CreatedAt.Local())

	queryCtx, cancel = context.WithCancel(ctx)
	err = service.TopicDelete(queryCtx, insertTopic.Id)
	cancel()
	require.NoError(t, err)

	queryCtx, cancel = context.WithCancel(ctx)
	_, err = service.TopicFindOneById(queryCtx, insertTopic.Id)
	cancel()
	assert.Equal(t, err, mongo.ErrNoDocuments)
}

func TestTopicCrud(t *testing.T) {
	mono.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_ = service.InitializeTable(ctx, true, true, false)

	now := time.Now()
	insertTopic := model.Topic{
		State:     types.StateTopicCreated,
		StatedAt:  &now,
		Name:      mono.GetMONOID(),
		Owner:     5,
		Private:   false,
		CreatedAt: &now,
		UpdatedAt: &now,
		//DeletedAt: nil,
		//Message:   []*model.Message{},
	}

	queryCtx, cancel := context.WithCancel(ctx)
	insertTopicId, err := service.TopicInsert(queryCtx, insertTopic)
	cancel()
	require.NoError(t, err)
	assert.Greater(t, insertTopicId, int64(0))

	insertTopic.Id = insertTopicId

	queryCtx, cancel = context.WithCancel(ctx)
	findTopic, err := service.TopicFindOneById(queryCtx, insertTopic.Id)
	cancel()
	require.NoError(t, err)

	assert.Equal(t, insertTopic.Id, findTopic.Id)
	assert.Equal(t, insertTopic.State, findTopic.State)
	timeEquals(t, insertTopic.StatedAt, findTopic.StatedAt)
	assert.Equal(t, insertTopic.Name, findTopic.Name)
	assert.Equal(t, insertTopic.Owner, findTopic.Owner)
	assert.Equal(t, insertTopic.Private, findTopic.Private)
	timeEquals(t, insertTopic.CreatedAt, findTopic.CreatedAt)
	timeEquals(t, insertTopic.UpdatedAt, findTopic.UpdatedAt)
	assert.Equal(t, insertTopic.Message, findTopic.Message)

	queryCtx, cancel = context.WithCancel(ctx)
	err = service.TopicDelete(queryCtx, insertTopic.Id)
	cancel()
	require.NoError(t, err)

	queryCtx, cancel = context.WithCancel(ctx)
	_, err = service.TopicFindOneById(queryCtx, insertTopic.Id)
	cancel()
	assert.Equal(t, err, mongo.ErrNoDocuments)
}

func TestTopicFindAllByPaging1(t *testing.T) {
	mono.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_ = service.InitializeTable(ctx, true, true, false)

	firstTime := time.Now().Add(-12 * time.Hour)
	topicCreate(ctx, t, 10, 1, firstTime)
	secondTime := time.Now().Add(-24 * time.Hour)
	topicCreate(ctx, t, 20, 2, secondTime)

	queryCtx, cancel := context.WithCancel(ctx)
	topics, err := service.TopicFindAll(queryCtx, model.Paging{})
	cancel()
	require.NoError(t, err)
	assert.Equal(t, len(topics), 30)

	queryCtx, cancel = context.WithCancel(ctx)
	findTime := time.Now().Add(-13 * time.Hour)
	topics, err = service.TopicFindAll(queryCtx, model.Paging{
		CreatedAt: &findTime,
		//UpdatedAt: nil,
		//Id:        0,
	})
	cancel()
	require.NoError(t, err)
	assert.Equal(t, len(topics), 10)

	queryCtx, cancel = context.WithCancel(ctx)
	findTime2 := time.Now().Add(-13 * time.Hour)
	topics, err = service.TopicFindAll(queryCtx, model.Paging{
		//CreatedAt: nil,
		UpdatedAt: &findTime2,
		//Id:        0,
	})
	cancel()
	require.NoError(t, err)
	assert.Equal(t, len(topics), 10)

	queryCtx, cancel = context.WithCancel(ctx)
	topics, err = service.TopicFindAll(queryCtx, model.Paging{
		//CreatedAt: nil,
		//UpdatedAt: nil,
		Id: 16,
	})
	cancel()
	require.NoError(t, err)
	assert.Equal(t, len(topics), 15)
}
