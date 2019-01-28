// VulcanizeDB
// Copyright © 2018 Vulcanize

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

package cmd

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/vulcanize/vulcanizedb/libraries/shared"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
	"github.com/vulcanize/vulcanizedb/pkg/transformers"
	shared2 "github.com/vulcanize/vulcanizedb/pkg/transformers/shared"
	"github.com/vulcanize/vulcanizedb/pkg/transformers/shared/constants"
)

// continuousLogSyncCmd represents the continuousLogSync command
var continuousLogSyncCmd = &cobra.Command{
	Use:   "continuousLogSync",
	Short: "Continuously sync logs at the head of the chain",
	Long: fmt.Sprintf(`Continously syncs logs based on the configured transformers.

vulcanizedb continousLogSync --config environments/local.toml
	
Available transformers for (optional) selection with --transformers:
%v

This command expects a light sync to have been run, and the presence of header records in the Vulcanize database.`,
		constants.AllTransformerLabels()),
	Run: func(cmd *cobra.Command, args []string) {
		syncMakerLogs()
	},
}

var transformerNames []string

func syncMakerLogs() {
	ticker := time.NewTicker(pollingInterval)
	defer ticker.Stop()

	blockChain := getBlockChain()
	db, err := postgres.NewDB(databaseConfig, blockChain.Node())
	if err != nil {
		log.Fatal("Failed to initialize database.")
	}

	initializers := getTransformerInitializers(transformerNames)

	watcher := shared.NewEventWatcher(db, blockChain)
	watcher.AddTransformers(initializers)

	for range ticker.C {
		err = watcher.Execute()
		if err != nil {
			// TODO Handle watcher errors in ContinuousLogSync
		}
	}
}

func getTransformerInitializers(transformerNames []string) []shared2.TransformerInitializer {
	var initializers []shared2.TransformerInitializer

	if transformerNames[0] == "all" {
		initializers = transformers.TransformerInitializers()
	} else {
		initializerMap := buildTransformerInitializerMap()
		for _, transformerName := range transformerNames {
			initializers = append(initializers, initializerMap[transformerName])
		}
	}
	return initializers
}

func buildTransformerInitializerMap() map[string]shared2.TransformerInitializer {
	initializerMap := make(map[string]shared2.TransformerInitializer)

	initializerMap[constants.BiteLabel] = transformers.GetBiteTransformer().NewTransformer
	initializerMap[constants.CatFileChopLumpLabel] = transformers.GetCatFileChopLumpTransformer().NewLogNoteTransformer
	initializerMap[constants.CatFileFlipLabel] = transformers.GetCatFileFlipTransformer().NewLogNoteTransformer
	initializerMap[constants.CatFilePitVowLabel] = transformers.GetCatFilePitVowTransformer().NewLogNoteTransformer
	initializerMap[constants.DealLabel] = transformers.GetDealTransformer().NewLogNoteTransformer
	initializerMap[constants.DentLabel] = transformers.GetDentTransformer().NewLogNoteTransformer
	initializerMap[constants.DripDripLabel] = transformers.GetDripDripTransformer().NewLogNoteTransformer
	initializerMap[constants.DripFileIlkLabel] = transformers.GetDripFileIlkTransformer().NewLogNoteTransformer
	initializerMap[constants.DripFileRepoLabel] = transformers.GetDripFileRepoTransformer().NewLogNoteTransformer
	initializerMap[constants.DripFileVowLabel] = transformers.GetDripFileVowTransformer().NewLogNoteTransformer
	initializerMap[constants.FlapKickLabel] = transformers.GetFlapKickTransformer().NewTransformer
	initializerMap[constants.FlipKickLabel] = transformers.GetFlipKickTransformer().NewTransformer
	initializerMap[constants.FlopKickLabel] = transformers.GetFlopKickTransformer().NewTransformer
	initializerMap[constants.FrobLabel] = transformers.GetFrobTransformer().NewTransformer
	initializerMap[constants.PitFileDebtCeilingLabel] = transformers.GetPitFileDebtCeilingTransformer().NewLogNoteTransformer
	initializerMap[constants.PitFileIlkLabel] = transformers.GetPitFileIlkTransformer().NewLogNoteTransformer
	initializerMap[constants.PriceFeedLabel] = transformers.GetPriceFeedTransformer().NewLogNoteTransformer
	initializerMap[constants.TendLabel] = transformers.GetTendTransformer().NewLogNoteTransformer
	initializerMap[constants.VatFluxLabel] = transformers.GetVatFluxTransformer().NewLogNoteTransformer
	initializerMap[constants.VatFoldLabel] = transformers.GetVatFoldTransformer().NewLogNoteTransformer
	initializerMap[constants.VatGrabLabel] = transformers.GetVatGrabTransformer().NewLogNoteTransformer
	initializerMap[constants.VatHealLabel] = transformers.GetVatHealTransformer().NewLogNoteTransformer
	initializerMap[constants.VatInitLabel] = transformers.GetVatInitTransformer().NewLogNoteTransformer
	initializerMap[constants.VatMoveLabel] = transformers.GetVatMoveTransformer().NewLogNoteTransformer
	initializerMap[constants.VatSlipLabel] = transformers.GetVatSlipTransformer().NewLogNoteTransformer
	initializerMap[constants.VatTollLabel] = transformers.GetVatTollTransformer().NewLogNoteTransformer
	initializerMap[constants.VatTuneLabel] = transformers.GetVatTuneTransformer().NewLogNoteTransformer
	initializerMap[constants.VowFlogLabel] = transformers.GetFlogTransformer().NewLogNoteTransformer

	return initializerMap
}

func init() {
	rootCmd.AddCommand(continuousLogSyncCmd)
	continuousLogSyncCmd.Flags().StringSliceVar(&transformerNames, "transformers", []string{"all"}, "transformer names to be run during this command")
}
