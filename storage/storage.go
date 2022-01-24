package storage

import "github.com/yaroslavklimuk/crazy-lottery/dto"

type (
	Storage interface {
		StoreSession(dto.Session) error
		GetSession(token string) (dto.Session, error)

		GetUserMoneyRewards(userId int64) (int64, error)
		StoreUserMoney(base dto.Reward, money dto.MoneyReward) error

		GetUserItemRewards(userId int64) (int64, error)
		StoreUserItem(base dto.Reward, item dto.ItemReward) error

		GetUnprocessedMoneyRewards() ([]dto.MoneyReward, error)
		SetMoneyRewardsProcessed(ids []int64) error

		GetUnprocessedItemsRewards() ([]dto.ItemReward, error)
		SetItemsRewardsProcessed(ids []int64) error
	}
)
