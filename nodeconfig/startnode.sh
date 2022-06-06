#!/bin/sh

# remove all sync data
pocket_core reset

# off-chain account commands
pocket_core accounts list
pocket_core accounts create
pocket_core accounts set-validator 169869f67cd3f78a722fb4795b69949fb4bc9084
pocket_core accounts show 169869f67cd3f78a722fb4795b69949fb4bc9084

# command to start node
pocket_core start

# transaction for sending tokens
pocket_core accounts send-tx 169869f67cd3f78a722fb4795b69949fb4bc9084 169869f67cd3f78a722fb4795b69949fb4bc9084 10000000 testnet 10000 xxx