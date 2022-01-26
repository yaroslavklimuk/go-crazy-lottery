package storage

import "github.com/yaroslavklimuk/crazy-lottery/dto"

type (
	Storage interface {
		StoreSession(dto.Session) error
		GetSession(token string) (dto.Session, error)

		GetUserMoneyRewards(userId int64) (int64, error)
		StoreUserMoneyReward(base dto.Reward, money dto.MoneyReward) error

		GetUserItemRewards(userId int64) (int64, error)
		StoreUserItemReward(base dto.Reward, item dto.ItemReward) error

		GetUnprocessedMoneyRewards() ([]dto.MoneyReward, error)
		SetMoneyRewardsProcessed(ids []int64) error

		GetUnprocessedItemsRewards() ([]dto.ItemReward, error)
		SetItemsRewardsProcessed(ids []int64) error

		StoreUser(user dto.User) (int64, error)
		GetUserById(id int64) (dto.User, error)
		GetUserByName(name string) (dto.User, error)
	}
)
