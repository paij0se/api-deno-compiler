# Deno Runner API

Welcome to the Deno Runner API! This API is designed to streamline the execution of your Deno code using the latest version and promptly deliver the output to you. With no limits per request, you have the flexibility to execute your code as needed without constraints.

## Example Request

### Endpoints

example with Deno

# Endpoint: POST /code

```typescript
const rawResponse = await fetch(
  "https://ad-c-9c338a775c74.herokuapp.com/code",
  {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      code: `console.log(await fetch("https://google.com"))`,
    }),
  }
);
const content = await rawResponse.json();
console.log(content);
```

# Endpoint: GET /code{ID}

example:

```typescript
const rawResponse = await fetch(
  "https://ad-c-9c338a775c74.herokuapp.com/code/tkwryp",
  {
    method: "GET",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
  }
);
const content = await rawResponse.json();
console.log(content);
```# deno-online-compiler
