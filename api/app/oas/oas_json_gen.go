// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
)

// Encode implements json.Marshaler.
func (s *LoginPostDef) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *LoginPostDef) encodeFields(e *jx.Encoder) {
	{
		if s.Message.Set {
			e.FieldStart("message")
			s.Message.Encode(e)
		}
	}
}

var jsonFieldsNameOfLoginPostDef = [1]string{
	0: "message",
}

// Decode decodes LoginPostDef from json.
func (s *LoginPostDef) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode LoginPostDef to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "message":
			if err := func() error {
				s.Message.Reset()
				if err := s.Message.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode LoginPostDef")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *LoginPostDef) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *LoginPostDef) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *LoginPostOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *LoginPostOK) encodeFields(e *jx.Encoder) {
	{
		if s.Message.Set {
			e.FieldStart("message")
			s.Message.Encode(e)
		}
	}
}

var jsonFieldsNameOfLoginPostOK = [1]string{
	0: "message",
}

// Decode decodes LoginPostOK from json.
func (s *LoginPostOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode LoginPostOK to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "message":
			if err := func() error {
				s.Message.Reset()
				if err := s.Message.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode LoginPostOK")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *LoginPostOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *LoginPostOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *LoginPostReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *LoginPostReq) encodeFields(e *jx.Encoder) {
	{
		if s.Name.Set {
			e.FieldStart("name")
			s.Name.Encode(e)
		}
	}
}

var jsonFieldsNameOfLoginPostReq = [1]string{
	0: "name",
}

// Decode decodes LoginPostReq from json.
func (s *LoginPostReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode LoginPostReq to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "name":
			if err := func() error {
				s.Name.Reset()
				if err := s.Name.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode LoginPostReq")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *LoginPostReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *LoginPostReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *TestGetDef) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *TestGetDef) encodeFields(e *jx.Encoder) {
	{
		if s.Message.Set {
			e.FieldStart("message")
			s.Message.Encode(e)
		}
	}
}

var jsonFieldsNameOfTestGetDef = [1]string{
	0: "message",
}

// Decode decodes TestGetDef from json.
func (s *TestGetDef) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TestGetDef to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "message":
			if err := func() error {
				s.Message.Reset()
				if err := s.Message.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TestGetDef")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *TestGetDef) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *TestGetDef) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *TestGetOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *TestGetOK) encodeFields(e *jx.Encoder) {
	{
		if s.Message.Set {
			e.FieldStart("message")
			s.Message.Encode(e)
		}
	}
}

var jsonFieldsNameOfTestGetOK = [1]string{
	0: "message",
}

// Decode decodes TestGetOK from json.
func (s *TestGetOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TestGetOK to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "message":
			if err := func() error {
				s.Message.Reset()
				if err := s.Message.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode TestGetOK")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *TestGetOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *TestGetOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
