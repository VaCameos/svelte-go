

### fe
 前置工作 npm i pnpm （无pnpm）
 
安装依赖 pnpm i
 
 跑 pnpm dev
 
 
 
 ### docker
 拉取postgres镜像 
 run的时候配置环境变量 POSTGRES_PASSWORD 为123456
 
 拉取pgadmin镜像 ，dpage那个
 接着创建服务、创建数据库参考以下
 https://www.cnblogs.com/xhznl/p/13155054.html
 创建的数据库名为sys
 跑go代码 
 
 ### 演示流程
 - 先创建账号 admin 密码admin  这个号是管理员号
 - 登陆后会跳转到管理账号页面， 演示添加删除更新
 - 演示普通账号的登录 会跳转到个人信息页面，这改页面中演示编辑功能 
