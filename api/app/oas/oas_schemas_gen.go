// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"fmt"
)

func (s *R500InternalServerErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

type Bearer struct {
	Token string
}

// GetToken returns the value of Token.
func (s *Bearer) GetToken() string {
	return s.Token
}

// SetToken sets the value of Token.
func (s *Bearer) SetToken(val string) {
	s.Token = val
}

// Ref: #/components/schemas/market
type Market struct {
	// The market ID(uuid).
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// The market images uuid.
	Images []string `json:"images"`
}

// GetID returns the value of ID.
func (s *Market) GetID() string {
	return s.ID
}

// GetName returns the value of Name.
func (s *Market) GetName() string {
	return s.Name
}

// GetDescription returns the value of Description.
func (s *Market) GetDescription() string {
	return s.Description
}

// GetImages returns the value of Images.
func (s *Market) GetImages() []string {
	return s.Images
}

// SetID sets the value of ID.
func (s *Market) SetID(val string) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *Market) SetName(val string) {
	s.Name = val
}

// SetDescription sets the value of Description.
func (s *Market) SetDescription(val string) {
	s.Description = val
}

// SetImages sets the value of Images.
func (s *Market) SetImages(val []string) {
	s.Images = val
}

// MarketMarketIdDeleteNoContent is response for MarketMarketIdDelete operation.
type MarketMarketIdDeleteNoContent struct{}

type MarketPostReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GetName returns the value of Name.
func (s *MarketPostReq) GetName() string {
	return s.Name
}

// GetDescription returns the value of Description.
func (s *MarketPostReq) GetDescription() string {
	return s.Description
}

// SetName sets the value of Name.
func (s *MarketPostReq) SetName(val string) {
	s.Name = val
}

// SetDescription sets the value of Description.
func (s *MarketPostReq) SetDescription(val string) {
	s.Description = val
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

type R200OK struct {
	Message string `json:"message"`
}

// GetMessage returns the value of Message.
func (s *R200OK) GetMessage() string {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *R200OK) SetMessage(val string) {
	s.Message = val
}

type R201Created struct {
	Message string `json:"message"`
}

// GetMessage returns the value of Message.
func (s *R201Created) GetMessage() string {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *R201Created) SetMessage(val string) {
	s.Message = val
}

func (*R201Created) signUpPostRes() {}

type R401Unauthorized struct {
	Message string `json:"message"`
}

// GetMessage returns the value of Message.
func (s *R401Unauthorized) GetMessage() string {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *R401Unauthorized) SetMessage(val string) {
	s.Message = val
}

func (*R401Unauthorized) signUpPostRes() {}

type R500InternalServerError struct {
	Message string `json:"message"`
}

// GetMessage returns the value of Message.
func (s *R500InternalServerError) GetMessage() string {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *R500InternalServerError) SetMessage(val string) {
	s.Message = val
}

// R500InternalServerErrorStatusCode wraps R500InternalServerError with StatusCode.
type R500InternalServerErrorStatusCode struct {
	StatusCode int
	Response   R500InternalServerError
}

// GetStatusCode returns the value of StatusCode.
func (s *R500InternalServerErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *R500InternalServerErrorStatusCode) GetResponse() R500InternalServerError {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *R500InternalServerErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *R500InternalServerErrorStatusCode) SetResponse(val R500InternalServerError) {
	s.Response = val
}

type SignUpPostReq struct {
	Name OptString `json:"name"`
}

// GetName returns the value of Name.
func (s *SignUpPostReq) GetName() OptString {
	return s.Name
}

// SetName sets the value of Name.
func (s *SignUpPostReq) SetName(val OptString) {
	s.Name = val
}
