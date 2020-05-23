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

- [GetOrders](#GetOrders)

### Работа с работниками:
- [GetWorkers](#GetWorkers)

### Работа с менеджерами:
- [GetManagers](#GetManagers)

___
### CreateModels
    http://fwqqq-backend.ddns.net:1323/CreateModels

Описание: 
Создает в базе данных все описанные модели (таблицы). Необходим для пересоздания базы данных. Применяется после метода [DropModels](#DropModels).
Ответ в случае успеха:

    Models Created
 
 Ответ в случае неудачи (модели уже созданны):
 
    {
        "message": "ОШИБКА #42P07 отношение \"workers\" уже существует"
    }
---
### DropModels
    http://fwqqq-backend.ddns.net:1323/DropModels
Ответ в случае успеха:

    Models Deleted/Dropped
 
 Ответ в случае неудачи (модели уже созданны):
 
    {
        "message": "ОШИБКА #42P01 таблица \"orders\" не существует"
    }
---
### NewManager
    http://fwqqq-backend.ddns.net:1323/api/auth/newmanager
    
Описание:
Создает в таблице нового менеджера. Менеджер имеет следующие параметры:
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
### NewWorker
    http://fwqqq-backend.ddns.net:1323/api/auth/newworker
    
Описание:
Создает в таблице нового работника. Работника имеет следующие параметры:
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
### Login
    http://fwqqq-backend.ddns.net:1323/api/auth/login
    
Описание:
Проверяет логин и пароль в базе данных. Если пароль и логин верны, то вернет токены, если нет то вернет пустые токены.
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

### NewOrder
    http://fwqqq-backend.ddns.net:1323/api/auth/neworder
    
Описание:
Создает новый заказ.

- `ID` - id заказа;
- `Status` - Статус заказа (Офис/);
- `ClientInitials` - Инициалы клиента;
- `ClientPhone` - Телефон клиента;
- `CurrentWorkerInitials` - Инициалы текущего работника.
- `CurrentWorkerPhone` - Телефон текущего работника.
- `СostManufacturing` - Цена производства.
- `CostPainting` - Цена покраски.
- `CostFinishing` - Цена производства.
- `CostFull` - Цена итоговая.
- `Params` - Массив с таблицами параметров для заказа. Имеет следующую структуру:

- `Title` - Заголовок или комментарий;
- `Height` - Высота;
- `Width` - Ширина;
- `Filenka` - Филёнка;


Пример тела запроса:

    {
        "status": "Office",
        "client_initials": "Clientov A.V.",
        "client_phone" : 79888563211,
        "current_worker_initials": "Ivanon I. I.",
        "current_worker_phone": 7988121212,
        "cost_manufacturing": 3000,
        "cost_painting": 2000,
        "cost_finishing": 1500,
        "cost_full": 7500,
        "params": [
            {
                "title": "Some Title or comment",
                "height": 12,
                "width": 15,
                "filenka": "filenka lalala"
            },
            {
                "title": "Some Comment2",
                "height": 13,
                "width": 1,
                "filenka": "some panel2"
            }]
    }

Ответ в случае успеха:

    {
        "message": "Order added"
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
                "ID": 2,
                "Date": "2020-05-23T15:30:07.688475+03:00",
                "status": "Office",
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
                        "title": "Some Title or comment",
                        "height": 12,
                        "width": 0,
                        "filenka": "filenka lalala"
                    },
                    {
                        "title": "Some Comment2",
                        "height": 13,
                        "width": 0,
                        "filenka": "some panel2"
                    }
                ]
            },
            {
                "ID": 3,
                "Date": "2020-05-23T18:15:37.800598+03:00",
                "status": "Office",
                "client_initials": "Pavfv A.P",
                "client_phone": 7988856111,
                "current_worker_initials": "Ivanon I. I.",
                "current_worker_phone": 7988121212,
                "cost_manufacturing": 1300,
                "cost_painting": 2000,
                "cost_finishing": 300,
                "cost_full": 3600,
                "params": [
                    {
                        "title": "Some Title again",
                        "height": 12,
                        "width": 0,
                        "filenka": "filenka123"
                    }
                ]
            }
        ]
    }

---

### GetManagers
    http://fwqqq-backend.ddns.net:1323/GetManagers
    
Описание:
Взвращает список всех менеджеров. Метод Get.


Пример ответа:



---
