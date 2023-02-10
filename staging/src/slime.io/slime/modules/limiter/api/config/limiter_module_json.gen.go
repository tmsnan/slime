// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: limiter_module.proto

package config

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/durationpb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for Limiter
func (this *Limiter) MarshalJSON() ([]byte, error) {
	str, err := LimiterModuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Limiter
func (this *Limiter) UnmarshalJSON(b []byte) error {
	return LimiterModuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RlsConfigMap
func (this *RlsConfigMap) MarshalJSON() ([]byte, error) {
	str, err := LimiterModuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RlsConfigMap
func (this *RlsConfigMap) UnmarshalJSON(b []byte) error {
	return LimiterModuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RateLimitService
func (this *RateLimitService) MarshalJSON() ([]byte, error) {
	str, err := LimiterModuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RateLimitService
func (this *RateLimitService) UnmarshalJSON(b []byte) error {
	return LimiterModuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	LimiterModuleMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	LimiterModuleUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{AllowUnknownFields: true}
)