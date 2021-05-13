package grpclient

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"time"
)

var defaultDialTimeout = time.Second * 2

type Option func(*option)

type option struct {
	credential       credentials.TransportCredentials
	keepalive        *keepalive.ClientParameters
	dialTimeout      time.Duration
	unaryInterceptor grpc.UnaryClientInterceptor
}

func WithCredential(credential credentials.TransportCredentials) Option {
	return func(o *option) {
		o.credential = credential
	}
}

func WithKeepAlive(keepalive *keepalive.ClientParameters) Option {
	return func(o *option) {
		o.keepalive = keepalive
	}
}

func WithDialTimeout(timeout time.Duration) Option {
	return func(o *option) {
		o.dialTimeout = timeout
	}
}

func WithUnaryInterceptor(unaryInterceptor grpc.UnaryClientInterceptor) Option {
	return func(o *option) {
		o.unaryInterceptor = unaryInterceptor
	}
}

func New(target string, options ...Option) (*grpc.ClientConn, error) {
	if target == "" {
		return nil, errors.New("target required")
	}

	opt := new(option)
	for _, f := range options {
		f(opt)
	}

	kp := defaultKeepAlive
	if opt.keepalive != nil {
		kp = opt.keepalive
	}

	dialTimeout := defaultDialTimeout
	if opt.dialTimeout > 0 {
		dialTimeout = opt.dialTimeout
	}

	dialOptions := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(*kp),
	}

	if opt.unaryInterceptor != nil {
		dialOptions = append(dialOptions, grpc.WithUnaryInterceptor(opt.unaryInterceptor))
	}

	if opt.credential == nil {
		dialOptions = append(dialOptions, grpc.WithInsecure())
	} else {
		dialOptions = append(dialOptions, grpc.WithTransportCredentials(opt.credential))
	}

	ctx, cancel := context.WithTimeout(context.Background(), dialTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, target, dialOptions...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return conn, nil
}
