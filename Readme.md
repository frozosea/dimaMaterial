# Материал для изучения

Для начала нужно понять какие пробелы присуствуют, я запихал сюда самый нормальный материал по нужным на проекте темам.

Все свитчеры дабы понять голанг юзают вот этот гайд от самого гугла - [Link](https://go.dev/tour/). Затем сразу советую пройти курс от МФТИ 
который здесь загружен ниже. 
Но освоить синтаксис это слишком мало для понимания, ниже я оставляю список того что обязательно надо знать если пишешь на го, 
я не очень помню что дают в курсе, но он прям суперский для начала:

- Web Frameworks and utils
    - [GIN (самый популярный)](https://github.com/gin-gonic/gin)
        - [Swagger generator](https://github.com/swaggo/swag)
    - HTTP (это стандартная либа, оч простая, но мощная)
    - [GRPC](https://github.com/MaksimDzhangirov/complete-gRPC/tree/main)
    - JSON
- Base
  - [Concurrency](https://go.dev/blog/pipelines)
  - Syntax
  - [System Design](https://habr.com/ru/articles/747112/)
- Databases
  - [SQL](https://proglib.io/p/sql-for-20-minutes)
    - Понимание реляционного построения (взаимосвязи и тд)
    - [Либа для взаимодействия с бд на го](https://go.dev/doc/tutorial/database-access)
  - [Redis](https://habr.com/ru/companies/wunderfund/articles/685894/)
    - Когда использовать, зачем и т.д. Очень советую пописать свои клиенты для редиса. 
- DevOps
  - [Docker](https://www.youtube.com/watch?v=n9uCgUzfeRQ&ab_channel=%D0%92%D0%BB%D0%B0%D0%B4%D0%B8%D0%BB%D0%B5%D0%BD%D0%9C%D0%B8%D0%BD%D0%B8%D0%BD)
    - [Docker compose/swarm](https://habr.com/ru/companies/ruvds/articles/450312/)
  - [CI/CD](https://habr.com/ru/articles/476368/)

  
После того как освоишь этот материал тебе нужно будет сделать следующие задания:

- [Решить тестовое по golang](https://github.com/CawaKharkov/golang-test-task/tree/master/meta). Его решение в папке `meta`, это то как его делал я
