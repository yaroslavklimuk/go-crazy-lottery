package dto

type (
	MoneyReward interface {
		GetId() uint
		SetId(id uint)
		GetAmount() int
		IsSent() bool
	}

	moneyRewardImpl struct {
		Id     uint
		Amount int
		Sent   bool
	}
)

func (m moneyRewardImpl) GetId() uint {
	return m.Id
}

func (m *moneyRewardImpl) SetId(id uint) {
	m.Id = id
}

func (m moneyRewardImpl) GetAmount() int {
	return m.Amount
}

func (m moneyRewardImpl) IsSent() bool {
	return m.Sent
}
