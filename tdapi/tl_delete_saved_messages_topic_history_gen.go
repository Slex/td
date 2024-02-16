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

// DeleteSavedMessagesTopicHistoryRequest represents TL type `deleteSavedMessagesTopicHistory#69df3d6a`.
type DeleteSavedMessagesTopicHistoryRequest struct {
	// Identifier of Saved Messages topic which messages will be deleted
	SavedMessagesTopicID int64
}

// DeleteSavedMessagesTopicHistoryRequestTypeID is TL type id of DeleteSavedMessagesTopicHistoryRequest.
const DeleteSavedMessagesTopicHistoryRequestTypeID = 0x69df3d6a

// Ensuring interfaces in compile-time for DeleteSavedMessagesTopicHistoryRequest.
var (
	_ bin.Encoder     = &DeleteSavedMessagesTopicHistoryRequest{}
	_ bin.Decoder     = &DeleteSavedMessagesTopicHistoryRequest{}
	_ bin.BareEncoder = &DeleteSavedMessagesTopicHistoryRequest{}
	_ bin.BareDecoder = &DeleteSavedMessagesTopicHistoryRequest{}
)

func (d *DeleteSavedMessagesTopicHistoryRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.SavedMessagesTopicID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DeleteSavedMessagesTopicHistoryRequest) String() string {
	if d == nil {
		return "DeleteSavedMessagesTopicHistoryRequest(nil)"
	}
	type Alias DeleteSavedMessagesTopicHistoryRequest
	return fmt.Sprintf("DeleteSavedMessagesTopicHistoryRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DeleteSavedMessagesTopicHistoryRequest) TypeID() uint32 {
	return DeleteSavedMessagesTopicHistoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*DeleteSavedMessagesTopicHistoryRequest) TypeName() string {
	return "deleteSavedMessagesTopicHistory"
}

// TypeInfo returns info about TL type.
func (d *DeleteSavedMessagesTopicHistoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "deleteSavedMessagesTopicHistory",
		ID:   DeleteSavedMessagesTopicHistoryRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SavedMessagesTopicID",
			SchemaName: "saved_messages_topic_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DeleteSavedMessagesTopicHistoryRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteSavedMessagesTopicHistory#69df3d6a as nil")
	}
	b.PutID(DeleteSavedMessagesTopicHistoryRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DeleteSavedMessagesTopicHistoryRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteSavedMessagesTopicHistory#69df3d6a as nil")
	}
	b.PutInt53(d.SavedMessagesTopicID)
	return nil
}

// Decode implements bin.Decoder.
func (d *DeleteSavedMessagesTopicHistoryRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteSavedMessagesTopicHistory#69df3d6a to nil")
	}
	if err := b.ConsumeID(DeleteSavedMessagesTopicHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode deleteSavedMessagesTopicHistory#69df3d6a: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DeleteSavedMessagesTopicHistoryRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteSavedMessagesTopicHistory#69df3d6a to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode deleteSavedMessagesTopicHistory#69df3d6a: field saved_messages_topic_id: %w", err)
		}
		d.SavedMessagesTopicID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DeleteSavedMessagesTopicHistoryRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteSavedMessagesTopicHistory#69df3d6a as nil")
	}
	b.ObjStart()
	b.PutID("deleteSavedMessagesTopicHistory")
	b.Comma()
	b.FieldStart("saved_messages_topic_id")
	b.PutInt53(d.SavedMessagesTopicID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DeleteSavedMessagesTopicHistoryRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteSavedMessagesTopicHistory#69df3d6a to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("deleteSavedMessagesTopicHistory"); err != nil {
				return fmt.Errorf("unable to decode deleteSavedMessagesTopicHistory#69df3d6a: %w", err)
			}
		case "saved_messages_topic_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode deleteSavedMessagesTopicHistory#69df3d6a: field saved_messages_topic_id: %w", err)
			}
			d.SavedMessagesTopicID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSavedMessagesTopicID returns value of SavedMessagesTopicID field.
func (d *DeleteSavedMessagesTopicHistoryRequest) GetSavedMessagesTopicID() (value int64) {
	if d == nil {
		return
	}
	return d.SavedMessagesTopicID
}

// DeleteSavedMessagesTopicHistory invokes method deleteSavedMessagesTopicHistory#69df3d6a returning error if any.
func (c *Client) DeleteSavedMessagesTopicHistory(ctx context.Context, savedmessagestopicid int64) error {
	var ok Ok

	request := &DeleteSavedMessagesTopicHistoryRequest{
		SavedMessagesTopicID: savedmessagestopicid,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
