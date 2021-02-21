#!/usr/bin/env bash/bin/bash
sudo dnf install postgresql postgresql-server # install client/server
sudo postgresql-setup  --initdb --unit postgresql   &&            # initialize PG cluster
sudo systemctl start postgresql  &&            # start cluster
sudo su - postgres &&
createuser igrid-iam -P &&
createdb igrid-iam-db --owner igrid-iam &&
sudo systemctl restart postgresql