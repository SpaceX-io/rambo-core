package grpclient

import (
	"google.golang.org/grpc/keepalive"
	"time"
)

var defaultKeepAlive = &keepalive.ClientParameters{
	Time:                10 * time.Second,
	Timeout:             time.Second,
	PermitWithoutStream: true,
}
