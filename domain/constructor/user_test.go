package constructor_test

import (
	"go_template/domain/constructor"
	"go_template/domain/entity"
	"go_template/domain/entity_const"
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
				Email: "email",
				HashedPassword: "hashed_password",
				UserType: "user",
			},
			want: entity.User{
				UserID: "user_id",
				Email: "email",
				HashedPassword: "hashed_password",
				UserType: entity_const.User,
			},
			wantErr: nil,
		},
		{
			name: "valid admin",
			args: constructor.NewUserCreateArgs{
				UserID: "user_id",
				Email: "email",
				HashedPassword: "hashed_password",
				UserType: "admin",
			},
			want: entity.User{
				UserID: "user_id",
				Email: "email",
				HashedPassword: "hashed_password",
				UserType: entity_const.Admin,
			},
			wantErr: nil,
		},
		{
			name: "invalid user id",
			args: constructor.NewUserCreateArgs{
				UserID: "",
				Email: "email",
				HashedPassword: "hashed_password",
				UserType: "user",
			},
			want: entity.User{},
			wantErr: entity_const.ErrInvalidUserID,
		},
		{
			name: "invalid email",
			args: constructor.NewUserCreateArgs{
				UserID: "user_id",
				Email: "",
				HashedPassword: "hashed_password",
				UserType: "user",
			},
			want: entity.User{},
			wantErr: entity_const.ErrInvalidEmail,
		},
		{
			name: "invalid hashed password",
			args: constructor.NewUserCreateArgs{
				UserID: "user_id",
				Email: "email",
				HashedPassword: "",
				UserType: "user",
			},
			want: entity.User{},
			wantErr: entity_const.ErrInvalidHashedPassword,
		},
		{
			name: "invalid user type",
			args: constructor.NewUserCreateArgs{
				UserID: "user_id",
				Email: "email",
				HashedPassword: "hashed_password",
				UserType: "invalid_user_type",
			},
			want: entity.User{},
			wantErr: entity_const.ErrInvalidUserType,
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