# Разворачивание и настройка сервера

Инструкция содержит список всех пакетов, которые необходимо установить, их установку и настройку.

## Компоненты:
## Postgresql:

- [Установка Postgresql](#Postgresql)

- [Настройка Postgresql](#Postgresql)

## Golang:

- [Установка Golang](#Golang)

- [Настройка Golang](#Golang)

## PgAdmin4:
- [Установка PgAdmin4](#PgAdmin4)

---
## Postgresql
### Установка:

    sudo sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt/ `lsb_release -cs`-pgdg main" >> /etc/apt/sources.list.d/pgdg.list'

    wget -q https://www.postgresql.org/media/keys/ACCC4CF8.asc -O - | sudo apt-key add -

    sudo apt-get update

    sudo apt-get install postgresql postgresql-contrib

### Настройка:
Вы можете либо переключиться в сессию учётной записи postgres и запустить там оболочку программы:

    sudo su - postgres
    psql
    
Попав тем или иным способом в командную строку psql, вам необходимо знать, как из неё выйти. Это можно сделать с помощью ввода команды выхода:

    \q
    
Создание новой роли (Имя указывается без кавычек, а пароль — в одинарных кавычках):

    create user testu1 with password 'testpass1';
    
Создание базы данных:

    create database vscale_db;
    
Назначение прав:

    grant all privileges on database vscale_db to testu1;
    
Вся минимально требующаяся предварительная настройка завершена. Выйти из psql можно введя "\q".

---
    
## Golang
### Установка:

    cd ~

    curl -O https://dl.google.com/go/go1.12.1.linux-amd64.tar.gz
    
    sudo tar -xvf go1.12.1.linux-amd64.tar.gz -C /usr/local
  
    sudo chown -R root:root /usr/local/go
  
Создание рабочего пространства Go

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
    
### Настройка:
Установка необходимых библиотек:

    go get -u github.com/labstack/echo/...
    
    go get github.com/go-pg/pg
    
    go get github.com/gofrs/uuid
    
    go get github.com/spf13/viper

---


