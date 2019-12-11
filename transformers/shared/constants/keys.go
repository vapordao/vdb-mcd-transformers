// VulcanizeDB
// Copyright © 2019 Vulcanize

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package constants

import (
	"github.com/makerdao/vulcanizedb/libraries/shared/factories/event"
	"github.com/makerdao/vulcanizedb/libraries/shared/storage/utils"
)

const (
	BidId     utils.Key = "bid_id"
	Cdpi      utils.Key = "cdpi"
	Flip      utils.Key = "flip"
	Guy       utils.Key = "guy"
	Ilk       utils.Key = "ilk"
	Owner     utils.Key = "owner"
	Timestamp utils.Key = "timestamp"
)

// TODO remove after transition to ColumnName
type ForeignKeyField string

const (
	AddressFK ForeignKeyField = "address_id"
	IlkFK     ForeignKeyField = "ilk_id"
	UrnFK     ForeignKeyField = "urn_id"
)

const (
	HeaderFK = "header_id"
	LogFK    = "log_id"
)

const (
	AddressColumn event.ColumnName = "address_id"
	ArtColumn     event.ColumnName = "art"
	BidColumn     event.ColumnName = "bid"
	BidIDColumn   event.ColumnName = "bid_id"
	DartColumn    event.ColumnName = "dart"
	DataColumn    event.ColumnName = "data"
	DinkColumn    event.ColumnName = "dink"
	DstColumn     event.ColumnName = "dst"
	EraColumn     event.ColumnName = "era"
	FlipColumn    event.ColumnName = "flip"
	GalColumn     event.ColumnName = "gal"
	IlkColumn     event.ColumnName = "ilk_id"
	InkColumn     event.ColumnName = "ink"
	LotColumn     event.ColumnName = "lot"
	PipColumn     event.ColumnName = "pip"
	RadColumn     event.ColumnName = "rad"
	RateColumn    event.ColumnName = "rate"
	SpotColumn    event.ColumnName = "spot"
	SrcColumn     event.ColumnName = "src"
	TabColumn     event.ColumnName = "tab"
	UColumn       event.ColumnName = "u"
	UrnColumn     event.ColumnName = "urn_id"
	UsrColumn     event.ColumnName = "usr"
	VColumn       event.ColumnName = "v"
	WColumn       event.ColumnName = "w"
	WadColumn     event.ColumnName = "wad"
	WhatColumn    event.ColumnName = "what"
	ValueColumn   event.ColumnName = "value"
)
