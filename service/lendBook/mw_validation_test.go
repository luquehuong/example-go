package lendBook

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/luquehuong/example-go/domain"
)

func Test_validationMiddleware_Create(t *testing.T) {
	serviceMock := &ServiceMock{
		CreateFunc: func(_ context.Context, p *domain.LendBook) error {
			return nil
		},
	}

	defaultCtx := context.Background()
	type args struct {
		p *domain.LendBook
	}

	bID := domain.MustGetUUIDFromString("92d16277-8796-4143-a9a2-253064a3af2a")
	uID := domain.MustGetUUIDFromString("b2ec8e9c-207e-4726-a4fa-e04d29df4bde")
	var timeLend time.Time
	var timeRuturn1 time.Time
	var timeRuturn2 time.Time
	json.Unmarshal([]byte("{2017-01-03T15:39:04.332168372+07:00}"), &timeLend)
	json.Unmarshal([]byte("{2018-07-03T15:39:04.332168372+07:00}"), &timeRuturn1)
	json.Unmarshal([]byte("{2019-07-03T15:39:04.332168372+07:00}"), &timeRuturn2)

	tests := []struct {
		name            string
		args            args
		wantErr         bool
		wantOutput      *domain.LendBook
		errorStatusCode int
	}{
		{
			name: "valid book",
			args: args{&domain.LendBook{
				Book_id: bID,
				User_id: uID,
			}},
			wantOutput: &domain.LendBook{
				Book_id: bID,
				User_id: uID,
				From:    timeLend,
				To:      timeRuturn1,
			},
		},
		{
			name: "invalid book by missing Book_id",
			args: args{&domain.LendBook{
				User_id: uID,
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by missing User_id",
			args: args{&domain.LendBook{
				Book_id: bID,
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "unavailable book",
			args: args{&domain.LendBook{
				User_id: uID,
				Book_id: bID,
				From:    timeLend,
				To:      timeRuturn2,
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: serviceMock,
			}
			err := mw.Create(defaultCtx, tt.args.p)
			if err != nil {
				if tt.wantErr == false {
					t.Errorf("validationMiddleware.Create() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				status, ok := err.(interface{ StatusCode() int })
				if !ok {
					t.Errorf("validationMiddleware.Create() error %v doesn't implement StatusCode()", err)
				}
				if tt.errorStatusCode != status.StatusCode() {
					t.Errorf("validationMiddleware.Create() status = %v, want status code %v", status.StatusCode(), tt.errorStatusCode)
					return
				}
				return
			}
		})
	}
}

func Test_validationMiddleware_FindAll(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		ctx context.Context
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
			mw := validationMiddleware{
				Service: tt.fields.Service,
			}
			got, err := mw.FindAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("validationMiddleware.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validationMiddleware.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validationMiddleware_Find(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		ctx      context.Context
		lendBook *domain.LendBook
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.LendBook
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: tt.fields.Service,
			}
			got, err := mw.Find(tt.args.ctx, tt.args.lendBook)
			if (err != nil) != tt.wantErr {
				t.Errorf("validationMiddleware.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validationMiddleware.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validationMiddleware_Delete(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		ctx      context.Context
		lendBook *domain.LendBook
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
			mw := validationMiddleware{
				Service: tt.fields.Service,
			}
			if err := mw.Delete(tt.args.ctx, tt.args.lendBook); (err != nil) != tt.wantErr {
				t.Errorf("validationMiddleware.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
