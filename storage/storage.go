package storage

import "github.com/yaroslavklimuk/crazy-lottery/dto"

type (
	Storage interface {
		StoreSession(dto.Session) error
		GetSession(token string) (dto.Session, error)

		GetUserMoneyRewards(userId int64) (int64, error)
		StoreUserMoneyReward(money dto.MoneyReward) (int64, error)

		GetUserItemRewards(userId int64) (int64, error)
		StoreUserItemReward(item dto.ItemReward) (int64, error)

		GetUnprocessedMoneyRewards() ([]dto.MoneyReward, error)
		SetMoneyRewardsProcessed(ids []int64) error

		GetUnprocessedItemsRewards() ([]dto.ItemReward, error)
		SetItemsRewardsProcessed(ids []int64) error

		StoreUser(user dto.User) (int64, error)
		UpdateBalance(userId int64, sum int) error
		GetUserById(id int64) (dto.User, error)
		GetUserByName(name string) (dto.User, error)
	}
)
