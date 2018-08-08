package lendBook

import (
	"context"
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
	testutil "github.com/luquehuong/example-go/config/database/pg/util"
	"github.com/luquehuong/example-go/domain"
)

func Test_pgService_Create(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	book := domain.Book{}
	user := domain.User{}
	errBook := testDB.Create(&book).Error
	errUser := testDB.Create(&user).Error
	if errBook != nil {
		t.Fatalf("Failed to create dummy table lendBook by error %v", errBook)
	}
	if errUser != nil {
		t.Fatalf("Failed to create dummy table lendBook by error %v", errUser)
	}

	fakeBookID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")
	fakeUserID := domain.MustGetUUIDFromString("1a98bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		p *domain.LendBook
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				&domain.LendBook{
					Book_id: book.ID,
					User_id: user.ID,
				},
			},
		},
		{
			name: "Book_id is not correct",
			args: args{
				&domain.LendBook{
					Book_id: fakeBookID,
					User_id: user.ID,
				},
			},
			wantErr: true,
		},
		{
			name: "User_id is not correct",
			args: args{
				&domain.LendBook{
					Book_id: book.ID,
					User_id: fakeUserID,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}
			if err := s.Create(context.Background(), tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("pgService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pgService_FindAll(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		in0 context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.LendBook
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: tt.fields.db,
			}
			got, err := s.FindAll(tt.args.in0)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgService.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pgService.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pgService_Delete(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		in0 context.Context
		p   *domain.LendBook
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
			s := &pgService{
				db: tt.fields.db,
			}
			if err := s.Delete(tt.args.in0, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("pgService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
