package dto

type (
	MoneyReward interface {
		GetId() int64
		SetId(id int64)
		GetUserId() int64
		GetAmount() int64
		IsSent() bool
	}

	moneyRewardImpl struct {
		Id     int64
		UserId int64
		Amount int64
		Sent   bool
	}
)

func (m moneyRewardImpl) GetId() int64 {
	return m.Id
}

func (m *moneyRewardImpl) SetId(id int64) {
	m.Id = id
}

func (m moneyRewardImpl) GetUserId() int64 {
	return m.UserId
}

func (m moneyRewardImpl) GetAmount() int64 {
	return m.Amount
}

func (m moneyRewardImpl) IsSent() bool {
	return m.Sent
}

func NewMoneyReward(userId int64, amount int64, sent bool, id int64) MoneyReward {
	return &moneyRewardImpl{
		Id:     id,
		UserId: userId,
		Amount: amount,
		Sent:   sent,
	}
}
