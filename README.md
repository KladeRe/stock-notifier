# Stock-Notifier

An application that gives notifications for selected stocks. 

**Any default values in this application should NEVER be taken as financial advice. This program is built to able to control your notifications and to give you complete control.**

Requirements:

* Go compiler https://go.dev/doc/install
* An API key for Alphavantage
* Node js version 18 or newer

Current backend server can be run with the following command in the server directory: 
```
go run .
```
The backend server is by default accessible at port 5050. 

The server expects and API key and MongoDB database credentials from a .env file.

Current frontend using Nuxt.js (based on Vue.js) can be run using the following command in the stock-ui directory:

```
npm run dev
```





