# Разворачивание и настройка сервера

Инструкция содержит список всех пакетов, которые необходимо установить, их установку и настройку.

## Компоненты:
### Postgresql:

- [Установка Postgresql](#Postgresql)

- [Настройка Postgresql](#Postgresql)

### Golang:

- [Установка Golang](#Golang)

- [Настройка Golang](#Golang)

### PgAdmin4:
- [Установка PgAdmin4](#PgAdmin4)

----

#Golang

- Установка

    cd ~
  
  curl -O https://dl.google.com/go/go1.12.1.linux-amd64.tar.gz
  
  sudo tar -xvf go1.12.1.linux-amd64.tar.gz -C /usr/local
  
  sudo chown -R root:root /usr/local/go
  
- Создание рабочего пространства Go

  mkdir -p $HOME/go/{bin,src}
  
Откройте файл `~/.profile` с помощью nano или другого предпочитаемого текстового редактора.Чтобы задать переменную $GOPATH, добавьте в файл следующую строку:

  export GOPATH=$HOME/go
  
При компиляции и установке инструментов Go помещает их в директорию $GOPATH/bin. Для удобства субдиректория /bin рабочего пространства обычно добавляется в переменную PATH в файле ~/.profile:

  export PATH=$PATH:$GOPATH/bin

Наконец, вам нужно добавить двоичный файл go в PATH. Для этого добавьте /usr/local/go/bin в конец строки:

  export PATH=$PATH:$GOPATH/bin:/usr/local/go/bin

  . ~/.profile
  
Для проверки установки проверьте текущую версию Go:

  go version

Теперь вы сможете создавать будущие проекты со следующей структурой директорий. В этом примере предполагается, что вы используете github.com в качестве репозитория:

  $GOPATH/src/github.com/username/project


