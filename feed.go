package gravity

import "fmt"

type FeedService struct {
	g *Gravity
}

func newFeedService(g *Gravity) *FeedService {
	return &FeedService{
		g: g,
	}
}

type RecommendFeedListParams struct {
	SID    string `json:"session_id"`
	LastID string `json:"last_id"`
	LogID  int    `json:"log_id"`
}

func (s *FeedService) RecommendFeedList() (st interface{}, err error) {
	resp, err := s.g.requestWithQuery("GET", EndpointFeedRecommendFeedListV2, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp)

	return resp, nil
}
