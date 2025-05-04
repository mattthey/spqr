```
docker compose -f ./tmp-mat up -d
```

# init shard settings
connect to spqr admin console
```
./tmp-mat/scripts/psql_spqr_admin.sh
```

```
CREATE DISTRIBUTION ds1 COLUMN TYPES integer;
ALTER DISTRIBUTION ds1 ATTACH RELATION orders DISTRIBUTION KEY id;
ALTER DISTRIBUTION ds1 ATTACH RELATION item DISTRIBUTION KEY order_id;

CREATE KEY RANGE krid1 FROM 1 ROUTE TO sh1 FOR DISTRIBUTION ds1;
CREATE KEY RANGE krid2 FROM 1000 ROUTE TO sh2 FOR DISTRIBUTION ds1;
CREATE KEY RANGE krid3 FROM 2000 ROUTE TO sh3 FOR DISTRIBUTION ds1;
```

connect to spqr and create table + fill data
```
psql postgresql://user1:password@localhost:8433/db1
```

```sql
CREATE TABLE orders (
    id SERIAL NOT NULL PRIMARY KEY,
    customer_id INT,
    order_date DATE
);

CREATE TABLE items (
   id SERIAL NOT NULL PRIMARY KEY,
   order_id SERIAL NOT NULL,
   name VARCHAR
);
```