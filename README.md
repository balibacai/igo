cd /data/go/src && bee api beego

bee generate scaffold user -fields="email:string,password:string" -driver="mysql" -conn="root:chao123@tcp(db:3306)/mydev"