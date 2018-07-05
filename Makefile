all:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o auth .

docker:
	- docker image rm reg.urantiatech.com/auth/auth
	docker build -t reg.urantiatech.com/auth/auth .
