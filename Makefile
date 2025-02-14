make server :
	cd backend && go mod tidy && go run .

make client:
	cd frontend && npm i && npm run dev
