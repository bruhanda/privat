### Privatbank api
[https://api.privatbank.ua/](https://api.privatbank.ua/)

Публичные методы:

`api := privat.NewPublicApi()`

1. Курсы валют ПриватБанка
   
   * Наличный курс ПриватБанка (в отделениях):
   `api.GetExchangeRatesCash()`
   
   * Безналичный курс ПриватБанка (конвертация по картам, Приват24, пополнение вкладов):
   `api.GetExchangeRatesCard()`
   
2. Курсы валют, драгоценных металлов НБУ и ЦБ РФ

   * `api.GetGold(country)`
   
3. Архив курсов валют ПриватБанка, НБУ

   * `api.GetExchangeArchive(date time.Time)`      