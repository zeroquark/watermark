package watermark

import (
	"context"
)

type Service interface {
	Get(ctx context.Context, filters ...internal.Filter) ([]internal.Document, error)
	Status(ctx context.Context, ticketID string) (internal.Status, error)
	Watermark(ctx context.Context, ticketID string, mark string) (int, error)
	AddDocument(ctx context.Context, doc *internal.Document) (string, error)
	ServiceStatus(ctx context.Context) (int, error)
}
