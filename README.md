## RestChat
### Функциональные требования
Регистрация (register)
Авторизация (login)
Один пользователь - одна сессия
Отображение списка онлайн пользователей
Отображение списка последних N сообщений
Автоматическое обновление списка пользователей
Автоматическое обновление сообщений
Отправка сообщения
Деавторизация (logout)
Ограничения формата пользователя
Ограничения формата сообщения

### Архитектурные требования
Клиент-серверная архитектура
REST api
Формат данных между клиентом и сервером Json
Бэкэнд - веб-свервер на Golang
Фронтэнд - SPA на React.js
Контейнеризация (Docker, Docker-compose)
? Интеграционные end to end тесты (selenium)
? CI/CD

### REST API

#### Ошибка формата ЗАПРОСА

(наименование полей в json запросе)
(например запрос обязательно должен содержать поля username и password)
400 {error: "error_message"} - ошибка формата данных

#### POST /api/login

Клиентские данные:

```json
 {username: "string", password: "string"}
```

Серверные данные:
 Успех: 200 OK
```json
{auth_token: "string", member: {id: 12, name: "vasya"}}
```
Ошибка: 401 Unauthorized, {error: "error_message"}

---

#### POST /api/user
Клиентские данные:
 {username: "string", password: "string"}
Серверные данные:
 Успех: 201 Created, {username: "string"}
	Ошибка:
  1) пользователь уже существует
   403 Forbidden, {error: "error_message"}
  2) не правильный формат логина/пароля
   400 Bad Request

---

#### POST /api/logout
Клиентские данные:
 {api_token: "string"}
Серверные данные:
 Успех: 200 OK, {}
	Ошибка:
  1) Неправильный токен
   400 Bad Request {error: "error_message"}

---


#### GET /api/members
Клиентские данные:
 {api_token: "string"}
Серверные данные:
 Успех: 200 OK:
 {members: [
  {id: 12, name: "vasya"},
  {id: 14, name: "petya"},
 ]}
   {members: []} // пустой список пользователей
    Ошибка:
  1) Неправильный токен
   400 Bad Request {error: "error_message"}

---


#### GET /api/messages
Клиентские данные:
 {api_token: "string"}
Серверные данные:
 Успех: 200 OK:
 {messages: [
  {id: 1002, member_name: "petya", text: "hello", time: "13:28:77"},
        {id: 1001, member_name: "vasya", text: "hello", time: "13:23:77"},
 ]}
   {messages: []} // Если сообщения отсутствуют
    Ошибка:
  1) Неправильный токен
   400 Bad Request {error: "error_message"}

---

#### POST /api/message
Клиентские данные:
 {api_token: "string", text: "string"}
Серверные данные:
 Успех: 200 OK:
 {message: {
  id: 1002,
  member_name: "petya",
  text: "hello",
  time: "13:28:77"
 },
    Ошибка:
  1) Неправильный токен
   400 Bad Request {error: "error_message"}

---
