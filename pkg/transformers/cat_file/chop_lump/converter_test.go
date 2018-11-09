// Copyright 2018 Vulcanize
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

package chop_lump_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vulcanize/vulcanizedb/pkg/transformers/cat_file/chop_lump"
	"github.com/vulcanize/vulcanizedb/pkg/transformers/test_data"
)

var _ = Describe("Cat file chop lump converter", func() {
	var converter chop_lump.CatFileChopLumpConverter

	BeforeEach(func() {
		converter = chop_lump.CatFileChopLumpConverter{}
	})

	Context("chop events", func() {
		It("converts a chop log to a model", func() {
			models, err := converter.ToModels([]types.Log{test_data.EthCatFileChopLog})

			Expect(err).NotTo(HaveOccurred())
			Expect(models).To(Equal([]interface{}{test_data.CatFileChopModel}))
		})
	})

	Context("lump events", func() {
		It("converts a lump log to a model", func() {
			models, err := converter.ToModels([]types.Log{test_data.EthCatFileLumpLog})

			Expect(err).NotTo(HaveOccurred())
			Expect(models).To(Equal([]interface{}{test_data.CatFileLumpModel}))
		})
	})

	It("returns err if log is missing topics", func() {
		badLog := types.Log{
			Data: []byte{1, 1, 1, 1, 1},
		}

		_, err := converter.ToModels([]types.Log{badLog})
		Expect(err).To(HaveOccurred())
	})

	It("returns err if log is missing data", func() {
		badLog := types.Log{
			Topics: []common.Hash{{}, {}, {}, {}},
		}

		_, err := converter.ToModels([]types.Log{badLog})
		Expect(err).To(HaveOccurred())
	})
})
