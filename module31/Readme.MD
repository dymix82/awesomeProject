# 31.5 Практическая работа
## Цель практической работы
Научиться:

работать с запросами POST, GET, PUT, DELETE;
применять принципы написания обработчиков HTTP-запросов.


## Что нужно сделать
Что нужно сделать
В прошлом домашнем задании вы писали приложение, которое принимает HTTP-запросы, создаёт пользователей, добавляет друзей и так далее. Давайте теперь приблизим наше приложение к реальному продукту.

Отрефакторьте приложение так, чтобы вы могли поднять две реплики данного приложения.
Используйте любую базу данных, чтобы сохранять информацию о пользователях, или можете сохранять информацию в файл, предварительно сереализуя в JSON.
### Ответ
Выбрана Postgres, схема таблицы для данных:
```sql
CREATE TABLE Users (
id SERIAL PRIMARY KEY,
users JSONB);
```
  

Напишите proxy или используйте, например, nginx.
### Ответ
/proxy/proxy.go

Протестируйте приложение.

1. Сделайте обработчик создания пользователя. У пользователя должны быть следующие поля: имя, возраст и массив друзей.
Пользователя необходимо сохранять в Базу данных.
Данный запрос должен возвращать ID пользователя и статус 201.
![img.png](img/img.png)
![img_1.png](img/img_1.png)

2. Сделайте обработчик, который делает друзей из двух пользователей. Например, если мы создали двух пользователей 
и нам вернулись их ID, то в запросе мы можем указать ID пользователя, который инициировал запрос на дружбу, 
и ID пользователя, который примет инициатора в друзья. 
![img_3.png](img/img_3.png)
![img_4.png](img/img_4.png)
3. Сделайте обработчик, который удаляет пользователя. Данный обработчик принимает ID пользователя и удаляет
   его из хранилища, а также стирает его из массива friends у всех его друзей
![img_6.png](img/img_6.png)
![img_7.png](img/img_7.png)
4. Сделайте обработчик, который возвращает всех друзей пользователя. 
![img_11.png](img/img_11.png)
5. Сделайте обработчик, который обновляет возраст пользователя. 
![img_9.png](img/img_9.png)
![img_10.png](img/img_10.png)