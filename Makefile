./client/node_modules:
	npm install --prefix ./client

client-dev: ./client/node_modules
	npm start --prefix ./client 

server-dev: 
	go1.18.3 run -ldflags="-X main.env=dev" *.go 

out:
	mkdir out

out/html: out ./client/node_modules
	npm run build --prefix ./client
	mv ./client/build ./out/html

out/app:
	astilectron-bundler
	rm -rf temp
	rm bind_*

clear: out
	rm -rf out