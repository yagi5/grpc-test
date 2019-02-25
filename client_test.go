package main

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	userpb "github.com/yagi5/grpc-test/pb"
	"google.golang.org/grpc"
)

type mockUsersClient struct {
	MockGetFunc func(context.Context, *userpb.GetRequest, ...grpc.CallOption) (*userpb.GetResponse, error)
}

func (c *mockUsersClient) Get(ctx context.Context, req *userpb.GetRequest, opts ...grpc.CallOption) (*userpb.GetResponse, error) {
	return c.MockGetFunc(ctx, req, opts...)
}

func TestGetUser(t *testing.T) {
	var tests = []struct {
		name    string
		getMock func(context.Context, *userpb.GetRequest, ...grpc.CallOption) (*userpb.GetResponse, error)
		id      string
		wantErr bool
		res     *user
	}{
		{
			name: "returns error",
			getMock: func(context.Context, *userpb.GetRequest, ...grpc.CallOption) (*userpb.GetResponse, error) {
				return nil, fmt.Errorf("")
			},
			id:      "1",
			wantErr: true,
		},
		{
			name: "returns non-error",
			getMock: func(context.Context, *userpb.GetRequest, ...grpc.CallOption) (*userpb.GetResponse, error) {
				return &userpb.GetResponse{User: &userpb.User{Id: "1", Name: "bob", Email: "email@example.com"}}, nil
			},
			id:  "1",
			res: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			c := &mockUsersClient{MockGetFunc: tt.getMock}
			cl := usersClient{grpcClient: c}
			u, err := cl.getUser(context.Background(), tt.id)
			if (err != nil) != tt.wantErr {
				t.Fatalf("failed expected: %#v but got %#v", tt.wantErr, err)
			}
			if reflect.DeepEqual(tt.res, u) {
				t.Errorf("failed expected: %#v but got %#v", tt.res, u)
			}
		})
	}
}
