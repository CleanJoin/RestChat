## RestChat

### Описание проекта

Данный проект создан в обучающих целях. Назначение - обучение простейшим навыкам проектирования, командной работы над проектом. Тренировка навыков программирования. Использования тестового окружения с применением контейнеризации.

Приложение представляет собой клиент-серверное приложение "веб-чат" с простейшими функциями регистрации, авторизации, отправки и приема сообщений, вывода списка онлайн участников.

Взаимодействие между клиентом и сервером производится посредством Rest-like HTTP API.

### Технологический стек

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![React](https://img.shields.io/badge/react-%2320232a.svg?style=for-the-badge&logo=react&logoColor=%2361DAFB) ![HTML5](https://img.shields.io/badge/html5-%23E34F26.svg?style=for-the-badge&logo=html5&logoColor=white) ![Bootstrap](https://img.shields.io/badge/bootstrap-%23563D7C.svg?style=for-the-badge&logo=bootstrap&logoColor=white)  ![JavaScript](https://img.shields.io/badge/javascript-%23323330.svg?style=for-the-badge&logo=javascript&logoColor=%23F7DF1E) ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white) ![Nginx](https://img.shields.io/badge/nginx-%23009639.svg?style=for-the-badge&logo=nginx&logoColor=white) ![Git](https://img.shields.io/badge/git-%23F05033.svg?style=for-the-badge&logo=git&logoColor=white) ![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)

### Функциональные требования

- Регистрация (register)
- Авторизация (login)
- Один пользователь - одна сессия
- Отображение списка онлайн пользователей
- Отображение списка последних N сообщений
- Автоматическое обновление списка пользователей
- Автоматическое обновление сообщений
- Отправка сообщения
- Деавторизация (logout)
- Ограничения формата пользователя
- Ограничения формата сообщения

### Архитектурные требования

- Клиент-серверная архитектура
- REST api
- Формат данных между клиентом и сервером JSON
- Бэкэнд - веб-свервер на Golang
- Фронтэнд - SPA на React.js
- Контейнеризация (Docker, Docker-compose)
- ? Интеграционные end to end тесты (selenium)
- ? CI/CD

### Макеты страниц

#### Страница авторизации

![img](./images/restchat_client_layout_1.svg)

- MessageBox - опциональное поле сообщений об ошибках
- login, password - текстовые поля ввода логина/пароля (общие для логина и регистрации)
- Login - кнопка запроса авторизации
- Register - кнопка запроса регистрации

#### Главное окно чата

![img](./images/restchat_client_layout_2.svg)

- (левая панель) - панель со списком онлайн участников чата
- (правая панель) - панель сообщений чата (последние внизу)
- Logout - кнопка деавторизации (выход из чата)
- Message - текстовое поле для отправки сообщения (отправка по нажатию Enter)

### REST API

#### Общая ошибка формата ЗАПРОСА

Общая ошибка для наименования или количества полей в json запросе от клиента. Например, запрос на авторизацию обязательно должен содержать поля username и password, ошибка выдается в случае нарушения данных требований.

Код ошибки: 400 Bad Request
Тело сообщения:

```json
{error: "error_message"}
```

---

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

```json
{username: "string", password: "string"}
```

Серверные данные:
 Успех: 201 Created,

```json
{username: "string"}
```

	Ошибка:
  1) пользователь уже существует
   403 Forbidden
```json
{error: "error_message"}
```
  2) не правильный формат логина/пароля
   400 Bad Request

---

#### POST /api/logout
Клиентские данные:
```json
{api_token: "string"}
```
Серверные данные:
 Успех: 200 OK, {}
	Ошибка:
  1) Неправильный токен
   400 Bad Request
```json
{error: "error_message"}
```

---

#### GET /api/members
Клиентские данные:
```json
{api_token: "string"}
```
Серверные данные:
 Успех: 200 OK:
```json
{
    members: [
        {id: 12, name: "vasya"},
        {id: 14, name: "petya"},
    ]
}
```
Либо, в случае пустого списка пользователей:
```json
{members: []}
```
Ошибка:
  1) Неправильный токен
   400 Bad Request
```json
{error: "error_message"}
```

---

#### GET /api/messages
Клиентские данные:
```json
{api_token: "string"}
```
Серверные данные:
 Успех: 200 OK:

```json
{
    messages: [
        {id: 1002, member_name: "petya", text: "hello", time: "13:28:77"},
        {id: 1001, member_name: "vasya", text: "hello", time: "13:23:77"},
    ]
}
```

Либо в случае отсутствия сообщений, пустой ответ:

```json
{messages: []}
```
    Ошибка:
  1) Неправильный токен
   400 Bad Request

```json
{error: "error_message"}
```

---

#### POST /api/message

Клиентские данные:

```json
{api_token: "string", text: "string"}
```

Серверные данные:
 Успех: 200 OK:
```json
{message: {
  id: 1002,
  member_name: "petya",
  text: "hello",
  time: "13:28:77"
 },
```
Ошибка:
1) Неправильный токен - 400 Bad Request

```json
{error: "error_message"}
```
