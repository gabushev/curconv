## Currency converter

### Description
This is a simple currency converter that converts the amount of money from one currency to another (with one restriction - the exchange comes within fiat and crypto!). The user can choose the currency available in the database (or modify it to make it available). The exchange rates are updated every 60 seconds.

### How to use
1. Clone the repository
2. Run "docker-compose build && docker-compose up" in the root directory
3. ...
4. ~~PROFIT!~~ You can check out the [Swagger doc](http://localhost:3000/swagger/index.html) to see the available endpoints and try them out.
5. Or just make a GET request to the endpoint
    ```bash
   $ curl -v 'http:/localhost:3000/convert?from=EUR&to=USDT&amount=1'
   ```