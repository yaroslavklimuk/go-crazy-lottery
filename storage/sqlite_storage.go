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

func (s *sqliteStorageImpl) GetUserMoneyRewards(userId int64) (int64, error) {
	row := s.conn.QueryRow(
		"SELECT SUM(mr.amount) FROM money_rewards as mr INNER JOIN rewards as r "+
			"ON r.reward_id = mr.id AND r.type = \"?\" AND r.user_id = ?",
		dto.Money, userId,
	)
	var amount int64
	err := row.Scan(&amount)
	if err != nil {
		return 0, err
	}
	return amount, nil
}

func (s *sqliteStorageImpl) StoreUserMoneyReward(base dto.Reward, money dto.MoneyReward) error {
	stmtMoney, err := s.conn.Prepare("INSERT INTO money_rewards (amount) VALUES (?)")
	if err != nil {
		return err
	}
	moneyRes, err := stmtMoney.Exec(money.GetAmount())
	moneyId, err := moneyRes.LastInsertId()
	if err != nil {
		return err
	}

	stmtBase, err := s.conn.Prepare("INSERT INTO rewards (reward_id, user_id, type) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmtBase.Exec(moneyId, base.GetUserId(), base.GetType())
	return err
}

func (s *sqliteStorageImpl) GetUserItemRewards(userId int64) (int64, error) {
	row := s.conn.QueryRow(
		"SELECT COUNT(ir.*) FROM item_rewards as ir INNER JOIN rewards as r "+
			"ON r.reward_id = ir.id AND r.type = \"?\" AND r.user_id = ?",
		dto.Item, userId,
	)
	var count int64
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *sqliteStorageImpl) StoreUserItemReward(base dto.Reward, item dto.ItemReward) error {
	stmtItem, err := s.conn.Prepare("INSERT INTO item_rewards (type) VALUES (?)")
	if err != nil {
		return err
	}
	itemRes, err := stmtItem.Exec(item.GetType())
	itemId, err := itemRes.LastInsertId()
	if err != nil {
		return err
	}

	stmtBase, err := s.conn.Prepare("INSERT INTO rewards (reward_id, user_id, type) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmtBase.Exec(itemId, base.GetUserId(), base.GetType())
	return err
}

func (s *sqliteStorageImpl) GetUnprocessedMoneyRewards() ([]dto.MoneyReward, error) {
	rows, err := s.conn.Query("SELECT id, amount FROM money_rewards WHERE sent = 0")
	if err != nil {
		return nil, err
	}
	var id, amount int64
	var result []dto.MoneyReward
	for rows.Next() {
		if err := rows.Scan(&id, &amount); err != nil {
			return nil, err
		}
		item := dto.NewMoneyReward(amount, false, id)
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
	rows, err := s.conn.Query("SELECT id, type FROM item_rewards WHERE sent = 0")
	if err != nil {
		return nil, err
	}
	var id int64
	var itemType dto.ItemRewardType
	var result []dto.ItemReward
	for rows.Next() {
		if err := rows.Scan(&id, &itemType); err != nil {
			return nil, err
		}
		item := dto.NewItemReward(itemType, false, id)
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

func (s *sqliteStorageImpl) StoreUser(user dto.User) (int64, error) {
	stmt, err := s.conn.Prepare(
		"INSERT INTO users (name, passwd, banc_acc, address, balance) VALUES (?, ?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(user.GetName(), user.GetPswdHash(), user.GetBankAcc(), user.GetAddress(), 0)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return id, err
}

func (s *sqliteStorageImpl) GetUserById(id int64) (dto.User, error) {
	row := s.conn.QueryRow("SELECT name, passwd, banc_acc, address, balance FROM users WHERE id = ?", id)
	var name, passwd, bancAcc, address string
	var balance int64
	err := row.Scan(&name, &passwd, &bancAcc, &address, &balance)
	if err != nil {
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
		return nil, err
	}
	return dto.NewUser(id, name, bancAcc, address, balance, passwd), nil
}
