package customers

import (
	"reflect"
	"testing"
)

/*
func TestNewUsecase(t *testing.T) {
	type args struct {
		repo RepositoryInterface
	}
	tests := []struct {
		name string
		args args
		want UsecaseInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUsecase(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}*/

func TestUsecase_Create(t *testing.T) {
	type fields struct {
		repo RepositoryInterface
	}
	type args struct {
		customers *Customers
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Testing Create Customer",
			fields: fields{
				repo: NewMockRepositoryInterface(t),
			},
			args: args{
				customers: &Customers{
					Firstname: "test",
					Lastname:  "good",
					Email: "test@gmail",
					Avatar: "test.jpg", 
				},
			},
			wantErr: false,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Usecase{
				repo: tt.fields.repo,
			}
			tt.fields.repo.(*MockRepositoryInterface).EXPECT().Save(tt.args.customers).Return(nil)
			if err := u.Create(tt.args.customers); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsecase_DeleteByID(t *testing.T) {
	type fields struct {
		repo RepositoryInterface
	}
	type args struct {
		customers *Customers
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
		name: "Testing Delete Customer",
		fields: fields{
			repo: NewMockRepositoryInterface(t),
		},
		args: args{
			customers: &Customers{
				Firstname: "test",
				Lastname:  "good",
				Email: "test@gmail",
				Avatar: "test.jpg",
			},
		},
		wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Usecase{
				repo: tt.fields.repo,
			}
			tt.fields.repo.(*MockRepositoryInterface).EXPECT().DeleteByID(tt.args.customers).Return(nil)
			if err := u.DeleteByID(tt.args.customers); (err != nil) != tt.wantErr {
				t.Errorf("DeleteByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsecase_FindAll(t *testing.T) {
	type fields struct {
		repo RepositoryInterface
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Customers
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Testing Find All Customer",
			fields: fields{
				repo: NewMockRepositoryInterface(t),
			},
			want: []Customers{
				{
					Firstname: "test",
					Lastname:  "good",
					Email: "test@gmail",
					Avatar: "test.jpg",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Usecase{
				repo: tt.fields.repo,
			}
			tt.fields.repo.(*MockRepositoryInterface).EXPECT().FindAll().Return(tt.want, nil)
			got, err := u.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_Read(t *testing.T) {
	type fields struct {
		repo RepositoryInterface
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Customers
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Testing Read Customer",
			fields: fields{
				repo: NewMockRepositoryInterface(t),
			},
			want: []Customers{
				{
					Firstname: "test",
					Lastname:  "good",
					Email: "test@gmail",
					Avatar: "test.jpg",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Usecase{
				repo: tt.fields.repo,
			}
			tt.fields.repo.(*MockRepositoryInterface).EXPECT().FindAll().Return(tt.want, nil)
			got, err := u.Read()
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//gagal dites
func TestUsecase_ReadByID(t *testing.T) {
	type fields struct {
		repo RepositoryInterface
	}
	type args struct {
		id any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Customers
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Testing Read By ID Customer",
			fields: fields{
				repo: NewMockRepositoryInterface(t),
			},
			args: args{
				id: 1,
			},
			want: Customers{
				Firstname: "test",
				Lastname:  "good",
				Email: "test@gmail",
				Avatar: "test.jpg",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Usecase{
				repo: tt.fields.repo,
			}
			//tt.fields.repo.(*MockRepositoryInterface).EXPECT().ReadByID(tt.args.id).Return(tt.want, nil)
			got, err := u.ReadByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//gagal dites
func TestUsecase_Save(t *testing.T) {
	type fields struct {
		repo RepositoryInterface
	}
	type args struct {
		customers *Customers
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
		name: "Testing Save Customer",
		fields: fields{
			repo: NewMockRepositoryInterface(t),
		},
		args: args{
			customers: &Customers{
				Firstname: "test",
				Lastname:  "good",
				Email: "test@gmail",
				Avatar: "test.jpg",
			},
		},
		wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Usecase{
				repo: tt.fields.repo,
			}
			tt.fields.repo.(*MockRepositoryInterface).EXPECT().Save(tt.args.customers).Return(nil)
			if err := u.Save(tt.args.customers); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

//gagal dites
func TestUsecase_UpdateByID(t *testing.T) {
	type fields struct {
		repo RepositoryInterface
	}
	type args struct {
		customers *Customers
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Testing Update Customer",
			fields: fields{
				repo: NewMockRepositoryInterface(t),
			},
			args: args{
				customers: &Customers{
					Firstname: "test",
					Lastname:  "good",
					Email: "test@gmail",
					Avatar: "test.jpg",
				},
			},
			wantErr: false,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Usecase{
				repo: tt.fields.repo,
			}
			tt.fields.repo.(*MockRepositoryInterface).EXPECT().UpdateByID(tt.args.customers).Return(nil)
			if err := u.UpdateByID(tt.args.customers); (err != nil) != tt.wantErr {
				t.Errorf("UpdateByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsecase_getByEmail(t *testing.T) {
	type fields struct {
		repo RepositoryInterface
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Customers
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Testing Get By Email Customer",
			fields: fields{
				repo: NewMockRepositoryInterface(t),
			},
			args: args{
				email: "test@gmail",
			},
			want: Customers{
				Firstname: "test",
				Lastname:  "good",
				Email: "test@gmail",
				Avatar: "test.jpg",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Usecase{
				repo: tt.fields.repo,
			}
			//tt.fields.repo.(*MockRepositoryInterface).EXPECT().getByEmail(tt.args.email).Return(tt.want, nil)
			got, err := u.getByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("getByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}
