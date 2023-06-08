package account

import (
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestNewAccountRepository(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *accountRepo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccountRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccountRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accountRepo_FindByUsername(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Actors
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := accountRepo{
				db: tt.fields.db,
			}
			got, err := r.FindByUsername(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByUsername() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accountRepo_Save(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		account *Actors
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
			r := accountRepo{
				db: tt.fields.db,
			}
			if err := r.Save(tt.args.account); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
