package constructor_test

import (
	"go_template/domain/constructor"
	"go_template/domain/entity"
	"go_template/domain/entity_const"
	"go_template/domain/validations"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewUser(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args constructor.NewUserCreateArgs
		want entity.User
		wantErr error
	}{
		{
			name: "valid user",
			args: constructor.NewUserCreateArgs{
				UserID: "user_id",
				UserDetailID: "user_detail_id",
				Email: "user1@example.com",
				Password: "password",
				HashedPassword: "hashed_password",
				UserType: "user",
				Name: "Test User",
			},
			want: entity.User{
				UserID: "user_id",
				Email: "user1@example.com",
				UserType: entity_const.User,
				HashedPassword: "hashed_password",
				UserDetail: &entity.UserDetail{
					UserDetailID: "user_detail_id",
					UserID: "user_id",
					Name: "Test User",
				},
			},
			wantErr: nil,
		},
		{
			name: "valid admin",
			args: constructor.NewUserCreateArgs{
				UserID: "user_id",
				UserDetailID: "user_detail_id",
				Email: "admin@example.com",
				Password: "password",
				HashedPassword: "hashed_password",
				UserType: "admin",
				Name: "Test Admin",
			},
			want: entity.User{
				UserID: "user_id",
				Email: "admin@example.com",
				UserType: entity_const.Admin,
				HashedPassword: "hashed_password",
				UserDetail: &entity.UserDetail{
					UserDetailID: "user_detail_id",
					UserID: "user_id",
					Name: "Test Admin",
				},
			},
			wantErr: nil,
		},
		{
			name: "invalid email",
			args: constructor.NewUserCreateArgs{
				UserID: "user_id",
				UserDetailID: "user_detail_id",
				Email: "invalid_email",
				Password: "password",
				HashedPassword: "hashed_password",
				UserType: "user",
				Name: "Test User",
			},
			want: entity.User{},
			wantErr: validations.ErrInvalidEmail,
		},
		{
			name: "invalid password",
			args: constructor.NewUserCreateArgs{
				UserID: "user_id",
				UserDetailID: "user_detail_id",
				Email: "user@example.com",
				Password: "pass",
				HashedPassword: "hashed_password",
				UserType: "user",
				Name: "Test User",
			},
			want: entity.User{},
			wantErr: validations.ErrInvalidPassword,
		},
		{
			name: "invalid user type",
			args: constructor.NewUserCreateArgs{
				UserID: "user_id",
				UserDetailID: "user_detail_id",
				Email: "user@example.com",
				Password: "password",
				HashedPassword: "hashed_password",
				UserType: "invalid_user_type",
				Name: "Test User",
			},
			want: entity.User{},
			wantErr: entity_const.ErrInvalidUserType,
		},
		{
			name: "invalid user detail id",
			args: constructor.NewUserCreateArgs{
				UserID: "user_id",
				UserDetailID: "",
				Email: "user@example.com",
				Password: "password",
				HashedPassword: "hashed_password",
				UserType: "user",
				Name: "Test User",
			},
			want: entity.User{},
			wantErr: validations.ErrInvalidUserDetailID,
		},
		{
			name: "invalid user id",
			args: constructor.NewUserCreateArgs{
				UserID: "",
				UserDetailID: "user_detail_id",
				Email: "user@example.com",
				Password: "password",
				HashedPassword: "hashed_password",
				UserType: "user",
				Name: "Test User",
			},
			want: entity.User{},
			wantErr: validations.ErrInvalidUserID,
		},
		{
			name: "invalid name",
			args: constructor.NewUserCreateArgs{
				UserID: "user_id",
				UserDetailID: "user_detail_id",
				Email: "user@example.com",
				Password: "password",
				HashedPassword: "hashed_password",
				UserType: "user",
				Name: "",
			},
			want: entity.User{},
			wantErr: validations.ErrInvalidName,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := constructor.NewUserCreate(tt.args)
			if diff := cmp.Diff(tt.wantErr, err, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("NewUserCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("NewUserCreate() got = %v, want %v", got, tt.want)
			}
		})
	}
}