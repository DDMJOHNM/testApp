# Full Stack Web Development With Go
- docker
- postgres
- sqlc
- makefile

```
docker run --name test-postgres \
-e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres


```
```
docker exec -it test-postgres psql -h localhost -p 5432 -U postgres -d postgres
\i /usr/share/chapter1/db/schema.sql 
(run schema code query in postegres)
\dt gowebapp.* 
sqlc generate 


```

