
#### 功能TODO

- [x] 注册
- [x] 登录
- [ ] 邀请另一个用户一起记账
- [x] 消费录入
    - [x] 消费金额
    - [x] 消费类型
    - [x] 支付渠道
    - [x] 消费时间
    - [x] 一句话描述
    - [ ] 转账、还款类型(支付宝转入银行卡,等等..)
    - [ ] 转账 到另一个用户
- [ ] 快速录入
    - [ ] 自定义消费模板
    - [ ] 模板直接录入
- [ ] 消费信息修改
- [ ] job录入（车贷房贷、工资等）
- [ ] 快速修改job录入的记录
- [ ] 预支付（信用卡花呗支出会计算到还款日）
- [ ] 账单查询
    - [ ] 日期范围查询
    - [ ] 单独用户查询
    - [ ] 消费类型查询
    - [ ] 账单分析
- [ ] 消费类型权重查询
- [ ] 余额录入、修改

#### 改进项
- [x] 各项的余额不应该靠每次加减维护,而是应该统计出来.按月统计
- [ ] 以用户id直接作为登陆token,后续可以加入混合时间戳都会形式,来实现单点登陆
- [ ] userController 有一些查询语句, 应统一代码风格
- [ ] 小伙伴一起记账邀请提醒,同意拒绝等状态. 需要一张消息表存放消息
- [ ] 用户之间的关联关系表