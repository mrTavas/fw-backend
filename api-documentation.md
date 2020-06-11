# Rest Api Документация

Документация содержит список все доступных методов, их описание.

## Список методов:
### Администрирование:

- [DropModels](#DropModels)

- [CreateModels](#CreateModels)

### Регистрация:

- [NewManager](#NewManager)

- [NewWorker](#NewWorker)

### Аутентификация:
- [Login](#Login)

### Работа с заказами:
- [NewOrder](#NewOrder)

- [DeleteOrder](#GetOrders)

- [GetOrders](#GetOrders)

- [GetOrderStatus](#GetOrderStatus)

- [NextStatus](#NextStatus)

- [DropStatus](#DropStatus)

- [GetPriceList](#GetPriceList)

- [NewPrice](#NewPrice)

- [DeletePrice](#DeletePrice)

- [ChangePrice](#ChangePrice)


### Работа с работниками:
- [GetWorkers](#GetWorkers)

- [DeleteWorker](#DeleteWorker)

- [GetWorkerOrders](#GetWorkerOrders)

### Работа с менеджерами:
- [GetManagers](#GetManagers)

- [DeleteManager](#DeleteManager)

___

## Администрирование
### CreateModels
    http://fwqqq-backend.ddns.net:1323/CreateModels

Описание: 
Создает в базе данных все описанные модели (таблицы). Необходим для пересоздания базы данных. Применяется после метода [DropModels](#DropModels). Метод Get.
Ответ в случае успеха:

    Models Created
 
Ответ в случае неудачи (модели уже созданны):
 
    {
        "message": "ОШИБКА #42P07 отношение \"workers\" уже существует"
    }
---
### DropModels
    http://fwqqq-backend.ddns.net:1323/DropModels

Описание: 
Удаляет в базе данных все описанные модели (таблицы). Метод Get
Ответ в случае успеха:

    Models Deleted/Dropped
 
 Пример ответа в случае неудачи (модели уже созданны):
 
    {
        "message": "ОШИБКА #42P01 таблица \"orders\" не существует"
    }
---

## Регистрация
### NewManager
    http://fwqqq-backend.ddns.net:1323/api/auth/newmanager
    
Описание:
Создает в таблице нового менеджера. Метод Post. Менеджер имеет следующие параметры:
- `ID` - id записи;
- `UUID` - Uuid генерируется случайноо;
- `Phone` - Номер телефона;
- `Password` - Пароль менеджера (md5);
- `Initials` - Фамилия и инициалы.

 `id` и `uuid` формируются автоматически, а `phone`, `password` и `initials` необходимо передать в json формате. Пример тела запроса:
 
     {
    	"phone": 89888793988, 
    	"initials": "Ivanon P. A.", 
    	"pass": "qwerty1"
     }

Ответ в случае успеха:

    {
        "user_uuid": "705c3d67-9bcb-4fa4-9f48-b3dac4576f24",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjowLCJsb2dpbiI6ODk4ODg3OTQ3NDcsImV4cCI6MTU4ODk4MDQ3NiwiaWF0IjoxNTg4OTc0NDc2fQ.AJ8Xk9i6JeM0SvHLTJP_uUAH6uXMMsrt83eUhcb_R2I",
        "refresh_token": "19ac79b3-6561-42c0-b9da-a19c7f7c6e88",
        "refresh_expiration": "2020-09-26T00:47:56.613403344+03:00"
    }
 
 Ответ в случае неудачи (пользователь уже есть в базе данных):
 
    {
        "message": "ОШИБКА #23505 повторяющееся значение ключа нарушает ограничение уникальности \"managers_phone_key\""
    }
---
|
### NewWorker
    http://fwqqq-backend.ddns.net:1323/api/auth/newworker
    
Описание:
Создает в таблице нового работника. Метод Post. Работника имеет следующие параметры:
- `ID` - id записи;
- `UUID` - Uuid генерируется случайноо;
- `Phone` - Номер телефона;
- `Password` - Пароль менеджера (md5);
- `Initials` - Фамилия и инициалы.
- `Сarpenter` - Столяр true/false.
- `Grinder` - Шлифовщик true/false.
- `Painter` - Маляр true/false.
- `Collector` - Сборщик true/false.

 `id` и `uuid` формируются автоматически, а `phone`, `password` и `initials` необходимо передать в json формате. Параметры `Сarpenter`, `Grinder`, `Painter`, `Collector` передаются если они необходимы (если значение не передается, то по умолчанию устанавливается false). Пример тела запроса:
 
    {
        "phone": 89888794747,
        "pass" : qwerty1
        "initials": "Ivanon I. I.",
        "carpenter": false,
        "grinder": false,
        "painter": false,
        "collector": true
    }

Ответ в случае успеха:

    {
        "user_uuid": "705c3d67-9bcb-4fa4-9f48-b3dac4576f24",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjowLCJsb2dpbiI6ODk4ODg3OTQ3NDcsImV4cCI6MTU4ODk4MDQ3NiwiaWF0IjoxNTg4OTc0NDc2fQ.AJ8Xk9i6JeM0SvHLTJP_uUAH6uXMMsrt83eUhcb_R2I",
        "refresh_token": "19ac79b3-6561-42c0-b9da-a19c7f7c6e88",
        "refresh_expiration": "2020-09-26T00:47:56.613403344+03:00"
    }
 
 Ответ в случае неудачи (пользователь уже есть в базе данных):
 
    {
        "message": "ОШИБКА #23505 повторяющееся значение ключа нарушает ограничение уникальности \"workers_phone_key\""
    }
---

## Аутентификация
### Login
    http://fwqqq-backend.ddns.net:1323/api/auth/login
    
Описание:
Проверяет логин и пароль в базе данных. Если пароль и логин верны, то вернет токены, если нет то вернет пустые токены. Метод Post.
Пример тела запроса:

    {
        "pass" : "qwerty1",
        "phone": 89888794747,
        "user": "Worker"
    }

Ответ в случае успеха:

    {
        "user_uuid":"c473d4d4-485b-4dc5-95df-799a222c6023",
           "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJsb2dpbiI6ODk4ODg3OTQ1NDUsImV4cCI6MTU4Njc4MDg3OCwiaWF0IjoxNTg2Nzc0ODc4fQ.NJ3cV2cf9AcBwmmBL25jJnNO1wLZlPJvrlZWkCUKCtg",
        "refresh_token":"3467563c-231c-46bc-a31d-f1d2eba44e1c",
        "refresh_expiration":"2020-08-31T13:47:58.39724323+03:00"
    }

Ответ в случае неудачи (логин или пароль неверны):

    {
        "user_uuid":"",
        "token":"",
        "refresh_token":"",
        "refresh_expiration":"0001-01-01T00:00:00Z"
    }
    
---

## Работа с заказами
### NewOrder
    http://fwqqq-backend.ddns.net:1323/api/auth/neworder
    
Описание:
Создает новый заказ. Метод Post.

- `ID` - id заказа;
- `Status` - Статус заказа:

    - `status_office` - Этап оффис;
    - `status_manufacturing` - Этап производтва;
    - `status_grinding` - Этап шлифовки;
    - `status_ready` - Этап готовности.

- `Title` - Название заказа:    
- `ClientInitials` - Инициалы клиента;
- `ClientPhone` - Телефон клиента;
- `CurrentWorkerInitials` - Инициалы текущего работника;
- `CurrentWorkerPhone` - Телефон текущего работника;
- `СostManufacturing` - Цена производства;
- `CostPainting` - Цена покраски;
- `CostFinishing` - Цена производства;
- `CostFull` - Цена итоговая;
- `Params` - Массив с таблицами параметров для заказа. Имеет следующую структуру:

    - `Title` - Заголовок или комментарий;
    - `Height` - Высота;
    - `Width` - Ширина;
    - `Filenka` - Филёнка.


Пример тела запроса:

    {
        "status": {
            "status_office": true,
            "status_manufacturing": false,
            "status_grinding": false,
            "status_printing": false,
            "status_ready": false
        },
        "title": "Title",
        "client_initials": "Clientov A.V.",
        "client_phone" : 79888563211,
        "current_worker_initials": "Ivanon I. I.",
        "current_worker_phone": 7988121212,
        "cost_manufacturing": 3000,
        "cost_painting": 2000,
        "cost_finishing": 1500,
        "cost_full": 7500,
        "color": "red",
        "patina": "patina",
        "fasad_article": "SomeArticle",
        "material": "tree", 
        "params": [
            {
                "title": "Some Title ",
                "height": 12,
                "weight": 15,
                "filenka": "filenka"
            },
            {
                "title": "Some Comment2",
                "height": 13,
                "weight": 1,
                "filenka": "panel2"               
            }]
    }

Ответ в случае успеха:

    {
        "message": "Order added"
    }


---

### DeleteOrder
    http://fwqqq-backend.ddns.net:1323/api/auth/DeleteOrder
    
Описание:
Удаляет заказ по его id. Метод Post.
Пример тела запроса:

    {
        "id": 1
    }

Пример ответа в случае успеха:

    {
        "message": "Order deleted"
    }

Пример ответа в случае если заказа с таким id несуществует:

    {
        "message": "pg: no rows in result set"
    }

---

### GetOrders
    http://fwqqq-backend.ddns.net:1323/GetOrders
    
Описание:
Взвращает список всех заказов. Метод Get.


Пример ответа:

    {
        "orders": [
            {
                "ID": 1,
                "Date": "2020-06-02T22:18:31.315815+03:00",
                "status": {
                    "data_office": "0001-01-01T00:00:00Z",
                    "data_manufacturing": "0001-01-01T00:00:00Z",
                    "data_grinding  ": "0001-01-01T00:00:00Z",
                    "data_printing  ": "0001-01-01T00:00:00Z",
                    "data_ready": "0001-01-01T00:00:00Z",
                    "status_office": true,
                    "status_manufacturing": false,
                    "status_grinding": false,
                    "status_printing": false,
                    "status_ready": false
                },
                "client_initials": "Clientov A.V.",
                "client_phone": 79888563211,
                "current_worker_initials": "Ivanon I. I.",
                "current_worker_phone": 7988121212,
                "cost_manufacturing": 3000,
                "cost_painting": 2000,
                "cost_finishing": 1500,
                "cost_full": 7500,
                "params": [
                    {
                        "title": "Some Title ",
                        "height": 12,
                        "width": 0,
                        "filenka": "filenka",
                        "color": "",
                        "patina": "",
                        "fasad_article": "",
                        "material": ""
                    },
                    {
                        "title": "Some Comment2",
                        "height": 13,
                        "width": 0,
                        "filenka": "panel2",
                        "color": "",
                        "patina": "",
                        "fasad_article": "",
                        "material": ""
                    }
                ]
            }
        ]
    }

---

### GetOrderStatus
    http://fwqqq-backend.ddns.net:1323/GetOrderStatus
    
Описание:
Взвращает статус заказа по его id. Метод Post.
Пример тела запроса:

    {
        "id": 1
    }

Пример ответа:

    {
        "data_office": "0001-01-01T00:00:00Z",
        "data_manufacturing": "0001-01-01T00:00:00Z",
        "data_grinding  ": "0001-01-01T00:00:00Z",
        "data_printing  ": "0001-01-01T00:00:00Z",
        "data_ready": "0001-01-01T00:00:00Z",
        "status_office": true,
        "status_manufacturing": false,
        "status_grinding": false,
        "status_printing": false,
        "status_ready": false
    }

Пример ответа в случае если заказа с таким id несуществует:

    {
        "message": "pg: no rows in result set"
    }

---

### NextStatus
    fwqqq-backend.ddns.net:1323/api/auth/NextStatus
    
Описание:
Устанавливает значение следующего статуса в `true` и устанавливает текущее время для данного статуса, по id заказа. Метод Post.
Пример тела запроса:

    {
        "id": 1
    }

Возвращает текущий статус заказа. Пример ответа:

    {
        "message": "ready"
    }

Пример ответа в случае если заказа с таким id несуществует:

    {
        "message": "pg: no rows in result set"
    }

---

### DropStatus
    fwqqq-backend.ddns.net:1323/api/auth/DropStatus
    
Описание:
Устанавливает значения статусов заказа в `false` (все кроме `status_office`) даты статусов не меняются, по id заказа. Метод Post.
Пример тела запроса:

    {
        "id": 1
    }

Пример ответа в случае успеха:

    {
        "message": "OK"
    }

Пример ответа в случае если заказа с таким id несуществует:

    {
        "message": "pg: no rows in result set"
    }

---

### GetPriceList
    fwqqq-backend.ddns.net:1323/GetPriceList
    
Описание:
Возвращает прайс лист. Метод Get.

Пример ответа:

    {
        "PriceList": [
            {
                "name": "f_01",
                "price": 4200
            },
            {
                "name": "f_02",
                "price": 4500
            },
            {
                "name": "f_03",
                "price": 5000
            },
            {
                "name": "f_04",
                "price": 5500
            },
            {
                "name": "a_01",
                "price": 7000
            },
            {
                "name": "a_02",
                "price": 8000
            },
            {
                "name": "b_01",
                "price": 7500
            },
            {
                "name": "Modern",
                "price": 6000
            },
            {
                "name": "Mausoleum",
                "price": 6500
            },
            {
                "name": "Massif",
                "price": 8500
            }
        ]
    }

---

### NewPrice
    http://fwqqq-backend.ddns.net:1323/api/auth/newPrice
    
Описание:
Добавляет запись в прайс лист. Метод Post.
Пример тела запроса:

    {
        "name": "Ф_01",
        "price": 5000
    }

Пример ответа в случае успеха:

    {
        "message": "OK"
    }

---

### DeletePrice
    http://fwqqq-backend.ddns.net:1323/api/auth/deletePrice
    
Описание:
Удаляет запись в прайс листе по имени оъекта. Метод Post.
Пример тела запроса:

    {
        "name": "Ф_01"
    }

Пример ответа в случае успеха:

    {
        "message": "Price deleted"
    }

Пример ответа в случае если работник с таким id несуществует:

    {
        "message": "pg: no rows in result set"
    }

---

### ChangePrice
    http://fwqqq-backend.ddns.net:1323/api/auth/changePrice
    
Описание:
Изменяет запись в прайс листе по имени оъекта. Метод Post.
Пример тела запроса:

    {
        "name": "Ф_01",
        "price": 5000
    }

Пример ответа в случае успеха:

    {
        "message": "OK"
    }

---

## Работа с работниками
### GetWorkers
    http://fwqqq-backend.ddns.net:1323/GetWorkers
    
Описание:
Взвращает список всех работников. Метод Get.

Пример ответа:

    {
        "workers": [
            {
                "ID": 1,
                "uuid": "e87016e5-85fb-43de-a675-5f1302701cfe",
                "phone": 898887947477,
                "pass": "qwerty1",
                "initials": "Worker1",
                "carpenter": false,
                "grinder": false,
                "painter": false,
                "collector": false
            },
            {
                "ID": 3,
                "uuid": "559fe949-167d-47fa-a112-b71834292693",
                "phone": 898887947472,
                "pass": "qwerty1",
                "initials": "Worker2",
                "carpenter": false,
                "grinder": false,
                "painter": false,
                "collector": false
            }
        ]
    }

---
### DeleteWorker
    http://fwqqq-backend.ddns.net:1323/api/auth/DeleteWorker
    
Описание:
Удаляет работника по его id. Метод Post.
Пример тела запроса:

    {
        "id": 1
    }

Пример ответа в случае успеха:

    {
        "message": "Manager deleted"
    }

Пример ответа в случае если работник с таким id несуществует:

    {
        "message": "pg: no rows in result set"
    }

---

### GetWorkerOrders
    http://fwqqq-backend.ddns.net:1323/api/auth/GetWorkerOrders
    
Описание:
Возвращает текущие заказы работника  по его номеру телефона. Метод Post.
Пример тела запроса:

    {
        "current_worker_phone": 7988121212
    }

Пример ответа в случае успеха:

    {
        "orders": [
            {
                "ID": 26,
                "title": "qwe",
                "Date": "2020-06-11T11:07:01.029196+03:00",
                "status": {
                    "data_office": "2020-06-11T11:07:01.028498471+03:00",
                    "data_manufacturing": "0001-01-01T00:00:00Z",
                    "data_grinding": "0001-01-01T00:00:00Z",
                    "data_printing": "0001-01-01T00:00:00Z",
                    "data_ready": "0001-01-01T00:00:00Z",
                    "status_office": true,
                    "status_manufacturing": false,
                    "status_grinding": false,
                    "status_printing": false,
                    "status_ready": false
                },
                "client_initials": "qwe",
                "client_phone": 0,
                "current_worker_initials": "Попов Петр",
                "current_worker_phone": 7988121212,
                "color": "qwwe",
                "patina": "qwe",
                "fasad_article": "",
                "material": "МДФ-16",
                "cost_manufacturing": 1,
                "cost_painting": 1,
                "cost_finishing": 1,
                "cost_full": 1,
                "params": [
                    {
                        "title": "Some Title or comment",
                        "height": 12,
                        "width": 15,
                        "filenka": "filenka lalala"
                    }
                ]
            },
            {
                "ID": 32,
                "title": "12",
                "Date": "2020-06-11T12:36:55.058282+03:00",
                "status": {
                    "data_office": "2020-06-11T12:36:55.057564962+03:00",
                    "data_manufacturing": "0001-01-01T00:00:00Z",
                    "data_grinding": "0001-01-01T00:00:00Z",
                    "data_printing": "0001-01-01T00:00:00Z",
                    "data_ready": "0001-01-01T00:00:00Z",
                    "status_office": true,
                    "status_manufacturing": false,
                    "status_grinding": false,
                    "status_printing": false,
                    "status_ready": false
                },
                "client_initials": "123",
                "client_phone": 123,
                "current_worker_initials": "Попов Петр",
                "current_worker_phone": 7988121212,
                "color": "123",
                "patina": "12",
                "fasad_article": "",
                "material": "МДФ-16",
                "cost_manufacturing": 1,
                "cost_painting": 1,
                "cost_finishing": 1,
                "cost_full": 1,
                "params": [
                    {
                        "title": "Some Title or comment",
                        "height": 12,
                        "width": 15,
                        "filenka": "filenka lalala"
                    }
                ]
            },
            {
                "ID": 33,
                "title": "123",
                "Date": "2020-06-11T13:10:21.917682+03:00",
                "status": {
                    "data_office": "2020-06-11T13:10:21.917146224+03:00",
                    "data_manufacturing": "0001-01-01T00:00:00Z",
                    "data_grinding": "0001-01-01T00:00:00Z",
                    "data_printing": "0001-01-01T00:00:00Z",
                    "data_ready": "0001-01-01T00:00:00Z",
                    "status_office": true,
                    "status_manufacturing": false,
                    "status_grinding": false,
                    "status_printing": false,
                    "status_ready": false
                },
                "client_initials": "123",
                "client_phone": 123,
                "current_worker_initials": "Попов Петр",
                "current_worker_phone": 7988121212,
                "color": "123",
                "patina": "123",
                "fasad_article": "",
                "material": "МДФ-16",
                "cost_manufacturing": 1,
                "cost_painting": 1,
                "cost_finishing": 1,
                "cost_full": 1,
                "params": [
                    {
                        "title": "Some Title or comment",
                        "height": 12,
                        "width": 15,
                        "filenka": "filenka lalala"
                    }
                ]
            }
        ]
    }

Пример ответа если ничего не найденно:

    {
        "orders": null
    }

---

## Работа с менеджерами
### GetManagers
    http://fwqqq-backend.ddns.net:1323/GetManagers
    
Описание:
Взвращает список всех менеджеров. Метод Get.

Пример ответа:

    {
        "managers": [
            {
                "ID": 2,
                "uuid": "4aaa1b6e-0ce5-4caf-bae0-342f056b345a",
                "phone": 898887947112,
                "pass": "qdadas12112",
                "initials": "Manager2"
            }
        ]
    }

---
### DeleteManager
    http://fwqqq-backend.ddns.net:1323/api/auth/DeleteManager
    
Описание:
Удаляет менеджера по его id. Метод Post.
Пример тела запроса:

    {
        "id": 1
    }

Пример ответа в случае успеха:

    {
        "message": "Manager deleted"
    }

Пример ответа в случае если менеджера с таким id несуществует:

    {
        "message": "pg: no rows in result set"
    }

---
