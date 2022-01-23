package dto

type (
	RewardType string
	Reward     interface {
		GetId() uint
		SetId(id uint)
		GetRewardId() uint
		GetType() RewardType
	}

	rewardImpl struct {
		Id       uint
		RewardId uint
		Type     RewardType
	}
)

const (
	Money RewardType = "money"
	Item  RewardType = "item"
)

func (r rewardImpl) GetId() uint {
	return r.Id
}

func (r *rewardImpl) SetId(id uint) {
	r.Id = id
}

func (r rewardImpl) GetRewardId() uint {
	return r.RewardId
}

func (r rewardImpl) GetType() RewardType {
	return r.Type
}
