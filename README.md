![RestChat](./client/src/Images/RestChat_logo.png)

---

### Описание проекта

Данный проект создан в обучающих целях. Назначение - обучение простейшим навыкам проектирования, командной работы над проектом. Тренировка навыков программирования. Использования тестового окружения с применением контейнеризации.

Приложение представляет собой клиент-серверное приложение "веб-чат" с простейшими функциями регистрации, авторизации, отправки и приема сообщений, вывода списка онлайн участников.

Взаимодействие между клиентом и сервером производится посредством RESTFUL HTTP API.

---

### Технологический стек

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![HTML5](https://img.shields.io/badge/html5-%23E34F26.svg?style=for-the-badge&logo=html5&logoColor=white) ![MUI](https://img.shields.io/badge/MUI-%230081CB.svg?style=for-the-badge&logo=material-ui&logoColor=white) ![JavaScript](https://img.shields.io/badge/javascript-%23323330.svg?style=for-the-badge&logo=javascript&logoColor=%23F7DF1E) ![React](https://img.shields.io/badge/react-%2320232a.svg?style=for-the-badge&logo=react&logoColor=%2361DAFB)  ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white) ![cypress](https://img.shields.io/badge/-cypress-%23E5E5E5?style=for-the-badge&logo=cypress&logoColor=058a5e) ![Nginx](https://img.shields.io/badge/nginx-%23009639.svg?style=for-the-badge&logo=nginx&logoColor=white) ![Git](https://img.shields.io/badge/git-%23F05033.svg?style=for-the-badge&logo=git&logoColor=white) ![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)

---

### Среда для разработки и тестирования

Тестовая среда организована с помощью контейнеризации и утилиты оркестрации Docker-Compose.

Для разворачивания среды разработки необходимо установить [Docker](https://docker.org/), [Docker-Compose](https://github.com/docker/compose) и [Git](https://git-scm.com/)

```bash
# Ubuntu 18.04, 20.04
sudo apt install docker.io docker-compose git
```

Далее клонировать репозиторий:

```bash
git clone https://github.com/CleanJoin/RestChat.git

# Текущая разрабатываемая версия приложения в ветке development
git checkout development
```

Запуск тестовой среды:

```bash
docker compose up --build
```

Полное веб-приложение (бэкенд + фронтенд) доступно по адресу http://localhost:10000 (адреса, включающие /api/ перенаправляются - бэкенду, остальное - фронтенду)


**Описание контейнеров тестовой среды**:

| Название контейнера | Адрес                      | Описание                                                                                       |
|---------------------|----------------------------|------------------------------------------------------------------------------------------------|
| proxy               | http://localhost:10000/    | nginx сервер, работающий в качестве reverse-proxy, перенаправляет запросы бэкенду и фронтенду. |
| restchat_client     | http://localhost:3000/     | Сервер разработки react с автоматической пересборкой проекта и перезапуском.                   |
| restchat_server     | http://localhost:8000/api/ | Сервер разработки codegangsta/gin с автоматической пересборкой проекта и перезапуском.         |

Тестовые серверы как клиента (фронтенд), так и сервера (бэкенд) настроены на автоматическую пересборку проекта и перезапуск сервера в случае изменений в исходных кодах в директории проекта. Директории с исходным кодом примонтированы внутри контейнеров тестового окружения.

**Запуск тестов**

**Тесты бэкенда**:

```bash
docker-compose run restchat_server go test restchat-server/restchat
```

**Тесты фронтенда**:

```bash
docker-compose run restchat_client yarn test
```

**Функциональные тесты**

```bash
docker-compose run functional_tests yarn cypress run
```

Запуск в интерактивном режиме без контейнера:

```bash
cd functional_tests
yarn install
yarn start
```

---

### Организация репозитория исходных кодов

В ветке **master** всегда должна быть текущая стабильная, полностью протестированная версия приложения.

Вся разработка ведется в ветке **development**. Решение о выпуске релизной версии посредством вливания ветки **development** в **master** принимается совместно командой разработки после проведения тестирования проекта.

Разработка новых функций (фичей) и исправление ошибок (фиксов) производится в отдельных ответвлениях от ветки **development**.

Для каждой новой фичи создается новая ветка с именем **feature_{feature_name}**.

```bash
git checkout -b feature_members_window
```

Для исправлений, требующих несколько коммитов создается ветка **fix_{fix_name}**. Допустимо небольшие исправления отправлять напрямую в **development**.

```bash
git checkout -b fix_some_bug
```

После окончания работы над новой веткой, она вливается обратно в **development** и изменения отправляются в главный репозиторий. Работа считается завершенной после того, как:

- Выполнена поставленная задача
- Написаны тесты, покрывающие новые функции
- Написаны тесты, учитывающие ситуацию с исправляемой ошибкой
- Тесты успешно выполняются

```bash
git checkout development
git merge feature_members_window
git merge fix_some_bug
git push
```

Каждый коммит должен быть атомарным (включать небольшие, завершенные по смыслу изменения) и должен быть сопровожден осмысленным комментарием (сообщением коммита), кратко излагающим суть изменений. В случае, если коммит содержит изменения, относящаяся к какой-либо задаче (issue) в [github](https://github.com/CleanJoin/RestChat/issues), номер задачи должен быть указан в сообщении коммита (#1, #2 #3 и т.д.).

---

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

---

### Архитектурные требования

- Клиент-серверная архитектура
- REST api
- Формат данных между клиентом и сервером JSON
- Бэкенд - веб-сервер на Golang
- Фронтенд - SPA на React.js
- Контейнеризация (Docker, Docker-compose) тестовой среды
- ? Интеграционные end to end тесты (selenium)
- ? CI/CD

---

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

---

### REST API

#### **Форматы данных и ограничения**

Формат даты и времени: ISO 8601, часовой пояс UTC

Пример:

```json
{"time": "2022-02-04T08:02:56Z"}
```

Максимальная длина логина: 16 символов.

Разрешенные символы для логина: [a-zA-Z0-9_-].

Разрешенные символы для пароля: [a-zA-Z0-9_-].

Максимальная длина пароля: 32 символа.

Максимальная длина сообщения: 1024 символа.

---

#### **Общая ошибка формата запроса от клиента**

Общая ошибка для наименования или количества полей в json запросе от клиента. Например, запрос на авторизацию обязательно должен содержать поля username и password, ошибка выдается в случае нарушения данных требований.

**Код ошибки**:

```http
400 Bad Request
```

**Тело сообщения**:

```json
{"error": "error_message"}
```

---

#### **POST /api/login**

**Описание**: Авторизация с помощью логина и пароля. В случае успеха сервер возвращает токен авторизации и информацию о пользователе (участнике чата). После авторизации создается новая сессия и пользователь считается онлайн до момента деавторизации (удаления сессии). Если ранее у пользователя была создана активная сессия, то она удаляется и заменяется новой (старый пользователь деавторизован).

**Данные запроса**:

```json
{"username": "string", "password": "string"}
```

**Успешная авторизация**:

```http
200 OK
```

```json
{"auth_token": "string", "member": {"id": 12, "name": "vasya"}}
```

**Ошибка - Неправильный логин/пароль**:

```http
401 Unauthorized
```

```json
{"error": "error_message"}
```

---

#### **POST /api/user**

**Описание**: Создание нового пользователя (регистрация) с заданным паролем. В случае успеха клиенту возвращается имя созданного пользователя.

**Данные запроса**:

```json
{"username": "string", "password": "string"}
```

**Пользователь успешно создан**:

```http
201 Created
```

```json
{"username": "string"}
```

**Ошибка 1 - Пользователь уже существует**:

```http
403 Forbidden
```

```json
{"error": "error_message"}
```

**Ошибка 2 - Неправильный формат логина/пароля**:

```http
400 Bad Request
```

```json
{"error": "error_message"}
```

---

#### **POST /api/logout**

**Описание**: Деавторизация пользователя (выход из чата, закрытие, удаление текущей сессии). После деавторизации пользователь больше не считается онлайн.

**Данные запроса**:

```json
{"api_token": "string"}
```

**Пользователь успешно деавторизован**:

```http
200 OK
```

```json
{}
```

**Ошибка - Неправильный токен**:

```http
400 Bad Request
```

```json
{"error": "error_message"}
```

---

#### **POST /api/members**

**Описание**: Получения списка участников чата (онлайн пользователей, авторизованных). Список передается только авторизованным пользователям, корректно предъявившим токен авторизации.

**Данные запроса**:

```json
{"api_token": "string"}
```

**Успешно получен список онлайн участников чата**:

```http
200 OK
```

```json
{
	"members": [
		{"id": 12, "name": "vasya"},
		{"id": 14, "name": "petya"}
	]
}
```
Успех, но в случае пустого списка пользователей:

```http
200 OK
```

```json
{"members": []}
```

**Ошибка - Неправильный токен авторизации**:

```http
400 Bad Request
```

```json
{"error": "error_message"}
```

---

#### **POST /api/messages**

**Описание**: Получения полного списка последних сообщений в чате. Доступен только для авторизованных пользователей, предъявивших корректный токен авторизации.

**Данные запроса**:

```json
{"api_token": "string"}
```

**Успешное получение списка последних сообщений**:

```http
200 OK
```

```json
{
	"messages": [
		{
			"id": 1002,
			"member_name": "petya",
			"text": "hello",
			"time": "2022-02-04T13:28:77Z"
		},
		{
			"id": 1001,
			"member_name": "vasya",
			"text": "hello",
			"time": "2022-02-04T13:23:77Z"
		}
	]
}
```

Либо в случае отсутствия сообщений, пустой список сообщений:

```http
200 OK
```

```json
{"messages": []}
```

**Ошибка - Неправильный токен авторизации**:

```http
400 Bad Request
```

```json
{"error": "error_message"}
```

---

#### **POST /api/message**

**Описание**: Отправка (создание) нового сообщения. Доступна только авторизованным пользователям. Клиент обязан передать токен авторизации и текст сообщения. В случае успеха клиенту возвращаются данные созданного сообщения. Текст сообщения не может быть пустым.

**Данные запроса**:

```json
{"api_token": "string", "text": "string"}
```

**Сообщение успешно создано**:

```http
201 CREATED
```

```json
{
	"message": {
        "id": 1002,
        "member_name": "petya",
        "text": "hello",
        "time": "2022-02-04T13:28:77Z"
	}
}
```

**Ошибка 1 - Неправильный токен**:

```http
400 Bad Request
```

```json
{"error": "error_message"}
```

**Ошибка 2 - Пустое сообщение**:

```http
400 Bad Request
```

```json
{"error": "error_message"}
```

---

#### **GET /api/health**

**Описание**: Проверка состояния бэкенда для автоматизированных средств мониторинга или интеграционных тестов.

**Сообщение успешно создано**:

```http
200 OK
```

```json
{ "success": true, "time": "2022-02-04T13:28:77Z" }
```

---
