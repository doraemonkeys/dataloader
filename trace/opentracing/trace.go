package opentracing

import (
	"context"
	"fmt"

	"github.com/doraemonkeys/dataloader/v7"

	"github.com/opentracing/opentracing-go"
)

// Tracer implements a tracer that can be used with the Open Tracing standard.
type Tracer[K comparable, V any] struct{}

// TraceLoad will trace a call to dataloader.LoadMany with Open Tracing.
func (Tracer[K, V]) TraceLoad(ctx context.Context, key K) (context.Context, dataloader.TraceLoadFinishFunc[V]) {
	span, spanCtx := opentracing.StartSpanFromContext(ctx, "Dataloader: load")

	span.SetTag("dataloader.key", fmt.Sprintf("%v", key))

	return spanCtx, func(thunk dataloader.Thunk[V]) {
		span.Finish()
	}
}

// TraceLoadMany will trace a call to dataloader.LoadMany with Open Tracing.
func (Tracer[K, V]) TraceLoadMany(ctx context.Context, keys []K) (context.Context, dataloader.TraceLoadManyFinishFunc[V]) {
	span, spanCtx := opentracing.StartSpanFromContext(ctx, "Dataloader: loadmany")

	span.SetTag("dataloader.keys", fmt.Sprintf("%v", keys))

	return spanCtx, func(thunk dataloader.ThunkMany[V]) {
		span.Finish()
	}
}

// TraceBatch will trace a call to dataloader.LoadMany with Open Tracing.
func (Tracer[K, V]) TraceBatch(ctx context.Context, keys []K) (context.Context, dataloader.TraceBatchFinishFunc) {
	span, spanCtx := opentracing.StartSpanFromContext(ctx, "Dataloader: batch")

	span.SetTag("dataloader.keys", fmt.Sprintf("%v", keys))

	return spanCtx, func() {
		span.Finish()
	}
}
