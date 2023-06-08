package account

import (
	"reflect"
	"testing"
)

func TestNewAccountUsecase(t *testing.T) {
	type args struct {
		repo AccountRepoInterface
	}
	tests := []struct {
		name string
		args args
		want AccountUsecaseInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccountUsecase(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccountUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accountUsecase_Create(t *testing.T) {
	type fields struct {
		repo AccountRepoInterface
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
			u := accountUsecase{
				repo: tt.fields.repo,
			}
			if err := u.Create(tt.args.account); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_accountUsecase_getByUsername(t *testing.T) {
	type fields struct {
		repo AccountRepoInterface
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
			u := accountUsecase{
				repo: tt.fields.repo,
			}
			got, err := u.getByUsername(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("getByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getByUsername() got = %v, want %v", got, tt.want)
			}
		})
	}
}
