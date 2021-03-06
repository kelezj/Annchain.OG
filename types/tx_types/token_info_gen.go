package tx_types

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/annchain/OG/common/math"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *TokenInfo) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
		return
	}
	err = z.PublicOffering.DecodeMsg(dc)
	if err != nil {
		return
	}
	err = z.Sender.DecodeMsg(dc)
	if err != nil {
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			return
		}
		z.CurrentValue = nil
	} else {
		if z.CurrentValue == nil {
			z.CurrentValue = new(math.BigInt)
		}
		err = z.CurrentValue.DecodeMsg(dc)
		if err != nil {
			return
		}
	}
	z.Destroyed, err = dc.ReadBool()
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *TokenInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 4
	err = en.Append(0x94)
	if err != nil {
		return
	}
	err = z.PublicOffering.EncodeMsg(en)
	if err != nil {
		return
	}
	err = z.Sender.EncodeMsg(en)
	if err != nil {
		return
	}
	if z.CurrentValue == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.CurrentValue.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	err = en.WriteBool(z.Destroyed)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TokenInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 4
	o = append(o, 0x94)
	o, err = z.PublicOffering.MarshalMsg(o)
	if err != nil {
		return
	}
	o, err = z.Sender.MarshalMsg(o)
	if err != nil {
		return
	}
	if z.CurrentValue == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.CurrentValue.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	o = msgp.AppendBool(o, z.Destroyed)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TokenInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
		return
	}
	bts, err = z.PublicOffering.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	bts, err = z.Sender.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.CurrentValue = nil
	} else {
		if z.CurrentValue == nil {
			z.CurrentValue = new(math.BigInt)
		}
		bts, err = z.CurrentValue.UnmarshalMsg(bts)
		if err != nil {
			return
		}
	}
	z.Destroyed, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TokenInfo) Msgsize() (s int) {
	s = 1 + z.PublicOffering.Msgsize() + z.Sender.Msgsize()
	if z.CurrentValue == nil {
		s += msgp.NilSize
	} else {
		s += z.CurrentValue.Msgsize()
	}
	s += msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TokensInfo) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(TokensInfo, zb0002)
	}
	for zb0001 := range *z {
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(TokenInfo)
			}
			err = (*z)[zb0001].DecodeMsg(dc)
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z TokensInfo) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		return
	}
	for zb0003 := range z {
		if z[zb0003] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z[zb0003].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z TokensInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0003 := range z {
		if z[zb0003] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z[zb0003].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TokensInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(TokensInfo, zb0002)
	}
	for zb0001 := range *z {
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(TokenInfo)
			}
			bts, err = (*z)[zb0001].UnmarshalMsg(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z TokensInfo) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0003 := range z {
		if z[zb0003] == nil {
			s += msgp.NilSize
		} else {
			s += z[zb0003].Msgsize()
		}
	}
	return
}
