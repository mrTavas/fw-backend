# Rest Api Документация

Документация содержит  список все доступных методов.

### Список методов:
[DropModels](#DropModels)

[CreateModels](#CreateModels)

[NewManager](#NewManager)


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
