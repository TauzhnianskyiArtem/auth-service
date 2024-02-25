#!/bin/bash
source .env

sleep 2 && goose postgres "${MIGRATION_DSN}" up -v