curl -X POST http://localhost:8080/user \
--data '{"Fullname": "Gildembeg Santos Gomes", "TaxIdentifier":"7f879879872394234", "Email":"gildeerg.santos4@gmail.com", "Password": "123456789"}'

curl -X POST http://localhost:8080/user \
--data '{"Fullname": "Gildembeg Santos Gomes", "TaxIdentifier":"7f879879872394235", "Email":"gildeerg.santos@gmail.com", "Password": "123456789"}'

curl -X PUT http://localhost:8080/user/1 \
--data '{"Fullname": "Gildembeg Santos Gomes", "TaxIdentifier":"7f879879872394234", "Email":"gildeerg.santos4@gmail.com", "Password": "123456789#0"}'

curl -X GET http://localhost:8080/users

curl -X GET http://localhost:8080/user/1

curl -X DELETE http://localhost:8080/user/1

curl -X DELETE http://localhost:8080/user/2
