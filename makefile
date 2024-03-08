request:
	curl -X POST http://localhost:8080/user \
	--data '{"fullname": "Gildembeg Santos Gomes", "tax_identifier":"7f879879872394234", "email":"gildeerg.santos4@gmail.com", "password": "123456789"}'

	curl -X POST http://localhost:8080/user \
	--data '{"fullname": "Gildembeg Santos Gomes", "tax_identifier":"7f879879872394235", "email":"gildeerg.santos@gmail.com", "password": "123456789"}'

	curl -X PUT http://localhost:8080/user/1 \
	--data '{"fullname": "Gildembeg Santos Gomes", "tax_identifier":"7f879879872394234", "email":"gildeerg.santos4@gmail.com", "password": "123456789#0"}'

	curl -X GET http://localhost:8080/users

	curl -X GET http://localhost:8080/user/1

	curl -X DELETE http://localhost:8080/user/1

	curl -X DELETE http://localhost:8080/user/2

	rm -rf db/test.db

run:
	go run cmd/server.go
