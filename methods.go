package main

import (
	"context"
	pb "github.com/jankrynauw/assa/protobufs"
)

// Create a Service object which we'll register with the Server
type myService struct {
	pb.UnimplementedMembersServer
}

func (s *myService) GetMember(ctx context.Context, req *pb.GetMemberRequest) (*pb.Member, error) {

	// Retrieve the member from the database
	// TODO: get member from database.
	member := pb.Member{
		Name:        req.GetName(),
		FamilyName:  "Example Family Name",
		Description: "Example description",
	}

	return &member, nil
}

func (s *myService) CreateMember(ctx context.Context, req *pb.CreateMemberRequest) (*pb.Member, error) {

	// Create a member in the database
	// TODO: create in database

	// Once created set the ID of the member
	req.GetMember().Name = "members/KRY001"

	return req.GetMember(), nil
}
