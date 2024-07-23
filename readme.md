# BE GOLANG GIN SQLX project structure

## Project Structure 
```
.
├── constant                     
    ├── default.go                        # constants for static things, default settings
    └── logger.go                         # constants for logger setup, add more files for another constants
├── dto                                     
    ├── migrations                        # migration folder for golang-migrate
    └── schema.sql                        # also leave the rootSchema/ initSchema here
├── infras
    ├── config                            # logic codes for viper config
    ├── connection                        # logic codes for making conenction to DB
    ├── logger                            # logic codes for making custom logger
    └── router                            # apis here    
        ├── controllers                   # business logics, validate, binding, error handling here
            ├── index.go                  # 
            └── sample_controllers.go
        ├── queries                       # query, intergrate with database
            ├── index.go                  #
            └── sample_queries.go
        └── server.go                     # root routing
├── models                      
    ├── models.go                         # place shema structs/ models here
    └── requests.go                       # place api request structs here
├── utils                      
    └── functions.go                      # common func for reuse
├── app.env                               # envroment keys
├── Dokcerfile                            # building to Docker
├── docker-compose.yaml                   #
├── main.go                               # Main func of the app
├── Makefile                              #
└── start.sh                              #
```

## This has not been completed yet, there are few features need TODO:
- [x] ~~**Update dockerfile, docker-compose**~~
- [ ] **Gin Middlewares**
- [ ] **Gin Authentication, Token,...**
- [ ] **Spliting services for microServices**
- [ ] **Redis**
- [ ] **Authorization..**