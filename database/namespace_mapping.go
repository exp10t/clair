// Copyright 2015 clair authors
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

package database

// DebianReleasesMapping translates Debian code names and class names to version numbers
var DebianReleasesMapping = map[string]string{
	// Code names
	"squeeze": "6",
	"wheezy":  "7",
	"jessie":  "8",
	"stretch": "9",
	"buster":  "10",
	"sid":     "unstable",

	// Class names
	"oldoldstable": "7",
	"oldstable":    "8",
	"stable":       "9",
	"testing":      "10",
	"unstable":     "unstable",
}

// UbuntuReleasesMapping translates Ubuntu code names to version numbers
var UbuntuReleasesMapping = map[string]string{
	"precise": "12.04",
	"quantal": "12.10",
	"raring":  "13.04",
	"trusty":  "14.04",
	"utopic":  "14.10",
	"vivid":   "15.04",
	"wily":    "15.10",
	"xenial":  "16.04",
	"yakkety": "16.10",
	"zesty":   "17.04",
	"artful":  "17.10",
}
