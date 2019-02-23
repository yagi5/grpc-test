package main

import (
	"context"
	"fmt"

	userpb "github.com/yagi5/grpc-test/pb"
	"google.golang.org/grpc"
)

type usersClient struct {
	client userpb.UsersClient
}

type user struct {
	ID    string
	Name  string
	Email string
}

func main() {
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := userpb.NewUsersClient(conn)

	c := usersClient{client: client}
	u, err := c.getUser(context.Background(), "1")
	if err != nil {
		panic(err)
	}

	fmt.Printf("user: %#v\n", u)
}

func (c *usersClient) getUser(ctx context.Context, id string) (*user, error) {
	res, err := c.client.Get(ctx, &userpb.GetRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &user{ID: res.User.Id, Name: res.User.Name, Email: res.User.Email}, nil
}
