package openpds

import (
	"context"
	"errors"
)

func (c Client) ListProviders(ctx context.Context, input *CreateProviderInput) (*CreateProviderOutput, error) {
	return nil, errors.New("not implemented")
}

func (c Client) CreateProvider(ctx context.Context, input *CreateProviderInput) (*CreateProviderOutput, error) {
	return nil, errors.New("not implemented")
}

func (c Client) UpdateProvider(ctx context.Context, input *CreateProviderInput) (*CreateProviderOutput, error) {
	return nil, errors.New("not implemented")
}

func (c Client) GetProvider(ctx context.Context, input *CreateProviderInput) (*CreateProviderOutput, error) {
	return nil, errors.New("not implemented")
}

func (c Client) DeleteProvider(ctx context.Context, input *CreateProviderInput) (*CreateProviderOutput, error) {
	return nil, errors.New("not implemented")
}

type CreateProviderOutput struct{}

type CreateProviderInput struct{}
