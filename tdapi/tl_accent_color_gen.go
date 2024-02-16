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

// AccentColor represents TL type `accentColor#53312690`.
type AccentColor struct {
	// Accent color identifier
	ID int32
	// Identifier of a built-in color to use in places, where only one color is needed; 0-6
	BuiltInAccentColorID int32
	// The list of 1-3 colors in RGB format, describing the accent color, as expected to be
	// shown in light themes
	LightThemeColors []int32
	// The list of 1-3 colors in RGB format, describing the accent color, as expected to be
	// shown in dark themes
	DarkThemeColors []int32
	// The minimum chat boost level required to use the color in a channel chat
	MinChannelChatBoostLevel int32
}

// AccentColorTypeID is TL type id of AccentColor.
const AccentColorTypeID = 0x53312690

// Ensuring interfaces in compile-time for AccentColor.
var (
	_ bin.Encoder     = &AccentColor{}
	_ bin.Decoder     = &AccentColor{}
	_ bin.BareEncoder = &AccentColor{}
	_ bin.BareDecoder = &AccentColor{}
)

func (a *AccentColor) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.ID == 0) {
		return false
	}
	if !(a.BuiltInAccentColorID == 0) {
		return false
	}
	if !(a.LightThemeColors == nil) {
		return false
	}
	if !(a.DarkThemeColors == nil) {
		return false
	}
	if !(a.MinChannelChatBoostLevel == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AccentColor) String() string {
	if a == nil {
		return "AccentColor(nil)"
	}
	type Alias AccentColor
	return fmt.Sprintf("AccentColor%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccentColor) TypeID() uint32 {
	return AccentColorTypeID
}

// TypeName returns name of type in TL schema.
func (*AccentColor) TypeName() string {
	return "accentColor"
}

// TypeInfo returns info about TL type.
func (a *AccentColor) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "accentColor",
		ID:   AccentColorTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "BuiltInAccentColorID",
			SchemaName: "built_in_accent_color_id",
		},
		{
			Name:       "LightThemeColors",
			SchemaName: "light_theme_colors",
		},
		{
			Name:       "DarkThemeColors",
			SchemaName: "dark_theme_colors",
		},
		{
			Name:       "MinChannelChatBoostLevel",
			SchemaName: "min_channel_chat_boost_level",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AccentColor) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode accentColor#53312690 as nil")
	}
	b.PutID(AccentColorTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AccentColor) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode accentColor#53312690 as nil")
	}
	b.PutInt32(a.ID)
	b.PutInt32(a.BuiltInAccentColorID)
	b.PutInt(len(a.LightThemeColors))
	for _, v := range a.LightThemeColors {
		b.PutInt32(v)
	}
	b.PutInt(len(a.DarkThemeColors))
	for _, v := range a.DarkThemeColors {
		b.PutInt32(v)
	}
	b.PutInt32(a.MinChannelChatBoostLevel)
	return nil
}

// Decode implements bin.Decoder.
func (a *AccentColor) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode accentColor#53312690 to nil")
	}
	if err := b.ConsumeID(AccentColorTypeID); err != nil {
		return fmt.Errorf("unable to decode accentColor#53312690: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AccentColor) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode accentColor#53312690 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode accentColor#53312690: field id: %w", err)
		}
		a.ID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode accentColor#53312690: field built_in_accent_color_id: %w", err)
		}
		a.BuiltInAccentColorID = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode accentColor#53312690: field light_theme_colors: %w", err)
		}

		if headerLen > 0 {
			a.LightThemeColors = make([]int32, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode accentColor#53312690: field light_theme_colors: %w", err)
			}
			a.LightThemeColors = append(a.LightThemeColors, value)
		}
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode accentColor#53312690: field dark_theme_colors: %w", err)
		}

		if headerLen > 0 {
			a.DarkThemeColors = make([]int32, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode accentColor#53312690: field dark_theme_colors: %w", err)
			}
			a.DarkThemeColors = append(a.DarkThemeColors, value)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode accentColor#53312690: field min_channel_chat_boost_level: %w", err)
		}
		a.MinChannelChatBoostLevel = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AccentColor) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode accentColor#53312690 as nil")
	}
	b.ObjStart()
	b.PutID("accentColor")
	b.Comma()
	b.FieldStart("id")
	b.PutInt32(a.ID)
	b.Comma()
	b.FieldStart("built_in_accent_color_id")
	b.PutInt32(a.BuiltInAccentColorID)
	b.Comma()
	b.FieldStart("light_theme_colors")
	b.ArrStart()
	for _, v := range a.LightThemeColors {
		b.PutInt32(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("dark_theme_colors")
	b.ArrStart()
	for _, v := range a.DarkThemeColors {
		b.PutInt32(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("min_channel_chat_boost_level")
	b.PutInt32(a.MinChannelChatBoostLevel)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AccentColor) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode accentColor#53312690 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("accentColor"); err != nil {
				return fmt.Errorf("unable to decode accentColor#53312690: %w", err)
			}
		case "id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode accentColor#53312690: field id: %w", err)
			}
			a.ID = value
		case "built_in_accent_color_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode accentColor#53312690: field built_in_accent_color_id: %w", err)
			}
			a.BuiltInAccentColorID = value
		case "light_theme_colors":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int32()
				if err != nil {
					return fmt.Errorf("unable to decode accentColor#53312690: field light_theme_colors: %w", err)
				}
				a.LightThemeColors = append(a.LightThemeColors, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode accentColor#53312690: field light_theme_colors: %w", err)
			}
		case "dark_theme_colors":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int32()
				if err != nil {
					return fmt.Errorf("unable to decode accentColor#53312690: field dark_theme_colors: %w", err)
				}
				a.DarkThemeColors = append(a.DarkThemeColors, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode accentColor#53312690: field dark_theme_colors: %w", err)
			}
		case "min_channel_chat_boost_level":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode accentColor#53312690: field min_channel_chat_boost_level: %w", err)
			}
			a.MinChannelChatBoostLevel = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (a *AccentColor) GetID() (value int32) {
	if a == nil {
		return
	}
	return a.ID
}

// GetBuiltInAccentColorID returns value of BuiltInAccentColorID field.
func (a *AccentColor) GetBuiltInAccentColorID() (value int32) {
	if a == nil {
		return
	}
	return a.BuiltInAccentColorID
}

// GetLightThemeColors returns value of LightThemeColors field.
func (a *AccentColor) GetLightThemeColors() (value []int32) {
	if a == nil {
		return
	}
	return a.LightThemeColors
}

// GetDarkThemeColors returns value of DarkThemeColors field.
func (a *AccentColor) GetDarkThemeColors() (value []int32) {
	if a == nil {
		return
	}
	return a.DarkThemeColors
}

// GetMinChannelChatBoostLevel returns value of MinChannelChatBoostLevel field.
func (a *AccentColor) GetMinChannelChatBoostLevel() (value int32) {
	if a == nil {
		return
	}
	return a.MinChannelChatBoostLevel
}
