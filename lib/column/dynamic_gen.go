// Licensed to ClickHouse, Inc. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. ClickHouse, Inc. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by make codegen DO NOT EDIT.
// source: lib/column/codegen/dynamic.tpl

package column

import (
	"database/sql"
	"encoding/json"
	"github.com/ClickHouse/ch-go/proto"
	"github.com/google/uuid"
	"github.com/paulmach/orb"
	"time"
)

// inferClickHouseTypeFromGoType takes a Go interface{} and converts it to a ClickHouse type.
// Returns empty string if type was not matched.
// This is best effort and does not work for all types.
// Optimally, users should provide a type using DynamicWithType.
func inferClickHouseTypeFromGoType(v any) string {
	switch v.(type) {
	case float32:
		return "Float32"
	case *float32:
		return "Float32"
	case float64:
		return "Float64"
	case *float64:
		return "Float64"
	case int8:
		return "Int8"
	case *int8:
		return "Int8"
	case int16:
		return "Int16"
	case *int16:
		return "Int16"
	case int32:
		return "Int32"
	case *int32:
		return "Int32"
	case int64:
		return "Int64"
	case *int64:
		return "Int64"
	case uint8:
		return "UInt8"
	case *uint8:
		return "UInt8"
	case uint16:
		return "UInt16"
	case *uint16:
		return "UInt16"
	case uint32:
		return "UInt32"
	case *uint32:
		return "UInt32"
	case uint64:
		return "UInt64"
	case *uint64:
		return "UInt64"
	case string:
		return "String"
	case *string:
		return "String"
	case json.RawMessage:
		return "String"
	case *json.RawMessage:
		return "String"
	case sql.NullString:
		return "String"
	case *sql.NullString:
		return "String"
	case bool:
		return "Bool"
	case *bool:
		return "Bool"
	case sql.NullBool:
		return "Bool"
	case *sql.NullBool:
		return "Bool"
	case time.Time:
		return "DateTime64(3)"
	case *time.Time:
		return "DateTime64(3)"
	case sql.NullTime:
		return "DateTime64(3)"
	case *sql.NullTime:
		return "DateTime64(3)"
	case uuid.UUID:
		return "UUID"
	case *uuid.UUID:
		return "UUID"
	case proto.IPv6:
		return "IPv6"
	case *proto.IPv6:
		return "IPv6"
	case orb.MultiPolygon:
		return "MultiPolygon"
	case *orb.MultiPolygon:
		return "MultiPolygon"
	case orb.Point:
		return "Point"
	case *orb.Point:
		return "Point"
	case orb.Polygon:
		return "Polygon"
	case *orb.Polygon:
		return "Polygon"
	case orb.Ring:
		return "Ring"
	case *orb.Ring:
		return "Ring"
	case []float32:
		return "Array(Float32)"
	case []*float32:
		return "Array(Float32)"
	case []float64:
		return "Array(Float64)"
	case []*float64:
		return "Array(Float64)"
	case []int8:
		return "Array(Int8)"
	case []*int8:
		return "Array(Int8)"
	case []int16:
		return "Array(Int16)"
	case []*int16:
		return "Array(Int16)"
	case []int32:
		return "Array(Int32)"
	case []*int32:
		return "Array(Int32)"
	case []int64:
		return "Array(Int64)"
	case []*int64:
		return "Array(Int64)"
	case []*uint8:
		return "Array(UInt8)"
	case []uint16:
		return "Array(UInt16)"
	case []*uint16:
		return "Array(UInt16)"
	case []uint32:
		return "Array(UInt32)"
	case []*uint32:
		return "Array(UInt32)"
	case []uint64:
		return "Array(UInt64)"
	case []*uint64:
		return "Array(UInt64)"
	case []string:
		return "Array(String)"
	case []*string:
		return "Array(String)"
	case []json.RawMessage:
		return "Array(String)"
	case []*json.RawMessage:
		return "Array(String)"
	case []sql.NullString:
		return "Array(String)"
	case []*sql.NullString:
		return "Array(String)"
	case []bool:
		return "Array(Bool)"
	case []*bool:
		return "Array(Bool)"
	case []sql.NullBool:
		return "Array(Bool)"
	case []*sql.NullBool:
		return "Array(Bool)"
	case []time.Time:
		return "Array(DateTime64(3))"
	case []*time.Time:
		return "Array(DateTime64(3))"
	case []sql.NullTime:
		return "Array(DateTime64(3))"
	case []*sql.NullTime:
		return "Array(DateTime64(3))"
	case []uuid.UUID:
		return "Array(UUID)"
	case []*uuid.UUID:
		return "Array(UUID)"
	case []proto.IPv6:
		return "Array(IPv6)"
	case []*proto.IPv6:
		return "Array(IPv6)"
	case []orb.MultiPolygon:
		return "Array(MultiPolygon)"
	case []*orb.MultiPolygon:
		return "Array(MultiPolygon)"
	case []orb.Point:
		return "Array(Point)"
	case []*orb.Point:
		return "Array(Point)"
	case []orb.Polygon:
		return "Array(Polygon)"
	case []*orb.Polygon:
		return "Array(Polygon)"
	case []orb.Ring:
		return "Array(Ring)"
	case []*orb.Ring:
		return "Array(Ring)"
	case map[string]float32:
		return "Map(String, Float32)"
	case map[string]float64:
		return "Map(String, Float64)"
	case map[string]int8:
		return "Map(String, Int8)"
	case map[string]int16:
		return "Map(String, Int16)"
	case map[string]int32:
		return "Map(String, Int32)"
	case map[string]int64:
		return "Map(String, Int64)"
	case map[string]uint8:
		return "Map(String, UInt8)"
	case map[string]uint16:
		return "Map(String, UInt16)"
	case map[string]uint32:
		return "Map(String, UInt32)"
	case map[string]uint64:
		return "Map(String, UInt64)"
	case map[string]string:
		return "Map(String, String)"
	case map[string]json.RawMessage:
		return "Map(String, String)"
	case map[string]sql.NullString:
		return "Map(String, String)"
	case map[string]bool:
		return "Map(String, Bool)"
	case map[string]sql.NullBool:
		return "Map(String, Bool)"
	case map[string]time.Time:
		return "Map(String, DateTime64(3))"
	case map[string]sql.NullTime:
		return "Map(String, DateTime64(3))"
	case map[string]uuid.UUID:
		return "Map(String, UUID)"
	case map[string]proto.IPv6:
		return "Map(String, IPv6)"
	case map[string]orb.MultiPolygon:
		return "Map(String, MultiPolygon)"
	case map[string]orb.Point:
		return "Map(String, Point)"
	case map[string]orb.Polygon:
		return "Map(String, Polygon)"
	case map[string]orb.Ring:
		return "Map(String, Ring)"
	default:
		return ""
	}
}
