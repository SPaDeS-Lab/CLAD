# CLAD

A Clustered BLockchain-based Authentication and Data Sharing Scheme for Resource Constrained Networks

## Architecture Overview

CLAD is implemented as a set of custom chaincode modules on Hyperledger Fabric, featuring:

- Offline cluster support using private data collections (TAL - Temporary Authentication Ledger)
- Geographical clustering through separate Fabric channels
- Challenge-response authentication protocol
- Threshold verification mechanism
- Inter-channel synchronization with Merkle root validation
- Timestamp-based conflict resolution

## Components

1. Core Authentication Chaincode

   - Challenge-response protocol implementation
   - Custom threshold verification logic
   - Integration with Fabric MSP

2. Cluster Management

   - Private data collections for TAL
   - Geographical channel separation
   - Offline operation support

3. Synchronization Mechanism
   - Inter-channel communication
   - Merkle root computation
   - Proof validation
   - Conflict resolution

## Setup Requirements

- Hyperledger Fabric v2.x
- Multiple peer nodes for cluster setup
- Network configuration for geographical distribution

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-org/clad.git
   cd clad
   ```

2. Install Prerequisites:

   - Docker v20.10.x or higher
   - Docker Compose v2.x
   - Go v1.17.x or higher
   - Node.js v16.x or higher
   - Python 3.8+

3. Set up the Hyperledger Fabric test network:

   ```bash
   cd fabric-samples/test-network
   ./network.sh up createChannel -ca
   ```

4. Deploy the chaincode modules:

   ```bash
   ./network.sh deployCCAAS -ccn authentication -ccp ../chaincode/authentication
   ./network.sh deployCCAAS -ccn synchronization -ccp ../chaincode/synchronization
   ```

5. Configure the monitoring stack:

   ```bash
   cd prometheus-grafana
   docker-compose up -d
   ```

6. Initialize the clusters:

   ```bash
   ./scripts/init-clusters.sh
   ```

7. Verify the installation:
   ```bash
   ./scripts/run-tests.sh
   ```

For detailed configuration options and cluster setup, refer to the Configuration section.

## License

This project is licensed under the MIT License.
