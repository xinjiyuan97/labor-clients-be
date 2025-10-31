# 品牌表联系人信息优化

## 概述

将品牌表中的联系人信息从直接存储字段（contact_person, contact_phone, contact_email）改为关联用户ID（contact_user_id），实现联系人信息的统一管理和复用。

## 修改内容

### 1. 数据库表结构修改

#### brands表新增字段

```sql
-- 基本信息
company_short_name VARCHAR(50) COMMENT '公司简称'
website VARCHAR(255) COMMENT '公司网站'

-- 公司信息
industry VARCHAR(50) COMMENT '所属行业'
company_size VARCHAR(20) COMMENT '公司规模'
credit_code VARCHAR(50) COMMENT '统一社会信用代码'
company_address VARCHAR(255) COMMENT '公司地址'
business_scope TEXT COMMENT '经营范围'
established_date DATE COMMENT '成立日期'
registered_capital DECIMAL(15,2) COMMENT '注册资本'

-- 联系人信息（存储user_id）
contact_user_id BIGINT COMMENT '联系人用户ID'
contact_position VARCHAR(50) COMMENT '联系人职位'

-- 证件信息
id_card_number VARCHAR(50) COMMENT '身份证号'
id_card_front VARCHAR(255) COMMENT '身份证正面照URL'
id_card_back VARCHAR(255) COMMENT '身份证反面照URL'
business_license VARCHAR(255) COMMENT '营业执照URL'
tax_certificate VARCHAR(255) COMMENT '税务登记证URL'
org_code_certificate VARCHAR(255) COMMENT '组织机构代码证URL'
bank_license VARCHAR(255) COMMENT '开户许可证URL'
other_certificates TEXT COMMENT '其他证件URL（JSON数组）'

-- 财务信息
bank_account VARCHAR(50) COMMENT '银行账号'
settlement_cycle VARCHAR(20) COMMENT '结算周期'
deposit_amount DECIMAL(15,2) COMMENT '保证金金额'

-- 状态信息
account_status ENUM('active', 'disabled', 'frozen') DEFAULT 'active' COMMENT '账号状态'
```

#### 数据库迁移文件

- **文件**: `migrations/update_brands_table.sql`
- **说明**: 使用ALTER TABLE语句添加所有新字段和索引

### 2. 数据模型更新

#### models/brand.go

更新Brand结构体，添加所有新字段：

```go
type Brand struct {
    BaseModel
    
    // 基本信息
    Name             string
    CompanyShortName string
    Logo             string
    Description      string
    Website          string
    
    // 公司信息
    Industry           string
    CompanySize        string
    CreditCode         string
    CompanyAddress     string
    BusinessScope      string
    EstablishedDate    *time.Time
    RegisteredCapital  float64
    
    // 联系人信息（存储user_id）
    ContactUserID   *int64
    ContactPosition string
    
    // 证件信息
    IDCardNumber        string
    IDCardFront         string
    IDCardBack          string
    BusinessLicense     string
    TaxCertificate      string
    OrgCodeCertificate  string
    BankLicense         string
    OtherCertificates   string
    
    // 财务信息
    BankAccount     string
    SettlementCycle string
    DepositAmount   float64
    
    // 状态信息
    AuthStatus    string
    AccountStatus string
}
```

### 3. Thrift接口定义更新

#### idls/admin.thrift

更新`UpdateBrandReq`结构体，添加所有可更新字段：

```thrift
struct UpdateBrandReq {
    1: i64 brand_id
    2: string company_name
    3: string company_short_name
    4: string logo
    5: string description
    6: string website
    7: string industry
    8: string company_size
    9: string credit_code
    10: string company_address
    11: string business_scope
    12: string established_date
    13: double registered_capital
    14: string contact_person
    15: string contact_phone
    16: string contact_email
    17: string contact_position
    18: string id_card_number
    19: string id_card_front
    20: string id_card_back
    21: string business_license
    22: string tax_certificate
    23: string org_code_certificate
    24: string bank_license
    25: string other_certificates
    26: string bank_account
    27: string settlement_cycle
    28: double deposit_amount
    29: string auth_status
    30: string account_status
}
```

### 4. 业务逻辑更新

#### biz/logic/admin/create_brand.go

**主要修改**：

1. 在创建品牌之前，先创建或获取联系人用户
2. 将联系人用户ID设置到品牌的`contact_user_id`字段
3. 品牌创建成功后，为联系人分配品牌管理员角色

**核心流程**：

```go
// 1. 创建或获取联系人用户
contactUser, err := createBrandAdminForContact(req.ContactPhone, req.ContactPerson, 0)

// 2. 创建品牌，关联联系人用户ID
brand := &models.Brand{
    Name:             req.CompanyName,
    CompanyShortName: req.CompanyShortName,
    ContactUserID:    &contactUser.ID,
    // ... 其他字段
}

// 3. 为联系人分配品牌管理员角色
assignBrandAdminRole(contactUser.ID, brandID)
```

**辅助函数**：

- `createBrandAdminForContact()`: 根据手机号创建或获取用户，返回用户对象
- `assignBrandAdminRole()`: 为用户分配品牌管理员角色

#### biz/logic/admin/update_brand.go

**主要修改**：

1. 支持更新所有新增字段
2. 如果提供了新的联系电话，自动创建或获取联系人用户并更新`contact_user_id`
3. 使用空值判断，只更新提供的字段

**关键代码**：

```go
// 更新联系人（如果提供了新的联系电话）
if req.ContactPhone != "" {
    contactUser, err := createBrandAdminForContact(req.ContactPhone, req.ContactPerson, 0)
    if err != nil {
        utils.Warnf("更新联系人失败: %v", err)
    } else {
        brand.ContactUserID = &contactUser.ID
    }
}
```

#### biz/logic/admin/get_brand_detail.go

**主要修改**：

1. 在返回品牌详情时，根据`contact_user_id`查询用户信息
2. 将用户的姓名、电话填充到响应对象中

**关键代码**：

```go
// 查询联系人信息
if brand.ContactUserID != nil && *brand.ContactUserID > 0 {
    contactUser, err := mysql.GetUserByID(mysql.DB, *brand.ContactUserID)
    if err == nil && contactUser != nil {
        brandDetail.ContactPerson = contactUser.Username
        brandDetail.ContactPhone = contactUser.Phone
    }
}
```

## 优势

### 1. 数据一致性
- 联系人信息统一存储在users表，避免数据冗余
- 联系人信息更新时，所有关联品牌自动同步

### 2. 权限管理
- 联系人自动关联用户账号
- 可以为联系人分配品牌管理员角色
- 支持多品牌场景下的权限管理

### 3. 扩展性
- 未来可以轻松添加联系人的其他信息（邮箱、地址等）
- 支持一个用户管理多个品牌
- 支持一个品牌有多个联系人（通过user_roles表）

### 4. 用户体验
- 联系人可以直接使用手机号登录
- 首次创建品牌时自动创建账号（默认密码123456）
- 账号信息可以在用户中心统一管理

## 使用示例

### 创建品牌

```bash
POST /api/v1/admin/brands
{
    "company_name": "示例科技有限公司",
    "company_short_name": "示例科技",
    "contact_person": "张三",
    "contact_phone": "13800138000",
    "contact_position": "总经理",
    "industry": "互联网",
    "company_size": "100-500人",
    "credit_code": "91110000XXXXXXXXXX",
    "company_address": "北京市朝阳区XXX",
    "registered_capital": 1000000.00,
    // ... 其他字段
}
```

**系统会自动**：
1. 检查手机号13800138000是否已注册
2. 如未注册，创建新用户（用户名：张三，密码：123456）
3. 创建品牌，将contact_user_id设置为该用户ID
4. 为该用户分配品牌管理员角色

### 更新品牌

```bash
PUT /api/v1/admin/brands/:brand_id
{
    "company_name": "新公司名称",
    "contact_phone": "13900139000",  // 更换联系人
    "contact_person": "李四",
    // ... 其他字段
}
```

**系统会自动**：
1. 查找或创建手机号13900139000对应的用户
2. 更新品牌的contact_user_id为新用户ID
3. 更新其他提供的字段

### 获取品牌详情

```bash
GET /api/v1/admin/brands/:brand_id
```

**响应包含**：
- 品牌的所有信息
- 从users表查询的联系人姓名和电话
- 其他关联信息

## 注意事项

1. **数据迁移**：
   - 需要执行`migrations/update_brands_table.sql`迁移文件
   - 旧数据的联系人信息需要手动迁移（可选）

2. **默认密码**：
   - 自动创建的用户默认密码为`123456`
   - 建议提醒用户首次登录后修改密码

3. **字段验证**：
   - contact_phone应进行手机号格式验证
   - credit_code应进行统一社会信用代码格式验证
   - 证件图片URL应验证可访问性

4. **权限控制**：
   - 品牌管理员只能修改自己管理的品牌
   - 系统管理员可以修改所有品牌

5. **联系人变更**：
   - 更换联系人时，不会自动删除原联系人的品牌管理员角色
   - 需要手动管理角色分配

## 相关文档

- [门店管理功能](./store_management.md)
- [品牌管理员管理](./brand_admin_management.md)
- [自动创建品牌管理员](./auto_create_brand_admin.md)
- [用户角色管理](./user_roles.md)

## 更新日志

- 2025-10-24: 初始版本，完成brands表结构优化和相关业务逻辑更新

