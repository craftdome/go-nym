![Static Badge](https://img.shields.io/badge/release-v1.0-blue)

# О репозитории
Эта библиотека призвана упростить взаимодействие с протоколом NYM для nym-client. Она реализует базовый набор команд для отправки и получения сообщений в/из mixnet.

## Миксноды разработчика

[![Dynamic JSON Badge](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fexplorer.nymtech.net%2Fapi%2Fv1%2Fmix-node%2F895&query=mix_node.identity_key&style=plastic&logo=numpy&logoColor=white&label=Advanced%20Engineering%201&color=%23136401&cacheSeconds=60&link=https%3A%2F%2Fexplorer.nymtech.net%2Fnetwork-components%2Fmixnode%2F895)](https://explorer.nymtech.net/network-components/mixnode/895)
[![Dynamic JSON Badge](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fexplorer.nymtech.net%2Fapi%2Fv1%2Fmix-node%2F895&query=stake_saturation&style=plastic&logo=myspace&logoColor=white&label=Stake&cacheSeconds=60&link=https%3A%2F%2Fexplorer.nymtech.net%2Fnetwork-components%2Fmixnode%2F895)](https://explorer.nymtech.net/network-components/mixnode/895)
[![Dynamic JSON Badge](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fexplorer.nymtech.net%2Fapi%2Fv1%2Fmix-node%2F895&query=mix_node.version&style=plastic&logo=git&logoColor=white&label=Version&cacheSeconds=60&link=https%3A%2F%2Fexplorer.nymtech.net%2Fnetwork-components%2Fmixnode%2F895)](https://explorer.nymtech.net/network-components/mixnode/895)

# Особенности

- [x] Используется [Gorilla WebSocket](https://github.com/gorilla/websocket)
- [x] Поддержка текстового протокола (JSON)
- [ ] Поддержка бинарного протокола (Binary)
- [ ] Поддержка пользовательского протокола в теле бинарного сообщения

# Установка
Версия [Go](https://go.dev/dl/): 1.21.4

Используйте стандартные инструменты Go для установки зависисмости:
```
go get github.com/craftdome/go-nym
```

Импорт базового пакета:
```go
import "github.com/craftdome/go-nym"
```

# Чтение сообщений

1. Для начала необходимо инициализировать клиент подключения с указанием адреса websocket-подключения.
```go
// Init the client via server's credentials
client := nym.NewClient("ws://127.0.0.1:1977")
```


2. Установка соединения c сервером (nym-client).
```go
// Dial a connection to the server
if err := client.Dial(); err != nil {
  panic(err)
}
// Closing the client upon completion of main()
defer client.Close()
```

3. Включаем прослушивание входящих сообщений, которые далее извлекаем через канал Messages().
```go
// Start reading WS messages
go client.ListenAndServe()

go func() {
  // Incoming Message Channel
  for message := range client.Messages() {
    switch message.Type() {
    case tags.Error:
      msg := message.(*response.Error)
      fmt.Printf("Error: %s\n", msg.Message)
    case tags.SelfAddress:
      msg := message.(*response.SelfAddress)
      fmt.Printf("SelfAddress: %s\n", msg.Address)
    case tags.Received:
      msg := message.(*response.Received)
      fmt.Printf("Received: %s, SenderTag: %s\n", msg.Message, msg.SenderTag)
    }
  }
}()
```

# Отправка сообщений

1. Получение собственного адреса (SelfAddress).
```go
// Request your own address
if err := client.SendRequestAsText(nym.NewGetSelfAddress()); err != nil {
  fmt.Fprintln(os.Stderr, err)
}
```

2. Отправка сообщения (Send).
```go
// Send a message
addr := "9wxVvU4mbk4ZpWLiKEM45Bw3ednUVsfTyooTU3jr6iDR.HcJdxJrPDqWfw6TiauLZEC3mKjFByzFGEtFvDPe2pKW3@Emswx6KXyjRfq1c2k4d4uD2e6nBSbH1biorCZUei8UNS"
r := nym.NewSend("Mix it up!", addr)
if err := client.SendRequestAsText(r); err != nil {
  fmt.Fprintln(os.Stderr, err)
}
```

3. Отправка сообщения с формированием SURB для получения анонимного ответа (SendAnonymous).
```go
// Send an anonymous message
addr = "9wxVvU4mbk4ZpWLiKEM45Bw3ednUVsfTyooTU3jr6iDR.HcJdxJrPDqWfw6TiauLZEC3mKjFByzFGEtFvDPe2pKW3@Emswx6KXyjRfq1c2k4d4uD2e6nBSbH1biorCZUei8UNS"
replySurbs := 1
r = nym.NewSendAnonymous("Enjoy your anonymous!", addr, replySurbs)
if err := client.SendRequestAsText(r); err != nil {
  fmt.Fprintln(os.Stderr, err)
}
```

4. Отправка ответа на сообщение SendAnonymous (Reply).
```go
// Reply to an anonymous message
senderTag := "HJNPCE41NncWGnNSC5w6Cq"
r = nym.NewReply("Pong.", senderTag)
if err := client.SendRequestAsText(r); err != nil {
  fmt.Fprintln(os.Stderr, err)
}
```
