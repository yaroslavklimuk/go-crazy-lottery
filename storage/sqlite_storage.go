package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yaroslavklimuk/crazy-lottery/dto"
)

var sqliteStorage Storage

type (
	sqliteStorageImpl struct {
		dbFile string
		conn   *sql.DB
	}
)

func GetStorage(dbFile string) (Storage, error) {
	if sqliteStorage == nil {
		sqliteStorage, err := createSqliteStorage(dbFile)
		return sqliteStorage, err
	}
	return sqliteStorage, nil
}

func createSqliteStorage(dbFile string) (*sqliteStorageImpl, error) {
	conn, err := connect(dbFile)
	if err != nil {
		return nil, err
	}
	return &sqliteStorageImpl{conn: conn, dbFile: dbFile}, nil
}

func connect(dbFile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (s *sqliteStorageImpl) StoreSession() error {
	//TODO implement me
	panic("implement me")
}

func (s *sqliteStorageImpl) CheckSession(token string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sqliteStorageImpl) GetUserMoneyRewards(userId uint) (float32, float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sqliteStorageImpl) StoreUserMoney(base dto.Reward, money dto.MoneyReward) error {
	//TODO implement me
	panic("implement me")
}

func (s *sqliteStorageImpl) GetUserItemRewards(userId uint) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sqliteStorageImpl) StoreUserItem(base dto.Reward, item dto.ItemReward) error {
	//TODO implement me
	panic("implement me")
}

func (s *sqliteStorageImpl) GetUnprocessedMoneyRewards() ([]dto.MoneyReward, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sqliteStorageImpl) SetMoneyRewardsProcessed(ids []uint) error {
	//TODO implement me
	panic("implement me")
}

func (s *sqliteStorageImpl) GetUnprocessedItemsRewards() ([]dto.ItemReward, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sqliteStorageImpl) SetItemsRewardsProcessed(ids []uint) error {
	//TODO implement me
	panic("implement me")
}
