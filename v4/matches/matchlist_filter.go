package matches

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/coltiebaby/glaw"
)

// Match Filter will help you sort through match history.
// To get a list of the IDs, you'll need to look at the [constants](*url)  supports
//
// url: https://developer.games.com/game-constants.html
type MatchFilter struct {
	Client glaw.ApiClient

	AccountID string
	Start     time.Time
	End       time.Time
	Champions []int
	Queues    []int
	Seasons   []int

	maxResults int
	page       int

	totalResults int
	begin        int
	end          int
}

func (filter *MatchFilter) Debug() {
	fmt.Printf("%+v\n", filter)
}

func getIndex(max, page int) (begin int, end int) {
	end = (page * max)
	begin = (end - max) + 1

	return begin, end
}

func NewMatchFilter(c glaw.ApiClient, start, end time.Time) (filter *MatchFilter) {
	filter = &MatchFilter{Client: c}

	filter.Start = start
	filter.End = end

	filter.SetMaxResults(20)

	return filter
}

// Confirm results are between 1 - 100
// TODO: log that we changes the result to something else
func (filter *MatchFilter) SetMaxResults(results int) {
	switch {
	case results <= 0:
		results = 1
	case results > 100:
		results = 100
	}

	filter.page = 0
	filter.maxResults = results
	filter.totalResults = results
}

func (filter *MatchFilter) Next() (matches MatchStorage, err error) {
	page := filter.page
	page++

	begin, end := getIndex(filter.maxResults, page)
	fmt.Println(begin, end)

	if begin > filter.totalResults {
		return matches, MaxResultsError
	} else {
		filter.page = page
	}

	filter.begin, filter.end = begin, end

	matches, err = GetMatchlists(filter.Client, filter.AccountID, filter.CreateValues())
	filter.totalResults = matches.TotalGames
	return matches, err
}

func (filter *MatchFilter) Prev() (matches MatchStorage, err error) {
	page := filter.page

	if page <= 1 {
		return matches, MinResultsError
	}

	page--
	filter.begin, filter.end = getIndex(filter.maxResults, page)
	filter.page = page

	matches, err = GetMatchlists(filter.Client, filter.AccountID, filter.CreateValues())
	filter.totalResults = matches.TotalGames
	return matches, err
}

func (filter *MatchFilter) GoTo(page int) (matches MatchStorage, err error) {
	wanted := (page-1)*filter.maxResults + 1
	if page < 1 || wanted > filter.totalResults {
		return matches, InvalidPageError
	}

	filter.begin, filter.end = getIndex(filter.maxResults, page)
	filter.page = page

	matches, err = GetMatchlists(filter.Client, filter.AccountID, filter.CreateValues())
	filter.totalResults = matches.TotalGames
	return matches, err
}

func (filter *MatchFilter) CreateValues() url.Values {
	params := url.Values{}

	if len(filter.Champions) > 0 {
		add(&params, `champion`, filter.Champions)
	}

	if len(filter.Queues) > 0 {
		add(&params, `queue`, filter.Queues)
	}

	if len(filter.Seasons) > 0 {
		add(&params, `season`, filter.Seasons)
	}

	if filter.begin != 0 && filter.end != 0 {
		params.Add(`startIndex`, strconv.Itoa(filter.begin))
		params.Add(`endIndex`, strconv.Itoa(filter.end))
	}

	// TODO: Not sure why these aren't working
	// from := timeToUnixMS(filter.Start)
	// to := timeToUnixMS(filter.End)

	// params.Add(`beginTime`, strconv.FormatInt(from, 10))
	// params.Add(`endTime`, strconv.FormatInt(to, 10))

	return params
}

func (filter *MatchFilter) TotalPages() int {
	return 0
}

// time to unix millisecond
func timeToUnixMS(t time.Time) int64 {
	return t.UnixNano() / int64(1000000)
}

func add(u *url.Values, key string, values []int) {
	for _, v := range values {
		u.Add(key, strconv.Itoa(v))
	}
}

var (
	MinResultsError  error = errors.New("Reached the min results")
	MaxResultsError  error = errors.New("Reached the max results")
	InvalidPageError error = errors.New("Page is invalid")
)
