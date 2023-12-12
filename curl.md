# CURL

## GetAll products by warehouse_id and page
```
curl -X 'GET' \
  'http://localhost:8080/warehouse/product?id=a494f72f-4044-43ed-b5c2-1145aa88352d&page=0' \
  -H 'accept: application/json'
  ```
## Reserve
```
curl -X 'POST' \
  'http://localhost:8080/product/reserve' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "reserveItems": [
    {
      "id": "OPT1",
      "quantity": 1
    },
    {
      "id": "OPT3",
      "quantity": 1
    }
  ],
  "warehouseId": "a494f72f-4044-43ed-b5c2-1145aa88352d"
}'
```


## Release

```
curl -X 'DELETE' \
  'http://localhost:8080/product/release' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "uniqueIds": [
    "OPT1",
    "OPT3"
  ],
  "warehouseId": "a494f72f-4044-43ed-b5c2-1145aa88352d"
}'
```
