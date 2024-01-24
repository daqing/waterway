About
=====

Waterway is a full-stack web framework written in Go, inspired by Ruby on Rails.

**[查看中文文档](https://github.com/daqing/waterway/blob/main/docs/zh-CN/README.md)**

Get Started
===========

## 1. Setup project skeleton

Use `gonew` to create a new project based on `waterway`:

```bash
$ gonew github.com/daqing/waterway example.com/foo/bar
```

Replace `example.com/foo/bar` with your real module name.

## 2. Setup local development environment

### Create `.env` file

```bash
$ cp .env.example .env
```

This file defines a few environment variables:

- WATERWAY_PG_URL
  **the URL string for connecting to PostgreSQL**
- WATERWAY_PORT
  **the port to listen on**
- WATERWAY_STORAGE_DIR
  **the full path to store the uploaded files**
- WATERWAY_ASSET_HOST
  **the CDN url for serving assets**
- WATERWAY_PWD
  **the path to current working directory**
