// Code generated by protoc-gen-redact. DO NOT EDIT.
// source: examples/user/pb/user.proto

package pb

import (
	context "context"
	redact "github.com/arrakis-digital/protoc-gen-redact/redact"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ grpc.Server
	_ context.Context
	_ redact.Redactor
	_ codes.Code
	_ status.Status
	_ redact.FieldRules
	_ empty.Empty
)

// RegisterRedactedChatServer wraps the ChatServer with the redacted server and registers the service in GRPC
func RegisterRedactedChatServer(s grpc.ServiceRegistrar, srv ChatServer, bypass redact.Bypass) {
	RegisterChatServer(s, RedactedChatServer(srv, bypass))
}

func RedactedChatServer(srv ChatServer, bypass redact.Bypass) ChatServer {
	if bypass == nil {
		bypass = redact.Falsy
	}
	return &redactedChatServer{srv: srv, bypass: bypass}
}

type redactedChatServer struct {
	UnsafeChatServer
	srv    ChatServer
	bypass redact.Bypass
}

// AddUser is the redacted wrapper for the actual ChatServer.AddUser method
func (s *redactedChatServer) AddUser(ctx context.Context, in *User) (*User, error) {
	if s.bypass.CheckInternal(ctx) {
		return s.srv.AddUser(ctx, in)
	}
	return nil, status.Error(codes.PermissionDenied, `Permission Denied. Method: "ChatServer.AddUser" has been redacted`)
}

// GetUser is the redacted wrapper for the actual ChatServer.GetUser method
func (s *redactedChatServer) GetUser(ctx context.Context, in *GetUserRequest) (*User, error) {
	res, err := s.srv.GetUser(ctx, in)
	if !s.bypass.CheckInternal(ctx) {
		// Apply redaction to the response
		redact.Apply(res)
	}
	return res, err
}

// ListUsers is the redacted wrapper for the actual ChatServer.ListUsers method
func (s *redactedChatServer) ListUsers(ctx context.Context, in *empty.Empty) (*ListUsersResponse, error) {
	if s.bypass.CheckInternal(ctx) {
		return s.srv.ListUsers(ctx, in)
	}
	return nil, status.Error(codes.Unavailable, `ChatServer.ListUsers unavailable`)
}

// Redact method implementation for User
func (x *User) Redact() {
	if x == nil {
		return
	}

	// Safe field: Username

	// Redacting field: Password
	x.Password = "REDACTED"

	// Redacting field: Email
	x.Email = `r*d@ct*d`

	// Safe field: Name

	// Safe field: Home
}

// Redact method implementation for GetUserRequest
func (x *GetUserRequest) Redact() {
	if x == nil {
		return
	}

	// Safe field: Username
}

// Redact method implementation for ListUsersResponse
func (x *ListUsersResponse) Redact() {
	if x == nil {
		return
	}

	// Safe field: Users
}

// Redact method implementation for User_Location
func (x *User_Location) Redact() {
	if x == nil {
		return
	}

	// Safe field: Lat

	// Safe field: Lng
}
