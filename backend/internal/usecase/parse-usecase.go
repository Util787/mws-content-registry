package usecase

func (p *ParseUseCase) ScrabData() error {
	YTVideosWithComments, err := p.YouTubeParseClient.ScrabVideosWithComments()
	if err != nil {
		return err
	}
	//...
}
