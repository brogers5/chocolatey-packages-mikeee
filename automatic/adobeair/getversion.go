package main

import (
	"fmt"
	"log"

	"github.com/mikeee/ChocoPackages/helpers/gohelpers"
)

// Copyright 2022 mikeee
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

func main() {
	url := "https://airsdk.harman.com/runtime"
	regexStr := "AIR\\sruntime\\s-\\sversion\\s[\\d\\.]+"
	matchIndex := 0

	rawString, err := gohelpers.RegexMatch(url, regexStr, matchIndex)

	if err != nil {
		log.Fatalf("failed to get latest version: %v", err)
	}

	fmt.Print(rawString)
}
