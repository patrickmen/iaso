[English](./README.md) | 简体中文

<h1 align="center">IASO</h1>

<div align="center">
网站设计方案
</div>

## 菜单

```
- 首页
- 关于我们
- 产品和技术
  - 概览
  - 产品
  - 服务
- 合作伙伴
- 技术资源
- 新闻资讯
- 职业发展
- 联系我们
```

## 使用

### 使用命令行
```bash
$ cd ui
$ npm install @material-ui/core
$ npm install
$ npm run start         # 访问 http://localhost:8000
```

### 使用 docker

```bash
# preview 
$ docker pull iaso-frontend
$ docker run -p 8080:80 iaso-frontend
# open http://localhost

# dev 
$ npm run docker:dev

# build 
$ npm run docker:build


# production dev 
$ npm run docker-prod:dev

// production build 
$ npm run docker-prod:build
```
