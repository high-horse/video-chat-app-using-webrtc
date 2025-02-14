# running server 
make server :
	cd backend && go mod tidy && go run .

	
# running client
make client:
	cd frontend && npm i && npm run dev
