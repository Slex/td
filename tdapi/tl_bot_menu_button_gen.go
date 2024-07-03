// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// BotMenuButton represents TL type `botMenuButton#c7b57ce6`.
type BotMenuButton struct {
	// Text of the button
	Text string
	// URL of a Web App to open when the button is pressed. If the link is of the type
	// internalLinkTypeWebApp, then it must be processed accordingly. Otherwise, the link
	// must be passed to openWebApp
	URL string
}

// BotMenuButtonTypeID is TL type id of BotMenuButton.
const BotMenuButtonTypeID = 0xc7b57ce6

// Ensuring interfaces in compile-time for BotMenuButton.
var (
	_ bin.Encoder     = &BotMenuButton{}
	_ bin.Decoder     = &BotMenuButton{}
	_ bin.BareEncoder = &BotMenuButton{}
	_ bin.BareDecoder = &BotMenuButton{}
)

func (b *BotMenuButton) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Text == "") {
		return false
	}
	if !(b.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotMenuButton) String() string {
	if b == nil {
		return "BotMenuButton(nil)"
	}
	type Alias BotMenuButton
	return fmt.Sprintf("BotMenuButton%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotMenuButton) TypeID() uint32 {
	return BotMenuButtonTypeID
}

// TypeName returns name of type in TL schema.
func (*BotMenuButton) TypeName() string {
	return "botMenuButton"
}

// TypeInfo returns info about TL type.
func (b *BotMenuButton) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botMenuButton",
		ID:   BotMenuButtonTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
		{
			Name:       "URL",
			SchemaName: "url",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BotMenuButton) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botMenuButton#c7b57ce6 as nil")
	}
	buf.PutID(BotMenuButtonTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BotMenuButton) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botMenuButton#c7b57ce6 as nil")
	}
	buf.PutString(b.Text)
	buf.PutString(b.URL)
	return nil
}

// Decode implements bin.Decoder.
func (b *BotMenuButton) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botMenuButton#c7b57ce6 to nil")
	}
	if err := buf.ConsumeID(BotMenuButtonTypeID); err != nil {
		return fmt.Errorf("unable to decode botMenuButton#c7b57ce6: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BotMenuButton) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botMenuButton#c7b57ce6 to nil")
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botMenuButton#c7b57ce6: field text: %w", err)
		}
		b.Text = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botMenuButton#c7b57ce6: field url: %w", err)
		}
		b.URL = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BotMenuButton) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode botMenuButton#c7b57ce6 as nil")
	}
	buf.ObjStart()
	buf.PutID("botMenuButton")
	buf.Comma()
	buf.FieldStart("text")
	buf.PutString(b.Text)
	buf.Comma()
	buf.FieldStart("url")
	buf.PutString(b.URL)
	buf.Comma()
	buf.StripComma()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BotMenuButton) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode botMenuButton#c7b57ce6 to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("botMenuButton"); err != nil {
				return fmt.Errorf("unable to decode botMenuButton#c7b57ce6: %w", err)
			}
		case "text":
			value, err := buf.String()
			if err != nil {
				return fmt.Errorf("unable to decode botMenuButton#c7b57ce6: field text: %w", err)
			}
			b.Text = value
		case "url":
			value, err := buf.String()
			if err != nil {
				return fmt.Errorf("unable to decode botMenuButton#c7b57ce6: field url: %w", err)
			}
			b.URL = value
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetText returns value of Text field.
func (b *BotMenuButton) GetText() (value string) {
	if b == nil {
		return
	}
	return b.Text
}

// GetURL returns value of URL field.
func (b *BotMenuButton) GetURL() (value string) {
	if b == nil {
		return
	}
	return b.URL
}
