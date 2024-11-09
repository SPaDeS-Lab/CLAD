package sync

import (
	"encoding/json"
	"fmt"
	"time"
)

type SyncChaincode struct {
	// SyncEvent represents a synchronization attempt from an offline cluster
	SyncEvent struct {
		ClusterID  string     // ID of the offline cluster initiating sync
		MerkleRoot []byte     // Merkle root of cluster's TAL
		Records    []TALEntry // Records to be synchronized
		Timestamp  int64      // Sync attempt timestamp
		Status     string     // Current sync status (pending/complete/failed)
	}

	// SyncConfig defines synchronization parameters
	SyncConfig struct {
		SyncInterval   int64  // Time between sync attempts in seconds
		RetryLimit     int    // Maximum sync retry attempts
		BatchSize      int    // Maximum records per sync batch
		VerifyTimeout  int64  // Timeout for Merkle root verification
		ConflictPolicy string // Policy for resolving timestamp conflicts
	}

	// SyncResult tracks the outcome of a synchronization attempt
	SyncResult struct {
		Success       bool   // Whether sync completed successfully
		ErrorMessage  string // Description of any sync failure
		UpdatedCount  int    // Number of records synchronized
		ConflictCount int    // Number of conflicts resolved
		CompletedAt   int64  // Timestamp of sync completion
	}
}

// ComputeMerkleRoot calculates the Merkle root of the TAL
func (sc *SyncChaincode) ComputeMerkleRoot(ctx contractapi.TransactionContextInterface) (*MerkleRoot, error) {
	// Get all records from TAL private collection
	talIterator, err := ctx.GetStub().GetPrivateDataByRange("TALCollection", "", "")
	if err != nil {
		return nil, fmt.Errorf("failed to get TAL records: %v", err)
	}
	defer talIterator.Close()

	// Build Merkle tree from records
	var records [][]byte
	for talIterator.HasNext() {
		record, err := talIterator.Next()
		if err != nil {
			return nil, err
		}
		records = append(records, record.Value)
	}

	merkleRoot := computeMerkleTreeRoot(records)
	return &MerkleRoot{Root: merkleRoot, Timestamp: time.Now().Unix()}, nil
}

// SynchronizeWithMainChannel handles the synchronization process
func (sc *SyncChaincode) SynchronizeWithMainChannel(ctx contractapi.TransactionContextInterface) error {
	// Get Merkle root
	merkleRoot, err := sc.ComputeMerkleRoot(ctx)
	if err != nil {
		return err
	}

	// Prepare synchronization event
	syncEvent := &SyncEvent{
		ChannelID:  ctx.GetStub().GetChannelID(),
		MerkleRoot: merkleRoot,
		Records:    sc.getAuthRecords(ctx),
	}

	// Emit synchronization event
	eventPayload, _ := json.Marshal(syncEvent)
	return ctx.GetStub().SetEvent("SyncRequired", eventPayload)
}
