
#### 优先改进项
- [x] 列表页筛选
- [ ] 消费大类
- [x] 重新设计 api 使用 beego
- [x] token加密
- [ ] 代码中的todo急需解决

#### 改进项
- [x] 各项的余额不应该靠每次加减维护,而是应该统计出来.按月统计
- [x] 以用户id直接作为登陆token,后续可以加入混合时间戳都会形式,来实现单点登陆
- [x] userController 有一些查询语句, 应统一代码风格
- [x] 小伙伴一起记账邀请提醒,同意拒绝等状态. 需要一张消息表存放消息
- [x] 用户之间的关联关系
- [ ] 头像压缩, 上传组件不够好用
- [ ] 销账金额还原
- [ ] 修改消费记录 余额也要随之变动, 计划是走销账流程重新添加消费记录
- [x] 首页仪表盘值过大过小都有bug
- [ ] 账单的备注功能太过鸡肋
- [x] 根据响应码给出统一的提示
- [ ] 废除各项余额的记录
- [x] 消费类型 前后端排序规则统一
- [ ] 账单的类型筛选加入"其他"类型

#### 功能TODO

- [x] 注册
- [x] 登录
- [x] 邀请另一个用户一起记账
- [x] 消费录入
    - [x] 消费金额
    - [x] 消费类型
    - [x] 支付渠道
    - [x] 消费时间
    - [x] 一句话描述
    - [ ] 转账、还款类型(支付宝转入银行卡,等等..)
    - [ ] 转账 到另一个用户
- [x] 消费记录修改
- [x] 销账
- [ ] 快速录入
    - [ ] 自定义消费模板
    - [ ] 模板直接录入
- [x] 消费信息修改
- [x] 消费信息发送给小伙伴
- [ ] job录入（车贷房贷、工资等）
- [ ] 快速修改job录入的记录
- [ ] 预支付（信用卡花呗支出会计算到还款日）
- [x] 账单查询
    - [x] 日期范围查询
    - [x] 单独用户查询
    - [x] 消费类型查询
    - [x] 账单分析
- [ ] 消费类型权重查询
- [x] 余额录入、修改
- [x] 头像修改裁剪功能

