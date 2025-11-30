package parseclients

import (
	"fmt"
	"strings"

	"github.com/Util787/mws-content-registry/internal/common"
	"github.com/Util787/mws-content-registry/internal/models"
)

func (ytpc *YouTubeParseClient) ScrabVideosWithComments() ([]models.YTVideosWithComments, error) {
	log := ytpc.log.With("op", common.GetOperationName())

	call := ytpc.ytService.Videos.List([]string{"snippet", "statistics"}).Chart(ytpc.chart).RegionCode(ytpc.regionCode).MaxResults(ytpc.videosLimit)
	response, err := call.Do()
	if err != nil {
		log.Error("Failed to fetch trending videos", "error", err)
		return nil, err
	}

	var videosWithComments = make([]models.YTVideosWithComments, 0, len(response.Items))

	for _, video := range response.Items {
		commentsCall := ytpc.ytService.CommentThreads.List([]string{"snippet"}).VideoId(video.Id).MaxResults(ytpc.commentsLimit).Order("relevance")
		commentsResponse, err := commentsCall.Do()
		if err != nil {
			log.Error("Failed to fetch comments for video", "videoId", video.Id, "error", err)
			continue
		}

		var comments = make([]models.YTComment, 0, len(commentsResponse.Items))

		for _, comment := range commentsResponse.Items {
			comments = append(comments, models.YTComment{
				Text:  comment.Snippet.TopLevelComment.Snippet.TextOriginal,
				Likes: comment.Snippet.TopLevelComment.Snippet.LikeCount,
			})
		}

		strbuilder := strings.Builder{}
		strbuilder.WriteString("https://www.youtube.com/watch?v=")
		strbuilder.WriteString(video.Id)

		videosWithComments = append(videosWithComments, models.YTVideosWithComments{
			Video: models.YTVideo{
				VideoURL:      strbuilder.String(),
				ChannelTitle:  video.Snippet.ChannelTitle,
				Title:         video.Snippet.Title,
				Description:   video.Snippet.Description,
				LikesCount:    video.Statistics.LikeCount,
				ViewsCount:    video.Statistics.ViewCount,
				CommentsCount: video.Statistics.CommentCount,
				PublishedAt:   video.Snippet.PublishedAt},
			Comments: comments,
		})
	}

	return videosWithComments, nil
}

func (ytpc *YouTubeParseClient) ScrabVideoByURL(videoURL string) (*models.YTVideosWithComments, error) {
	log := ytpc.log.With("op", common.GetOperationName())

	// извлекаем videoId из URL
	parts := strings.Split(videoURL, "v=")
	if len(parts) < 2 {
		log.Error("Invalid YouTube URL", "url", videoURL)
		return nil, fmt.Errorf("invalid YouTube URL: %s", videoURL)
	}
	videoID := strings.Split(parts[1], "&")[0]

	// запрос информации о видео
	call := ytpc.ytService.Videos.List([]string{"snippet", "statistics"}).Id(videoID)
	response, err := call.Do()
	if err != nil {
		log.Error("Failed to fetch video info", "videoId", videoID, "error", err)
		return nil, err
	}
	if len(response.Items) == 0 {
		log.Error("Video not found", "videoId", videoID)
		return nil, fmt.Errorf("video not found: %s", videoID)
	}
	video := response.Items[0]

	// запрос комментариев
	commentsCall := ytpc.ytService.CommentThreads.List([]string{"snippet"}).VideoId(videoID).MaxResults(ytpc.commentsLimit).Order("relevance")
	commentsResponse, err := commentsCall.Do()
	if err != nil {
		log.Error("Failed to fetch comments for video", "videoId", videoID, "error", err)
		return nil, err
	}

	comments := make([]models.YTComment, 0, len(commentsResponse.Items))
	for _, comment := range commentsResponse.Items {
		comments = append(comments, models.YTComment{
			Text:  comment.Snippet.TopLevelComment.Snippet.TextOriginal,
			Likes: comment.Snippet.TopLevelComment.Snippet.LikeCount,
		})
	}

	return &models.YTVideosWithComments{
		Video: models.YTVideo{
			VideoURL:      videoURL,
			ChannelTitle:  video.Snippet.ChannelTitle,
			Title:         video.Snippet.Title,
			Description:   video.Snippet.Description,
			LikesCount:    video.Statistics.LikeCount,
			ViewsCount:    video.Statistics.ViewCount,
			CommentsCount: video.Statistics.CommentCount,
			PublishedAt:   video.Snippet.PublishedAt,
		},
		Comments: comments,
	}, nil
}
