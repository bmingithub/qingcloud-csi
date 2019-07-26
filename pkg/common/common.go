/*
Copyright (C) 2018 Yunify, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this work except in compliance with the License.
You may obtain a copy of the License in the LICENSE file, or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package common

import (
	"fmt"
	"hash/fnv"
	"time"
)

// EntryFunction print timestamps
func EntryFunction(functionName string) (info string, hash string) {
	current := time.Now()
	hash = GenerateHashInEightBytes(current.UTC().String())
	return fmt.Sprintf("*************** enter %s at %s hash %s ***************", functionName,
		current.Format(DefaultTimeFormat), hash), hash
}

// ExitFunction print timestamps
func ExitFunction(functionName, hash string) (info string) {
	current := time.Now()
	return fmt.Sprintf("=============== exit %s at %s hash %s ===============", functionName,
		current.Format(DefaultTimeFormat), hash)
}

// GenerateRandIdSuffix generates a random resource id.
func GenerateHashInEightBytes(input string) string {
	h := fnv.New32a()
	h.Write([]byte(input))
	return fmt.Sprintf("%.8x", h.Sum32())
}
