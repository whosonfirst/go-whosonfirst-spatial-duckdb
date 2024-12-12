// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatbuf

import "strconv"

type TimeUnit int16

const (
	TimeUnitSECOND      TimeUnit = 0
	TimeUnitMILLISECOND TimeUnit = 1
	TimeUnitMICROSECOND TimeUnit = 2
	TimeUnitNANOSECOND  TimeUnit = 3
)

var EnumNamesTimeUnit = map[TimeUnit]string{
	TimeUnitSECOND:      "SECOND",
	TimeUnitMILLISECOND: "MILLISECOND",
	TimeUnitMICROSECOND: "MICROSECOND",
	TimeUnitNANOSECOND:  "NANOSECOND",
}

var EnumValuesTimeUnit = map[string]TimeUnit{
	"SECOND":      TimeUnitSECOND,
	"MILLISECOND": TimeUnitMILLISECOND,
	"MICROSECOND": TimeUnitMICROSECOND,
	"NANOSECOND":  TimeUnitNANOSECOND,
}

func (v TimeUnit) String() string {
	if s, ok := EnumNamesTimeUnit[v]; ok {
		return s
	}
	return "TimeUnit(" + strconv.FormatInt(int64(v), 10) + ")"
}
