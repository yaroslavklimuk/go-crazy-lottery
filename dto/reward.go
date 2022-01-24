package dto

type (
	RewardType string
	Reward     interface {
		GetId() int64
		SetId(id int64)
		GetUserId() int64
		GetRewardId() int64
		GetType() RewardType
	}

	rewardImpl struct {
		Id       int64
		UserId   int64
		RewardId int64
		Type     RewardType
	}
)

const (
	Money RewardType = "money"
	Item  RewardType = "item"
)

func (r rewardImpl) GetId() int64 {
	return r.Id
}

func (r *rewardImpl) SetId(id int64) {
	r.Id = id
}

func (r rewardImpl) GetUserId() int64 {
	return r.UserId
}

func (r rewardImpl) GetRewardId() int64 {
	return r.RewardId
}

func (r rewardImpl) GetType() RewardType {
	return r.Type
}
