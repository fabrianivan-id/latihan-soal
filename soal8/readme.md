# valid
curl -s -X POST localhost:8080/products \
  -H 'Content-Type: application/json' \
  -d '{"sku":"SKU-98765432","productName":"Wireless Mouse","quantityInStock":150,"price":29.99,"category":"Electronics"}' -i

# invalid (shows ordered messages)
curl -s -X POST localhost:8080/products \
  -H 'Content-Type: application/json' \
  -d '{"sku":"foo","productName":"","quantityInStock":-1,"price":0,"category":"Toys"}'
# -> ["The sku is a mandatory field"/format depending, "The productName is a mandatory field", "The quantityInStock cannot be negative", "The price must be greater than zero", "Invalid product category"]
