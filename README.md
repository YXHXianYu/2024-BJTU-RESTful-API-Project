# 2024-BJTU-RESTful-API-Project
> 大三下 ZLG老师 实训课 个人作业

## 1. 简介



## 2. 作业一，RESTful API 设计

### 2.1 用户API

#### 2.1.1 用户注册

* 资源标识：`/api/v1/users`

* 操作：POST

* 权限：无须权限

* 请求Payload

  ```json
  {
      "username": "exampleUser",
      "password": "password",
      "email": "example@email.com",
      "telephone": "10123456789"
  }
  ```

* 响应Payload

  ```json
  {
      "code": 200,
      "message": "Success.",
      "data": {
          "uuid": "00000000-0000-0000-0000-000000000000",
          "username": "exampleUser",
          "password": "password",
          "email": "example@email.com",
          "telephone": "10123456789"
      }
  }
  ```

#### 2.1.2 用户登入

* 资源表示：`/api/v1/sessions`

* 操作：POST

* 权限：无需权限

* 请求Payload

  ```json
  {
      "username": "exampleUser",
      "password": "password"
  }
  ```

* 响应Payload

  ```json
  {
      "code": 200,
      "message": "Success.",
      "data": {
          "uuid": "00000000-0000-0000-0000-000000000000",
          "username": "exampleUser",
          "password": "password",
          "email": "example@email.com",
          "telephone": "10123456789",
          "token": "abcdefg1234567890"
      }
  }
  ```

#### 2.1.3 查询所有用户信息

* 资源表示：`/api/v1/users`

* 操作：GET

* 权限：Admin

* 请求Payload

  * 无

* 响应Payload

  ```json
  {
      "code": 200,
      "message": "Success.",
      "data": [
          {
              "uuid": "00000000-0000-0000-0000-000000000000",
              "username": "exampleUser",
              "password": "password",
              "email": "example@email.com",
              "telephone": "10123456789",
              "token": "abcdefg1234567890"
          }, {
              ...
          }
      ]
  }
  ```

#### 2.1.4 查询单个用户信息

* 资源表示：`/api/v1/users/{uuid}`

* 操作：GET

* 权限：Admin

* 参数

  ```json
  "uuid": "00000000-0000-0000-0000-000000000000"
  ```
  
* 请求Payload

  * 无

* 响应Payload

  ```json
  {
      "code": 200,
      "message": "Success.",
      "data": {
          "uuid": "00000000-0000-0000-0000-000000000000",
          "username": "exampleUser",
          "password": "password",
          "email": "example@email.com",
          "telephone": "10123456789"
      }
  }
  ```

### 2.2 博客内容API

#### 2.2.1 添加博客

* 资源表示：`/api/v1/blogs`

* 操作：POST

* 权限：User

* 请求Payload

  ```json
  {
      "title": "title",
      "content": "content"
  }
  ```

* 响应Payload

  ```json
  {
      "code": 200,
      "message": "Success.",
      "data": {
          "uuid": "00000000-0000-0000-0000-000000000000",
          "title": "title",
          "content": "content",
          "createdUser": "username"
      }
  }
  ```

#### 2.2.2 查询单个博客

* 资源表示：`/api/v1/blogs/{uuid}`

* 操作：GET

* 权限：User

* 参数

  ```json
  "uuid": "00000000-0000-0000-0000-000000000000"
  ```

* 请求Payload

  * 无

* 响应Payload

  ```json
  {
      "code": 200,
      "message": "Success.",
      "data": {
          "uuid": "00000000-0000-0000-0000-000000000000",
          "title": "title",
          "content": "content",
          "createdUser": "username"
      }
  }
  ```

#### 2.2.3 查询所有博客

* 资源表示：`/api/v1/blogs`

* 操作：POST

* 权限：User

* 请求Payload

  * 无

* 响应Payload

  ```json
  {
      "code": 200,
      "message": "Success.",
      "data": [
          {
              "uuid": "00000000-0000-0000-0000-000000000000",
              "title": "title",
              "content": "content",
              "createdUser": "username"
          }, {
              ...
          }
      ]
  }
  ```

#### 2.2.4 更新单个博客

* 资源表示：`/api/v1/blogs/{uuid}`

* 操作：PATCH

* 权限：对应User

* 参数

  ```json
  "uuid": "00000000-0000-0000-0000-000000000000"
  ```

* 请求Payload

  ```json
  {
      "title": "New title",
      "content": "Updated content"
  }
  ```

* 响应Payload

  ```json
  {
      "code": 200,
      "message": "Success.",
      "data": {
          "uuid": "00000000-0000-0000-0000-000000000000",
          "title": "New title",
          "content": "Updated content",
          "createdUser": "username"
      }
  }
  ```

### 2.3 第三方鉴权API

### 2.4 API工具说明

* 我使用了ChatGPT4
* 我的Prompt：作业要求 + “请给你给我一个大概的思路和文档框架。并举一个接口作为例子”
* **GPT4提供了用户注册和登入的接口**
* 通过观察GPT4提供的文档框架、例子，**我自己补全了其他的接口**

## 3. 作业二，RESTful API 后端设计与实现



## 4. 作业三，博客 前端的实现

## 5. 作业四，作业二/三 集成测试的实现
