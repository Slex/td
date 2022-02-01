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

// Reaction represents TL type `reaction#ff6c56f1`.
type Reaction struct {
	// Text representation of the reaction
	Reaction string
	// Reaction title
	Title string
	// True, if the reaction can be added to new messages and enabled in chats
	IsActive bool
	// Static icon for the reaction
	StaticIcon Sticker
	// Appear animation for the reaction
	AppearAnimation Sticker
	// Select animation for the reaction
	SelectAnimation Sticker
	// Activate animation for the reaction
	ActivateAnimation Sticker
	// Effect animation for the reaction
	EffectAnimation Sticker
	// Around animation for the reaction; may be null
	AroundAnimation Sticker
	// Center animation for the reaction; may be null
	CenterAnimation Sticker
}

// ReactionTypeID is TL type id of Reaction.
const ReactionTypeID = 0xff6c56f1

// Ensuring interfaces in compile-time for Reaction.
var (
	_ bin.Encoder     = &Reaction{}
	_ bin.Decoder     = &Reaction{}
	_ bin.BareEncoder = &Reaction{}
	_ bin.BareDecoder = &Reaction{}
)

func (r *Reaction) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Reaction == "") {
		return false
	}
	if !(r.Title == "") {
		return false
	}
	if !(r.IsActive == false) {
		return false
	}
	if !(r.StaticIcon.Zero()) {
		return false
	}
	if !(r.AppearAnimation.Zero()) {
		return false
	}
	if !(r.SelectAnimation.Zero()) {
		return false
	}
	if !(r.ActivateAnimation.Zero()) {
		return false
	}
	if !(r.EffectAnimation.Zero()) {
		return false
	}
	if !(r.AroundAnimation.Zero()) {
		return false
	}
	if !(r.CenterAnimation.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *Reaction) String() string {
	if r == nil {
		return "Reaction(nil)"
	}
	type Alias Reaction
	return fmt.Sprintf("Reaction%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Reaction) TypeID() uint32 {
	return ReactionTypeID
}

// TypeName returns name of type in TL schema.
func (*Reaction) TypeName() string {
	return "reaction"
}

// TypeInfo returns info about TL type.
func (r *Reaction) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reaction",
		ID:   ReactionTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Reaction",
			SchemaName: "reaction",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "IsActive",
			SchemaName: "is_active",
		},
		{
			Name:       "StaticIcon",
			SchemaName: "static_icon",
		},
		{
			Name:       "AppearAnimation",
			SchemaName: "appear_animation",
		},
		{
			Name:       "SelectAnimation",
			SchemaName: "select_animation",
		},
		{
			Name:       "ActivateAnimation",
			SchemaName: "activate_animation",
		},
		{
			Name:       "EffectAnimation",
			SchemaName: "effect_animation",
		},
		{
			Name:       "AroundAnimation",
			SchemaName: "around_animation",
		},
		{
			Name:       "CenterAnimation",
			SchemaName: "center_animation",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *Reaction) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reaction#ff6c56f1 as nil")
	}
	b.PutID(ReactionTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *Reaction) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reaction#ff6c56f1 as nil")
	}
	b.PutString(r.Reaction)
	b.PutString(r.Title)
	b.PutBool(r.IsActive)
	if err := r.StaticIcon.Encode(b); err != nil {
		return fmt.Errorf("unable to encode reaction#ff6c56f1: field static_icon: %w", err)
	}
	if err := r.AppearAnimation.Encode(b); err != nil {
		return fmt.Errorf("unable to encode reaction#ff6c56f1: field appear_animation: %w", err)
	}
	if err := r.SelectAnimation.Encode(b); err != nil {
		return fmt.Errorf("unable to encode reaction#ff6c56f1: field select_animation: %w", err)
	}
	if err := r.ActivateAnimation.Encode(b); err != nil {
		return fmt.Errorf("unable to encode reaction#ff6c56f1: field activate_animation: %w", err)
	}
	if err := r.EffectAnimation.Encode(b); err != nil {
		return fmt.Errorf("unable to encode reaction#ff6c56f1: field effect_animation: %w", err)
	}
	if err := r.AroundAnimation.Encode(b); err != nil {
		return fmt.Errorf("unable to encode reaction#ff6c56f1: field around_animation: %w", err)
	}
	if err := r.CenterAnimation.Encode(b); err != nil {
		return fmt.Errorf("unable to encode reaction#ff6c56f1: field center_animation: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *Reaction) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reaction#ff6c56f1 to nil")
	}
	if err := b.ConsumeID(ReactionTypeID); err != nil {
		return fmt.Errorf("unable to decode reaction#ff6c56f1: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *Reaction) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reaction#ff6c56f1 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode reaction#ff6c56f1: field reaction: %w", err)
		}
		r.Reaction = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode reaction#ff6c56f1: field title: %w", err)
		}
		r.Title = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode reaction#ff6c56f1: field is_active: %w", err)
		}
		r.IsActive = value
	}
	{
		if err := r.StaticIcon.Decode(b); err != nil {
			return fmt.Errorf("unable to decode reaction#ff6c56f1: field static_icon: %w", err)
		}
	}
	{
		if err := r.AppearAnimation.Decode(b); err != nil {
			return fmt.Errorf("unable to decode reaction#ff6c56f1: field appear_animation: %w", err)
		}
	}
	{
		if err := r.SelectAnimation.Decode(b); err != nil {
			return fmt.Errorf("unable to decode reaction#ff6c56f1: field select_animation: %w", err)
		}
	}
	{
		if err := r.ActivateAnimation.Decode(b); err != nil {
			return fmt.Errorf("unable to decode reaction#ff6c56f1: field activate_animation: %w", err)
		}
	}
	{
		if err := r.EffectAnimation.Decode(b); err != nil {
			return fmt.Errorf("unable to decode reaction#ff6c56f1: field effect_animation: %w", err)
		}
	}
	{
		if err := r.AroundAnimation.Decode(b); err != nil {
			return fmt.Errorf("unable to decode reaction#ff6c56f1: field around_animation: %w", err)
		}
	}
	{
		if err := r.CenterAnimation.Decode(b); err != nil {
			return fmt.Errorf("unable to decode reaction#ff6c56f1: field center_animation: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *Reaction) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode reaction#ff6c56f1 as nil")
	}
	b.ObjStart()
	b.PutID("reaction")
	b.Comma()
	b.FieldStart("reaction")
	b.PutString(r.Reaction)
	b.Comma()
	b.FieldStart("title")
	b.PutString(r.Title)
	b.Comma()
	b.FieldStart("is_active")
	b.PutBool(r.IsActive)
	b.Comma()
	b.FieldStart("static_icon")
	if err := r.StaticIcon.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode reaction#ff6c56f1: field static_icon: %w", err)
	}
	b.Comma()
	b.FieldStart("appear_animation")
	if err := r.AppearAnimation.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode reaction#ff6c56f1: field appear_animation: %w", err)
	}
	b.Comma()
	b.FieldStart("select_animation")
	if err := r.SelectAnimation.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode reaction#ff6c56f1: field select_animation: %w", err)
	}
	b.Comma()
	b.FieldStart("activate_animation")
	if err := r.ActivateAnimation.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode reaction#ff6c56f1: field activate_animation: %w", err)
	}
	b.Comma()
	b.FieldStart("effect_animation")
	if err := r.EffectAnimation.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode reaction#ff6c56f1: field effect_animation: %w", err)
	}
	b.Comma()
	b.FieldStart("around_animation")
	if err := r.AroundAnimation.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode reaction#ff6c56f1: field around_animation: %w", err)
	}
	b.Comma()
	b.FieldStart("center_animation")
	if err := r.CenterAnimation.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode reaction#ff6c56f1: field center_animation: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *Reaction) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode reaction#ff6c56f1 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("reaction"); err != nil {
				return fmt.Errorf("unable to decode reaction#ff6c56f1: %w", err)
			}
		case "reaction":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode reaction#ff6c56f1: field reaction: %w", err)
			}
			r.Reaction = value
		case "title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode reaction#ff6c56f1: field title: %w", err)
			}
			r.Title = value
		case "is_active":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode reaction#ff6c56f1: field is_active: %w", err)
			}
			r.IsActive = value
		case "static_icon":
			if err := r.StaticIcon.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode reaction#ff6c56f1: field static_icon: %w", err)
			}
		case "appear_animation":
			if err := r.AppearAnimation.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode reaction#ff6c56f1: field appear_animation: %w", err)
			}
		case "select_animation":
			if err := r.SelectAnimation.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode reaction#ff6c56f1: field select_animation: %w", err)
			}
		case "activate_animation":
			if err := r.ActivateAnimation.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode reaction#ff6c56f1: field activate_animation: %w", err)
			}
		case "effect_animation":
			if err := r.EffectAnimation.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode reaction#ff6c56f1: field effect_animation: %w", err)
			}
		case "around_animation":
			if err := r.AroundAnimation.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode reaction#ff6c56f1: field around_animation: %w", err)
			}
		case "center_animation":
			if err := r.CenterAnimation.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode reaction#ff6c56f1: field center_animation: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetReaction returns value of Reaction field.
func (r *Reaction) GetReaction() (value string) {
	if r == nil {
		return
	}
	return r.Reaction
}

// GetTitle returns value of Title field.
func (r *Reaction) GetTitle() (value string) {
	if r == nil {
		return
	}
	return r.Title
}

// GetIsActive returns value of IsActive field.
func (r *Reaction) GetIsActive() (value bool) {
	if r == nil {
		return
	}
	return r.IsActive
}

// GetStaticIcon returns value of StaticIcon field.
func (r *Reaction) GetStaticIcon() (value Sticker) {
	if r == nil {
		return
	}
	return r.StaticIcon
}

// GetAppearAnimation returns value of AppearAnimation field.
func (r *Reaction) GetAppearAnimation() (value Sticker) {
	if r == nil {
		return
	}
	return r.AppearAnimation
}

// GetSelectAnimation returns value of SelectAnimation field.
func (r *Reaction) GetSelectAnimation() (value Sticker) {
	if r == nil {
		return
	}
	return r.SelectAnimation
}

// GetActivateAnimation returns value of ActivateAnimation field.
func (r *Reaction) GetActivateAnimation() (value Sticker) {
	if r == nil {
		return
	}
	return r.ActivateAnimation
}

// GetEffectAnimation returns value of EffectAnimation field.
func (r *Reaction) GetEffectAnimation() (value Sticker) {
	if r == nil {
		return
	}
	return r.EffectAnimation
}

// GetAroundAnimation returns value of AroundAnimation field.
func (r *Reaction) GetAroundAnimation() (value Sticker) {
	if r == nil {
		return
	}
	return r.AroundAnimation
}

// GetCenterAnimation returns value of CenterAnimation field.
func (r *Reaction) GetCenterAnimation() (value Sticker) {
	if r == nil {
		return
	}
	return r.CenterAnimation
}
