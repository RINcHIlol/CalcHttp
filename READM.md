# CalcHttp

**CalcHttp** — это HTTP-сервис для вычисления математических выражений, передаваемых в формате JSON. Проект позволяет отправлять выражения через POST-запросы и получать результат вычислений или соответствующие ошибки.

---

## Возможности

- **Корректный ввод**: Выражение вычислено успешно.
- **Ошибка 422**: Входные данные не соответствуют требованиям приложения.
- **Ошибка 500**: Случай какой-либо иной ошибки («Что-то пошло не так»).

---

## Примеры использования

### 1. **Успешное выполнение запроса**

Вычисление выражения `2+2*2`:

```bash
curl -i -X POST http://localhost:8080/api/v1/calculate \
-H "Content-Type: application/json" \
-d '{"expression": "2+2*2"}'
```

**Ответ**:
```text
HTTP/1.1 200 OK
Date: Tue, 17 Dec 2024 20:44:04 GMT
Content-Length: 12
Content-Type: text/plain; charset=utf-8

{"result":6}
```

---

### 2. **Ошибка 422: Некорректное выражение**

Ввод некорректного выражения (`ощзф`):

```bash
curl -i -X POST http://localhost:8080/api/v1/calculate \
-H "Content-Type: application/json" \
-d '{"expression": "buba"}'
```

**Ответ**:
```text
HTTP/1.1 422 Unprocessable Entity
Date: Tue, 17 Dec 2024 20:45:00 GMT
Content-Length: 35
Content-Type: text/plain; charset=utf-8

{"error":"Expression is not valid"}%
```

---

### 3. **Ошибка 500: Пустое выражение**

Передача пустого выражения:

```bash
curl -i -X POST http://localhost:8080/api/v1/calculate \
-H "Content-Type: application/json" \
-d '{"expression": ""}'
```

**Ответ**:
```text
HTTP/1.1 500 Internal Server Error
Date: Tue, 17 Dec 2024 20:45:47 GMT
Content-Length: 33
Content-Type: text/plain; charset=utf-8

{“error”:“Internal server error”}
```

---

## Инструкция по запуску проекта

1. **Клонируйте репозиторий**:
   ```bash
   git clone https://github.com/RINcHIlol/CalcHttp.git
   cd CalcHttp
   ```

2. **Запустите проект** с помощью команды:
   ```bash
   go run main.go
   ```

3. **Сервер запустится на порту `8080`**.

Теперь можно отправлять запросы на эндпоинт:  
`http://localhost:8080/api/v1/calculate`

---

## Инструкция по запуску тестов

1. **Введите команду**:
   ```bash
    go test -v ./pkg/calculation/tests
   ```

2. **Получите результат**:
   ```text
   === RUN   TestCalcAPI
    === RUN   TestCalcAPI/Valid_Expression
    === RUN   TestCalcAPI/Invalid_Expression
    === RUN   TestCalcAPI/Internal_server_error
    --- PASS: TestCalcAPI (0.00s)
    --- PASS: TestCalcAPI/Valid_Expression (0.00s)
    --- PASS: TestCalcAPI/Invalid_Expression (0.00s)
    --- PASS: TestCalcAPI/Internal_server_error (0.00s)
    PASS
    ok      calc_http/pkg/calculation/tests (cached)
   ```

---