help:
	@echo '---------- 環境構築に関するコマンド ----------'
	@echo 'init           -- プロジェクト初期のセットアップを行います※基本的にクローンしてきて1回目のみ実行'
	@echo 'remake         -- 環境を作り直します※dockerの構成等変更になったらこのコマンドを実行してください'
	@echo ''
	@echo '---------- Goに関するコマンド ----------'
	@echo 'run            -- goサーバーを起動します'
	@echo 'migrate        -- DBのAutoMigrateを実行します'
	@echo 'test           -- Goのテストをカバレッジ表示で実行します'
	@echo ''
	@echo '---------- Gitに関するコマンド ----------'
	@echo 'gs             -- git status'

# ---------- 環境構築に関するコマンド ----------
init:
	go mod tidy
	cp .env.example .env
	@make migrate

remake:
	@make init

# ---------- Goに関するコマンド ----------
run:
	go run ./cmd/main.go

migrate-up:
	go run ./migrate/migrate.go

test:
	go test -v -cover ./...

cover:
	go test -cover ./... -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html
	open cover.html

# ---------- Gitに関するコマンド ----------
gs:
	git status
