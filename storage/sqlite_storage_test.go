package storage

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yaroslavklimuk/crazy-lottery/dto"
	"reflect"
	"testing"
)

type (
	testSessionDto struct {
		Token     string
		UserId    int64
		ExpiredAt int64
	}
	testUserDto struct {
		Name     string
		BancAcc  string
		Address  string
		Balance  int64
		PswdHash string
	}
	testUserMoneyDto struct {
		Id     int64
		UserId int64
		Amount int64
		Sent   bool
	}
)

func TestGetStorage(t *testing.T) {
	type args struct {
		dbFile string
	}
	tests := []struct {
		name    string
		args    args
		want    Storage
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStorage(tt.args.dbFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStorage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStorage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_connect(t *testing.T) {
	type args struct {
		dbFile string
	}
	tests := []struct {
		name    string
		args    args
		want    *sql.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := connect(tt.args.dbFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createSqliteStorage(t *testing.T) {
	type args struct {
		dbFile string
	}
	tests := []struct {
		name    string
		args    args
		want    *sqliteStorageImpl
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createSqliteStorage(tt.args.dbFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("createSqliteStorage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createSqliteStorage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqliteStorageImpl_StoreSession(t *testing.T) {
	tests := []struct {
		name    string
		want    testSessionDto
		wantErr bool
	}{
		{
			name: "first case",
			want: testSessionDto{
				Token:     "74ry4378tfhurh78rty784",
				UserId:    3534,
				ExpiredAt: 16987694,
			},
		},
	}

	dbFile := "test_db.sqlite"
	conn, err := connect(dbFile)
	if err != nil {
		t.Error(err)
		return
	}
	sqliteStorage := &sqliteStorageImpl{conn: conn, dbFile: dbFile}

	driver, err := sqlite3.WithInstance(conn, &sqlite3.Config{})
	if err != nil {
		t.Error(err)
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://../migrations/sqlite",
		"ql",
		driver,
	)
	if err != nil {
		t.Error(err)
		return
	}

	for _, tt := range tests {
		err = m.Up()
		if err != nil {
			t.Error(err)
			return
		}

		t.Run(tt.name, func(t *testing.T) {
			sess := dto.NewSession(tt.want.Token, tt.want.UserId, tt.want.ExpiredAt)

			if err := sqliteStorage.StoreSession(sess); (err != nil) != tt.wantErr {
				t.Errorf("StoreSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			row := conn.QueryRow("SELECT user_id, expired_at FROM sessions WHERE token = ?", tt.want.Token)
			var userId, expiredAt int64
			err := row.Scan(&userId, &expiredAt)
			if err != nil {
				t.Error(err)
				return
			}

			if tt.want.UserId != userId || tt.want.ExpiredAt != expiredAt {
				t.Error("Stored data is incorrect")
				return
			}
		})

		err = m.Down()
		if err != nil {
			t.Error(err)
			return
		}
	}
}

func Test_sqliteStorageImpl_GetSession(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    testSessionDto
		wantErr bool
	}{
		{
			name: "first case",
			args: args{
				token: "f49574895f48f45f4f4545",
			},
			want: testSessionDto{
				Token:     "f49574895f48f45f4f4545",
				UserId:    1234,
				ExpiredAt: 23435345,
			},
			wantErr: false,
		},
	}

	dbFile := "test_db.sqlite"
	conn, err := connect(dbFile)
	if err != nil {
		t.Error(err)
		return
	}
	sqliteStorage := &sqliteStorageImpl{conn: conn, dbFile: dbFile}

	driver, err := sqlite3.WithInstance(conn, &sqlite3.Config{})
	if err != nil {
		t.Error(err)
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://../migrations/sqlite",
		"ql",
		driver,
	)
	if err != nil {
		t.Error(err)
		return
	}

	for _, tt := range tests {
		err = m.Up()
		if err != nil {
			t.Error(err)
			return
		}

		insertStmt, err := conn.Prepare("INSERT INTO sessions (token, user_id, expired_at) VALUES (?, ?, ?)")
		if err != nil {
			t.Error(err)
			return
		}

		_, err = insertStmt.Exec(tt.want.Token, tt.want.UserId, tt.want.ExpiredAt)
		if err != nil {
			t.Error(err)
			return
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := sqliteStorage.GetSession(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.GetToken() != tt.want.Token {
				t.Errorf("GetSession().GetToken() got = %s, want %s", got.GetToken(), tt.want.Token)
			}
			if got.GetUserId() != tt.want.UserId {
				t.Errorf("GetSession().GetUserId() got = %d, want %d", got.GetUserId(), tt.want.UserId)
			}
			if got.GetExpiredAt() != tt.want.ExpiredAt {
				t.Errorf("GetSession().GetExpiredAt() got = %d, want %d", got.GetExpiredAt(), tt.want.ExpiredAt)
			}
		})

		err = m.Down()
		if err != nil {
			t.Error(err)
			return
		}
	}
}

func Test_sqliteStorageImpl_StoreUser(t *testing.T) {
	tests := []struct {
		name    string
		want    testUserDto
		wantErr bool
	}{
		{
			name: "first case",
			want: testUserDto{
				Name:     "5895f74895f45f",
				BancAcc:  "5fk8490385kf490358",
				Address:  "7c87j89457f4895k9058",
				Balance:  12564,
				PswdHash: "75j498574d3475d3w",
			},
		},
	}

	dbFile := "test_db.sqlite"
	conn, err := connect(dbFile)
	if err != nil {
		t.Error(err)
		return
	}
	sqliteStorage := &sqliteStorageImpl{conn: conn, dbFile: dbFile}

	driver, err := sqlite3.WithInstance(conn, &sqlite3.Config{})
	if err != nil {
		t.Error(err)
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://../migrations/sqlite",
		"ql",
		driver,
	)
	if err != nil {
		t.Error(err)
		return
	}

	for _, tt := range tests {
		err = m.Up()
		if err != nil {
			t.Error(err)
			return
		}

		t.Run(tt.name, func(t *testing.T) {
			u := tt.want
			user := dto.NewUser(0, u.Name, u.BancAcc, u.Address, u.Balance, u.PswdHash)

			var id int64
			if id, err = sqliteStorage.StoreUser(user); (err != nil) != tt.wantErr {
				t.Errorf("StoreUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			row := conn.QueryRow("SELECT name, passwd, banc_acc, address, balance FROM users WHERE id = ?", id)
			var name, passwd, bancAcc, address string
			var balance int64
			err = row.Scan(&name, &passwd, &bancAcc, &address, &balance)
			if err != nil {
				t.Error()
				return
			}

			if u.Name != name {
				t.Errorf("Name is incorrect. Want - %s, got  - %s\n", u.Name, name)
			}
			if u.BancAcc != bancAcc {
				t.Errorf("BancAcc is incorrect. Want - %s, got  - %s\n", u.BancAcc, bancAcc)
			}
			if u.Address != address {
				t.Errorf("Address is incorrect. Want - %s, got  - %s\n", u.Address, address)
			}
			if u.Balance != balance {
				t.Errorf("Balance is incorrect. Want - %d, got  - %d\n", u.Balance, balance)
			}
			if u.PswdHash != passwd {
				t.Errorf("PswdHash is incorrect. Want - %s, got  - %s\n", u.PswdHash, passwd)
			}
		})

		err = m.Down()
		if err != nil {
			t.Error(err)
			return
		}
	}
}

func Test_sqliteStorageImpl_GetUserById(t *testing.T) {
	tests := []struct {
		name    string
		want    testUserDto
		wantErr bool
	}{
		{
			name: "first case",
			want: testUserDto{
				Name:     "89f384f389f34f3",
				BancAcc:  "9834j9834fj3849fj39",
				Address:  "8924j8934fj398fij3489f",
				Balance:  12434,
				PswdHash: "sd8ud89udf89fj8f9djf89sd",
			},
			wantErr: false,
		},
	}

	dbFile := "test_db.sqlite"
	conn, err := connect(dbFile)
	if err != nil {
		t.Error(err)
		return
	}
	sqliteStorage := &sqliteStorageImpl{conn: conn, dbFile: dbFile}

	driver, err := sqlite3.WithInstance(conn, &sqlite3.Config{})
	if err != nil {
		t.Error(err)
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://../migrations/sqlite",
		"ql",
		driver,
	)
	if err != nil {
		t.Error(err)
		return
	}

	for _, tt := range tests {
		err = m.Up()
		if err != nil {
			t.Error(err)
			return
		}

		stmt, err := conn.Prepare(
			"INSERT INTO users (name, passwd, banc_acc, address, balance) VALUES (?, ?, ?, ?, ?)",
		)
		if err != nil {
			t.Error(err)
			return
		}

		u := tt.want
		_, err = stmt.Exec(u.Name, u.PswdHash, u.BancAcc, u.Address, u.Balance)
		if err != nil {
			t.Error(err)
			return
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := sqliteStorage.GetUserById(1)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.GetName() != u.Name {
				t.Errorf("GetUserById().GetName() got = %s, want %s", got.GetName(), u.Name)
			}
			if got.GetPswdHash() != u.PswdHash {
				t.Errorf("GetUserById().GetPswdHash() got = %s, want %s", got.GetPswdHash(), u.PswdHash)
			}
			if got.GetBankAcc() != u.BancAcc {
				t.Errorf("GetUserById().GetBankAcc() got = %s, want %s", got.GetBankAcc(), u.BancAcc)
			}
			if got.GetAddress() != u.Address {
				t.Errorf("GetUserById().GetAddress() got = %s, want %s", got.GetAddress(), u.Address)
			}
			if got.GetBalance() != u.Balance {
				t.Errorf("GetUserById().GetBalance() got = %d, want %d", got.GetBalance(), u.Balance)
			}
		})

		err = m.Down()
		if err != nil {
			t.Error(err)
			return
		}
	}
}

func Test_sqliteStorageImpl_GetUserByName(t *testing.T) {
	tests := []struct {
		name    string
		want    testUserDto
		wantErr bool
	}{
		{
			name: "first case",
			want: testUserDto{
				Name:     "89f384f389f34f3",
				BancAcc:  "9834j9834fj3849fj39",
				Address:  "8924j8934fj398fij3489f",
				Balance:  12434,
				PswdHash: "sd8ud89udf89fj8f9djf89sd",
			},
			wantErr: false,
		},
	}

	dbFile := "test_db.sqlite"
	conn, err := connect(dbFile)
	if err != nil {
		t.Error(err)
		return
	}
	sqliteStorage := &sqliteStorageImpl{conn: conn, dbFile: dbFile}

	driver, err := sqlite3.WithInstance(conn, &sqlite3.Config{})
	if err != nil {
		t.Error(err)
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://../migrations/sqlite",
		"ql",
		driver,
	)
	if err != nil {
		t.Error(err)
		return
	}

	for _, tt := range tests {
		err = m.Up()
		if err != nil {
			t.Error(err)
			return
		}

		stmt, err := conn.Prepare(
			"INSERT INTO users (name, passwd, banc_acc, address, balance) VALUES (?, ?, ?, ?, ?)",
		)
		if err != nil {
			t.Error(err)
			return
		}

		u := tt.want
		_, err = stmt.Exec(u.Name, u.PswdHash, u.BancAcc, u.Address, u.Balance)
		if err != nil {
			t.Error(err)
			return
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := sqliteStorage.GetUserByName(u.Name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.GetName() != u.Name {
				t.Errorf("GetUserById().GetName() got = %s, want %s", got.GetName(), u.Name)
			}
			if got.GetPswdHash() != u.PswdHash {
				t.Errorf("GetUserById().GetPswdHash() got = %s, want %s", got.GetPswdHash(), u.PswdHash)
			}
			if got.GetBankAcc() != u.BancAcc {
				t.Errorf("GetUserById().GetBankAcc() got = %s, want %s", got.GetBankAcc(), u.BancAcc)
			}
			if got.GetAddress() != u.Address {
				t.Errorf("GetUserById().GetAddress() got = %s, want %s", got.GetAddress(), u.Address)
			}
			if got.GetBalance() != u.Balance {
				t.Errorf("GetUserById().GetBalance() got = %d, want %d", got.GetBalance(), u.Balance)
			}
		})

		err = m.Down()
		if err != nil {
			t.Error(err)
			return
		}
	}
}

func Test_sqliteStorageImpl_StoreUserMoneyReward(t *testing.T) {
	tests := []struct {
		name    string
		want    testUserMoneyDto
		wantErr bool
	}{
		{
			name: "first case",
			want: testUserMoneyDto{
				Id:     0,
				UserId: 1245,
				Amount: 143532,
				Sent:   false,
			},
		},
	}

	dbFile := "test_db.sqlite"
	conn, err := connect(dbFile)
	if err != nil {
		t.Error(err)
		return
	}
	sqliteStorage := &sqliteStorageImpl{conn: conn, dbFile: dbFile}

	driver, err := sqlite3.WithInstance(conn, &sqlite3.Config{})
	if err != nil {
		t.Error(err)
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://../migrations/sqlite",
		"ql",
		driver,
	)
	if err != nil {
		t.Error(err)
		return
	}

	for _, tt := range tests {
		err = m.Up()
		if err != nil {
			t.Error(err)
			return
		}

		t.Run(tt.name, func(t *testing.T) {
			reward := dto.NewMoneyReward(tt.want.UserId, tt.want.Amount, tt.want.Sent, tt.want.Id)
			if _, err = sqliteStorage.StoreUserMoneyReward(reward); (err != nil) != tt.wantErr {
				t.Errorf("StoreSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			row := conn.QueryRow("SELECT user_id, amount, sent FROM money_rewards WHERE id = ?", 1)
			var userId, amount int64
			var sent bool
			err = row.Scan(&userId, &amount, &sent)
			if err != nil {
				t.Error(err)
				return
			}

			if tt.want.UserId != userId {
				t.Errorf("UserId is incorrect. Want - %d, got  - %d\n", tt.want.UserId, userId)
			}
			if tt.want.Amount != amount {
				t.Errorf("Amount is incorrect. Want - %d, got  - %d\n", tt.want.Amount, amount)
			}
			if tt.want.Sent != sent {
				t.Errorf("Sent is incorrect. Want - %v, got  - %v\n", tt.want.Sent, sent)
			}
		})

		err = m.Down()
		if err != nil {
			t.Error(err)
			return
		}
	}
}

// #######################################

// #######################################

func Test_sqliteStorageImpl_GetUnprocessedItemsRewards(t *testing.T) {
	type fields struct {
		dbFile string
		conn   *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []dto.ItemReward
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqliteStorageImpl{
				dbFile: tt.fields.dbFile,
				conn:   tt.fields.conn,
			}
			got, err := s.GetUnprocessedItemsRewards()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUnprocessedItemsRewards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUnprocessedItemsRewards() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqliteStorageImpl_GetUnprocessedMoneyRewards(t *testing.T) {
	type fields struct {
		dbFile string
		conn   *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []dto.MoneyReward
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqliteStorageImpl{
				dbFile: tt.fields.dbFile,
				conn:   tt.fields.conn,
			}
			got, err := s.GetUnprocessedMoneyRewards()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUnprocessedMoneyRewards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUnprocessedMoneyRewards() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqliteStorageImpl_GetUserItemRewards(t *testing.T) {
	type fields struct {
		dbFile string
		conn   *sql.DB
	}
	type args struct {
		userId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqliteStorageImpl{
				dbFile: tt.fields.dbFile,
				conn:   tt.fields.conn,
			}
			got, err := s.GetUserItemRewards(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserItemRewards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUserItemRewards() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqliteStorageImpl_GetUserMoneyRewards(t *testing.T) {
	type fields struct {
		dbFile string
		conn   *sql.DB
	}
	type args struct {
		userId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqliteStorageImpl{
				dbFile: tt.fields.dbFile,
				conn:   tt.fields.conn,
			}
			got, err := s.GetUserMoneyRewards(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserMoneyRewards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUserMoneyRewards() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqliteStorageImpl_SetItemsRewardsProcessed(t *testing.T) {
	type fields struct {
		dbFile string
		conn   *sql.DB
	}
	type args struct {
		ids []int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqliteStorageImpl{
				dbFile: tt.fields.dbFile,
				conn:   tt.fields.conn,
			}
			if err := s.SetItemsRewardsProcessed(tt.args.ids); (err != nil) != tt.wantErr {
				t.Errorf("SetItemsRewardsProcessed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_sqliteStorageImpl_SetMoneyRewardsProcessed(t *testing.T) {
	type fields struct {
		dbFile string
		conn   *sql.DB
	}
	type args struct {
		ids []int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqliteStorageImpl{
				dbFile: tt.fields.dbFile,
				conn:   tt.fields.conn,
			}
			if err := s.SetMoneyRewardsProcessed(tt.args.ids); (err != nil) != tt.wantErr {
				t.Errorf("SetMoneyRewardsProcessed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_sqliteStorageImpl_StoreUserItemReward(t *testing.T) {
	type fields struct {
		dbFile string
		conn   *sql.DB
	}
	type args struct {
		item dto.ItemReward
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqliteStorageImpl{
				dbFile: tt.fields.dbFile,
				conn:   tt.fields.conn,
			}
			if _, err := s.StoreUserItemReward(tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("StoreUserItemReward() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
