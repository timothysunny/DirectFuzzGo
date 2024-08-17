# DirectFuzzGo
**Version 0.01**

Direct FUZZ is a simple yet powerful directory Fuzzer written in Golang. 


- [Installation](https://github.com/timothysunny/DirectFuzzGo#installation)
- [Usage](https://github.com/timothysunny/DirectFuzzGo#usage)
- [Features](https://github.com/timothysunny/DirectFuzzGo#Features)
- [Todo](https://github.com/timothysunny/DirectFuzzGo#Todo)


## Installation

1.Clone the repository:
  ```sh
  git clone https://github.com/timothysunny/DirectFuzzGo
 ```

  ```sh
  cd DirectFuzzGo 
  go build -o dfuzz
  ```

 ## Usage
```bash
./dfuzz -u <URL> -w <wordlist>
```


## Features

- ✔️ Basic directory brute-forcing functionality
- ✔️ HTTP status code color coding for easy identification


## Todo

- `[ ]` Implement multi-threading for faster directory brute-forcing
- `[ ]` Add support for custom HTTP headers
- `[ ]` Implement timeout handling for requests
- `[ ]` Improve error handling and user messages
- `[ ]` Add more customization options

## License

DirectFuzzGO is released under MIT license. See [LICENSE](https://github.com/timothysunny/DirectFuzzGo/blob/main/LICENSE).



