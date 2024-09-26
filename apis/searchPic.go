package apis

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func SearchPic(searchText string) []byte {
	searchText = url.QueryEscape(searchText)

	url := fmt.Sprintf("https://contextualwebsearch-websearch-v1.p.rapidapi.com/api/Search/ImageSearchAPI?q=%s&pageNumber=1&pageSize=10&autoCorrect=true", searchText)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "926b24827cmshed83bbafe91e8dcp1e5886jsn04d052c0f2ca")
	req.Header.Add("X-RapidAPI-Host", "contextualwebsearch-websearch-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)
	if res == nil || res.Body == nil {
		return nil
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	return body
}
