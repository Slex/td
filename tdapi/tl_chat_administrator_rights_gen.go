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

// ChatAdministratorRights represents TL type `chatAdministratorRights#5f4f9044`.
type ChatAdministratorRights struct {
	// True, if the administrator can access the chat event log, get boost list, see hidden
	// supergroup and channel members, report supergroup spam messages and ignore slow mode.
	// Implied by any other privilege; applicable to supergroups and channels only
	CanManageChat bool
	// True, if the administrator can change the chat title, photo, and other settings
	CanChangeInfo bool
	// True, if the administrator can create channel posts or view channel statistics;
	// applicable to channels only
	CanPostMessages bool
	// True, if the administrator can edit messages of other users and pin messages;
	// applicable to channels only
	CanEditMessages bool
	// True, if the administrator can delete messages of other users
	CanDeleteMessages bool
	// True, if the administrator can invite new users to the chat
	CanInviteUsers bool
	// True, if the administrator can restrict, ban, or unban chat members or view supergroup
	// statistics; always true for channels
	CanRestrictMembers bool
	// True, if the administrator can pin messages; applicable to basic groups and
	// supergroups only
	CanPinMessages bool
	// True, if the administrator can create, rename, close, reopen, hide, and unhide forum
	// topics; applicable to forum supergroups only
	CanManageTopics bool
	// True, if the administrator can add new administrators with a subset of their own
	// privileges or demote administrators that were directly or indirectly promoted by them
	CanPromoteMembers bool
	// True, if the administrator can manage video chats
	CanManageVideoChats bool
	// True, if the administrator can create new chat stories, or edit and delete posted
	// stories; applicable to supergroups and channels only
	CanPostStories bool
	// True, if the administrator can edit stories posted by other users, pin stories and
	// access story archive; applicable to supergroups and channels only
	CanEditStories bool
	// True, if the administrator can delete stories posted by other users; applicable to
	// supergroups and channels only
	CanDeleteStories bool
	// True, if the administrator isn't shown in the chat member list and sends messages
	// anonymously; applicable to supergroups only
	IsAnonymous bool
}

// ChatAdministratorRightsTypeID is TL type id of ChatAdministratorRights.
const ChatAdministratorRightsTypeID = 0x5f4f9044

// Ensuring interfaces in compile-time for ChatAdministratorRights.
var (
	_ bin.Encoder     = &ChatAdministratorRights{}
	_ bin.Decoder     = &ChatAdministratorRights{}
	_ bin.BareEncoder = &ChatAdministratorRights{}
	_ bin.BareDecoder = &ChatAdministratorRights{}
)

func (c *ChatAdministratorRights) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.CanManageChat == false) {
		return false
	}
	if !(c.CanChangeInfo == false) {
		return false
	}
	if !(c.CanPostMessages == false) {
		return false
	}
	if !(c.CanEditMessages == false) {
		return false
	}
	if !(c.CanDeleteMessages == false) {
		return false
	}
	if !(c.CanInviteUsers == false) {
		return false
	}
	if !(c.CanRestrictMembers == false) {
		return false
	}
	if !(c.CanPinMessages == false) {
		return false
	}
	if !(c.CanManageTopics == false) {
		return false
	}
	if !(c.CanPromoteMembers == false) {
		return false
	}
	if !(c.CanManageVideoChats == false) {
		return false
	}
	if !(c.CanPostStories == false) {
		return false
	}
	if !(c.CanEditStories == false) {
		return false
	}
	if !(c.CanDeleteStories == false) {
		return false
	}
	if !(c.IsAnonymous == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatAdministratorRights) String() string {
	if c == nil {
		return "ChatAdministratorRights(nil)"
	}
	type Alias ChatAdministratorRights
	return fmt.Sprintf("ChatAdministratorRights%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatAdministratorRights) TypeID() uint32 {
	return ChatAdministratorRightsTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatAdministratorRights) TypeName() string {
	return "chatAdministratorRights"
}

// TypeInfo returns info about TL type.
func (c *ChatAdministratorRights) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatAdministratorRights",
		ID:   ChatAdministratorRightsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CanManageChat",
			SchemaName: "can_manage_chat",
		},
		{
			Name:       "CanChangeInfo",
			SchemaName: "can_change_info",
		},
		{
			Name:       "CanPostMessages",
			SchemaName: "can_post_messages",
		},
		{
			Name:       "CanEditMessages",
			SchemaName: "can_edit_messages",
		},
		{
			Name:       "CanDeleteMessages",
			SchemaName: "can_delete_messages",
		},
		{
			Name:       "CanInviteUsers",
			SchemaName: "can_invite_users",
		},
		{
			Name:       "CanRestrictMembers",
			SchemaName: "can_restrict_members",
		},
		{
			Name:       "CanPinMessages",
			SchemaName: "can_pin_messages",
		},
		{
			Name:       "CanManageTopics",
			SchemaName: "can_manage_topics",
		},
		{
			Name:       "CanPromoteMembers",
			SchemaName: "can_promote_members",
		},
		{
			Name:       "CanManageVideoChats",
			SchemaName: "can_manage_video_chats",
		},
		{
			Name:       "CanPostStories",
			SchemaName: "can_post_stories",
		},
		{
			Name:       "CanEditStories",
			SchemaName: "can_edit_stories",
		},
		{
			Name:       "CanDeleteStories",
			SchemaName: "can_delete_stories",
		},
		{
			Name:       "IsAnonymous",
			SchemaName: "is_anonymous",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatAdministratorRights) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatAdministratorRights#5f4f9044 as nil")
	}
	b.PutID(ChatAdministratorRightsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatAdministratorRights) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatAdministratorRights#5f4f9044 as nil")
	}
	b.PutBool(c.CanManageChat)
	b.PutBool(c.CanChangeInfo)
	b.PutBool(c.CanPostMessages)
	b.PutBool(c.CanEditMessages)
	b.PutBool(c.CanDeleteMessages)
	b.PutBool(c.CanInviteUsers)
	b.PutBool(c.CanRestrictMembers)
	b.PutBool(c.CanPinMessages)
	b.PutBool(c.CanManageTopics)
	b.PutBool(c.CanPromoteMembers)
	b.PutBool(c.CanManageVideoChats)
	b.PutBool(c.CanPostStories)
	b.PutBool(c.CanEditStories)
	b.PutBool(c.CanDeleteStories)
	b.PutBool(c.IsAnonymous)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatAdministratorRights) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatAdministratorRights#5f4f9044 to nil")
	}
	if err := b.ConsumeID(ChatAdministratorRightsTypeID); err != nil {
		return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatAdministratorRights) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatAdministratorRights#5f4f9044 to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_manage_chat: %w", err)
		}
		c.CanManageChat = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_change_info: %w", err)
		}
		c.CanChangeInfo = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_post_messages: %w", err)
		}
		c.CanPostMessages = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_edit_messages: %w", err)
		}
		c.CanEditMessages = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_delete_messages: %w", err)
		}
		c.CanDeleteMessages = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_invite_users: %w", err)
		}
		c.CanInviteUsers = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_restrict_members: %w", err)
		}
		c.CanRestrictMembers = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_pin_messages: %w", err)
		}
		c.CanPinMessages = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_manage_topics: %w", err)
		}
		c.CanManageTopics = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_promote_members: %w", err)
		}
		c.CanPromoteMembers = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_manage_video_chats: %w", err)
		}
		c.CanManageVideoChats = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_post_stories: %w", err)
		}
		c.CanPostStories = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_edit_stories: %w", err)
		}
		c.CanEditStories = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_delete_stories: %w", err)
		}
		c.CanDeleteStories = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field is_anonymous: %w", err)
		}
		c.IsAnonymous = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatAdministratorRights) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatAdministratorRights#5f4f9044 as nil")
	}
	b.ObjStart()
	b.PutID("chatAdministratorRights")
	b.Comma()
	b.FieldStart("can_manage_chat")
	b.PutBool(c.CanManageChat)
	b.Comma()
	b.FieldStart("can_change_info")
	b.PutBool(c.CanChangeInfo)
	b.Comma()
	b.FieldStart("can_post_messages")
	b.PutBool(c.CanPostMessages)
	b.Comma()
	b.FieldStart("can_edit_messages")
	b.PutBool(c.CanEditMessages)
	b.Comma()
	b.FieldStart("can_delete_messages")
	b.PutBool(c.CanDeleteMessages)
	b.Comma()
	b.FieldStart("can_invite_users")
	b.PutBool(c.CanInviteUsers)
	b.Comma()
	b.FieldStart("can_restrict_members")
	b.PutBool(c.CanRestrictMembers)
	b.Comma()
	b.FieldStart("can_pin_messages")
	b.PutBool(c.CanPinMessages)
	b.Comma()
	b.FieldStart("can_manage_topics")
	b.PutBool(c.CanManageTopics)
	b.Comma()
	b.FieldStart("can_promote_members")
	b.PutBool(c.CanPromoteMembers)
	b.Comma()
	b.FieldStart("can_manage_video_chats")
	b.PutBool(c.CanManageVideoChats)
	b.Comma()
	b.FieldStart("can_post_stories")
	b.PutBool(c.CanPostStories)
	b.Comma()
	b.FieldStart("can_edit_stories")
	b.PutBool(c.CanEditStories)
	b.Comma()
	b.FieldStart("can_delete_stories")
	b.PutBool(c.CanDeleteStories)
	b.Comma()
	b.FieldStart("is_anonymous")
	b.PutBool(c.IsAnonymous)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatAdministratorRights) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatAdministratorRights#5f4f9044 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatAdministratorRights"); err != nil {
				return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: %w", err)
			}
		case "can_manage_chat":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_manage_chat: %w", err)
			}
			c.CanManageChat = value
		case "can_change_info":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_change_info: %w", err)
			}
			c.CanChangeInfo = value
		case "can_post_messages":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_post_messages: %w", err)
			}
			c.CanPostMessages = value
		case "can_edit_messages":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_edit_messages: %w", err)
			}
			c.CanEditMessages = value
		case "can_delete_messages":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_delete_messages: %w", err)
			}
			c.CanDeleteMessages = value
		case "can_invite_users":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_invite_users: %w", err)
			}
			c.CanInviteUsers = value
		case "can_restrict_members":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_restrict_members: %w", err)
			}
			c.CanRestrictMembers = value
		case "can_pin_messages":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_pin_messages: %w", err)
			}
			c.CanPinMessages = value
		case "can_manage_topics":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_manage_topics: %w", err)
			}
			c.CanManageTopics = value
		case "can_promote_members":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_promote_members: %w", err)
			}
			c.CanPromoteMembers = value
		case "can_manage_video_chats":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_manage_video_chats: %w", err)
			}
			c.CanManageVideoChats = value
		case "can_post_stories":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_post_stories: %w", err)
			}
			c.CanPostStories = value
		case "can_edit_stories":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_edit_stories: %w", err)
			}
			c.CanEditStories = value
		case "can_delete_stories":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field can_delete_stories: %w", err)
			}
			c.CanDeleteStories = value
		case "is_anonymous":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatAdministratorRights#5f4f9044: field is_anonymous: %w", err)
			}
			c.IsAnonymous = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCanManageChat returns value of CanManageChat field.
func (c *ChatAdministratorRights) GetCanManageChat() (value bool) {
	if c == nil {
		return
	}
	return c.CanManageChat
}

// GetCanChangeInfo returns value of CanChangeInfo field.
func (c *ChatAdministratorRights) GetCanChangeInfo() (value bool) {
	if c == nil {
		return
	}
	return c.CanChangeInfo
}

// GetCanPostMessages returns value of CanPostMessages field.
func (c *ChatAdministratorRights) GetCanPostMessages() (value bool) {
	if c == nil {
		return
	}
	return c.CanPostMessages
}

// GetCanEditMessages returns value of CanEditMessages field.
func (c *ChatAdministratorRights) GetCanEditMessages() (value bool) {
	if c == nil {
		return
	}
	return c.CanEditMessages
}

// GetCanDeleteMessages returns value of CanDeleteMessages field.
func (c *ChatAdministratorRights) GetCanDeleteMessages() (value bool) {
	if c == nil {
		return
	}
	return c.CanDeleteMessages
}

// GetCanInviteUsers returns value of CanInviteUsers field.
func (c *ChatAdministratorRights) GetCanInviteUsers() (value bool) {
	if c == nil {
		return
	}
	return c.CanInviteUsers
}

// GetCanRestrictMembers returns value of CanRestrictMembers field.
func (c *ChatAdministratorRights) GetCanRestrictMembers() (value bool) {
	if c == nil {
		return
	}
	return c.CanRestrictMembers
}

// GetCanPinMessages returns value of CanPinMessages field.
func (c *ChatAdministratorRights) GetCanPinMessages() (value bool) {
	if c == nil {
		return
	}
	return c.CanPinMessages
}

// GetCanManageTopics returns value of CanManageTopics field.
func (c *ChatAdministratorRights) GetCanManageTopics() (value bool) {
	if c == nil {
		return
	}
	return c.CanManageTopics
}

// GetCanPromoteMembers returns value of CanPromoteMembers field.
func (c *ChatAdministratorRights) GetCanPromoteMembers() (value bool) {
	if c == nil {
		return
	}
	return c.CanPromoteMembers
}

// GetCanManageVideoChats returns value of CanManageVideoChats field.
func (c *ChatAdministratorRights) GetCanManageVideoChats() (value bool) {
	if c == nil {
		return
	}
	return c.CanManageVideoChats
}

// GetCanPostStories returns value of CanPostStories field.
func (c *ChatAdministratorRights) GetCanPostStories() (value bool) {
	if c == nil {
		return
	}
	return c.CanPostStories
}

// GetCanEditStories returns value of CanEditStories field.
func (c *ChatAdministratorRights) GetCanEditStories() (value bool) {
	if c == nil {
		return
	}
	return c.CanEditStories
}

// GetCanDeleteStories returns value of CanDeleteStories field.
func (c *ChatAdministratorRights) GetCanDeleteStories() (value bool) {
	if c == nil {
		return
	}
	return c.CanDeleteStories
}

// GetIsAnonymous returns value of IsAnonymous field.
func (c *ChatAdministratorRights) GetIsAnonymous() (value bool) {
	if c == nil {
		return
	}
	return c.IsAnonymous
}
