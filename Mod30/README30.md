Написать HTTP-сервис, который принимает входящие соединения с JSON-данными и обрабатывает их следующим образом:
1. Сделайте обработчик создания пользователя. У пользователя должны быть следующие поля: имя, возраст и массив друзей. Пользователя необходимо сохранять в мапу.
2. Сделайте обработчик, который делает друзей из двух пользователей.
3. Сделайте обработчик, который удаляет пользователя. Данный обработчик принимает ID пользователя и удаляет его из хранилища
4. Сделайте обработчик, который возвращает всех друзей пользователя
5. Сделайте обработчик, который обновляет возраст пользователя