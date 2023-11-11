#!/bin/sh -e

OPERATION="${@:-up}"
# operation變成給這個腳本所有參數,如果沒有參數會變成up

echo "正在執行資料庫變遷，方向：$OPERATION"

if [ "$DATABASE_URL" = "" ]; then
    echo "FATAL: 未設定DATABASE_URL環境變數"
    exit 1
fi

migrate -database $DATABASE_URL -path db/migrations $OPERATION