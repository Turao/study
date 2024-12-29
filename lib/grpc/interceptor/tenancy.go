package interceptor

import (
	"context"
	"log"

	"github.com/turao/topics/metadata"
	"google.golang.org/grpc"
	grpcmetadata "google.golang.org/grpc/metadata"
)

func WithTenancyInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		key := HeaderTenancy.String()
		m, ok := grpcmetadata.FromIncomingContext(ctx)
		if !ok {
			ctx = context.WithValue(ctx, HeaderTenancy, metadata.TenancyTesting)
			return handler(ctx, req)
		}

		values := m.Get(key)
		if len(values) == 0 {
			ctx = context.WithValue(ctx, HeaderTenancy, metadata.TenancyTesting)
			return handler(ctx, req)
		}

		tenancy, err := metadata.NewTenancy(values[0])
		if err != nil {
			log.Println(err)
		}

		ctx = context.WithValue(ctx, HeaderTenancy, tenancy)
		return handler(ctx, req)
	}
}
