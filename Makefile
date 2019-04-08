all:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o auth .

prod:
	GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o auth .
	- docker image rm reg.urantiatech.com/auth/auth 
	docker build -t reg.urantiatech.com/auth/auth .
	docker push reg.urantiatech.com/auth/auth

dev:
	GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o auth .
	- docker image rm localhost:5000/auth/auth 
	docker build -t localhost:5000/auth/auth .
	docker push localhost:5000/auth/auth
