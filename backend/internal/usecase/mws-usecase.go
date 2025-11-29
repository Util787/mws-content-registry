package usecase

import (
	"strconv"
	"strings"
	"time"

	"github.com/Util787/mws-content-registry/internal/models"
)

const YTTimeFormat = time.RFC3339

func (m *MWSTablesUsecase) TakeRecords(pageNum int, pageSize int, sort []map[string]string, recordId string, fields []string) ([]models.MWSTableRecord, error) {
	return m.MWSTablesClient.TakeRecords(pageNum, pageSize, sort, recordId, fields)
}

func (m *MWSTablesUsecase) AddLLMContentAnalyze(recordId string) error {
	records, err := m.MWSTablesClient.TakeRecords(0, 0, nil, recordId, nil)
	if err != nil {
		return err
	}

	rec := records[0]

	analyzeData, err := m.LLMClient.GenerateContentAnalyze(rec)
	if err != nil {
		return err
	}

	var recUpdate models.MWSTableUpdateRecord
	recUpdate.RecordID = rec.RecordID
	recUpdate.Fields.Topic = &analyzeData.Topic
	recUpdate.Fields.Recomendations = &analyzeData.Recomendations
	recUpdate.Fields.CommentsSummary = &analyzeData.CommentsSummary
	recUpdate.Fields.CommentsTone = &analyzeData.CommentsTone

	now := int(time.Now().UnixMilli())
	recUpdate.Fields.AiAnalyzeDate = &now

	err = m.MWSTablesClient.UpdateRecords([]models.MWSTableUpdateRecord{recUpdate})
	if err != nil {
		return err
	}
	return nil
}

func (m *MWSTablesUsecase) AddYTVideoByURL(url string) error {
	vid, err := m.YouTubeParseClient.ScrabVideoByURL(url)
	if err != nil {
		return err
	}

	t, err := time.Parse(YTTimeFormat, vid.Video.PublishedAt)
	if err != nil {
		return err
	}

	pub := t.UnixMilli()

	comments := fmtYTComments(vid.Comments)

	rec := models.MWSTableNewRecord{
		Fields: models.MWSTableAddRecordFields{
			URL:             &vid.Video.VideoURL,
			Views:           &vid.Video.ViewsCount,
			Topic:           nil, // to be filled by LLM
			Likes:           &vid.Video.LikesCount,
			Comments:        &comments,
			CommentsCount:   &vid.Video.CommentsCount,
			Description:     &vid.Video.Description,
			Author:          &vid.Video.ChannelTitle,
			Recomendations:  nil, // to be filled by LLM
			CommentsSummary: nil, // to be filled by LLM
			CommentsTone:    nil, // to be filled by LLM
		},
	}

	// time validation
	if pub == 0 || pub == 946674000000 {
		rec.Fields.PublishedAt = nil
	} else {
		rec.Fields.PublishedAt = &pub
	}

	err = m.MWSTablesClient.AddRecords([]models.MWSTableNewRecord{rec})
	if err != nil {
		return err
	}

	return nil
}

func (m *MWSTablesUsecase) AddRecentYTVideos() error {
	vids, err := m.YouTubeParseClient.ScrabVideosWithComments()
	if err != nil {
		return err
	}

	var recs = make([]models.MWSTableNewRecord, 0, len(vids))

	for _, vid := range vids {
		t, err := time.Parse(YTTimeFormat, vid.Video.PublishedAt)
		if err != nil {
			return err
		}

		pub := t.UnixMilli()

		comments := fmtYTComments(vid.Comments)

		rec := models.MWSTableNewRecord{
			Fields: models.MWSTableAddRecordFields{
				URL:             &vid.Video.VideoURL,
				PublishedAt:     &pub,
				Views:           &vid.Video.ViewsCount,
				Topic:           nil, // to be filled by LLM
				Likes:           &vid.Video.LikesCount,
				Comments:        &comments,
				CommentsCount:   &vid.Video.CommentsCount,
				Description:     &vid.Video.Description,
				Author:          &vid.Video.ChannelTitle,
				Recomendations:  nil, // to be filled by LLM
				CommentsSummary: nil, // to be filled by LLM
				CommentsTone:    nil, // to be filled by LLM
			},
		}

		recs = append(recs, rec)
	}

	err = m.MWSTablesClient.AddRecords(recs)
	if err != nil {
		return err
	}

	return nil
}

// formatting comments in format: "<comment_text>(likes:<likes_count>)"
func fmtYTComments(rawComments []models.YTComment) string {
	sb := strings.Builder{}
	for _, cmnt := range rawComments {
		sb.WriteString(cmnt.Text)
		sb.WriteString("(likes:")
		sb.WriteString(strconv.Itoa(int(cmnt.Likes)))
		sb.WriteString(")")
	}
	return sb.String()
}
