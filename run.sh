docker container rm -f backend
docker container rm -f mongo

docker container run -d --name mongo mongo:3.5.13-jessie

make docker
docker container run -d --name backend --link mongo -p 3000:3000 muchrm/go-echo:latest