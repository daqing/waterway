dev:
  overmind start -f Procfile.dev

install-deps:
  go install github.com/cosmtrek/air@latest
  brew install tmux
  brew install overmind
  curl -fsSL https://bun.sh/install | bash

build-cli:
  cd ./cli && go build && cd ..
  mv ./cli/cli ./bin/

build-cli-docker:
  cd ./cli && GOOS=linux GOARCH=amd64 go build -o ../bin/cli_amd .

cli +args:
  ./bin/cli {{args}}

packjs:
  rm -f ./public/js/*
  cd ./frontend/javascripts/ && bun build --minify --splitting --outdir=../../public/js ./src/*.jsx

build: packjs
  GOOS=linux GOARCH=amd64 go build -o ./bin .

bun:
  cd ./frontend/javascripts && bun install

docker: build build-cli-docker
  docker build -t waterway -f Dockerfile --platform linux/amd64 .

dbdocker:
  docker build -t waterway_db -f Dockerfile.db  --platform linux/amd64 .

push:
  docker tag waterway reg.appmz.cn/daqing/waterway
  docker push reg.appmz.cn/daqing/waterway

  docker tag waterway_db reg.appmz.cn/daqing/waterway_db
  docker push reg.appmz.cn/daqing/waterway_db

migrate_pg:
  find db/pg/*.sql | xargs -I{} psql -U $POSTGRES_USER -d waterway -f {}

migrate_sq:
  find db/sqlite3/*.sql -exec sh -c 'sqlite3 waterway.sqlite3 < {}' \;

create_pg:
  psql -U $POSTGRES_USER -d postgres -c "create database waterway"

create_sq:
  touch waterway.sqlite3

drop_pg:
  psql -U $POSTGRES_USER -d postgres -c "drop database waterway"

drop_sq:
  rm waterway.sqlite3

reset_pg: drop_pg create_pg migrate_pg

reset_sq: drop_sq create_sq migrate_sq

db_pg:
  psql -U $POSTGRES_USER -d waterway

db_sq:
  sqlite3 waterway.sqlite3
