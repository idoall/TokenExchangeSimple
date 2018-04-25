// Copyright 2016 mshk.top, lion@mshk.top
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

// HmacSigner uses HMAC SHA256 for signing payloads.
type HuobiHmacSigner struct {
	Key []byte
}

// Sign signs provided payload and returns encoded string sum.
func (hs *HuobiHmacSigner) Sign(payload []byte) string {
	mac := hmac.New(sha256.New, []byte(hs.Key))
	mac.Write(payload)
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
