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

