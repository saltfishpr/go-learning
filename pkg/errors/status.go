package errors

import (
	"fmt"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/saltfishpr/go-learning/pkg/util"
)

type StatusImpl struct {
	code     codes.Code
	message  string
	reason   string
	metadata map[string]string
}

func NewStatus(code codes.Code, message string) *StatusImpl {
	return &StatusImpl{
		code:    code,
		message: message,
	}
}

func (s *StatusImpl) String() string {
	return fmt.Sprintf("code = %d, message = %s reason = %s, metadata = %v", s.code, s.message, s.reason, s.metadata)
}

func (s *StatusImpl) GRPCStatus() *status.Status {
	_s := status.New(s.code, s.message)
	if s.reason == "" && s.metadata == nil {
		return _s
	}

	ds, _ := _s.WithDetails(&errdetails.ErrorInfo{
		Reason:   s.reason,
		Metadata: s.metadata,
	})
	return ds
}

func (s *StatusImpl) Clone() *StatusImpl {
	return &StatusImpl{
		code:     s.code,
		message:  s.message,
		reason:   s.reason,
		metadata: util.CloneMap(s.metadata),
	}
}

func (s *StatusImpl) WithMessage(message string) *StatusImpl {
	_s := s.Clone()
	_s.message = message
	return _s
}

func (s *StatusImpl) WithReason(reason string) *StatusImpl {
	_s := s.Clone()
	_s.reason = reason
	return _s
}

func (s *StatusImpl) WithMetadataPair(key, value string) *StatusImpl {
	_s := s.Clone()
	if _s.metadata == nil {
		_s.metadata = make(map[string]string)
	}
	_s.metadata[key] = value
	return _s
}

func (s *StatusImpl) WithMetadataMap(metadata map[string]string) *StatusImpl {
	_s := s.Clone()
	if _s.metadata == nil {
		_s.metadata = make(map[string]string)
	}
	for k, v := range metadata {
		_s.metadata[k] = v
	}
	return _s
}
