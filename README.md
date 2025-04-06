# game-rpc
游戏服务 RPC。

# 模块介绍

本模块主要提供基本的会员服务，包括会员注册、登录、注销、修改密码、修改个人信息等。以及会员等级、积分、余额等相关服务。

同时提供了第三方登录管理，token 黑名单等功能。

## 目前支持

任意 smtp 邮箱

短信服务： 阿里云，腾讯云， 合一短信，短信宝

使用短信宝作为服务方时，coreAPI的配置文件中短信验证的SmsTemplateId的值需为captcha
---

Simple Admin Message Center module. Simple Admin Message Center module, responsible for SMS, internal announcements, email sending, etc.

## Currently supported

Any SMTP mailbox

SMS service: Alibaba Cloud, Tencent Cloud, Uni SMS,SMSBao

When using SMSBao as a service provider, the SmsTemplateId for SMS verification in the coreAPI configuration file should be captcha