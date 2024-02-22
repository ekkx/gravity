package gravity

type LoginData struct {
	Pnum  string   `json:"pnum"`
	Token string   `json:"token"`
	User  UserInfo `json:"user"`
}

type UserInfo struct {
	AnchorLevel         string    `json:"anchor_level"`
	AnchorLevelBadgeURL string    `json:"anchor_level_badge_url"`
	AnchorLevelURL      string    `json:"anchor_level_url"`
	BadgeList           []Badge   `json:"badge_list"`
	BadgeListWall       []Badge   `json:"badge_list_wall"`
	Birthdate           int       `json:"birthdate"`
	Center2DBackground  string    `json:"2d_bg_img"`
	CrownRank           int       `json:"crown_rank"`
	CustomID            string    `json:"custom_id"`
	EncrypUserID        string    `json:"encryp_user_id"`
	EncryptEmail        string    `json:"encrypt_email"`
	Ext                 ExtraInfo `json:"ext"`
	Gender              int       `json:"gender"`
	IsOld3D             int       `json:"is_3D"`
	LoginType           int       `json:"login_type"`
	Name                string    `json:"name"`
	PlatoGender         int       `json:"plato_gender"`
	Portrait            string    `json:"portrait"`
	PortraitFrame       string    `json:"portrait_frame"`
	Profile             string    `json:"profile"`
	RegisterTime        int       `json:"db_time"`
	RegisterVersion     string    `json:"app_version"`
	ShouldHidden        int       `json:"should_hidden"`
	UserCountry         string    `json:"user_country"`
	UserID              string    `json:"user_id"`
	VirtualTypeV2       int       `json:"virtual_type_v2"`
	VoiceCreatorStatus  int       `json:"voice_creator_status"`

	BGImg               string `json:"bg_img"`
	CompressPortrait    string `json:"compress_portrait"`
	CreatorStatus       int    `json:"creator_status"`
	GroupLevel          int    `json:"group_level"`
	GuardStatus         int    `json:"guard_status"`
	IsMinor             int    `json:"is_minor"`
	LimitSerialNum      int    `json:"limit_serial_num"`
	LoginTime           int    `json:"login_time"`
	LowerCustomID       string `json:"lower_custom_id"`
	PortraitColor       string `json:"portrait_color"`
	PortraitFrameExpire int    `json:"portrait_frame_expire"`
	PortraitFrameID     int    `json:"portrait_frame_id"`
	PortraitImage       string `json:"portrait_image"`
	RefreshToken        string `json:"refres_token"`
	Status              int    `json:"status"`
	VirtualImage        string `json:"virtual_image"`
}

type Badge struct {
	BadgeType int    `json:"badge_type"`
	Desc      string `json:"desc"`
	DialogURL string `json:"dialog_url"`
	IntroURL  string `json:"intro_url"`
	Name      string `json:"name"`
	Status    int    `json:"status"`
	Title     string `json:"title"`
	IconURL   string `json:"icon_url"`
}

type ExtraInfo struct {
	CenterBackground        CenterBackgroundData      `json:"center_background"`
	CharacterAttribute      string                    `json:"character_attribute"`
	Charm                   Charm                     `json:"charm_info"`
	CreatorStatus           int                       `json:"creator_status"`
	EntryEffects            EntrySpecialEffectData    `json:"entry_effects"`
	Exam                    ExamBean                  `json:"exam"`
	Family                  Family                    `json:"family"`
	Fans                    int                       `json:"fans"`
	Follow                  int                       `json:"follow"`
	Friends                 int                       `json:"friends"`
	GoodAnchorStatus        int                       `json:"voice_creator_status"`
	GuardStatus             int                       `json:"guard_status"`
	HasDeepFriend           int                       `json:"is_widget"`
	HaveVoiceUserTag        int                       `json:"have_voice_user_tag"`
	InterestTag             []InterestTag             `json:"interest_tag"`
	IsIn3D                  int                       `json:"is_in_3D"`
	IsPlato                 int                       `json:"is_plato"`
	MessageBubble           UserMessageBubble         `json:"message_bubble"`
	MicrophoneInfo          VoiceMicrophoneData       `json:"microphone_info"`
	Mute                    int                       `json:"behavior"`
	OnlineStatus            int                       `json:"online_status"`
	Pendant                 DecorationPendantData     `json:"feed_pendant"`
	PortraitFrame           PortraitFrame             `json:"portrait_frame"`
	PortraitInfo            PortraitData              `json:"portrait_info"`
	ProfileCard             DecorationProfileCardData `json:"profile_card_decoration"`
	RecallUser              int                       `json:"recall_user"`
	Relation                int                       `json:"relation"`
	RelationType            int                       `json:"relation_type"`
	Setting                 Setting                   `json:"setting"`
	ShowLotteryNewUserGuide int                       `json:"user_is_lottery_new"`
	SnsInfo                 SnsInfo                   `json:"sns_info"`
	SquareUpdateTime        int64                     `json:"square_update_time"`
	UserAuth                UserAuth                  `json:"user_authority"`
	UserTagData             []UserTagData             `json:"user_tag_list"`
	UserStatistic           UserFans                  `json:"user_statistic"`
	VIP                     VIPInfo                   `json:"vip"`
	VirtualBG               string                    `json:"virtual_bg"`
	VirtualBGImage          string                    `json:"bg_img"`
	VirtualImage            string                    `json:"virtualImage"`
	VoiceChannelId          string                    `json:"voice_channel_id"`
	VoiceOnline             int                       `json:"voice_online"`
	VoiceUserTag            UserTagData               `json:"voice_user_tag"`
}

type VIPInfo struct {
	ExpireDate string `json:"expire_date"`
	ExpireTime int64  `json:"expire_time"`
	IsVIP      int    `json:"is_vip"`
	VIPLevel   int    `json:"vip_level"`
}

type UserFans struct {
	FansCount   int `json:"fans_count"`
	FollowCount int `json:"follow_count"`
	VoteCount   int `json:"vote_count"`
}

type UserTagData struct{}

type UserAuth struct{}

type SnsInfo struct{}

type Setting struct{}

type DecorationProfileCardData struct{}

type PortraitData struct{}

type PortraitFrame struct {
	ExpireTime         int64  `json:"expire_time"`
	FrameID            int    `json:"frame_id"`
	FrameImage         string `json:"frame_image"`
	FrameType          int    `json:"frame_type"`
	HideStyle          int    `json:"is_hide"`
	LimitNum           int    `json:"limit_num"`
	LimitSerialNum     int    `json:"limit_serial_num"`
	LimitStatus        int    `json:"limit_status"`
	PortraitFrameLevel string `json:"portrait_frame_level"`
	PortraitFrameName  string `json:"portrait_frame_name"`
}

type DecorationPendantData struct {
	ActivityBeginTime int64  `json:"activity_begin_time"`
	ActivityEndTime   int64  `json:"activity_end_time"`
	ActivitySource    int64  `json:"activity_source"`
	AndroidURL        string `json:"android_url"`
	Count             int    `json:"count"`
	Desc              string `json:"desc"`
	ExpireTime        int64  `json:"expire_time"`
	GoodsType         int    `json:"goods_type"`
	ID                int    `json:"id"`
	IsGet             int    `json:"is_get"`
	Level             string `json:"level"`
	Name              string `json:"name"`
	PendantUrl        string `json:"pendant_url"`
	PreviewURL        string `json:"preview_url"`
	ValidDays         string `json:"valid_days"`
	VipLevel          int    `json:"vip_level"`
}

type VoiceMicrophoneData struct {
	Color     string `json:"color"`
	GetStatus int    `json:"is_get"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	VIPLevel  int    `json:"vip_level"`
}

type UserMessageBubble struct {
	BubbleID   int   `json:"bubble_id"`
	ExpireTime int64 `json:"expire_time"`
}

type InterestTag struct {
	Title string `json:"title"`
	Type  int    `json:"type"`
}

type Family struct {
	BadgeText string `json:"badge_text"`
	BadgeURL  string `json:"badge_url"`
	FamilyID  int    `json:"family_id"`
	Level     int    `json:"level"`
}

type CenterBackgroundData struct {
	ActivityBeginTime int64  `json:"activity_begin_time"`
	ActivityEndTime   int64  `json:"activity_end_time"`
	AndroidUrl        string `json:"android_url"`
	Anim              string `json:"anim"`
	Count             int    `json:"count"`
	ExpireTime        int64  `json:"expire_time"`
	GetStatus         int    `json:"is_get"`
	GoodsType         int    `json:"goods_type"`
	ID                int    `json:"id"`
	Image             string `json:"image"`
	Level             string `json:"level"`
	Name              string `json:"name"`
	Preview           string `json:"preview"`
	TagType           string `json:"tag_type"`
	ValidDays         int    `json:"valid_days"`
	VipLevel          int    `json:"vip_level"`
}

type Charm struct {
	CanCharmLike int   `json:"can_charm_like"`
	Level        int   `json:"charm_level"`
	UpgradeNeed  int64 `json:"upgrade_need_value"`
	Value        int64 `json:"charm_value"`
}

type EntrySpecialEffectData struct {
	ActivityBeginTime int64  `json:"activity_begin_time"`
	ActivityEndTime   int64  `json:"activity_end_time"`
	ActivitySource    string `json:"activity_source"`
	AndroidURL        string `json:"android_url"`
	ColorValue        string `json:"color_value"`
	Count             int    `json:"count"`
	EntryDesc         string `json:"entry_desc"`
	EntryURL          string `json:"entry_url"`
	ExpireTime        int64  `json:"expire_time"`
	GoodsType         int    `json:"goods_type"`
	ID                int    `json:"id"`
	IsGet             int    `json:"is_get"`
	Level             string `json:"level"`
	Name              string `json:"name"`
	PreviewURL        string `json:"preview_url"`
	ValidDays         string `json:"valid_days"`
	VideoURL          string `json:"video_url"`
	VipLevel          int    `json:"vip_level"`
}

type ExamBean struct {
	AIExam   AIExam       `json:"ai_exam"`
	External ExternalBean `json:"external"`
	Internal InternalBean `json:"internal"`
	MBTIExam MBTIExam     `json:"mbti_exam"`
}

type ExternalBean struct {
	Character   string `json:"character"`
	Detail      string `json:"detail"`
	Explanation string `json:"explanation"`
}

type InternalBean struct {
	Character   string `json:"character"`
	Detail      string `json:"detail"`
	Explanation string `json:"explanation"`
}

type AIExam struct {
	Character   string `json:"character"`
	Explanation string `json:"explanation"`
	IsFinish    bool   `json:"is_finish"`
}

type MBTIExam struct {
	CharacterAttribute string `json:"character_attribute"`
	CompleteStatus     int    `json:"complete_status"`
}
