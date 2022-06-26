./client/node_modules:
	npm install --prefix ./client

client-dev: ./client/node_modules
	npm start --prefix ./client 

server-dev: 
	cd ./server && go1.18.3 run -ldflags="-X main.env=dev" *.go 

out: 
	mkdir out

out/client: out
	npm run build --prefix ./client
	mv ./client/build ./out/client

out/server: out
	mkdir out/server
	cd ./server && go1.18.3 build -o run main.go 
	mv server/run out/server/run

all: out/client out/server

clean: out
	rm -rf ./out