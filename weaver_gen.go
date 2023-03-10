package weaverutils

// Code generated by "weaver generate". DO NOT EDIT.
import (
	"context"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
	"time"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:     "gopkg.eu.org/weaverutils/WeaverFastCache",
		Iface:    reflect.TypeOf((*WeaverFastCache)(nil)).Elem(),
		New:      func() any { return &weaverFastCache{} },
		ConfigFn: func(i any) any { return i.(*weaverFastCache).WithConfig.Config() },
		Routed:   true,
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return weaverFastCache_local_stub{impl: impl.(WeaverFastCache), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return weaverFastCache_client_stub{stub: stub, getMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "gopkg.eu.org/weaverutils/WeaverFastCache", Method: "Get"}), setMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "gopkg.eu.org/weaverutils/WeaverFastCache", Method: "Set"}), delMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "gopkg.eu.org/weaverutils/WeaverFastCache", Method: "Del"}), hasMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "gopkg.eu.org/weaverutils/WeaverFastCache", Method: "Has"}), hasGetMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "gopkg.eu.org/weaverutils/WeaverFastCache", Method: "HasGet"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return weaverFastCache_server_stub{impl: impl.(WeaverFastCache), addLoad: addLoad}
		},
	})
}

// Local stub implementations.

type weaverFastCache_local_stub struct {
	impl   WeaverFastCache
	tracer trace.Tracer
}

func (s weaverFastCache_local_stub) Get(ctx context.Context, a0 []byte) (r0 []byte, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "weaverutils.WeaverFastCache.Get", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Get(ctx, a0)
}

func (s weaverFastCache_local_stub) Set(ctx context.Context, a0 []byte, a1 []byte) (err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "weaverutils.WeaverFastCache.Set", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Set(ctx, a0, a1)
}

func (s weaverFastCache_local_stub) Del(ctx context.Context, a0 []byte) (err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "weaverutils.WeaverFastCache.Del", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Del(ctx, a0)
}

func (s weaverFastCache_local_stub) Has(ctx context.Context, a0 []byte) (r0 bool, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "weaverutils.WeaverFastCache.Has", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Has(ctx, a0)
}

func (s weaverFastCache_local_stub) HasGet(ctx context.Context, a0 []byte) (r0 []byte, r1 bool, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "weaverutils.WeaverFastCache.HasGet", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.HasGet(ctx, a0)
}

// Client stub implementations.

type weaverFastCache_client_stub struct {
	stub          codegen.Stub
	getMetrics    *codegen.MethodMetrics
	setMetrics    *codegen.MethodMetrics
	delMetrics    *codegen.MethodMetrics
	hasMetrics    *codegen.MethodMetrics
	hasGetMetrics *codegen.MethodMetrics
}

func (s weaverFastCache_client_stub) Get(ctx context.Context, a0 []byte) (r0 []byte, err error) {
	// Update metrics.
	start := time.Now()
	s.getMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "weaverutils.WeaverFastCache.Get", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
		err = s.stub.WrapError(err)

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.getMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.getMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + (len(a0) * 1))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	serviceweaver_enc_slice_byte_87461245(enc, a0)

	// Set the shardKey.
	var r weaverFastCacheRouter
	shardKey := _hashWeaverFastCache(r.Get(ctx, a0))

	// Call the remote method.
	s.getMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 1, enc.Data(), shardKey)
	if err != nil {
		return
	}
	s.getMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_byte_87461245(dec)
	err = dec.Error()
	return
}

func (s weaverFastCache_client_stub) Set(ctx context.Context, a0 []byte, a1 []byte) (err error) {
	// Update metrics.
	start := time.Now()
	s.setMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "weaverutils.WeaverFastCache.Set", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
		err = s.stub.WrapError(err)

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.setMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.setMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + (len(a0) * 1))
	size += (4 + (len(a1) * 1))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	serviceweaver_enc_slice_byte_87461245(enc, a0)
	serviceweaver_enc_slice_byte_87461245(enc, a1)

	// Set the shardKey.
	var r weaverFastCacheRouter
	shardKey := _hashWeaverFastCache(r.Set(ctx, a0, a1))

	// Call the remote method.
	s.setMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 4, enc.Data(), shardKey)
	if err != nil {
		return
	}
	s.setMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s weaverFastCache_client_stub) Del(ctx context.Context, a0 []byte) (err error) {
	// Update metrics.
	start := time.Now()
	s.delMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "weaverutils.WeaverFastCache.Del", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
		err = s.stub.WrapError(err)

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.delMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.delMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + (len(a0) * 1))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	serviceweaver_enc_slice_byte_87461245(enc, a0)

	// Set the shardKey.
	var r weaverFastCacheRouter
	shardKey := _hashWeaverFastCache(r.Del(ctx, a0))

	// Call the remote method.
	s.delMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		return
	}
	s.delMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s weaverFastCache_client_stub) Has(ctx context.Context, a0 []byte) (r0 bool, err error) {
	// Update metrics.
	start := time.Now()
	s.hasMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "weaverutils.WeaverFastCache.Has", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
		err = s.stub.WrapError(err)

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.hasMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.hasMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + (len(a0) * 1))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	serviceweaver_enc_slice_byte_87461245(enc, a0)

	// Set the shardKey.
	var r weaverFastCacheRouter
	shardKey := _hashWeaverFastCache(r.Has(ctx, a0))

	// Call the remote method.
	s.hasMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 2, enc.Data(), shardKey)
	if err != nil {
		return
	}
	s.hasMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.Bool()
	err = dec.Error()
	return
}

func (s weaverFastCache_client_stub) HasGet(ctx context.Context, a0 []byte) (r0 []byte, r1 bool, err error) {
	// Update metrics.
	start := time.Now()
	s.hasGetMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "weaverutils.WeaverFastCache.HasGet", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
		err = s.stub.WrapError(err)

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.hasGetMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.hasGetMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + (len(a0) * 1))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	serviceweaver_enc_slice_byte_87461245(enc, a0)

	// Set the shardKey.
	var r weaverFastCacheRouter
	shardKey := _hashWeaverFastCache(r.HasGet(ctx, a0))

	// Call the remote method.
	s.hasGetMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 3, enc.Data(), shardKey)
	if err != nil {
		return
	}
	s.hasGetMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_byte_87461245(dec)
	r1 = dec.Bool()
	err = dec.Error()
	return
}

// Server stub implementations.

type weaverFastCache_server_stub struct {
	impl    WeaverFastCache
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s weaverFastCache_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Get":
		return s.get
	case "Set":
		return s.set
	case "Del":
		return s.del
	case "Has":
		return s.has
	case "HasGet":
		return s.hasGet
	default:
		return nil
	}
}

func (s weaverFastCache_server_stub) get(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 []byte
	a0 = serviceweaver_dec_slice_byte_87461245(dec)
	var r weaverFastCacheRouter
	s.addLoad(_hashWeaverFastCache(r.Get(ctx, a0)), 1.0)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Get(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_byte_87461245(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s weaverFastCache_server_stub) set(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 []byte
	a0 = serviceweaver_dec_slice_byte_87461245(dec)
	var a1 []byte
	a1 = serviceweaver_dec_slice_byte_87461245(dec)
	var r weaverFastCacheRouter
	s.addLoad(_hashWeaverFastCache(r.Set(ctx, a0, a1)), 1.0)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.Set(ctx, a0, a1)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s weaverFastCache_server_stub) del(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 []byte
	a0 = serviceweaver_dec_slice_byte_87461245(dec)
	var r weaverFastCacheRouter
	s.addLoad(_hashWeaverFastCache(r.Del(ctx, a0)), 1.0)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.Del(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s weaverFastCache_server_stub) has(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 []byte
	a0 = serviceweaver_dec_slice_byte_87461245(dec)
	var r weaverFastCacheRouter
	s.addLoad(_hashWeaverFastCache(r.Has(ctx, a0)), 1.0)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Has(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Bool(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s weaverFastCache_server_stub) hasGet(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 []byte
	a0 = serviceweaver_dec_slice_byte_87461245(dec)
	var r weaverFastCacheRouter
	s.addLoad(_hashWeaverFastCache(r.HasGet(ctx, a0)), 1.0)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, r1, appErr := s.impl.HasGet(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_byte_87461245(enc, r0)
	enc.Bool(r1)
	enc.Error(appErr)
	return enc.Data(), nil
}

// Router methods.

// _hashWeaverFastCache returns a 64 bit hash of the provided value.
func _hashWeaverFastCache(r uint64) uint64 {
	var h codegen.Hasher
	h.WriteUint64(uint64(r))
	return h.Sum64()
}

// _orderedCodeWeaverFastCache returns an order-preserving serialization of the provided value.
func _orderedCodeWeaverFastCache(r uint64) codegen.OrderedCode {
	var enc codegen.OrderedEncoder
	enc.WriteUint64(uint64(r))
	return enc.Encode()
}

// Encoding/decoding implementations.

func serviceweaver_enc_slice_byte_87461245(enc *codegen.Encoder, arg []byte) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		enc.Byte(arg[i])
	}
}

func serviceweaver_dec_slice_byte_87461245(dec *codegen.Decoder) []byte {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = dec.Byte()
	}
	return res
}
