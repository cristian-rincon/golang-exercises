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

If you are using [httpie](https://httpie.io/), you can test by using this line: `http -f POST http://localhost:8081/fibo name=fibo25 delay=3s number=25`