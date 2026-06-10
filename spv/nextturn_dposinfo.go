package spv

import (
	"bytes"
	"errors"
	"math"
	"sort"

	"github.com/elastos/Elastos.ELA.SideChain.ECO/common"
	"github.com/elastos/Elastos.ELA.SideChain.ECO/log"

	spv "github.com/elastos/Elastos.ELA.SPV/interface"

	"github.com/elastos/Elastos.ELA/core/types/payload"
)

type NextTurnDPOSInfo struct {
	*payload.NextTurnDPOSInfo
}

var (
	nextTurnDposInfo *NextTurnDPOSInfo
	zero             = common.Hex2Bytes("000000000000000000000000000000000000000000000000000000000000000000")
	backupProducers  [][]byte
)

func GetTotalProducersCount() int {
	if nextTurnDposInfo == nil {
		return 0
	}
	count, err := SafeAdd(len(nextTurnDposInfo.CRPublicKeys), len(nextTurnDposInfo.DPOSPublicKeys))
	if err != nil {
		log.Error("SafeAdd error", "error", err)
		return 0
	}
	return count
}

func SpvIsWorkingHeight() bool {
	if nextTurnDposInfo != nil {
		return SpvService.GetBlockListener().BlockHeight() > nextTurnDposInfo.WorkingHeight
	}
	return false
}

func MainChainIsPowMode() bool {
	return consensusMode == spv.POW
}

func GetProducers(elaHeight uint64) ([][]byte, int, error) {
	producers := make([][]byte, 0)
	totalCount := 0
	if SpvService == nil {
		return producers, totalCount, errors.New("spv is not start")
	}
	if GetCurrentConsensusMode() == spv.POW {
		return producers, totalCount, nil
	}
	if elaHeight == math.MaxUint64 { //defaults producers
		producers = GetDefaultProducers()
		totalCount = len(producers)
		return producers, totalCount, nil
	}
	crcArbiters, normalArbitrs, err := SpvService.GetArbiters(uint32(elaHeight))
	if err != nil {
		return producers, totalCount, err
	}
	if IsOnlyCRConsensus {
		normalArbitrs = make([][]byte, 0)
	}

	for _, arbiter := range crcArbiters {
		if len(arbiter) > 0 && bytes.Compare(zero, arbiter) != 0 {
			producers = append(producers, arbiter)
		}
	}
	totalCount, err = SafeAdd(len(crcArbiters), len(normalArbitrs))
	if err != nil {
		return nil, totalCount, err
	}
	return producers, totalCount, nil
}

func SetBackupProducers(producers []string) {
	backupProducers = make([][]byte, 0, len(producers))
	for _, p := range producers {
		key := common.Hex2Bytes(p)
		if len(key) > 0 {
			backupProducers = append(backupProducers, key)
		}
	}
	log.Info("Backup producers configured", "count", len(backupProducers))
}

func GetDefaultProducers() [][]byte {
	defaultsProducers := backupProducers
	producers := make([][]byte, 0)
	for _, producer := range defaultsProducers {
		producers = append(producers, producer)
	}
	sort.Slice(producers, func(i, j int) bool {
		return bytes.Compare(producers[i], producers[j]) < 0
	})
	return producers
}

func GetSpvHeight() uint64 {
	if SpvService != nil && SpvService.GetBlockListener() != nil {
		header, err := SpvService.HeaderStore().GetBest()
		if err != nil {
			log.Error("SpvService getBest error", "error", err)
			return uint64(SpvService.GetBlockListener().BlockHeight())
		}
		return uint64(header.Height)
	}
	return 0
}

func GetWorkingHeight() uint32 {
	if nextTurnDposInfo != nil {
		return nextTurnDposInfo.WorkingHeight
	}
	return 0
}

func GetSpvService() *Service {
	return SpvService
}
