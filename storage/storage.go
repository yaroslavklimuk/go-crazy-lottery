package storage

import "github.com/yaroslavklimuk/crazy-lottery/dto"

type (
	Storage interface {
		StoreSession() error
		CheckSession(token string) (bool, error)

		GetUserMoneyRewards(userId uint) (float32, float32, error)
		StoreUserMoney(base dto.Reward, money dto.MoneyReward) error

		GetUserItemRewards(userId uint) (uint, error)
		StoreUserItem(base dto.Reward, item dto.ItemReward) error

		GetUnprocessedMoneyRewards() ([]dto.MoneyReward, error)
		SetMoneyRewardsProcessed(ids []uint) error

		GetUnprocessedItemsRewards() ([]dto.ItemReward, error)
		SetItemsRewardsProcessed(ids []uint) error
	}
)
