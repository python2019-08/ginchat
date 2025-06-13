# 1.  github.com/jinzhu/gorm和gorm.io/gorm 选哪个

### GORM 版本选择指南：`jinzhu/gorm` vs `gorm.io/gorm`

#### **一、历史背景与版本演进**
1. **`jinzhu/gorm`（v1 版本）**  
   - 早期 GORM 版本，由创始人 Jinzhu 维护。
   - 最后更新于 2020 年，目前处于维护模式（仅修复关键 bug）。
   - 社区活跃度逐渐降低，但仍有大量存量项目使用。

2. **`gorm.io/gorm`（v2 版本）**  
   - 2020 年发布的重构版本，由 GORM 官方团队（gorm.io）维护。
   - 完全兼容 Go Module，采用新的包路径 `gorm.io/gorm`。
   - 设计更现代化，性能优化，支持更多数据库（如 ClickHouse、TiDB）。


#### **二、核心差异对比**
| **特性**               | `jinzhu/gorm` (v1)        | `gorm.io/gorm` (v2)        |
|------------------------|---------------------------|---------------------------|
| **包路径**             | `github.com/jinzhu/gorm`  | `gorm.io/gorm`            |
| **Go Module 支持**     | 不友好（需 `replace`）    | 原生支持                   |
| **查询语法**           | 链式调用为主              | 更灵活（支持 `clause`）    |
| **事务处理**           | 手动 `Begin`/`Commit`     | 支持 `Transaction()` 方法  |
| **钩子（Hook）**       | 简单接口                  | 更强大的接口（支持更多事件）|
| **错误处理**           | 返回单个 `error`          | 支持 `errors.Is()` 判断    |
| **批量操作**           | 性能一般                  | 优化批量插入/更新          |
| **新数据库支持**       | 有限（MySQL、PostgreSQL 等）| 新增 ClickHouse、SQLite3 等 |
| **活跃维护**           | 否                        | 是                        |


#### **三、选择建议**

##### 1. **新项目推荐使用 `gorm.io/gorm`**
- **理由**：
  - 官方推荐版本，持续更新与维护。
  - 更好的 Go Module 支持，避免依赖冲突。
  - 性能优化和新特性（如 `clause` 语法）。
  - 更完善的错误处理和钩子系统。

- **示例依赖配置**：
  ```go
  // go.mod
  require (
      gorm.io/gorm v2.5.2  // 最新稳定版
      gorm.io/driver/mysql v1.5.0  // 或其他数据库驱动
  )
  ```

##### 2. **旧项目继续使用 `jinzhu/gorm`**
- **理由**：
  - 迁移成本高（API 不兼容，需重构代码）。
  - 功能已满足需求，无需新特性。
  - 稳定性优先，不愿承担升级风险。

- **迁移挑战示例**：
  ```go
  // v1 查询
  db.Where("name = ?", "jinzhu").First(&user)
  
  // v2 查询（更灵活，但语法不同）
  db.Where("name = ?", "jinzhu").First(&user)  // 兼容写法
  db.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&users)  // 新增特性
  ```


#### **四、如何迁移（从 v1 到 v2）**
1. **更新依赖**：
   ```bash
   go get -u gorm.io/gorm
   go get -u gorm.io/driver/[数据库类型]
   ```

2. **修改导入路径**：
   ```go
   // 旧
   import "github.com/jinzhu/gorm"
   
   // 新
   import "gorm.io/gorm"
   ```

3. **适配 API 变更**：
   - 钩子（Hook）接口调整（如 `BeforeSave` 参数变化）。
   - 错误处理方式（如 `RecordNotFound` 改为 `errors.Is(err, gorm.ErrRecordNotFound)`）。
   - 批量操作（如 `CreateInBatches` 参数调整）。

4. **参考官方迁移指南**：  
   官方提供了详细的迁移文档：[gorm.io/docs/migration.html](https://gorm.io/docs/migration.html)


#### **五、实际项目案例**
- **选择 v2 的场景**：
  - 微服务架构（需要频繁更新依赖）。
  - 需要新数据库支持（如 ClickHouse）。
  - 对批量操作性能有要求。

- **选择 v1 的场景**：
  - 稳定运行的单体应用（无重大功能更新计划）。
  - 团队对 v1 熟悉，且迁移成本高。
  - 项目依赖 v1 特定功能（如旧版钩子系统）。


#### **六、总结**
- **新项目**：无脑选 `gorm.io/gorm`（v2），避免未来迁移成本。
- **旧项目**：评估迁移收益与成本，若功能稳定可继续使用 v1。

GORM v2 代表了 Go ORM 的未来趋势，建议尽早适应其设计理念，尤其是在新项目中。
=============================================================
# 2. TableName()在AutoMigrate时的作用
```go
type UserBasic struct {
	gorm.Model
	Name          string 
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
```
TableName()在AutoMigrate时的作用?

### GORM 中 `TableName()` 方法在 `AutoMigrate` 中的作用

在 GORM 中，`TableName()` 方法是一个**表名钩子（Hook）**，它的作用是**自定义数据库表名**。当你调用 `AutoMigrate` 时，GORM 会根据这个方法的返回值来创建或更新对应的数据库表。


#### **一、核心作用**
1. **表名映射**  
   - GORM 默认会将结构体名转为复数形式作为表名（如 `UserBasic` → `user_basics`）。
   - 通过 `TableName()` 方法，可以强制指定表名为 `user_basic`（单数形式）。

2. **`AutoMigrate` 执行时**  
   - 当调用 `db.AutoMigrate(&UserBasic{})` 时，GORM 会：
     1. 检查 `UserBasic` 是否实现了 `TableName()` 方法。
     2. 如果实现了，则使用返回值 `user_basic` 作为表名。
     3. 创建或更新名为 `user_basic` 的表，而非默认的 `user_basics`。


#### **二、示例验证**
```go
// 定义模型
type UserBasic struct {
    gorm.Model
    Name string
}

// 自定义表名
func (table *UserBasic) TableName() string {
    return "user_basic"
}

// 执行迁移
func main() {
    db, _ := gorm.Open(mysql.Open("dsn"), &gorm.Config{})
    
    // 自动迁移时，会创建名为 user_basic 的表
    db.AutoMigrate(&UserBasic{})
    
    // 验证表名
    var count int64
    db.Raw("SHOW TABLES LIKE 'user_basic'").Count(&count)
    fmt.Println("Table exists:", count > 0)  // 输出 true
}
```


#### **三、与其他表名配置方式的对比**
| **配置方式**               | **作用域**           | **优先级**         | **示例**                         |
|----------------------------|----------------------|--------------------|----------------------------------|
| `TableName()` 方法         | 单个模型             | 最高               | `func (t *User) TableName() { return "users" }` |
| `gorm:"table_name:xxx"` 标签 | 单个字段（已废弃）   | 中等               | `// v1 中使用，但 v2 已废弃`      |
| `NamingStrategy` 全局配置   | 所有模型             | 最低               | `db, _ := gorm.Open(..., &gorm.Config{ NamingStrategy: schema.NamingStrategy{ TablePrefix: "t_", SingularTable: true } })` |


#### **四、常见应用场景**
1. **保持表名一致性**  
   - 强制使用单数形式表名（如 `user` 而非 `users`）。
   - 兼容已有数据库设计（如历史表名 `user_basic`）。

2. **多租户场景**  
   - 为不同租户动态返回不同表名：
     ```go
     func (t *UserBasic) TableName() string {
         if tenantID := GetCurrentTenantID(); tenantID != "" {
             return fmt.Sprintf("%s_user_basic", tenantID)  // 动态表名：tenant1_user_basic
         }
         return "user_basic"
     }
     ```

3. **表名前缀**  
   - 为所有表添加统一前缀：
     ```go
     func (t *UserBasic) TableName() string {
         return "prefix_" + gorm.DefaultTableNameHandler(db, t)  // prefix_user_basic
     }
     ```


#### **五、注意事项**
1. **大小写敏感**  
   - 在 Linux/macOS 系统中，数据库表名通常**区分大小写**，而 Windows 不区分。
   - 建议统一使用小写表名（如 `user_basic`）以避免兼容性问题。

2. **性能影响**  
   - 每个查询都会调用 `TableName()` 方法，对于高频操作可能有轻微性能影响。
   - 若表名固定，可考虑使用全局 `NamingStrategy` 代替。

3. **跨数据库兼容性**  
   - 不同数据库对表名长度、特殊字符的限制不同，需确保表名合法。


#### **六、总结**
`TableName()` 方法是 GORM 中控制表名的最灵活方式，特别是在 `AutoMigrate` 时：
- **若无 `TableName()`**：GORM 默认将结构体名转为复数形式（如 `UserBasic` → `user_basics`）。
- **若实现 `TableName()`**：强制使用返回值作为表名（如 `user_basic`）。

在设计数据库模型时，建议明确指定表名，避免因 GORM 默认规则变化导致的兼容性问题。

======================================================
# 3.Incorrect datetime value: '0000-00-00' for column 'login_out_time' 

```sh
test_gorm_mysql.go:27 Error 1292 (22007): Incorrect datetime value: '0000-00-00' for column 'login_out_time' at row 1
[0.884ms] [rows:0] INSERT INTO `user_basic` (`created_at`,`updated_at`,`deleted_at`,`name`,`pass_word`,`phone`,`email`,`identity`,`client_ip`,`client_port`,`login_time`,`heartbeat_time`,`login_out_time`,`is_logout`,`device_info`) VALUES ('2025-06-13 22:24:34.184','2025-06-13 22:24:34.184',NULL,'申专','','','','','','',0,0,'0000-00-00 00:00:00',false,'')
```

### MySQL "Incorrect datetime value" 错误解决方案

#### **一、错误原因**
这个错误是由于 MySQL 不接受 `'0000-00-00'` 作为合法的 `DATETIME` 值导致的。从 MySQL 5.7 开始，默认配置下：
- `strict_mode` 被启用，禁止插入无效的日期值。
- `DATETIME` 字段不允许为 `NULL` 时，必须有合法值。

在你的案例中，`login_out_time` 字段被设置为 `'0000-00-00 00:00:00'`，这在严格模式下是不合法的。


#### **二、解决方案**

##### 1. **修改数据库连接参数**
在连接 MySQL 时，添加 `parseTime=true&loc=Local&sql_mode=ALLOW_INVALID_DATES` 参数：
```go
dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?parseTime=true&loc=Local&sql_mode=ALLOW_INVALID_DATES"
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
```
- `sql_mode=ALLOW_INVALID_DATES`：允许插入 `'0000-00-00'` 这样的零值。

##### 2. **将字段设置为可空（NULL）**
在模型中，将 `login_out_time` 类型改为指针：
```go
type UserBasic struct {
    gorm.Model
    // 其他字段...
    LoginOutTime *time.Time `gorm:"column:login_out_time"`
}
```
- 插入时，若值为 `nil`，数据库会存储为 `NULL`。

##### 3. **使用默认值**
在模型中为 `login_out_time` 设置默认值：
```go
type UserBasic struct {
    gorm.Model
    // 其他字段...
    LoginOutTime time.Time `gorm:"column:login_out_time;default:CURRENT_TIMESTAMP"`
}
```
- 若插入时未指定值，会自动使用当前时间。

##### 4. **插入合法的日期值**
确保插入时提供合法的日期，例如使用 `time.Time` 类型：
```go
user := UserBasic{
    // 其他字段...
    LoginOutTime: time.Now(),  // 或其他合法时间
}
db.Create(&user)
```


#### **三、验证与测试**
1. **检查 MySQL 版本**：
   ```sql
   SELECT VERSION();
   ```
   - MySQL 5.7+ 默认启用严格模式，需特别注意。

2. **查看当前 SQL_MODE**：
   ```sql
   SELECT @@sql_mode;
   ```
   - 若包含 `STRICT_TRANS_TABLES`，则需要调整。

3. **测试插入语句**：
   ```sql
   SET sql_mode='';  -- 临时禁用严格模式
   INSERT INTO user_basic (login_out_time) VALUES ('0000-00-00 00:00:00');
   
   SET sql_mode='STRICT_TRANS_TABLES';  -- 恢复严格模式
   INSERT INTO user_basic (login_out_time) VALUES ('0000-00-00 00:00:00');  -- 这会报错
   ```


#### **四、最佳实践**
1. **优先使用可空字段（NULL）**  
   - 对于可能没有值的日期字段（如登录退出时间），使用 `*time.Time` 类型。

2. **避免依赖零值日期**  
   - `'0000-00-00'` 是 MySQL 特有的表示方式，不具备跨数据库兼容性。

3. **统一日期格式**  
   - 使用 Go 的 `time.Time` 类型与数据库交互，避免手动处理字符串格式。

4. **配置连接参数**  
   - 在连接字符串中显式设置 `sql_mode`，确保环境一致性。


#### **五、修复示例代码**
```go
type UserBasic struct {
    gorm.Model
    // 其他字段...
    LoginOutTime *time.Time `gorm:"column:login_out_time"`  // 改为指针类型
}

// 创建用户时
user := UserBasic{
    // 其他字段...
    LoginOutTime: nil,  // 初始未退出，存储为 NULL
}
db.Create(&user)

// 用户退出时
now := time.Now()
db.Model(&user).Update("login_out_time", &now)  // 更新为当前时间
```


通过以上方法，你可以解决 `'0000-00-00'` 日期值错误，并建立更健壮的日期处理机制。

=============================================================
# 4.Invalid default value for 'login_out_time '

```go
type UserBasic struct {
	gorm.Model
	Name          string 
	LoginOutTime  time.Time `gorm:"column:login_out_time ;default:CURRENT_TIMESTAMP"  json:"login_out_time"`
}
```
但是
```sh
Error 1067 (42000): Invalid default value for 'login_out_time '
[0.083ms] [rows:0] CREATE TABLE `user_basic` (`id` bigint unsigned AUTO_INCREMENT,`created_at` datetime(3) NULL,`updated_at` datetime(3) NULL,`deleted_at` datetime(3) NULL,`name` longtext,`pass_word` longtext,`phone` longtext,`email` longtext,`identity` longtext,`client_ip` longtext,`client_port` longtext,`login_time` datetime(3) NULL,`heartbeat_time` datetime(3) NULL,`login_out_time ` datetime(3) NULL DEFAULT CURRENT_TIMESTAMP,`is_logout` boolean,`device_info` longtext,PRIMARY KEY (`id`),INDEX `idx_user_basic_deleted_at` (`deleted_at`))
2025/06/13 22:45:25 db.AutoMigrate() error: Error 1067 (42000): Invalid default value for 'login_out_time '
```


### MySQL "Invalid default value" 错误解决方案

#### **一、错误原因**
这个错误是由于 MySQL 对 `DATETIME` 类型的默认值有严格限制导致的：

1. **版本限制**：  
   - MySQL 5.7 及以下版本：`DATETIME` 默认不支持 `CURRENT_TIMESTAMP` 作为默认值（`TIMESTAMP` 类型支持）。
   - MySQL 8.0+：允许 `DATETIME` 使用 `DEFAULT CURRENT_TIMESTAMP`，但需显式设置 `explicit_defaults_for_timestamp=1`。

2. **字段类型冲突**：  
   - 你的模型中 `LoginOutTime` 定义为 `time.Time`（非指针），GORM 会生成 `NOT NULL` 约束，但未显式指定 `NOT NULL`，导致冲突。

3. **空格问题**：  
   - 标签中 `column:login_out_time ;default` 存在多余空格，可能导致 SQL 解析错误。


#### **二、解决方案**

##### 1. **修改字段类型为 `TIMESTAMP`**
`TIMESTAMP` 类型在 MySQL 中支持 `DEFAULT CURRENT_TIMESTAMP`：
```go
type UserBasic struct {
    gorm.Model
    // 其他字段...
    LoginOutTime time.Time `gorm:"column:login_out_time;type:timestamp;default:CURRENT_TIMESTAMP" json:"login_out_time"`
}
```

##### 2. **显式设置 `NOT NULL` 约束**
确保字段非空，并使用 `DEFAULT CURRENT_TIMESTAMP`：
```go
type UserBasic struct {
    gorm.Model
    // 其他字段...
    LoginOutTime time.Time `gorm:"column:login_out_time;not null;default:CURRENT_TIMESTAMP" json:"login_out_time"`
}
```

##### 3. **使用 `DATETIME` 并调整 MySQL 配置**
若必须使用 `DATETIME`，需修改 MySQL 配置（`my.cnf`）：
```ini
[mysqld]
explicit_defaults_for_timestamp=1
```
然后重启 MySQL 服务。

##### 4. **删除默认值，由代码控制**
移除默认值，在创建记录时显式设置时间：
```go
type UserBasic struct {
    gorm.Model
    // 其他字段...
    LoginOutTime time.Time `gorm:"column:login_out_time" json:"login_out_time"`
}

// 创建时设置默认值
user := UserBasic{
    // 其他字段...
    LoginOutTime: time.Now(),
}
db.Create(&user)
```

##### 5. **使用可空类型（推荐）**
将字段改为指针类型，允许 `NULL` 值：
```go
type UserBasic struct {
    gorm.Model
    // 其他字段...
    LoginOutTime *time.Time `gorm:"column:login_out_time" json:"login_out_time"`
}

// 创建时未退出，默认为 NULL
user := UserBasic{
    // 其他字段...
    LoginOutTime: nil,
}
db.Create(&user)
```


#### **三、验证与测试**
1. **检查 MySQL 版本**：
   ```sql
   SELECT VERSION();
   ```
   - 若为 5.7 及以下，优先使用 `TIMESTAMP` 或可空 `DATETIME`。

2. **查看当前 SQL_MODE**：
   ```sql
   SELECT @@sql_mode;
   ```
   - 确保不包含 `NO_ZERO_DATE` 或 `STRICT_TRANS_TABLES` 限制。

3. **测试表创建**：
   ```sql
   CREATE TABLE test (
       login_out_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- 有效
       -- 或
       login_out_time DATETIME DEFAULT CURRENT_TIMESTAMP   -- 仅 MySQL 8.0+ 有效
   );
   ```


#### **四、最佳实践**
1. **优先使用 `TIMESTAMP` 类型**  
   - 若需要自动更新时间戳，`TIMESTAMP` 是更好的选择。

2. **明确字段约束**  
   - 避免隐式约束，显式指定 `NOT NULL` 或 `NULL`。

3. **保持数据库兼容性**  
   - 若项目可能迁移到其他数据库（如 PostgreSQL），避免依赖 MySQL 特定特性。

4. **代码与数据库解耦**  
   - 时间逻辑尽量在代码层处理，而非依赖数据库默认值。


#### **五、修复示例代码**
```go
// 方案 1：使用 TIMESTAMP 类型
type UserBasic struct {
    gorm.Model
    // 其他字段...
    LoginOutTime time.Time `gorm:"column:login_out_time;type:timestamp;default:CURRENT_TIMESTAMP" json:"login_out_time"`
}

// 方案 2：使用可空 DATETIME
type UserBasic struct {
    gorm.Model
    // 其他字段...
    LoginOutTime *time.Time `gorm:"column:login_out_time" json:"login_out_time"`
}

// 创建用户时
user := UserBasic{
    // 其他字段...
    LoginOutTime: nil,  // 初始未退出
}
db.Create(&user)

// 用户退出时
now := time.Now()
db.Model(&user).Update("login_out_time", &now)
```


通过以上方法，你可以解决 `Invalid default value` 错误，并建立更健壮的日期字段处理机制。