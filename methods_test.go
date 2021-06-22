package main

import (
	"context"
	"fmt"
	"testing"

	pb "github.com/jankrynauw/assa/protobufs"
)

// Simulate a client object
// without going through the gRPC connection
var client pb.MembersServer

// This init() function will only run when running Go tests.
func init() {
	client = &myService{}
}

func TestMyService_CreateMember(t *testing.T) {
	// Construct request
	req := pb.CreateMemberRequest{Member: &pb.Member{
		FamilyName:  "Krynauw",
		Description: "Some description",
	}}

	res, err := client.CreateMember(context.Background(), &req)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
}
