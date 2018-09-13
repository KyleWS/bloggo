if ! [[ -e ./TLS/fullchain.pem && -e ./TLS/privkey.pem ]]; then 
	echo "local certs not found. generating..."
	openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj "/CN=localhost" -keyout $(pwd)/TLS/privkey.pem -out $(pwd)/TLS/fullchain.pem
fi

if [ "$1" = "tls" ]; then
	export SERVER_ADDRESS=:443
else
	export SERVER_ADDRESS=:2555
fi

export TLSCERT=$(pwd)/TLS/fullchain.pem
export TLSKEY=$(pwd)/TLS/privkey.pem

# turned off so I don't have to keep adding dummy data
#docker rm -f mongo-server
output=$(docker ps --filter name="mongo-server" -q)
if [[ -z "$output" ]]; then
	docker run -d --name mongo-server -p 27017:27017 mongo
fi
go build
./bloggo
