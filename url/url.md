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