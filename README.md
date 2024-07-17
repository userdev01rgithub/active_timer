### README.md (Russian)

# Таймер Активности

## Описание

Таймер Активности - это простое приложение на Go, которое помогает отслеживать ваше рабочее время и делать перерывы. Приложение использует библиотеку [Fyne](https://fyne.io) для создания графического интерфейса пользователя.

## Установка

1. Установите Go (если еще не установлен). Инструкции можно найти на [официальном сайте Go](https://golang.org/doc/install).
2. Склонируйте репозиторий:

    ```bash
    git clone https://github.com/userdev01rgithub/active_timer.git
    ```
3. Перейдите в каталог проекта:

    ```bash
    cd active_timer
    ```
4. Установите зависимости:

    ```bash
    go mod tidy
    ```

## Запуск

Для запуска приложения выполните следующую команду:

```bash
go run main.go
```
## Билд приложения

Для билда приложения выполните следующую команду:

```bash
go build -o active_timer cmd/main.go
```

## Использование

После запуска приложения, вы увидите графический интерфейс с кнопками для запуска и остановки таймера. Таймер будет отслеживать ваше рабочее время и предлагать делать перерывы через определенные интервалы времени.

## Зависимости

Проект использует следующие зависимости:

- [fyne.io/fyne/v2](https://pkg.go.dev/fyne.io/fyne/v2)
- [github.com/BurntSushi/toml](https://pkg.go.dev/github.com/BurntSushi/toml)
- и другие (см. go.mod)

## Лицензия

Этот проект лицензирован под лицензией MIT. Подробности см. в файле LICENSE.

---

### README.md (English)

# Activity Timer

## Description

Activity Timer is a simple Go application that helps you track your work time and take breaks. The application uses the [Fyne](https://fyne.io) library to create a graphical user interface.

## Installation

1. Install Go (if not already installed). Instructions can be found on the [official Go website](https://golang.org/doc/install).
2. Clone the repository:

    ```bash
    git clone https://github.com/userdev01rgithub/active_timer.git
    ```
3. Navigate to the project directory:

    ```bash
    cd active_timer
    ```
4. Install dependencies:

    ```bash
    go mod tidy
    ```

## Running

To run the application, execute the following command:

```bash
go run main.go
```

## Usage

After starting the application, you will see a graphical interface with buttons to start and stop the timer. The timer will track your work time and suggest taking breaks at regular intervals.

## Dependencies

The project uses the following dependencies:

- [fyne.io/fyne/v2](https://pkg.go.dev/fyne.io/fyne/v2)
- [github.com/BurntSushi/toml](https://pkg.go.dev/github.com/BurntSushi/toml)
- and others (see go.mod)

## License

This project is licensed under the MIT License. See the LICENSE file for details.
