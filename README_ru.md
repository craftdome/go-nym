![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/craftdome/go-nym?style=flat-square)
![GitHub release (with filter)](https://img.shields.io/github/v/release/craftdome/go-nym?style=flat-square)

# О репозитории
Эта библиотека призвана упростить взаимодействие с протоколом NYM для [nym-client](https://nymtech.net/docs/clients/overview.html#the-websocket-client). Она реализует базовый набор команд для отправки и получения сообщений в/из mixnet.

# Особенности

- [x] Используется [Gorilla WebSocket](https://github.com/gorilla/websocket)
- [x] Поддержка текстового протокола (JSON)
- [ ] Поддержка бинарного протокола (Binary)
- [ ] Поддержка пользовательского протокола в теле бинарного сообщения

# Подготовка к использованию

1. Для работы библиотеки требуется активное подключение к websocket-клиенту. Простая установка и запуск `nym-client` описаны в оффициальной документации -> [link](https://nymtech.net/docs/clients/websocket/setup.html)

2. Протестировано с версией nym-client `1.1.32`

3. Импорт зависимости

Используйте стандартные инструменты Go для установки зависисмости:
```
go get github.com/craftdome/go-nym
```

Импорт базового пакета:
```go
import "github.com/craftdome/go-nym"
```

# Использование

### Советы перед использованием

- Помните, если вы планируете дать доступ для подключения к `nym-client` извне (для этого вам следует указать ip внешнего интерфейса вашей машины), `nym-client` не имеет функции авторизации подключения.
- Единовременно разрешено только 1 подключение.
- Если вам нужно внешнее подключение, используйте локальную сеть ([10.0.0.0/8, 192.168.0.0/16, 172.16.0.0/12](https://en.wikipedia.org/wiki/Private_network)) вместо глобальной для повышения безопасности.

### Инициализация

1. Для начала, необходимо инициализировать клиент подключения с указанием адреса `nym-client`. Адрес и порт подключения копируем из консоли после запуска `nym-client`. По умолчанию это `localhost:1977`, а в моём случае `192.168.88.4:1977`.

https://github.com/craftdome/go-nym/blob/17a1c03713c834ca65295c0c5818b0d1dba364e2/example/main.go#L21-L22

2. Установка соединения c `nym-client`.

https://github.com/craftdome/go-nym/blob/17a1c03713c834ca65295c0c5818b0d1dba364e2/example/main.go#L24-L27

### Чтение сообщений

3. Включаем прослушивание входящих сообщений, которые далее извлекаем через канал `Messages()`.

https://github.com/craftdome/go-nym/blob/17a1c03713c834ca65295c0c5818b0d1dba364e2/example/main.go#L30-L54

### Отправка сообщений

4. Получение собственного адреса (SelfAddress).

https://github.com/craftdome/go-nym/blob/17a1c03713c834ca65295c0c5818b0d1dba364e2/example/main.go#L57-L59

5. Отправка сообщения (Send).

https://github.com/craftdome/go-nym/blob/17a1c03713c834ca65295c0c5818b0d1dba364e2/example/main.go#L61-L66

6. Отправка сообщения с формированием SURB для получения анонимного ответа (SendAnonymous).

https://github.com/craftdome/go-nym/blob/17a1c03713c834ca65295c0c5818b0d1dba364e2/example/main.go#L68-L74

7. Отправка ответа на сообщение SendAnonymous (Reply).

https://github.com/craftdome/go-nym/blob/17a1c03713c834ca65295c0c5818b0d1dba364e2/example/main.go#L76-L81

### Корректное закрытие соединения

8. Закрываем соединение с `nym-client` после завершения работы (interrupt/kill signal) и ожидаем сигнал `done` от читающей горутины.

https://github.com/craftdome/go-nym/blob/17a1c03713c834ca65295c0c5818b0d1dba364e2/example/main.go#L83-L91

# Поддержать разработчика (Миксноды)

Ниже приведён список микснод разработчика. Если вы в поисках ноды для делегирования токенов, можете присмотреться к моим вариантам. Комиссия владельца ноды всего 4%, что является низким показателем среди других нод, где комиссия может доходить до 20%, а то и выше, если прибавить стоимость обслуживания ноды (встречается 4000 токенов - это уже около 28% о максимальной доходности ноды).

[![Dynamic JSON Badge](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fexplorer.nymtech.net%2Fapi%2Fv1%2Fmix-node%2F895&query=mix_node.identity_key&style=flat-square&logo=numpy&logoColor=white&label=Advanced%20Engineering%201&color=%23136401&cacheSeconds=60)](https://explorer.nymtech.net/network-components/mixnode/895)
[![Dynamic JSON Badge](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fexplorer.nymtech.net%2Fapi%2Fv1%2Fmix-node%2F895&query=stake_saturation&style=flat-square&logo=myspace&logoColor=white&label=Stake&cacheSeconds=60)](https://explorer.nymtech.net/network-components/mixnode/895)
[![Dynamic JSON Badge](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fexplorer.nymtech.net%2Fapi%2Fv1%2Fmix-node%2F895&query=profit_margin_percent&style=flat-square&logo=buymeacoffee&logoColor=white&label=Owner%20Profit)](https://explorer.nymtech.net/network-components/mixnode/895)
[![Dynamic JSON Badge](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fexplorer.nymtech.net%2Fapi%2Fv1%2Fmix-node%2F895&query=mix_node.version&style=flat-square&logo=git&logoColor=white&label=Version&cacheSeconds=60)](https://explorer.nymtech.net/network-components/mixnode/895)

[![Dynamic JSON Badge](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fexplorer.nymtech.net%2Fapi%2Fv1%2Fmix-node%2F1227&query=mix_node.identity_key&style=flat-square&logo=numpy&logoColor=white&label=Advanced%20Engineering%202&color=%23136401&cacheSeconds=60)](https://explorer.nymtech.net/network-components/mixnode/1227)
[![Dynamic JSON Badge](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fexplorer.nymtech.net%2Fapi%2Fv1%2Fmix-node%2F1227&query=stake_saturation&style=flat-square&logo=myspace&logoColor=white&label=Stake&cacheSeconds=60)](https://explorer.nymtech.net/network-components/mixnode/1227)
[![Dynamic JSON Badge](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fexplorer.nymtech.net%2Fapi%2Fv1%2Fmix-node%2F1227&query=profit_margin_percent&style=flat-square&logo=buymeacoffee&logoColor=white&label=Owner%20Profit)](https://explorer.nymtech.net/network-components/mixnode/1227)
[![Dynamic JSON Badge](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fexplorer.nymtech.net%2Fapi%2Fv1%2Fmix-node%2F1227&query=mix_node.version&style=flat-square&logo=git&logoColor=white&label=Version&cacheSeconds=60)](https://explorer.nymtech.net/network-components/mixnode/1227)
