package openpds

import (
	"context"
	"errors"
)

func (c Client) CheckReceivingMethods(ctx context.Context, input *CreateReceivingMethodInput) (*CreateReceivingMethodOutput, error) {
	return nil, errors.New("not implemented")
}

func (c Client) FetchReceivingMethod(ctx context.Context, input *CreateReceivingMethodInput) (*CreateReceivingMethodOutput, error) {
	return nil, errors.New("not implemented")
}

type CreateReceivingMethodOutput struct{}

type CreateReceivingMethodInput struct{}
