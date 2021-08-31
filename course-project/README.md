# Simple server using workers

Web server created to handle workers & jobs by using concurrent operations from go.

## How to test

You can send a `POST` request to `/fibo` by using any REST Client. the structure of the body is:

```
{
    "name": "a job name",
    "delay": 4s,
    "number": 25,
}
```