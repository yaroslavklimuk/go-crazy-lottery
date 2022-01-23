package dto

type (
	ItemRewardType string
	ItemReward     interface {
		GetId() uint
		SetId(id uint)
		GetType() ItemRewardType
		IsSent() bool
	}

	itemRewardImpl struct {
		Id   uint
		Type ItemRewardType
		Sent bool
	}
)

const (
	BicycleItem ItemRewardType = "bicycle"
	CrutchItem  ItemRewardType = "crutch"
)

func (m itemRewardImpl) GetId() uint {
	return m.Id
}

func (m *itemRewardImpl) SetId(id uint) {
	m.Id = id
}

func (m itemRewardImpl) GetType() ItemRewardType {
	return m.Type
}

func (m itemRewardImpl) IsSent() bool {
	return m.Sent
}
