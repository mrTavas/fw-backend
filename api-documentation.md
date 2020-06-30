# Rest Api Документация

Документация содержит список все доступных методов, их описание.

## Список методов:
### Администрирование:

- [DropModels](#DropModels)

- [CreateModels](#CreateModels)

### Регистрация:

- [NewManager](#NewManager)

- [NewWorker](#NewWorker)

- [NewClient](#NewClient)


### Аутентификация:
- [Login](#Login)

### Работа с заказами:
- [NewOrder](#NewOrder)

- [DeleteOrder](#GetOrders)

- [GetOrders](#GetOrders)

- [GetOrderStatus](#GetOrderStatus)

- [StartStep](#StartStep)

- [EndStep](#EndStep)

- [DropStatus](#DropStatus)

- [EditOrder](#EditOrder)

- [GetOrderAllChanges](#GetOrderAllChanges)

- [GetOrderLastChanges](#GetOrderLastChanges)


### Работа с прайс листами:

- [GetPriceList](#GetPriceList)

- [NewPrice](#NewPrice)

- [DeletePrice](#DeletePrice)

- [ChangePrice](#ChangePrice)


### Работа с работниками:

- [GetWorkers](#GetWorkers)

- [GetWorkerCurrentOrders](#GetWorkerCurrentOrders)

- [GetWorkerOldOrders](#GetWorkerOldOrders)

- [GetAllCarpenters](#GetAllCarpenters)

- [GetAllGrinders](#GetAllGrinders)

- [GetAllPainters](#GetAllPainters)

- [GetAllCollectors](#GetAllCollectors)

- [EditWorker](#EditWorker)

- [DeleteWorker](#DeleteWorker)

### Работа с менеджерами:
- [GetManagers](#GetManagers)

- [DeleteManager](#DeleteManager)

### Работа с клиентами:
- [GetClients](#GetClients)

- [DeleteClient](#DeleteClient)

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
- `ID` - id записи.
- `UUID` - Uuid генерируется случайноо.
- `Phone` - Номер телефона.
- `Password` - Пароль менеджера (md5).
- `Initials` - Фамилия и инициалы.
- `Сarpenter` - Столяр (manufacturing) true/false.
- `Grinder` - Шлифовщик true/false.
- `Painter` - Маляр true/false.
- `Collector` - Сборщик true/false.
- `СurrentBalance` - Текущий баланс работника.

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

### NewClient
    http://fwqqq-backend.ddns.net:1323/api/auth/NewClient
    
Описание:
Создает в таблице нового клиента. Метод Post. Менеджер имеет следующие параметры:

- `ID` - id записи;
- `Phone` - Номер телефона;
- `Password` - Пароль клиента (код активации из смс, формируется автоматически);
- `Score` - Скидочный счет клиента (поумолчанию = 1. При значении 0.9 клиент будет иметь скидку в 10%);

Пример тела запроса:
 
    {
        "phone": 898887947499, 
        "initials": "Мединцев А. С." 
    }

Ответ:

    {
        "message": "OK"
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
Структура заказа:

- `ID` - id заказа;
- `Status` - Статус заказа:

    - `status_office_start` - Этап оффис начат (default: true);
    - `status_office_end` - Этап оффис завершен;
    - `status_manufacturing_start` - Этап производтва начат;
    - `status_manufacturing_end` - Этап производтва завершен;
    - `status_grinding_start` - Этап шлифовки начат;
    - `status_grinding_end` - Этап шлифовки завершен;
    - `status_collectind_start` - Этап сборки начат;
    - `status_collecting_end` - Этап сборки завершен;
    - `status_ready` - Заказ завершен.

- `ClientID` - id клиента (если клиент не зарегистрирован передавать client_id = 0);
- `ClientInitials` - инициалы клиента (передавать если client_id = 0);
- `ClientPhone` - телефон клиента (передавать если client_id = 0);
- `CurrentWorkerID` - id работника (инициалы и телефон работника автоматически добавятся в заказ);
- `Title` - Название заказа:    
- `ClientInitials` - Инициалы клиента;
- `ClientPhone` - Телефон клиента;
- `CurrentWorkerInitials` - Инициалы текущего работника (добавятся автоматически по id);
- `CurrentWorkerPhone` - Телефон текущего работника (добавится автоматически по id);
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
        "title": "Title",
        "client_id": 0,
        "client_initials": "Clientov A.V.",
        "client_phone" : 79888563211,
        "current_worker_id": 1,
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

Ответ:

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

Пример ответа:

    {
        "message": "Order deleted"
    }

Пример ответа в случае если заказа с таким id несуществует:

    {
        "message": "pg: no rows in result set"
    }

---

### GetOrders
    http://fwqqq-backend.ddns.net:1323/api/auth/GetOrders
    
Описание:
Взвращает список всех незавершенных заказов. Метод Get.


Пример ответа:

    {
        "orders": [
            {
                "id": 1,
                "title": "Title",
                "Date": "2020-06-30T17:33:32.59735+03:00",
                "status": {
                    "data_office_start": "2020-06-30T17:33:32.596472371+03:00",
                    "data_office_end": "0001-01-01T00:00:00Z",
                    "data_manufacturing_start": "0001-01-01T00:00:00Z",
                    "data_manufacturing_end": "0001-01-01T00:00:00Z",
                    "data_grinding_start": "0001-01-01T00:00:00Z",
                    "data_grinding_end": "0001-01-01T00:00:00Z",
                    "data_printing_start": "0001-01-01T00:00:00Z",
                    "data_printing_end": "0001-01-01T00:00:00Z",
                    "data_collecting_start": "0001-01-01T00:00:00Z",
                    "data_collecting_end": "0001-01-01T00:00:00Z",
                    "data_ready": "0001-01-01T00:00:00Z",
                    "status_office_start": true,
                    "status_office_end": false,
                    "status_manufacturing_start": false,
                    "status_manufacturing_end": false,
                    "status_grinding_start": false,
                    "status_grinding_end": false,
                    "status_printing_start": false,
                    "status_printing_end": false,
                    "status_collecting_start": false,
                    "status_collecting_end": false,
                    "status_ready": false
                },
                "client_id": 0,
                "client_initials": "Clientov A.V.",
                "client_phone": 79888563211,
                "current_worker_id": 1,
                "current_worker_initials": " Плотный П. В.",
                "current_worker_phone": 8988879411111,
                "color": "red",
                "patina": "patina",
                "fasad_article": "SomeArticle",
                "material": "tree",
                "cost_carpenter": 3000,
                "cost_grinder": 2000,
                "cost_painter": 1500,
                "cost_collector": 2000,
                "params": [
                    {
                        "title": "Some Title ",
                        "height": 12,
                        "width": 0,
                        "filenka": "filenka"
                    },
                    {
                        "title": "Some Comment2",
                        "height": 13,
                        "width": 0,
                        "filenka": "panel2"
                    }
                ]
            }
        ]
    }

---

### GetOrderStatus
    http://fwqqq-backend.ddns.net:1323/api/auth/GetOrderStatus
    
Описание:
Взвращает статус заказа по его id. Метод Post.
Пример тела запроса:

    {
        "id": 1
    }

Пример ответа:

    {
        "data_office_start": "2020-06-30T17:33:32.596472371+03:00",
        "data_office_end": "0001-01-01T00:00:00Z",
        "data_manufacturing_start": "0001-01-01T00:00:00Z",
        "data_manufacturing_end": "0001-01-01T00:00:00Z",
        "data_grinding_start": "0001-01-01T00:00:00Z",
        "data_grinding_end": "0001-01-01T00:00:00Z",
        "data_printing_start": "0001-01-01T00:00:00Z",
        "data_printing_end": "0001-01-01T00:00:00Z",
        "data_collecting_start": "0001-01-01T00:00:00Z",
        "data_collecting_end": "0001-01-01T00:00:00Z",
        "data_ready": "0001-01-01T00:00:00Z",
        "status_office_start": true,
        "status_office_end": false,
        "status_manufacturing_start": false,
        "status_manufacturing_end": false,
        "status_grinding_start": false,
        "status_grinding_end": false,
        "status_printing_start": false,
        "status_printing_end": false,
        "status_collecting_start": false,
        "status_collecting_end": false,
        "status_ready": false
    }

Пример ответа в случае если заказа с таким id несуществует:

    {
        "message": "pg: no rows in result set"
    }

---

### StartStep
    fwqqq-backend.ddns.net:1323/api/auth/StartStep
    
Описание:
Устанавливает значение следующего статуса заказа с тегом _start  в `true` и устанавливает текущее время для данного статуса. Устанавливает нового работника заказа. Метод Post. При использовании данного метода в лог запишутся изменения в заказе, например: "Заказ начал этап manufacturing. Назначенный работник: Иванов И. И."
Пример тела запроса:

    {
        "order_id": 1,
        "new_worker_id": 2
    }


Пример ответа:

    {
        "message": "Начался этап manufacturing."
    }

Пример ответа в случая если предыдущий этап еще не завершен:

    {
        "message": "Есть незавершенные этапы."
    }

Пример ответа в случае если заказа с таким id несуществует:

    {
        "message": "pg: no rows in result set"
    }

---

### EndStep
    fwqqq-backend.ddns.net:1323/api/auth/EndStep
    
Описание:
Устанавливает значение следующего статуса заказа с тегом _end  в `true` и устанавливает текущее время для данного статуса. Добавляет к балансу работника соответвующее значение. Сохраняет в таблицу сохраненых стадий. Метод Post. При использовании данного метода в лог запишутся изменения в заказе, например: "Заказ завершил этап manufacturing"
Пример тела запроса:

    {
        "id": 1
    }


Пример ответа:

    {
        "message": "Этап manufacturing завершился."
    }

Пример ответа в случая если нет начатых этапов:

    {
        "message": "Нет начатых этапов."
    }

Пример ответа в случае если заказа с таким id несуществует:

    {
        "message": "pg: no rows in result set"
    }

---

### DropStatus
    fwqqq-backend.ddns.net:1323/api/auth/DropStatus
    
Описание:
Устанавливает значения статусов заказа в `false` (все кроме `status_office_start`) и даты (все кроме `data_office_start`) в начальные значения, по id заказа. Метод Post.
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

### EditOrder
    http://fwqqq-backend.ddns.net:1323/api/EditOrder
    
Описание:
Редактирование заказа. Метод схож с NewOrder, только в данном случе необходимо передать еще и id заказа который подлежит редактированию. Метод Post.
Данный метод требует авторизацию через bearer token (токен менеджера). При выполнении данного метода все изменения будут записаны в лог, по токену будет определен менеджер, совершивший изменения, что также будет зафиксировано в логе.
Пример тела запроса:


    {
        "id": 1,
        "title": "Title",
        "Date": "2020-06-30T17:33:32.59735+03:00",
        "status": {
            "data_office_start": "2020-06-30T17:33:32.596472371+03:00",
            "data_office_end": "0001-01-01T00:00:00Z",
            "data_manufacturing_start": "0001-01-01T00:00:00Z",
            "data_manufacturing_end": "0001-01-01T00:00:00Z",
            "data_grinding_start": "0001-01-01T00:00:00Z",
            "data_grinding_end": "0001-01-01T00:00:00Z",
            "data_printing_start": "0001-01-01T00:00:00Z",
            "data_printing_end": "0001-01-01T00:00:00Z",
            "data_collecting_start": "0001-01-01T00:00:00Z",
            "data_collecting_end": "0001-01-01T00:00:00Z",
            "data_ready": "0001-01-01T00:00:00Z",
            "status_office_start": true,
            "status_office_end": false,
            "status_manufacturing_start": false,
            "status_manufacturing_end": false,
            "status_grinding_start": false,
            "status_grinding_end": false,
            "status_printing_start": false,
            "status_printing_end": false,
            "status_collecting_start": false,
            "status_collecting_end": false,
            "status_ready": false
        },
        "client_id": 0,
        "client_initials": "Clientov A.V.",
        "client_phone": 79888563211,
        "current_worker_id": 1,
        "current_worker_initials": " Плотный П. В.",
        "current_worker_phone": 8988879411111,
        "color": "red",
        "patina": "patina",
        "fasad_article": "SomeArticle",
        "material": "tree",
        "cost_carpenter": 3000,
        "cost_grinder": 2000,
        "cost_painter": 1500,
        "cost_collector": 2000,
        "params": [
            {
                "title": "Some Title ",
                "height": 12,
                "width": 0,
                "filenka": "filenka"
            },
            {
                "title": "Some Comment2",
                "height": 13,
                "width": 0,
                "filenka": "panel2"
            }
        ]
    }

Пример ответа:

    {
        "message": "OK"
    }

Пример ответа если неправельный или просроченный токен:

    {
        "message": "invalid or expired jwt"
    }

---

### GetOrderAllChanges
    http://fwqqq-backend.ddns.net:1323/api/auth/GetOrderAllChanges
    
Описание:
Возвращает все изменения, которые были совершены над заказом (по id заказа). Метод Post.

Пример тела запроса:

    {
        "id": 1
    }


Пример ответа:

    {
        "changes": [
            {
                "ID": 1,
                "order_id": 1,
                "Date": "2020-06-23T00:24:19.322464+03:00",
                "manager_id": 0,
                "initials": "",
                "changes": "Заказ переведен на этап manufacturing. Назначенный работник: Иванов D. И."
            },
            {
                "ID": 2,
                "order_id": 1,
                "Date": "2020-06-23T01:03:38.340254+03:00",
                "manager_id": 1,
                "initials": "Mr Manager",
                "changes": "Изменено Title с Title на i change it. Изменено StatusManufacturing с true на false. Изменено Комментарий к параметрам с Some Title  на HH."
            }
        ]
    }

---

### GetOrderLastChanges
    http://fwqqq-backend.ddns.net:1323/api/auth/GetOrderLastChanges
    
Описание:
Возвращает последнее изменение, которое было совершено над заказом (по id заказа). Метод Post.

Пример тела запроса:

    {
        "id": 1
    }


Пример ответа:

    {
        "ID": 2,
        "order_id": 1,
        "Date": "2020-06-23T01:03:38.340254+03:00",
        "manager_id": 1,
        "initials": "Mr Manager",
        "changes": "Изменено Title с Title на i change it. Изменено StatusManufacturing с true на false. Изменено Комментарий к параметрам с Some Title  на HH."
    }

---

## Работа с прайс листами
### GetPriceList
    fwqqq-backend.ddns.net:1323/api/auth/GetPriceList
    
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
    http://fwqqq-backend.ddns.net:1323/api/auth/NewPrice
    
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
    http://fwqqq-backend.ddns.net:1323/api/auth/DeletePrice
    
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
    http://fwqqq-backend.ddns.net:1323/api/auth/ChangePrice
    
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
    http://fwqqq-backend.ddns.net:1323/api/auth/GetWorkers
    
Описание:
Взвращает список всех работников. Метод Get.

Пример ответа:

    {
        "workers": [
            {
                "ID": 1,
                "uuid": "7af7b074-afa0-4e42-a68d-1495b8042fe5",
                "CurrentBalance": 0,
                "phone": 8988879409999,
                "pass": "qwerty1",
                "initials": "Иванов F. И.",
                "carpenter": false,
                "grinder": false,
                "painter": false,
                "collector": false
            },
            {
                "ID": 2,
                "uuid": "9825fe4c-4732-429f-aa8b-7f5f0b64b69f",
                "CurrentBalance": 0,
                "phone": 8988879400000,
                "pass": "qwerty1",
                "initials": "Иванов D. И.",
                "carpenter": false,
                "grinder": false,
                "painter": false,
                "collector": false
            }
        ]
    }

---

### GetAllCarpenters
    http://fwqqq-backend.ddns.net:1323/api/auth/GetAllCarpenters
    
Описание:
Взвращает список всех работников плотников (на этап manufacturing). Метод Get.

Пример ответа:

    {
        "workers": [
            {
                "ID": 2,
                "uuid": "cce66873-c0ae-4dd6-b0eb-53a6883c615f",
                "CurrentBalance": 0,
                "phone": 89888794111111,
                "pass": "qwerty1",
                "initials": " Плотный П. В.2",
                "carpenter": true,
                "grinder": false,
                "painter": false,
                "collector": false
            },
            {
                "ID": 1,
                "uuid": "43d2fb58-fe73-484f-bce5-a46c74f43ebf",
                "CurrentBalance": 3000,
                "phone": 8988879411111,
                "pass": "qwerty1",
                "initials": " Плотный П. В.",
                "carpenter": true,
                "grinder": false,
                "painter": false,
                "collector": false
            }
        ]
    }

---

### GetAllGrinders
    http://fwqqq-backend.ddns.net:1323/api/auth/GetAllGrinders
    
Описание:
Взвращает список всех работников шлифовщиков (на этап grinding). Метод Get.

Пример ответа:

    {
        "workers": [
            {
                "ID": 1,
                "uuid": "43d2fb58-fe73-484f-bce5-a46c74f43ebf",
                "CurrentBalance": 0,
                "phone": 8988879411111,
                "pass": "qwerty1",
                "initials": " Плотный П. В.",
                "carpenter": true,
                "grinder": true,
                "painter": false,
                "collector": false
            }
        ]
    }

---

### GetAllPainters
    http://fwqqq-backend.ddns.net:1323/api/auth/GetAllPainters
    
Описание:
Взвращает список всех работников маляров (на этап painting). Метод Get.

Пример ответа:

    {
        "workers": [
            {
                "ID": 1,
                "uuid": "43d2fb58-fe73-484f-bce5-a46c74f43ebf",
                "CurrentBalance": 0,
                "phone": 8988879411111,
                "pass": "qwerty1",
                "initials": " Плотный П. В.",
                "carpenter": false,
                "grinder": false,
                "painter": true,
                "collector": false
            }
        ]
    }

---

### GetAllCollectors
    http://fwqqq-backend.ddns.net:1323/api/auth/GetAllCollectors
    
Описание:
Взвращает список всех работников сборщиков (на этап collectind). Метод Get.

Пример ответа:

    {
        "workers": [
            {
                "ID": 1,
                "uuid": "43d2fb58-fe73-484f-bce5-a46c74f43ebf",
                "CurrentBalance": 0,
                "phone": 8988879411111,
                "pass": "qwerty1",
                "initials": " Плотный П. В.",
                "carpenter": false,
                "grinder": false,
                "painter": false,
                "collector": true
            }
        ]
    }

---

### EditWorker
    http://fwqqq-backend.ddns.net:1323/api/auth/EditWorker
    
Описание:
Редактирование работника. Метод схож с NewWorker, только в данном случе необходимо передать еще и id работника который подлежит редактированию. Метод Post. Если какой-либо из статуссов (carpenter, grinder, painter или collector) не передается, то он установится в `false`.

Пример тела запроса:


    {
        "ID": 1,
        "uuid": "43d2fb58-fe73-484f-bce5-a46c74f43ebf",
        "CurrentBalance": 0,
        "phone": 8988879411111,
        "pass": "qwerty1",
        "initials": " Плотный П. В.",
        "carpenter": true,
        "grinder": true,
        "painter": false,
        "collector": false
    }

Пример ответа:

    {
        "message": "OK"
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

### GetWorkerCurrentOrders
    http://fwqqq-backend.ddns.net:1323/api/auth/GetWorkerCurrentOrders
    
Описание:
Возвращает текущие заказы работника  по его id. Метод Post.
Пример тела запроса:

    {
        "id": 1
    }

Пример ответа:

    {
        "orders": [
            {
                "id": 1,
                "title": "Title",
                "Date": "2020-06-30T17:33:32.59735+03:00",
                "status": {
                    "data_office_start": "2020-06-30T17:33:32.596472371+03:00",
                    "data_office_end": "0001-01-01T00:00:00Z",
                    "data_manufacturing_start": "2020-06-30T17:44:16.847558907+03:00",
                    "data_manufacturing_end": "0001-01-01T00:00:00Z",
                    "data_grinding_start": "0001-01-01T00:00:00Z",
                    "data_grinding_end": "0001-01-01T00:00:00Z",
                    "data_printing_start": "0001-01-01T00:00:00Z",
                    "data_printing_end": "0001-01-01T00:00:00Z",
                    "data_collecting_start": "0001-01-01T00:00:00Z",
                    "data_collecting_end": "0001-01-01T00:00:00Z",
                    "data_ready": "0001-01-01T00:00:00Z",
                    "status_office_start": true,
                    "status_office_end": false,
                    "status_manufacturing_start": false,
                    "status_manufacturing_end": false,
                    "status_grinding_start": false,
                    "status_grinding_end": false,
                    "status_printing_start": false,
                    "status_printing_end": false,
                    "status_collecting_start": false,
                    "status_collecting_end": false,
                    "status_ready": false
                },
                "client_id": 0,
                "client_initials": "Clientov A.V.",
                "client_phone": 79888563211,
                "current_worker_id": 1,
                "current_worker_initials": " Плотный П. В.",
                "current_worker_phone": 8988879411111,
                "color": "red",
                "patina": "patina",
                "fasad_article": "SomeArticle",
                "material": "tree",
                "cost_carpenter": 3000,
                "cost_grinder": 2000,
                "cost_painter": 1500,
                "cost_collector": 2000,
                "params": [
                    {
                        "title": "Some Title ",
                        "height": 12,
                        "width": 0,
                        "filenka": "filenka"
                    },
                    {
                        "title": "Some Comment2",
                        "height": 13,
                        "width": 0,
                        "filenka": "panel2"
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

### GetWorkerOldOrders
    http://fwqqq-backend.ddns.net:1323/api/auth/GetWorkerOldOrders
    
Описание:
Возвращает завершенные заказы работника  по его id. Метод Post.
Пример тела запроса:

    {
        "id": 1
    }

Пример ответа:

    {
        "saved_orders": [
            {
                "ID": 1,
                "order_id": 1,
                "title": "Title",
                "Date": "2020-06-30T17:33:32.59735+03:00",
                "status": {
                    "data_office_start": "2020-06-30T17:33:32.596472371+03:00",
                    "data_office_end": "2020-06-30T17:44:14.56757536+03:00",
                    "data_manufacturing_start": "2020-06-30T17:44:16.847558907+03:00",
                    "data_manufacturing_end": "2020-06-30T17:52:09.693244359+03:00",
                    "data_grinding_start": "0001-01-01T00:00:00Z",
                    "data_grinding_end": "0001-01-01T00:00:00Z",
                    "data_printing_start": "0001-01-01T00:00:00Z",
                    "data_printing_end": "0001-01-01T00:00:00Z",
                    "data_collecting_start": "0001-01-01T00:00:00Z",
                    "data_collecting_end": "0001-01-01T00:00:00Z",
                    "data_ready": "0001-01-01T00:00:00Z",
                    "status_office_start": true,
                    "status_office_end": true,
                    "status_manufacturing_start": true,
                    "status_manufacturing_end": true,
                    "status_grinding_start": false,
                    "status_grinding_end": false,
                    "status_printing_start": false,
                    "status_printing_end": false,
                    "status_collecting_start": false,
                    "status_collecting_end": false,
                    "status_ready": false
                },
                "current_worker_id": 0,
                "client_initials": "Clientov A.V.",
                "client_phone": 79888563211,
                "CurrentWorkerID": 1,
                "current_worker_initials": " Плотный П. В.",
                "current_worker_phone": 8988879411111,
                "color": "red",
                "patina": "patina",
                "fasad_article": "SomeArticle",
                "material": "tree",
                "cost_carpenter": 3000,
                "cost_grinder": 2000,
                "cost_painter": 1500,
                "cost_collector": 2000,
                "params": [
                    {
                        "title": "Some Title ",
                        "height": 12,
                        "width": 0,
                        "filenka": "filenka"
                    },
                    {
                        "title": "Some Comment2",
                        "height": 13,
                        "width": 0,
                        "filenka": "panel2"
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
    http://fwqqq-backend.ddns.net:1323/api/auth/GetManagers
    
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

Пример ответа:

    {
        "message": "Manager deleted"
    }

Пример ответа в случае если менеджера с таким id несуществует:

    {
        "message": "pg: no rows in result set"
    }

---

## Работа с клиентами
### GetClients
    http://fwqqq-backend.ddns.net:1323/api/auth/GetClients
    
Описание:
Взвращает список всех зарегистрированных клиентов. Метод Get.

Пример ответа:

    {
        "clients": [
            {
                "ID": 1,
                "phone": 898887947499,
                "Password": 426113,
                "initials": "Мединцев А. С.",
                "Score": 1
            }
        ]
    }

---
### DeleteClient
    http://fwqqq-backend.ddns.net:1323/api/auth/DeleteClient
    
Описание:
Удаляет клиента по его id. Метод Post.
Пример тела запроса:

    {
        "id": 1
    }

Пример ответа:

    {
        "message": "Client deleted"
    }

Пример ответа в случае если клиента с таким id несуществует:

    {
        "message": "pg: no rows in result set"
    }

---
