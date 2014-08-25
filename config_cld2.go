//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

// +build cld2 full

package bleve

import (
	// cld2 token filter
	_ "github.com/couchbaselabs/bleve/analysis/token_filters/cld2"

	// detect language analyzer
	_ "github.com/couchbaselabs/bleve/analysis/analyzers/detect_lang_analyzer"
)
