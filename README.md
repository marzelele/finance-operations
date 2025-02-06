# REST API сервис для финансовых операций
## Запуск
```shell
make run
```
## API 
### POST /api/finance/replenish
Пополнение баланса пользователя

##### Example Input:
```json
{
    "user_id": "a0f80990-7feb-4626-b275-997040051c1e",
    "amount": 1000
}
```
##### Example Output:
```json
{
    "operation_id": "8b0a2277-fada-434d-bdf6-d0e1f24052d2"
}
```

### POST /api/finance/transfer
Перевод денег от одного пользователя к другому
##### Example Input:
```json
{
    "source_user_id": "70ad97be-a3a3-4e0b-9b46-fb8a0d23b48a",
    "destination_user_id": "d54a3df5-cabd-4862-9b26-871398e5ffa2",
    "amount": 100
}
```
##### Example Output:
```json
{
    "operation_id": "bf6d10a7-0418-49a6-96eb-ec38b4de5553"
}
```
### GET /api/finance/operations/last?user_id=70ad97be-a3a3-4e0b-9b46-fb8a0d23b48a
Просмотр 10 последних операций пользователя
##### Example Output:
```json
{
  "operations": [
    {
      "id": "bf6d10a7-0418-49a6-96eb-ec38b4de5553",
      "request_time": "2025-02-06T09:40:47.647517+10:00",
      "type": 2,
      "details": {
        "source_user_id": "70ad97be-a3a3-4e0b-9b46-fb8a0d23b48a",
        "destination_user_id": "d54a3df5-cabd-4862-9b26-871398e5ffa2",
        "amount": 100
      }
    },
    {
      "id": "8948c2b0-a877-43e3-86c7-09f46ea92bbd",
      "request_time": "2025-02-06T09:40:03.302254+10:00",
      "type": 1,
      "details": {
        "source_user_id": "70ad97be-a3a3-4e0b-9b46-fb8a0d23b48a",
        "destination_user_id": null,
        "amount": 100
      }
    }
  ]
}
```
