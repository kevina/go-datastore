// Copyright 2013 Matthew Baird
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import (
	"encoding/json"
	"fmt"
	"github.com/ipfs/go-datastore/Godeps/_workspace/src/github.com/mattbaird/elastigo/api"
)

// MGet allows the caller to get multiple documents based on an index, type (optional) and id (and possibly routing).
// The response includes a docs array with all the fetched documents, each element similar in structure to a document
// provided by the get API.
// see http://www.elasticsearch.org/guide/reference/api/multi-get.html
func MGet(pretty bool, index string, _type string, mgetRequest MGetRequestContainer) (MGetResponseContainer, error) {
	var url string
	var retval MGetResponseContainer
	if len(index) <= 0 {
		url = fmt.Sprintf("/_mget?%s", api.Pretty(pretty))
	}
	if len(_type) > 0 && len(index) > 0 {
		url = fmt.Sprintf("/%s/%s/_mget?%s", index, _type, api.Pretty(pretty))
	} else if len(index) > 0 {
		url = fmt.Sprintf("/%s/_mget?%s", index, api.Pretty(pretty))
	}
	body, err := api.DoCommand("GET", url, nil)
	if err != nil {
		return retval, err
	}
	if err == nil {
		// marshall into json
		jsonErr := json.Unmarshal(body, &retval)
		if jsonErr != nil {
			return retval, jsonErr
		}
	}
	fmt.Println(body)
	return retval, err
}

type MGetRequestContainer struct {
	Docs []MGetRequest `json:"docs"`
}

type MGetRequest struct {
	Index  string   `json:"_index"`
	Type   string   `json:"_type"`
	ID     string   `json:"_id"`
	IDS    []string `json:"_ids,omitempty"`
	Fields []string `json:"fields,omitempty"`
}

type MGetResponseContainer struct {
	Docs []api.BaseResponse `json:"docs"`
}
