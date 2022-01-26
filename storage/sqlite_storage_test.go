package storage

import (
	"database/sql"
	"github.com/yaroslavklimuk/crazy-lottery/dto"
	"reflect"
	"testing"
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

func Test_sqliteStorageImpl_GetSession(t *testing.T) {
	type fields struct {
		dbFile string
		conn   *sql.DB
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.Session
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
			got, err := s.GetSession(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSession() got = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func Test_sqliteStorageImpl_GetUserById(t *testing.T) {
	type fields struct {
		dbFile string
		conn   *sql.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.User
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
			got, err := s.GetUserById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqliteStorageImpl_GetUserByName(t *testing.T) {
	type fields struct {
		dbFile string
		conn   *sql.DB
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.User
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
			got, err := s.GetUserByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByName() got = %v, want %v", got, tt.want)
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

func Test_sqliteStorageImpl_StoreSession(t *testing.T) {
	type fields struct {
		dbFile string
		conn   *sql.DB
	}
	type args struct {
		sess dto.Session
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
			if err := s.StoreSession(tt.args.sess); (err != nil) != tt.wantErr {
				t.Errorf("StoreSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_sqliteStorageImpl_StoreUser(t *testing.T) {
	type fields struct {
		dbFile string
		conn   *sql.DB
	}
	type args struct {
		user dto.User
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
			got, err := s.StoreUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("StoreUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StoreUser() got = %v, want %v", got, tt.want)
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
		base dto.Reward
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
			if err := s.StoreUserItemReward(tt.args.base, tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("StoreUserItemReward() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_sqliteStorageImpl_StoreUserMoneyReward(t *testing.T) {
	type fields struct {
		dbFile string
		conn   *sql.DB
	}
	type args struct {
		base  dto.Reward
		money dto.MoneyReward
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
			if err := s.StoreUserMoneyReward(tt.args.base, tt.args.money); (err != nil) != tt.wantErr {
				t.Errorf("StoreUserMoneyReward() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
