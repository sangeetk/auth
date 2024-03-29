all:
	go build -o auth .

dev:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o auth .
	- docker image rm localhost:32000/auth/auth 
	docker build -t localhost:32000/auth/auth .
	docker push localhost:32000/auth/auth

prod:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o auth .
	- docker image rm reg.urantiatech.com/auth/auth 
	docker build -t reg.urantiatech.com/auth/auth .
	docker push reg.urantiatech.com/auth/auth
