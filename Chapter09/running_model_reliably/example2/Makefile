all: compile docker push clean

compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o goregtrain

docker:
	sudo docker build --force-rm=true -t dwhitena/goregtrain:multi .

push:
	sudo docker push dwhitena/goregtrain:multi

clean:
	rm goregtrain
