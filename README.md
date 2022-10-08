# belajar-golang

## Cara Menjakankan

1. Jalankan docker compose terlebih dahulu dengan perintah `docker-compose up`
2. Jalankan perintah `go mod download` untuk mendownload dependency yang dibutuhkan oleh project
3. Kemudian jalankan perintah `go run main.go` untuk menjalakan aplikasi

## Akses REST API

### Get Employee

```shell
curl --location --request GET 'http://localhost:3000/api/employee'
```

### Get One Employee

```shell
curl --location --request GET 'http://localhost:3000/api/employee/1'
```

### Create Employee

```shell
curl --location --request POST 'http://localhost:3000/api/employee' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 1,
    "name": "rudi",
    "age": 10,
    "address": "jakarta",
    "level": "junior",
    "salary": 1000
}'
```

### Update Employee

```shell
curl --location --request PUT 'http://localhost:3000/api/employee' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 1,
    "name": "rudi",
    "age": 10,
    "address": "depok",
    "level": "junior",
    "salary": 1000
}'
```

### Delete Employee

```shell
curl --location --request DELETE 'http://localhost:3000/api/employee/1'
```
