all: compile docker push clean

compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o goregpredict

docker:
	sudo docker build --force-rm=true -t dwhitena/goregpredict .

push:
	sudo docker push dwhitena/goregpredict

clean:
	rm goregpredict
