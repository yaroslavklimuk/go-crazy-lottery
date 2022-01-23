package dto

type (
	User interface {
		GetId() uint
		SetId(id uint)
		GetName() string
		GetBankAcc() string
		GetAddress() string
		GetBalance() int
	}

	userImpl struct {
		Id      uint
		Name    string
		BankAcc string
		Address string
		Balance int
	}
)

func (u userImpl) GetId() uint {
	return u.Id
}

func (u *userImpl) SetId(id uint) {
	u.Id = id
}

func (u userImpl) GetName() string {
	return u.Name
}

func (u userImpl) GetBankAcc() string {
	return u.BankAcc
}

func (u userImpl) GetBalance() int {
	return u.Balance
}

func (u userImpl) GetAddress() string {
	return u.Address
}

func NewUser(id uint, name string, bankAcc string, address string, balance int) User {
	return &userImpl{
		Id:      id,
		Name:    name,
		BankAcc: bankAcc,
		Address: address,
		Balance: balance,
	}
}
