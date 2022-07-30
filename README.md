# go-btc-rate-api
Golang API that works with bitcoin rate from [Binance API]("https://api.binance.com/api/v3/ticker/price?symbol=BTCUAH")

### Running with Docker 
1. Clone repository
2. Create ```.env``` file in the root folder and create EMAIL_ADDRESS and EMAIL_PASSWORD variables with email address data that will be used for service work file.
    >For email accounts with **two-factor authentication**:
    >-[generate](https://support.google.com/mail/answer/185833?hl=en) application password
    > -set EMAIL_PASSWORD value to application password 

3. Build and run Docker image:
  - run ```docker build --tag docker-btc-rate-api .```
  - run ```docker run -d -p 8080:8080 docker-btc-rate-api```
  Service will be accessible by http://localhost:8080 url


### Endpoints
 **GET** /rate - returns current BTC-UAH rate
 **POST** /subscription - signs an email to receive information on exchange rate
 **POST** /sendEmail - sends current BTC-UAH rate to all subscribers


### Endpoints statuses
##### **GET** /rate
- `200` - returns BTC-UAH rate
##### **POST** /subscription
- `200` - email signed
- `400` - email not valid
- `409` - emails is already subcribed
##### **POST** /sendEmail
- `200` - messages were successfully sent