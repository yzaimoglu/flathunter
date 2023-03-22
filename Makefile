OUT = flathunter
BACKEND_VERSION = 0.1
FRONTEND_VERSION = 0.1

run:
	tmux new-session -d -s flathunter-dev
	tmux send-keys -t flathunter-dev 'make dev-server' C-m
	tmux split-window -v -t flathunter-dev
	tmux send-keys -t flathunter-dev 'make dev-client' C-m
	tmux split-window -h -t flathunter-dev
	tmux send-keys -t flathunter-dev 'make dev-frontend' C-m
	tmux attach -t flathunter-dev

docker-up:
	docker compose --env-file backend/.env up -d

docker-down:
	docker compose --env-file backend/.env down

dev-client:
	cd backend && go run ./cmd/client

dev-server:
	cd backend && go run ./cmd/server

dev-frontend:
	cd frontend && npm run dev

prod-client:
	cd backend && go build -o ../$(OUT)-client ./cmd/client

prod-server:
	cd backend && go build -o ../$(OUT)-server ./cmd/server

prod-frontend:
	cd frontend && npm run build

prod-backend:
	make prod-client
	make prod-server

clean:
	rm $(OUT)-client $(OUT)-server $(OUT)-frontend

docker-server:
	docker build -t flathunter-server:$(BACKEND_VERSION) backend -f Dockerfile.server

docker-client:
	docker build -t flathunter-client:$(BACKEND_VERSION) backend -f Dockerfile.client

docker-frontend:
	docker build -t flathunter-frontend:$(FRONTEND_VERSION) frontend -f Dockerfile.frontend
