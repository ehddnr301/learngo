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