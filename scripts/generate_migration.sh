#!/bin/sh
# 以上是指用sh執行這個腳本

migrate create -ext sql -dir db/migrations "$@" #$@指給這個腳本所有的參數