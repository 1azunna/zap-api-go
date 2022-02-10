// Zed Attack Proxy (ZAP) and its related class files.
//
// ZAP is an HTTP/HTTPS proxy for assessing web application security.
//
// Copyright 2017 the ZAP development team
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// *** This file was automatically generated. ***
//

package zap

type Reveal struct {
	c *Client
}

// Tells if shows hidden fields and enables disabled fields
//
// This component is optional and therefore the API will only work if it is installed
func (r Reveal) Reveal() (map[string]interface{}, error) {
	return r.c.Request("reveal/view/reveal/", nil)
}

// Sets if shows hidden fields and enables disabled fields
//
// This component is optional and therefore the API will only work if it is installed
func (r Reveal) SetReveal(reveal string) (map[string]interface{}, error) {
	m := map[string]string{
		"reveal": reveal,
	}
	return r.c.Request("reveal/action/setReveal/", m)
}
