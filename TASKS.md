# Go Harness — учебный трек

Этот файл — источник истины для заданий трека. Новые утренние и вечерние
итерации добавляются сюда; после проверки кода обновляются статусы и чекбоксы.

Статусы:

- `done` — задание выполнено;
- `in progress` — выполнена только часть требований;
- `backlog` — задание ещё не начато.

## №1 — базовый CLI (`done`)

### Утро — аргументы командной строки

Требования:

- создать Go-модуль и исполняемую программу;
- прочитать команду и prompt из `os.Args`;
- не обращаться к отсутствующему аргументу;
- выводить понятную ошибку, если команда или prompt не переданы.

### Вечер — `validate` и `stats`

Требования:

- поддержать команды `validate` и `stats`;
- считать слова через `strings.Fields`;
- считать Unicode-символы (runes), а не байты;
- корректно обрабатывать неизвестную команду.

Критерии готовности:

- [x] CLI запускается через `go run .`;
- [x] отсутствие аргументов не приводит к panic;
- [x] `validate` и `stats` работают;
- [x] статистика корректна для Unicode-строки.

Документация: [A Tour of Go](https://go.dev/tour/),
[`os.Args`](https://pkg.go.dev/os#pkg-variables),
[`strings.Fields`](https://pkg.go.dev/strings#Fields),
[`utf8.RuneCountInString`](https://pkg.go.dev/unicode/utf8#RuneCountInString).

## №2 — `Prompt` как предметная модель (`done`)

### Утро — struct и методы

Требования:

- представить prompt отдельным `struct`;
- перенести подсчёт статистики и валидацию в функции или методы, работающие с
  `Prompt`;
- считать prompt валидным, только если в нём не меньше трёх слов;
- сохранить рабочие команды CLI.

### Вечер — ошибки и граница с CLI

Требования:

- предметная ошибка валидации возвращается как `error`, а не печатается внутри
  метода;
- CLI решает, как показать ошибку или сообщение об успехе пользователю;
- вызывающий код действительно проверяет возвращённую ошибку;
- пустая строка, пробелы и строка из двух слов не проходят валидацию.

Текущее состояние после review:

- [x] добавлен `Prompt`;
- [x] статистика и валидация оформлены методами;
- [x] минимум — три слова;
- [x] команды `stats` и `validate` остаются рабочими;
- [x] `validatePrompt` возвращает `error` вместо `bool` и не печатает;
- [x] `main` проверяет ошибку и отвечает за вывод результата валидации.

Критерии готовности:

- [x] предметная валидация ничего не печатает;
- [x] невалидный prompt представлен ненулевым `error`;
- [x] валидный prompt возвращает `nil`;
- [x] CLI корректно показывает оба результата;
- [x] `go vet ./...` и `go test ./...` проходят.

Документация: [Go by Example — Structs](https://gobyexample.com/structs),
[Go by Example — Methods](https://gobyexample.com/methods),
[Error handling and Go](https://go.dev/blog/error-handling-and-go).

## №3 — package `prompt` и unit-тесты (`backlog`)

### Утро — выделение package

Требования:

- вынести `Prompt`, валидацию и статистику из `main` в отдельный package
  `prompt`;
- оставить ввод/вывод в `main`;
- определить явный API статистики со счётчиками слов и Unicode-символов;
- не допустить зависимости package `prompt` от терминала.

### Вечер — table-driven tests

Требования:

- `prompt` возвращает ошибки и ничего не печатает в stdout;
- написать table-driven tests с `t.Run` для валидации строк:
  `"Explain RAG architecture"`, `"Explain RAG"`, `""`, `"      "` и
  `"Объясни архитектуру RAG"`;
- отдельно проверить статистику ASCII- и Unicode-строк;
- тестировать API package напрямую, не запуская CLI, shell или executable.

Критерии готовности:

- [ ] `main` отвечает только за CLI и отображение результата;
- [ ] package `prompt` не вызывает `fmt.Print*`;
- [ ] ошибки валидации передаются через `error`;
- [ ] есть table-driven tests для валидации и статистики;
- [ ] проходят `gofmt -w .`, `go vet ./...`, `go test ./...` и
  `go test -v ./...`.

Документация: [package `testing`](https://pkg.go.dev/testing),
[Strings, bytes, runes and characters in Go](https://go.dev/blog/strings),
[Error handling and Go](https://go.dev/blog/error-handling-and-go).

## №4 — HTTP-клиент (`backlog`)

### Утро — первый package `client`

Требования:

- добавить команду `send` и отдельный package `client`;
- отправлять `POST` с JSON вида `{"prompt":"..."}` через стандартную
  библиотеку;
- передавать endpoint через конфигурацию, не зашивать URL в `Send`;
- устанавливать `Content-Type: application/json`;
- закрывать `response.Body` через `defer`;
- передавать HTTP- и transport errors вызывающему коду;
- тестировать без настоящего LLM через `httptest.Server`.

### Вечер — status codes и typed response

Требования:

- вернуть из `Send` типизированный response, декодированный через
  `encoding/json`;
- различать transport error, HTTP `4xx/5xx` и успешный `2xx`;
- превращать malformed JSON при `200 OK` в ошибку;
- добавлять контекст к ошибкам через `fmt.Errorf` и `%w`;
- table-driven test покрывает valid JSON, HTTP error, invalid JSON и transport
  failure; успешный тест проверяет содержимое response;
- дополнительная часть: включить небольшое тело HTTP-error в диагностику.

Критерии готовности:

- [ ] HTTP-детали не находятся в `main`;
- [ ] endpoint конфигурируется, request/response типизированы;
- [ ] status code проверяется до обработки успешного response;
- [ ] body всегда закрывается;
- [ ] тесты не используют интернет;
- [ ] проходят `gofmt -w .`, `go vet ./...`, `go test -v ./...` и
  `go test -race ./...`.

Документация: [`net/http`](https://pkg.go.dev/net/http),
[`encoding/json`](https://pkg.go.dev/encoding/json),
[`net/http/httptest`](https://pkg.go.dev/net/http/httptest),
[Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors).

## №5 — OpenAI-compatible LLM client (`backlog`)

### Утро — Chat Completions contract

Требования:

- реализовать минимальный `POST /v1/chat/completions` только на стандартной
  библиотеке, без OpenAI SDK;
- представить request вложенными Go-структурами: `model` и `messages` с
  `role`/`content`;
- получать model и URL из конфигурации, а пользовательский текст — из CLI;
- декодировать нужную часть response: `choices[].message.role/content`;
- вернуть в `main` текст модели, а не raw JSON;
- проверить длину `choices` перед обращением к `choices[0]`;
- через `httptest.Server` проверить model, messages, role и content;
- покрыть valid response, empty choices и malformed response, сохранив тесты
  transport error и HTTP `4xx/5xx`.

Критерии готовности:

- [ ] request соответствует минимальному Chat Completions contract;
- [ ] используются JSON tags, вложенные structs и slices;
- [ ] неизвестные поля response не мешают декодированию;
- [ ] пустой `choices` возвращает ошибку без panic;
- [ ] CLI печатает текст модели;
- [ ] тесты не требуют LLM или интернета;
- [ ] проходят `gofmt -w .`, `go vet ./...`, `go test -v ./...` и
  `go test -race ./...`.

Документация: [`encoding/json`](https://pkg.go.dev/encoding/json),
[A Tour of Go — Structs](https://go.dev/tour/moretypes/2),
[A Tour of Go — Slices](https://go.dev/tour/moretypes/7).

### Вечер

Ещё не выдано. Следующую итерацию нужно добавить сюда перед началом работы.
