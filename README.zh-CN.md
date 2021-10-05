[English](./README.md) | 简体中文

<h1 align="center">IASO</h1>

  一个网站的后台管理系统。前端采用[Ant Design Pro v2.2.0](https://github.com/ant-design/ant-design-pro/releases/tag/2.2.0)作为脚手架，后端基于[Gin框架](https://github.com/gin-gonic/gin)开发。

## 菜单

```
- 首页
- 关于我们
- 技术平台
  - 靶点验证平台
  - SBDD平台
  - Biomarker开发平台
- 研发管线
- 合作伙伴
  - 科研机构
  - 生物制药公司
- 新闻资讯
- 职业发展
- 联系我们
```

## 前提条件
1. [安装和配置golang开发环境](https://golang.org/)
3. 安装 npm & nodejs
4. cd ui & npm install @material-ui/core & npm install @material-ui/icons & npm install --save rc-texty

## 贡献
#### 后端
```bash
go run main.go -verbose
```
#### 前端
```bash
$ cd ui
$ npm run start         # visit http://localhost:8000
```

### 编译和打包
```bash
$ cd ui/
$ npm run build
$ cd iaso/
$ make docker-build
```
