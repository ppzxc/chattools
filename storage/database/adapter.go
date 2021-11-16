package database

import (
	"context"
	"github.com/ppzxc/chattools/storage/database/model"
)

type Service interface {
	InitializeTable(ctx context.Context, dropTableOnStart bool, createTableOnStart bool, testUserInsertOnStart bool) error
	Close(context.Context) error

	UserInsert(ctx context.Context, user model.User) (id int64, err error)
	UserFindAllByPaging(ctx context.Context, paging model.Paging) (users []model.User, err error)
	UserFindAllByTopicIdAndPaging(ctx context.Context, topicId int64, paging model.Paging) ([]model.User, error)
	UserFindOneById(ctx context.Context, id int64) (user model.User, err error)
	UserFindOneByEmail(ctx context.Context, email string) (user model.User, err error)
	UserUpdate(ctx context.Context, user model.User) (err error)
	UserDeleteByUserId(ctx context.Context, userId int64) (err error)
	UserLogout(ctx context.Context, userId int64) (err error)

	TopicInsert(ctx context.Context, topic model.Topic) (topicId int64, err error)
	TopicFindAll(ctx context.Context, paging model.Paging) (topics []model.Topic, err error)
	TopicFindAllByUserId(ctx context.Context, userId int64, paging model.Paging) (topics []model.Topic, err error)
	TopicFindOneById(ctx context.Context, topicId int64) (topic model.Topic, err error)
	TopicDelete(ctx context.Context, topicId int64) (err error)

	SubscriptionInsert(ctx context.Context, subscription model.Subscription) (subscriptionId int64, err error)
	SubscriptionFindAllByTopicId(ctx context.Context, topicId int64) (subscriptions []model.Subscription, err error)
	SubscriptionFindAllByUserId(ctx context.Context, userId int64) (subscriptions []model.Subscription, err error)
	SubscriptionFindOneByUserIdAndTopicId(ctx context.Context, userId int64, topicId int64) (subscription model.Subscription, err error)
	SubscriptionExistsByUserIdAndTopicId(ctx context.Context, userId int64, topicId int64) (subscriptionId int64, err error)
	SubscriptionUpdateAck(ctx context.Context, subscription model.Subscription) (err error)
	SubscriptionDeleteAllByUserId(ctx context.Context, userId int64) (err error)
	SubscriptionDeleteByTopicIdAndUserId(ctx context.Context, topicId int64, userId int64) (err error)

	MessageInsert(ctx context.Context, message model.Message) (sequenceId int64, err error)
	MessageFindAllByTopicIdAndMoreThanSequenceId(ctx context.Context, topicId int64, sequenceId int64) ([]model.Message, error)
	MessageFindMaxSequenceIdByTopicId(ctx context.Context, topicId int64) (int64, error)

	FileFindOneById(ctx context.Context, fileId int64) (model.File, error)
	FileInsert(ctx context.Context, file model.File) (fileId int64, err error)

	ProfileImageUpdate(ctx context.Context, profile model.File) (fileId int64, err error)
	ProfileUpdateByUserId(ctx context.Context, profile model.Profile) (err error)

	NotifyInsertOne(ctx context.Context, notify model.Notify) (int64, error)
	NotifyInsertMany(ctx context.Context, notify []*model.Notify) (err error)
	NotifyUpdate(ctx context.Context, notify *model.Notify) (err error)
	NotifyFindAllByReceiveUserId(ctx context.Context, receiveUserId int64, paging model.Paging) ([]*model.Notify, error)
	NotifyFindOneReceiveUserIdById(ctx context.Context, notifyId int64) (int64, error)
}
