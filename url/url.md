# URL

### 0

- somethings를 for문으로 돌게 되면 slice 이기때문에
- ` for index, something := range somethings ` 이렇게 적어줄수있다.

#### url request

- request 가 go standard library 에 있기 때문에 사용하면 된다.
- http 인데 이것은 res, err 를 리턴한다.


### 1

- 선언만하고 initialize 하지 않은 map 에다가 값을 넣을수는 없다.
- ` var results = map[string]string ` 에다가 값을 넣을수는 없다.
- ` var results = map[string]string{} ` 라면 가능
- 혹은 `make` function을 이용한다. (map 을 만들어주는 함수)
- `make(map[string]string)` 처럼 이용한다.


### 2

- 순서대로가 아니라 동시에 처리하는 방법.

#### Goroutines

- Goroutines 는 프로그램이 작동하는동안만 유효하다. (메인함수가 유효한 동안)
- main func 은 다른 goroutines를 기다려 주지않는다.
- 다른함수와 동시에 실행시키는 함수.
- 함수 호출 앞에 `go`를 붙여주면 된다.
- `go countBoy()`
- `go countGirl()` 이렇게 두개만 있으면 함수가 종료되지만
- `go countBoy()`
- `  countGirl()` 이렇게 두개가 있으면 countGirl을 하는동안 에는 countBoy 도 살아있기 때문에 두개다 실행이 된다.

#### time

- time.Sleep(time.Second) 하면 1초를 쉰다.
- time.Sleep(time.Second * 5) 하면 5초를 쉰다.


### 3

#### Channel

- goroutine이랑 메인함수 사이에 정보를 전달하기 위한 방법.
- goroutine 끼리도 정보전달 가능.
- pipe같은 것으로 메시지를 보낼수도 있고 받을수도 있다.

#### Channel 만들기.

- 어떤 타입의 데이터를 보낼건지 chan 뒤에 적어줘야한다.
- `c := make(chan bool)`
- 그리고 goroutine 에서는 어떤 채널을 사용하는지 알려줘야한다.
- `go isSexy(person, c)`
- 그리고 function 에서는 어떤 routine을 사용하는지 표시한다.
- `func isSexy(person string, c chan bool){}`
- 그리고 function 에서는 return이 아니라 `<-` 을 사용한다.
- `c <- true`
- main함수에서 function에서 보낸 true를 찾으려면 ?
- `result := <-c` 해준다
  - 채널로 부터 뭔가 받을 때 메인함수는 어떤 값이 올때까지 기다린다.
  - `fmt.Println(<-c)` 도 가능하다.