# govendor
```
govendor add +external
govendor remove +unused

```

# init api structure
```
export MYSQL_USERNAME=root
export MYSQL_PASSWORD=
export MYSQL_DATABASE=test

cd /data/go/src/beego

bee api . -tables="users" -driver=mysql -conn="$MYSQL_USERNAME:$MYSQL_PASSWORD@tcp(db:3306)/$MYSQL_DATABASE"

```

# scaffold
```
bee generate scaffold user -fields="email:string,password:string" -driver="mysql" -conn="$MYSQL_USERNAME:$MYSQL_PASSWORD@tcp(db:3306)/$MYSQL_DATABASE"

```

# migrate
```
bee migrate -driver="mysql" -conn="$MYSQL_USERNAME:$MYSQL_PASSWORD@tcp(db:3306)/$MYSQL_DATABASE"

```

# test
```
cd test

go test
  or 
go test -v jwt_test.go

```

# add by develop branch

