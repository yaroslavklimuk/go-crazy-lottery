package dto

type (
	User interface {
		GetId() int64
		SetId(id int64)
		GetName() string
		GetBankAcc() string
		GetAddress() string
		GetBalance() int64
		GetPswdHash() string
	}

	userImpl struct {
		Id       int64
		Name     string
		BankAcc  string
		Address  string
		Balance  int64
		PswdHash string
	}
)

func (u userImpl) GetId() int64 {
	return u.Id
}

func (u *userImpl) SetId(id int64) {
	u.Id = id
}

func (u userImpl) GetName() string {
	return u.Name
}

func (u userImpl) GetBankAcc() string {
	return u.BankAcc
}

func (u userImpl) GetBalance() int64 {
	return u.Balance
}

func (u userImpl) GetAddress() string {
	return u.Address
}

func (u userImpl) GetPswdHash() string {
	return u.PswdHash
}

func NewUser(id int64, name string, bankAcc string, address string, balance int64, pswdHash string) User {
	return &userImpl{
		Id:       id,
		Name:     name,
		BankAcc:  bankAcc,
		Address:  address,
		Balance:  balance,
		PswdHash: pswdHash,
	}
}
