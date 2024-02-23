# Deno Runner API

Welcome to the Deno Runner API! This API is designed to streamline the execution of your Deno code using the latest version and promptly deliver the output to you. With no limits per request, you have the flexibility to execute your code as needed without constraints.

## Example Request

### Deno Example

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

# Endpoint: POST /code

Feel free to explore and integrate this API into your projects! If you have any questions or feedback, please don't hesitate to reach out.