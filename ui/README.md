English | [简体中文](./README.zh-CN.md)

<h1 align="center">IASO</h1>

<div align="center">
An solution for website design.
</div>

## Menu

```
- HOME
- ABOUT US
- PRODUCTS
  - SUMMARY
  - PRODUCTS
  - SERVICES
- PARTNERING
- RESOURCES
- NEWS
- CAREERS
- CONTACT US
```

## Usage

### Use bash

```bash
$ cd ui
$ npm install @material-ui/core
$ npm install
$ npm start         # visit http://localhost:8000
```

### Use by docker

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

# production build
$ npm run docker-prod:build
```