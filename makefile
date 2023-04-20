cicd_build:
	rm -rf app
	mkdir app
	make bb

bb:
	go build -o app/api  