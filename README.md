### Simple Example of Dependency Injection / User Repository

#### Uses gorm.io Library for PostgreSQL

#### Verify SQL Connection

```bash
psql -P expanded=auto -h postgres-instance.company.com --username=testuser --dbname=testdb
```

#### Output

```bash
[~/git/goworkspace/src/di]$ go run *.go

2022/10/15 14:07:06 /Users/gbhujan/git/goworkspace/pkg/mod/gorm.io/driver/postgres@v1.4.4/migrator.go:178
[96.160ms] [rows:1] SELECT count(*) FROM information_schema.tables WHERE table_schema = CURRENT_SCHEMA() AND table_name = 'user_records' AND table_type = 'BASE TABLE'

2022/10/15 14:07:06 /Users/gbhujan/git/goworkspace/pkg/mod/gorm.io/driver/postgres@v1.4.4/migrator.go:151
[94.129ms] [rows:0] CREATE TABLE "user_records" ("userid" bigint,"name" text)

2022/10/15 14:07:06 /Users/gbhujan/git/goworkspace/src/di/user_repository.go:27
[242.399ms] [rows:1] INSERT INTO "user_records" ("userid","name") VALUES (100,'Jordan Peterson')

2022/10/15 14:07:07 /Users/gbhujan/git/goworkspace/src/di/user_repository.go:35
[95.588ms] [rows:1] SELECT * FROM "user_records" WHERE "userid" = 100
2022/10/15 14:07:07 user fetched : {100 Jordan Peterson}
2022/10/15 14:07:07 user-id : 100
```

### Verify SQL Table

```sql
testdb=> select * from user_records;
 userid |      name
--------+-----------------
    100 | Jordan Peterson
(1 row)
```