migrate create -ext sql -dir db/migration -seq init_schema

docker exec -it postgres12 createdb --username=root --owner=root simple_bank
docker exec -it postgres12 psql -U root simple_bank


Lesson 9 : MySQL bilan Postgres farqini ko'rildi
 
Lesson 10 : Github# SimpleBank Action... githubda proyektni test qilishni o'rganildi

Lesson 11 : Create Get List

Lesson 12 : viper, env file o'rnatish va sozlash

Lesson 13: mockgen -destination db/mock/store.go  github.com/teachschool/simplebank/db/sqlc Store