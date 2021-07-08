English | [简体中文](./README.zh-CN.md)

<h1 align="center">IASO</h1>

<div>
A backend management system for website design. The frontend UI uses Ant Design Pro v2.0 as scaffolding and backend service is based on Gin Framework.
</div>

## Menu

```
- Home
- About US
- Technology
  - Gene To Protein Platform
  - CADD Platform
  - SBDD Platform
  - DEL Platform
- Pipeline
- Partnering
- News
- Careers
- Contact US
```

## Prerequest
1. [Install/Setup Golang development environment](https://golang.org/)
2. Setup Go envs
3. Install/Setup npm & nodejs
4. cd ui & npm install @material-ui/core & npm install @material-ui/icons

## Contributing
#### backend
```bash
go run main.go -verbose
```
#### frontend
```bash
$ cd ui
$ npm run start         # visit http://localhost:8000
```

### Build & Compile
```bash
$ cd ui/
$ npm run build
$ cd iaso/
$ make docker-build
```
