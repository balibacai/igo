# 用户身份验证


## 1）身份验证接口

| 事项 | 描述 | 备注 |
| :--- | :--- | :--- |
| 接口地址 | POST /v1/login |  无 |
| 对接人 |  mist |  无 |

**需求变更历史 :**

| 日期 | 需求说明 | 备注 |
| :--- | :--- | :--- | 
| 20181010 | 支持邮箱密码方式登录   | 无 |


**请求参数 :**

| 参数名 | 字段类型 | 字段说明 | 备注 |
| :--- | :--- | :--- | :--- |
| email  | string | 邮箱 | 必填 |
| password  | string | 密码 | 必填 |

**响应参数**：

| 参数名 | 字段类型 | 字段说明 | 备注 |
| :--- | :--- | :--- | :--- |
| error | int | 错误代码 | 为0表示检测通过；大于0表示错误信息 |
| msg | string | 错误摘要 | 大于0表示错误信息 |
| data | object | 数据/错误详情 | 当error为0时，此处为数据信息 |


**data字段协议**：

| 字段 | 字段类型 | 字段说明 | 备注 |
| :--- | :--- | :--- | :--- |
| token | string | 验证成功后的token凭证 |  |
| expiredAt | int | token过期时间，绝对时间，精确到秒 |  |


demo如下：

```js
{
    "code": 0,
    "msg": ""
    "data": {
        "token": "eyJhbGfiOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEyMzQ1NjcsImV4cCI6MTU0MTY4MjQ2NiwiaXNzIjoiaWdvIn0.kZsh9SXaDxZRTQIbO07hXpyyhsw3WDPQ6QF2q3HgZxE",
        "expiredAt": 1541682466
    }
}
```

**code码说明 :**

| CODE ID | 说明 |
| :--- | :--- |
| 100001 | error occurs when parsing login form |
| 100002 | error occurs when validating login credentials | 
| 100003 | login credentials invalid |
| 100004 | error occurs when generating login token |