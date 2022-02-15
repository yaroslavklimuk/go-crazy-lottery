package dto

import (
	"fmt"
	"math/rand"
)

const (
	PhysItemRewardType                = "item"
	BicycleItem        ItemRewardType = "bicycle"
	CrutchItem         ItemRewardType = "crutch"
	MaxItems                          = 5
)

type (
	ItemRewardType string
	ItemReward     interface {
		SerializableReward
		GetId() int64
		SetId(id int64)
		GetUserId() int64
		GetType() ItemRewardType
		IsSent() bool
	}

	itemRewardImpl struct {
		Id     int64
		UserId int64
		Type   ItemRewardType
		Sent   bool
	}
)

func (m itemRewardImpl) GetId() int64 {
	return m.Id
}

func (m *itemRewardImpl) SetId(id int64) {
	m.Id = id
}

func (m itemRewardImpl) GetUserId() int64 {
	return m.UserId
}

func (m itemRewardImpl) GetType() ItemRewardType {
	return m.Type
}

func (m itemRewardImpl) IsSent() bool {
	return m.Sent
}

func (m itemRewardImpl) Serialize() string {
	return fmt.Sprintf("{\"type\":\"%s\",\"amount\":\"%s\"}", PhysItemRewardType, m.Type)
}

func NewItemReward(userId int64, itemType ItemRewardType, sent bool, id int64) ItemReward {
	return &itemRewardImpl{
		Id:     id,
		UserId: userId,
		Type:   itemType,
		Sent:   sent,
	}
}

func GenerateItemType() ItemRewardType {
	t := rand.Intn(2)
	if t == 0 {
		return BicycleItem
	} else {
		return CrutchItem
	}
}
