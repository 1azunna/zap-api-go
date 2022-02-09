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

type Soap struct {
	c *Client
}

// Import a WSDL definition from local file.
//
// This component is optional and therefore the API will only work if it is installed
func (s Soap) ImportFile(file string) (map[string]interface{}, error) {
	m := map[string]string{
		"file": file,
	}
	return s.c.Request("soap/action/importFile/", m)
}

// Import a WSDL definition from a URL.
//
// This component is optional and therefore the API will only work if it is installed
func (s Soap) ImportUrl(url string) (map[string]interface{}, error) {
	m := map[string]string{
		"url": url,
	}
	return s.c.Request("soap/action/importUrl/", m)
}

