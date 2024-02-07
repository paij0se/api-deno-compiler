<h1>this a simple api that execute your deno code and send you the output, has not limit per request</h1>

example request:

in deno:

```ts
const rawResponse = await fetch(
  "https://api-deno-compiler.herokuapp.com/code",
  {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      code: `console.log(await fetch("https://api-deno-compiler.herokuapp.com/code"))`,
    }),
  }
);
const content = await rawResponse.json();
console.log(content);

```

other example with deno, with more requests:

```ts
const code = [
  `console.log("hello world")`,
  `console.log(Deno.version)`,
  `console.log("🍱 🦕")`,
  `for(let i=0;i<10;i++){console.log("number:",i)}`,
  `this would have an error`,
];

for (let i = 0; i < 10; i++) {
  const rawResponse = await fetch(
    "https://api-deno-compiler.herokuapp.com/code",
    {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        code: `${code[Math.floor(Math.random() * code.length)]}`,
      }),
    },
  );
  const content = await rawResponse.json();
  console.log(content);
}
```

in python:

```py
import requests

code = """
console.log(Deno.memoryUsage()
"""

r = requests.post("https://api-deno-compiler.herokuapp.com/code",
                  json={"code": code})
print(r.text)
```

in go:
```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	postBody, _ := json.Marshal(map[string]string{
		"code": "console.log(Deno.version)",
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("https://api-deno-compiler.herokuapp.com/code", "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	fmt.Printf(sb)
	duration := time.Since(start)

	fmt.Printf("API Response Time: %d%s\n", duration.Milliseconds(), "ms")
}


```
<h1>Used in:</h1>

- Deno online compiler : https://deno-online-compiler.herokuapp.com/

- dino-bot : https://github.com/ELPanaJose/dino-bot
# api-deno-compiler
# api-deno-compiler
