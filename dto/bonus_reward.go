package dto

import (
	"fmt"
	"math/rand"
)

const BonusRewardType = "bonus"

type (
	BonusReward interface {
		SerializableReward
		GetUserId() int64
		GetAmount() int64
	}

	bonusRewardImpl struct {
		UserId int64
		Amount int64
	}
)

func (m bonusRewardImpl) GetUserId() int64 {
	return m.UserId
}

func (m bonusRewardImpl) GetAmount() int64 {
	return m.Amount
}

func (m bonusRewardImpl) Serialize() string {
	return fmt.Sprintf("{\"type\":\"%s\",\"amount\":%d}", BonusRewardType, m.Amount)
}

func NewBonusReward(userId int64, amount int64) BonusReward {
	return &bonusRewardImpl{
		UserId: userId,
		Amount: amount,
	}
}

func GenerateBonusAmount(max int64) int64 {
	return rand.Int63n(max)
}
