package types

type Update struct {
	UpdateID int32 `json:"update_id,omitempty"`
	Message  Message
}

type Message struct {
	MessageID             int32 `json:"message_id,omitempty"`
	From                  *User
	Date                  int32
	Chat                  *Chat
	ForwardFrom           *User
	ForwardFromChat       *Chat
	ForwardFromMessageID  int32
	ForwardSignature      string
	ForwardDate           int32
	ReplyToMessage        *Message
	EditDate              int32
	MediaGroupID          string
	AuthorSignature       string
	Text                  string
	Entities              []MessageEntity
	CaptionEntities       []MessageEntity
	Audio                 *Audio
	Document              *Document
	Game                  *Game
	Photo                 []PhotoSize
	Sticker               *Sticker
	Video                 *Video
	Voice                 *Voice
	VideoNote             *VideoNote
	Caption               string
	Contact               *Contact
	Location              *Location
	Venue                 *Venue
	NewChatMembers        []User
	LeftCharMember        *User
	NewChatTitle          string
	NewChatPhoto          []PhotoSize
	DeleteChatPhoto       bool
	GroupChatCreated      bool
	SupergroupChatCreated bool
	ChannelChatCreated    bool
	MigrateToChatID       int32
	MigrateFromChatID     int32
	PinnedMessage         *Message
	Invoice               *Invoice
	SuccessfulPayment     *SuccessfulPayment
	ConnectedWebsite      string
}

type User struct {
	ID           int32  `json:"id,omitempty"`
	IsBot        bool   `json:"is_bot,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Username     string `json:"username,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

type Chat struct {
	ID                          int32      `json:"id,omitempty"`
	Type                        string     `json:"type,omitempty"`
	Title                       string     `json:"title,omitempty"`
	Username                    string     `json:"username,omitempty"`
	FirstName                   string     `json:"first_name,omitempty"`
	LastName                    string     `json:"last_name,omitempty"`
	AllMembersAreAdministrators bool       `json:"all_members_are_administrators,omitempty"`
	Photo                       *ChatPhoto `json:"photo,omitempty"`
	Description                 string     `json:"description,omitempty"`
	InviteLink                  string     `json:"invite_link,omitempty"`
	PinnedMessage               *Message   `json:"pinned_message,omitempty"`
	StickerSetName              string     `json:"sticker_set_name,omitempty"`
	CanSetStickerName           bool       `json:"can_set_sticker_name,omitempty"`
}

type ChatPhoto struct {
	SmallFileID string `json:"small_file_id,omitempty"`
	BigFileID   string `json:"big_file_id,omitempty"`
}

type MessageEntity struct {
	Type   string `json:"type,omitempty"`
	Offset int32  `json:"offset,omitempty"`
	Length int32  `json:"length,omitempty"`
	URL    string `json:"url,omitempty"`
	User   *User  `json:"user,omitempty"`
}

type Audio struct {
	FileID    string `json:"file_id,omitempty"`
	Duration  int32  `json:"duration,omitempty"`
	Performer string `json:"performer,omitempty"`
	Title     string `json:"title,omitempty"`
	MIMEType  string `json:"mime_type,omitempty"`
	FileSize  int32  `json:"file_size,omitempty"`
}

// Document represents a general file (as opposed to photos, voice messages and audio files)
type Document struct {
	FileID   string     `json:"file_id,omitempty"`   // unique file identifier
	Thumb    *PhotoSize `json:"thumb,omitempty"`     // optional, document thumbnail as defined by sender
	FileName string     `json:"file_name,omitempty"` // optional, original filename as defined by sender
	MIMEType string     `json:"mime_type,omitempty"` // optional, MIME type of the file as defined by sender
	FileSize int32      `json:"file_size,omitempty"` // optional, file size
}

// PhotoSize represents one size of a photo or a file/sticker thumbnail
type PhotoSize struct {
	FileID   string `json:"file_id,omitempty"`   // unique identifier for this file
	Width    int32  `json:"width,omitempty"`     // photo width
	Height   int32  `json:"height,omitempty"`    // photo height
	FileSize int32  `json:"file_size,omitempty"` // optional, file size
}

// Game represents a game. Use BotFather to create and edit games,
// their short names will act as unique identifiers.
type Game struct {
	Title       string      // title of the game.
	Description string      // description of the game.
	Photo       []PhotoSize // photo that will be displayed in the game message in chats.
	// optional, brief description of the game or high scores included in the game message.
	// can be automatically edited to include current high scores for the game when the bot calls setGameScore,
	// or manually edited using editMessageText. 0-4096 characters.
	Text         string
	TextEntities []MessageEntity // optional, special entities that appear in text, such as usernames, URLs, bot commands, etc.
	Animation    *Animation      // optional. Animation that will be displayed in the game message in chats. Upload via BotFather.
}

// Animation represents an animation file to be displayed in the message containing a game.
// You can provide an animation for your game so that it looks stylish in chats.
type Animation struct {
	FileID   string     `json:"file_id,omitempty"`   // unique file identifier
	Thumb    *PhotoSize `json:"thumb,omitempty"`     // optional, animation thumbnail as defined by sender
	FileName string     `json:"file_name,omitempty"` // optional, original animation filename as defined by sender
	MIMEType string     `json:"mime_type,omitempty"` // optional, MIME type of the file as defined by sender
	FileSize int32      `json:"file_size,omitempty"` // optional, file size
}

type Sticker struct {
	FileID       string
	Width        int32
	Height       int32
	Thumb        *PhotoSize
	Emoji        string
	SetName      string
	MaskPosition *MaskPosition
	FileSize     int32
}

// MaskPosition describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	// The part of the face relative to which the mask should be placed.
	// One of “forehead”, “eyes”, “mouth”, or “chin”.
	Point string `json:"point,omitempty"`
	// Shift by X-axis measured in widths of the mask scaled to the face size, from left to right.
	// For example, choosing -1.0 will place mask just to the left of the default mask position.
	XShift float32 `json:"x_shift,omitempty"`
	// Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom.
	// For example, 1.0 will place the mask just below the default mask position.
	YShift float32 `json:"y_shift,omitempty"`
	Scale  float32 `json:"scale,omitempty"` // Mask scaling coefficient. For example, 2.0 means double size.
}

type Video struct {
	FileID   string     `json:"file_id,omitempty"`
	Width    int32      `json:"width,omitempty"`
	Height   int32      `json:"height,omitempty"`
	Duration int32      `json:"duration,omitempty"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	MIMEType string     `json:"mime_type,omitempty"`
	FileSize int32      `json:"file_size,omitempty"`
}

type Voice struct {
	FileID   string
	Duration int32
	MIMEType string
	FileSize int32
}

type VideoNote struct {
	FileID   string
	Length   int32
	Duration int32
	Thumb    *PhotoSize
	FileSize int32
}

type Contact struct {
	PhoneNumber string
	FirstName   string
	LastName    string
	UserID      int32
}

type Location struct {
	Longitude float32
	Latitude  float32
}

type Venue struct {
	Location     *Location
	Title        string
	Address      string
	FoursquareID string
}

type Invoice struct {
	Title          string
	Description    string
	StartParameter string
	Currency       string
	TotalAmount    int32
}

type SuccessfulPayment struct {
	Currency                string
	TotalAmount             int32
	InvoicePayload          string
	ShippingOptionID        string
	OrderInfo               *OrderInfo
	TelegramPaymentChargeID string
	ProviderPaymentChargeID string
}

type OrderInfo struct {
	Name            string
	PhoneNumber     string
	Email           string
	ShippingAddress *ShippingAddress
}

type ShippingAddress struct {
	CountryCode string
	State       string
	City        string
	StreetLine1 string
	StreetLine2 string
	PostCode    string
}
