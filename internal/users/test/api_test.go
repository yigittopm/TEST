package usecase

import (
	"context"
	"testing"

	"github.com/yigittopm/test/database"

	"github.com/yigittopm/test/internal/users/dtos"
	usersRepository "github.com/yigittopm/test/internal/users/repository"
	usersUsecase "github.com/yigittopm/test/internal/users/usecase"
)

var usecase usersUsecase.Usecase

func init() {
	db, _ := database.Start()
	repo := usersRepository.New(db)
	usecase = usersUsecase.New(repo)
}

func TestRegister(t *testing.T) {
	tests := []struct {
		name string
		req  dtos.RegisterRequest
		want dtos.RegisterResponse
	}{
		{
			name: "TestRegister",
			req: dtos.RegisterRequest{
				Username: "yigittopm",
				Email:    "yigittopm@hotmail.com",
				Password: "password",
			},
			want: dtos.RegisterResponse{ID: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := usecase.Register(context.Background(), tt.req)
			if got.ID != tt.want.ID {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}

}
