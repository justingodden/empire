# The Empire Programming Language
<p align="center">
🚀🚀🚀
</p>
Empire is an interpretted language written 100% in Go, created to learn how programming languages are built, but ended up being pretty feature-full!

Including the following:
* C-like syntax
* REPL
* variable bindings
* integers and booleans
* arithmetic expressions
* built-in functions
* first-class and higher-order functions
* closures
* a string data structure
* an array data structure
* a hash data structure
* loops
* library imports
* small but easily extendible standard library

Empire interprets **.emp** files.

## Build
### Linux
#### 1. Clone repo
```bash
git clone https://github.com/justingodden/empire.git

cd empire
```

#### 2. Build binary and make it executable
```bash
mkdir -p ./bin

go build -o ./bin/empire ./cmd/empire/main.go

chmod +x ./bin/empire
```

#### 3. (Recommended) add binary and stdlib to /usr/local
```bash
sudo mkdir -p /usr/local/empire/bin && sudo cp ./bin/empire /usr/local/empire/bin

sudo mkdir -p /usr/local/empire/stdlib && sudo cp -r ./stdlib /usr/local/empire
```
 #### 4. Add executable to PATH (note this will not be permanent beyond the current terminal)
 ```bash
export PATH="$PATH:/usr/local/empire/bin"
 ```

#### 5a. Run the Empire REPL using the **'empire'** command
```bash
➜  [user] $ empire

▓█████  ███▄ ▄███▓ ██▓███   ██▓ ██▀███  ▓█████ 
▓█   ▀ ▓██▒▀█▀ ██▒▓██░  ██▒▓██▒▓██ ▒ ██▒▓█   ▀ 
▒███   ▓██    ▓██░▓██░ ██▓▒▒██▒▓██ ░▄█ ▒▒███   
▒▓█  ▄ ▒██    ▒██ ▒██▄█▓▒ ▒░██░▒██▀▀█▄  ▒▓█  ▄ 
░▒████▒▒██▒   ░██▒▒██▒ ░  ░░██░░██▓ ▒██▒░▒████▒
░░ ▒░ ░░ ▒░   ░  ░▒▓▒░ ░  ░░▓  ░ ▒▓ ░▒▓░░░ ▒░ ░
 ░ ░  ░░  ░      ░░▒ ░      ▒ ░  ░▒ ░ ▒░ ░ ░  ░
   ░   ░      ░   ░░        ▒ ░  ░░   ░    ░   
   ░  ░       ░             ░     ░        ░  ░
                                               
Hello justin! Welcome to the Empire progamming Language.
Let's get started 🚀
>> 
```

#### 5b. Run an .emp src file
```bash
empire ./filename.emp
```



## Syntax
### Fibonacci example
```rust
let fibonacci = fn(x) {
    if (x == 0) {
        0
    } else {
        if (x == 1) {
            1
        } else {
            fibonacci(x - 1) + fibonacci(x - 2);
        }
    }
};
```

### Library imports
```go
import ("math",)

import (
    "math",
    "otherlib",
)
```

### Print function
```python
print("Hello World!")
```

### Variables
```rust
let age = 20;
let name = "Empire";
let result = 10 * (10 / 2)
```

### Functions
```rust
// Create a function using the 'fn' keyword
let add = fn(a, b) {
    return a + b
};

// First-class functions - passing a function, as a function argument
let twice = fn(f, x) {
    return f(f(x));
};
let addTwo = fn(x) {
    return x + 2;
};

twice(addTwo, 2);
```

### Conditionals
```rust
!true
!false
!!!!!!true

foo == bar
foo != bar
foo < bar
foo > bar

if (true) { return true } else { return false };

let result = if (10 > 5) { true } else { false };

let factorial = fn(n) { 
    if (n == 0) {
        1
    } else {
        n * factorial(n - 1)
    }
};
```

### While loops
```python
while (true) {
    doFunc()
}
```

### 'For' loops (modified while loop)
```js
let arr = [1, 2, 3];
let i = 0;

while (i < len(arr)) {
    print(arr[i])
    let i = i + 1
}
```

### Arrays
```js
let myArr = [1, 2, 3, 4, 5]

len(myArr)
first(myArr)
last(myArr)
rest(myArr)
push(myArr)
[1, 2, 3 + 3, fn(x) { x }, add(2, 2)][3](100)

myArr[2];
```

### Hash maps
```rust
let myHash = {"name": "John", age: 22, true: "hello"};
print(myHash["name"])
```

### Standard library
```go
import ("math",)

let res = mathPow(2, 4);

print(res)
```

