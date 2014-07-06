/*
Copyright 2013 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package reddit

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

type Subreddit struct {
	Name string
	Items []Item
}

type Item struct {
	Title    string
	URL      string
	Comments int `json:"num_comments"`
}

type Comment struct {

}

type Submission struct {
	Title string
	LinkURL string
	Comments []Comment
}

//func GetSubmission(url string) (Submission, error) 

func GetSubreddit(name string) (Subreddit, error) {

	var subreddit Subreddit

	url := fmt.Sprintf("http://reddit.com/r/%s.json", name)

	r, err := Get(url)

	if err != nil {
		return subreddit, err
	}

	items := make([]Item, len(r.Data.Children))
	for i, child := range r.Data.Children {
		items[i] = child.Data
	}

	

	subreddit.Name = name
	subreddit.Items = items

	return subreddit, nil
}



func (i Item) String() string {
	com := ""
	switch i.Comments {
	case 0:
		// nothing
	case 1:
		com = " (1 comment)"
	default:
		com = fmt.Sprintf(" (%d comments)", i.Comments)
	}
	return fmt.Sprintf("%s%s\n%s", i.Title, com, i.URL)
}


func Get(request string) (*Response, error) {
	
	resp, err := http.Get(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	r := new(Response)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	
	return r, nil
}
