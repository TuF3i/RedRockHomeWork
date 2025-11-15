package services

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

func BuildServiceHFunc(ctx context.Context, c *app.RequestContext) *Service {
	return &Service{
		ctx: ctx,
		c:   c,
	}
}

type Service struct {
	ctx context.Context
	c   *app.RequestContext
}
