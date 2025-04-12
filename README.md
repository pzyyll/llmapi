# LLMAPI

## 数据库连接 (DSN)

支持的数据库连接字符串格式：

### SQLite
- **格式**:  
  `sqlite://{uri}`  
- **示例**:  
  `sqlite://test.db` 或 `sqlite://file:test.db`  
- **参考**:  
  [SQLite URI 格式](https://sqlite.org/uri.html)

---

### PostgreSQL
- **示例**:  
  `postgres://user:secret@localhost/mydb`  
- **参考**:  
  [PostgreSQL 连接字符串](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING-URIS)

---

### MySQL
- **格式**:  
  `mysql://{dsn}`  
- **示例**:  
  `mysql://user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local`  
- **参考**:  
  [MySQL DSN 格式](https://github.com/go-sql-driver/mysql#dsn-data-source-name)

---

### SQL Server
- **示例**:  
  `sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm`  
- **参考**:  
  [SQL Server 连接字符串](https://github.com/go-gorm/sqlserver)