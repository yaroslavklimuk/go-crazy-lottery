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

func (s *sqliteStorageImpl) StoreSession(sess dto.Session) error {
	stmt, err := s.conn.Prepare("INSERT INTO sessions (token, user_id, expired_at) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sess.GetToken(), sess.GetUserId(), sess.GetExpiredAt())
	return err
}

func (s *sqliteStorageImpl) GetSession(token string) (dto.Session, error) {
	row := s.conn.QueryRow("SELECT user_id, expired_at FROM sessions WHERE token = ?", token)
	var userId, expiredAt int64
	err := row.Scan(&userId, &expiredAt)
	if err != nil {
		return nil, err
	}
	return dto.NewSession(token, userId, expiredAt), nil
}

func (s *sqliteStorageImpl) StoreUser(user dto.User) (int64, error) {
	stmt, err := s.conn.Prepare(
		"INSERT INTO users (name, passwd, banc_acc, address, balance) VALUES (?, ?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(user.GetName(), user.GetPswdHash(), user.GetBankAcc(), user.GetAddress(), user.GetBalance())
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return id, err
}

func (s *sqliteStorageImpl) UpdateBalance(userId int64, sum int) error {
	stmt, err := s.conn.Prepare("UPDATE users SET balance = balance + ? WHERE id = ?;")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sum, userId)
	return err
}

func (s *sqliteStorageImpl) GetUserMoneyRewards(userId int64) (int64, error) {
	row := s.conn.QueryRow("SELECT SUM(amount) FROM money_rewards WHERE user_id = ? GROUP BY user_id", userId)
	var amount int64
	err := row.Scan(&amount)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return amount, nil
}

func (s *sqliteStorageImpl) StoreUserMoneyReward(reward dto.MoneyReward) (int64, error) {
	stmtMoney, err := s.conn.Prepare("INSERT INTO money_rewards (user_id, amount) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	res, err := stmtMoney.Exec(reward.GetUserId(), reward.GetAmount())
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (s *sqliteStorageImpl) GetUserItemRewards(userId int64) (int64, error) {
	row := s.conn.QueryRow("SELECT COUNT(*) FROM item_rewards WHERE user_id = ?", userId)
	var count int64
	err := row.Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return count, nil
}

func (s *sqliteStorageImpl) StoreUserItemReward(item dto.ItemReward) (int64, error) {
	stmtItem, err := s.conn.Prepare("INSERT INTO item_rewards (user_id, type) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	itemRes, err := stmtItem.Exec(item.GetUserId(), item.GetType())
	if err != nil {
		return 0, err
	}
	return itemRes.LastInsertId()
}

func (s *sqliteStorageImpl) GetUnprocessedMoneyRewards() ([]dto.MoneyReward, error) {
	rows, err := s.conn.Query("SELECT id, user_id, amount FROM money_rewards WHERE sent = 0")
	if err != nil {
		return nil, err
	}
	var id, userId, amount int64
	var result []dto.MoneyReward
	for rows.Next() {
		if err := rows.Scan(&id, &userId, &amount); err != nil {
			return nil, err
		}
		item := dto.NewMoneyReward(userId, amount, false, id)
		result = append(result, item)
	}
	return result, err
}

func (s *sqliteStorageImpl) SetMoneyRewardsProcessed(ids []int64) error {
	stmt, err := s.conn.Prepare("UPDATE money_rewards SET sent = 1 WHERE id = ?")
	if err != nil {
		return err
	}
	for _, id := range ids {
		_, err = stmt.Exec(id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *sqliteStorageImpl) GetUnprocessedItemsRewards() ([]dto.ItemReward, error) {
	rows, err := s.conn.Query("SELECT id, user_id, type FROM item_rewards WHERE sent = 0")
	if err != nil {
		return nil, err
	}
	var id, userId int64
	var itemType dto.ItemRewardType
	var result []dto.ItemReward
	for rows.Next() {
		if err := rows.Scan(&id, &userId, &itemType); err != nil {
			return nil, err
		}
		item := dto.NewItemReward(userId, itemType, false, id)
		result = append(result, item)
	}
	return result, err
}

func (s *sqliteStorageImpl) SetItemsRewardsProcessed(ids []int64) error {
	stmt, err := s.conn.Prepare("UPDATE item_rewards SET sent = 1 WHERE id = ?")
	if err != nil {
		return err
	}
	for _, id := range ids {
		_, err = stmt.Exec(id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *sqliteStorageImpl) GetUserById(id int64) (dto.User, error) {
	row := s.conn.QueryRow("SELECT name, passwd, banc_acc, address, balance FROM users WHERE id = ?", id)
	var name, passwd, bancAcc, address string
	var balance int64
	err := row.Scan(&name, &passwd, &bancAcc, &address, &balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return dto.NewUser(id, name, bancAcc, address, balance, passwd), nil
}

func (s *sqliteStorageImpl) GetUserByName(name string) (dto.User, error) {
	row := s.conn.QueryRow("SELECT id, passwd, banc_acc, address, balance FROM users WHERE name = ?", name)
	var passwd, bancAcc, address string
	var id, balance int64
	err := row.Scan(&id, &passwd, &bancAcc, &address, &balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return dto.NewUser(id, name, bancAcc, address, balance, passwd), nil
}
