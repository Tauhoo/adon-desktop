./client/node_modules:
	npm install --prefix ./client

client-dev: ./client/node_modules
	npm start --prefix ./client 

server-dev: 
	ADON_ENV=dev go1.18.3 run main.go

out:
	mkdir out

out/html: out ./client/node_modules
	npm run build --prefix ./client
	mv ./client/build ./out/html

out/app: out/html
	go1.18.3 mod tidy
	astilectron-bundler
	./build.sh

all: out/app

clear: out
	rm -rf out
	rm -rf temp
	rm bind_*