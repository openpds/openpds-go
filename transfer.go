package openpds

import (
	"context"
	"errors"
)

func (c Client) ListTransfers(ctx context.Context, input *CreateTransferInput) (*CreateTransferOutput, error) {
	return nil, errors.New("not implemented")
}

func (c Client) CreateTransfer(ctx context.Context, input *CreateTransferInput) (*CreateTransferOutput, error) {
	return nil, errors.New("not implemented")
}

func (c Client) UpdateTransfer(ctx context.Context, input *CreateTransferInput) (*CreateTransferOutput, error) {
	return nil, errors.New("not implemented")
}

func (c Client) ConfirmTransfer(ctx context.Context, input *CreateTransferInput) (*CreateTransferOutput, error) {
	return nil, errors.New("not implemented")
}

func (c Client) CancelTransfer(ctx context.Context, input *CreateTransferInput) (*CreateTransferOutput, error) {
	return nil, errors.New("not implemented")
}

func (c Client) GetTransfer(ctx context.Context, input *CreateTransferInput) (*CreateTransferOutput, error) {
	return nil, errors.New("not implemented")
}

func (c Client) DeleteTransfer(ctx context.Context, input *CreateTransferInput) (*CreateTransferOutput, error) {
	return nil, errors.New("not implemented")
}

type CreateTransferOutput struct{}

type CreateTransferInput struct{}
