# Toko Ijah Inventory Web App - API

An API to store actual stock of products, record of the products going out of inventory, record of the products going in to inventory, and generate report based on inventory value / omzet / selling / profit.  [Demo](http://tokoijah.kmaguswira.me)

## Getting Started

### Prerequisites

You must have `glide` installed to run this repo on your local machine. Please check [glide](https://github.com/Masterminds/glide) for more information.

### Installing

Clone this repository to your `$GOPATH` directory

```bash
cd $GOPATH/src/github.com/kmaguswira
git clone https://github.com/kmaguswira/salestock-api.git
```

Go to repository and install dependencies.

```bash
cd salestock-api
glide install
glide up
```

Migrate the database, with `e` flag and value `migrate`. There are two type of environment `migrate` and `development` by default the application is running on `development`
```bash
go run main.go -e migrate
```

To start the application `go run main.go` or `go run main.go -e development`
```bash
go run main.go
```
and it's running on port `7011`. You can change the port by editing `yaml` file in `config` directory.

## Database Schema
![Schema](https://raw.githubusercontent.com/kmaguswira/salestock-api/master/dbschema.png)

*   `Product` the actual quantity of the product
*   `Order` the product that will be going in to the inventory from supplier. `Order status` will be complete if the sum of `Order Progress quantityReceived` is equal `Order orderQuantity`
*   `Order Progress` the progress of particular order that supplier have been send to the inventory.
*   `Product Out` the products that are going out inventory neither they are sold or broken
*   `Sales` the transaction for the products that are going out inventory because they are sold.

## Endpoint

Base path for the endpoint is `localhost:7011/v1/`.
Every models have their own `blueprint` endpoint to `CRUD` the data. `blueprint` consist 5 operation endpoint such as

* `/` with `POST` method to create the model
* `/all` with `GET` method to retrieve all data
* `/get/:id` with `GET` method to retrieve a data based on id
* `/update/:id` with `PUT` method to update a data based on id
* `/delete/:id` with `DELETE` method to delete a data based on id

For `/all` with `GET` method can receive query parameter to filter the data. The available query for this request `where` `order` `limit` `offset`. By default `where` is `unset`, `order` is `created_at desc`, `limit` is `1000`, and `offset` is `0`. For now to filter the data with`where`, it just work to filter `string` value. More detail with query `where` see the [`gorm`](http://gorm.io/docs/query.html) query documentation.
### Example:
* Filter product name with exact string
    ```
    http://localhost:7011/v1/product/all?where={"name =": "baju"}
    ```
* Filter product name that consist the string
    ```
    http://localhost:7011/v1/product/all?where={"name LIKE": "%baju%"}
    ```
### List:
```
GET    /public/*filepath                [serve static file]
POST   /v1/product/create               [create a product]
GET    /v1/product/all                  [retrieve all product data]
GET    /v1/product/get/:id              [retrieve a product based on id]
PUT    /v1/product/update/:id           [update a product based on id]
DELETE /v1/product/delete/:id           [delete a product based on id]
POST   /v1/order/create                 [create a order]
GET    /v1/order/all                    [retrieve all order data]
GET    /v1/order/get/:id                [retrieve an order based on id]
PUT    /v1/order/update/:id             [update an order based on id]
DELETE /v1/order/delete/:id             [delete an order based on id]
POST   /v1/order-progress/create        [create an order progress]
GET    /v1/order-progress/all           [retrieve all order progress data]
GET    /v1/order-progress/get/:id       [retrieve an order progress based on id]
PUT    /v1/order-progress/update/:id    [update an order progress based on id]
DELETE /v1/order-progress/delete/:id    [delete an order progress based on id]
GET    /v1/sales/all                    [retrieve all sales data]
GET    /v1/sales/get/:id                [retrieve a sales based on id]
DELETE /v1/sales/delete/:id             [delete a sales based on id]
POST   /v1/sales/new-sales              [create a transaction that consist of sales record and their product out]
POST   /v1/product-out/create           [create a product out]
GET    /v1/product-out/all              [retrieve all product out data]
GET    /v1/product-out/get/:id          [retrieve a product out based on id]
PUT    /v1/product-out/update/:id       [update a product out based on id]
DELETE /v1/product-out/delete/:id       [delete a product out based on id]
POST   /v1/csv/import                   [migrate the database by csv]
GET    /v1/csv/export                   [export insight report to csv, and download it]
GET    /v1/insight/value-product        [get a report for inventory value]
GET    /v1/insight/sales                [get a report for sales filtered by date range]
```

### Query Param
* `GET    /v1/csv/export` 
    To download the csv you must request `GET    /v1/insight/value-product` or `GET    /v1/insight/sales` first because csv is generated by their operation. This request just send file that created by two request above.
    ```
        ?type=(sales|product)
    ```
* `GET    /v1/insight/sales`
    ```
        ?start=(yyyy-mm-dd)&&end=(yyyy-mm-dd)
    ```
### Form/JSON Request
* `POST   /v1/product/create` 
    ```
    Content-Type : application/json
    body: {
        "sku" string required
	    "name" string required
	    "quantity" int required
    }
    ```
* `PUT    /v1/product/update/:id`
    ```
    Content-Type : application/json
    body: {
        "sku" string
	    "name" string
	    "quantity" int
    }
    ```
* `POST   /v1/order/create`
    ```
    Content-Type : application/json
    body: {
        "productId" uint required
	    "orderQuantity" int required
	    "basePrice" int required
	    "invoice" string
    }
    ```
* `PUT    /v1/product/update/:id`
    ```
    Content-Type : application/json
    body: {
        "productId" uint 
	    "orderQuantity" int
	    "basePrice" int 
	    "totalPrice" int
	    "invoice" string
    }
    ```
* `POST   /v1/order-progress/create`
    ```
    Content-Type : application/json
    body: {
        "orderId" uint required
	    "quantityReceived" int required
    }
    ```
* `PUT    /v1/order-progress/update/:id`
    ```
    Content-Type : application/json
    body: {
        "quantityReceived" int
    }
    ```
* `POST   /v1/sales/new-sales`
    ```
    Content-Type : application/json
    body: {
        "note" string required
	    "products": [
	        ...
	        {
                "productId" int required
    			"quantity" int required
    			"sellPrice" int required
	        }
	        ...
	    ]
    }
    ```
* `POST   /v1/csv/import`
    ```
    Content-Type : multipart/formdata
    form: {
        "type" string required (product, order, sales)
        "file" file required
    }
    ```