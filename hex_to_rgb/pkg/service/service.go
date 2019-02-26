package service

import (
	"context"

	log "github.com/sirupsen/logrus"
	colors "gopkg.in/go-playground/colors.v1"
)

// HexToRgbService describes the service.
type HexToRgbService interface {
	// Add your methods here
	Transform(ctx context.Context, s string) (rs string, err error)
}

type basicHexToRgbService struct{}

func (b *basicHexToRgbService) Transform(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Trasform

	color, err := colors.ParseHEX("#" + s)
	log.Infoln("the color is", color, s)
	return color.ToRGB().String(), err
}

// NewBasicHexToRgbService returns a naive, stateless implementation of HexToRgbService.
func NewBasicHexToRgbService() HexToRgbService {
	return &basicHexToRgbService{}
}

// New returns a HexToRgbService with all of the expected middleware wired in.
func New(middleware []Middleware) HexToRgbService {
	var svc HexToRgbService = NewBasicHexToRgbService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
