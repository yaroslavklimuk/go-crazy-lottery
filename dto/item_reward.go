package dto

type (
	ItemRewardType string
	ItemReward     interface {
		GetId() int64
		SetId(id int64)
		GetType() ItemRewardType
		IsSent() bool
	}

	itemRewardImpl struct {
		Id   int64
		Type ItemRewardType
		Sent bool
	}
)

const (
	BicycleItem ItemRewardType = "bicycle"
	CrutchItem  ItemRewardType = "crutch"
)

func (m itemRewardImpl) GetId() int64 {
	return m.Id
}

func (m *itemRewardImpl) SetId(id int64) {
	m.Id = id
}

func (m itemRewardImpl) GetType() ItemRewardType {
	return m.Type
}

func (m itemRewardImpl) IsSent() bool {
	return m.Sent
}

func NewItemReward(itemType ItemRewardType, sent bool, id int64) ItemReward {
	return &itemRewardImpl{
		Id:   id,
		Type: itemType,
		Sent: sent,
	}
}
