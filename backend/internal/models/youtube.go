package models

type YTVideosWithComments struct {
	Video    YTVideo
	Comments []YTComment
}

type YTVideo struct {
	VideoURL      string
	ChannelTitle     string
	Title         string
	Description   string
	LikesCount    uint64
	ViewsCount    uint64
	CommentsCount uint64
	PublishedAt   string
}

type YTComment struct {
	Text  string
	Likes int64
}
