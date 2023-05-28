#!/bash/bin

go run greet.go -f ./etc/greet-api.yaml

curl -i -X GET http://localhost:9999/from/you