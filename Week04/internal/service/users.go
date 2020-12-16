package service

import (
	"context"
	"fmt"
	pb "github.com/neverhover/Go-001/tree/main/Week04/api/users"
	"github.com/neverhover/Go-001/tree/main/Week04/internal/data"
	"github.com/neverhover/Go-001/tree/main/Week04/internal/pkg/service"
)

type UsersService struct {
	pb.UnimplementedUsersServiceServer

	srv service.Service
}

func NewUserService(srv service.Service) *UsersService {
	return &UsersService{
		srv: srv,
	}
}

func (u UsersService) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	userObj := in.GetUser()
	u.srv.Logger().Info(fmt.Sprintf("Create user %+v", userObj))
	dbUsers := new(data.Users)
	dbUsers.ID = userObj.GetId()
	dbUsers.Password = userObj.GetPassword()
	dbUsers.DialString = userObj.GetDialString()
	dbUsers.Domain = userObj.GetDomain()
	dbUsers.Mailbox = userObj.GetMailbox()
	dbUsers.NumberAlias = userObj.GetNumberAlias()
	dbUsers.UserContext = userObj.GetUserContext()
	dbUsers.Create(u.srv.DB())
	return &pb.CreateUserResponse{User: in.User}, nil
}
