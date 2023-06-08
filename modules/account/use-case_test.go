package account

import (
	"reflect"
	"testing"
)

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
		{
			name: "Testing create aktor",
			fields: fields{
				repo: NewMockAccountRepoInterface(t),
			},
			args: args{
				account: &Actors{Username: "test", Password: "test", Role_id: 1},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := accountUsecase{
				repo: tt.fields.repo,
			}
			tt.fields.repo.(*MockAccountRepoInterface).EXPECT().Save(tt.args.account).Return(nil)
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
		{
			name: "GetBy Username",
			fields: fields{
				repo: NewMockAccountRepoInterface(t),
			},
			args: args{
				username: "test",
			},
			want: Actors{
				Username: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := accountUsecase{
				repo: tt.fields.repo,
			}
			tt.fields.repo.(*MockAccountRepoInterface).EXPECT().FindByUsername(tt.args.username).Return(tt.want, nil)
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
