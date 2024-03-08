#!/usr/bin/env bash
/commands/wait-for-it.sh -t 30 testlocal:3306
/commands/run_migrations.sh
/commands/run.sh
