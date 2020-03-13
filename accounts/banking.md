# Banking

### 0

#### export
- export 해줄때는 대문자로 시작해야한다.
- export 해줄때는 위에 export 하려는것의 이름으로 시작하는 주석을 달아야한다.
- 안에 있는 요소도 대문자로 시작해야한다.

#### struct를 private 하게
- struct making function 으로 검색
- struct을 private으로 만들고
- public인 function이 struct를 만든다.

#### return &account
- 새로운 object를 리턴하고 싶은데 또 object를 만들긴 싫으니까
- 복사본 자체를 return 하는것.

### 1

#### method 만드는 방법
- function이랑 똑같이 적는다.
- receiver 추가
  - `func`와 `func의name` 사이에 괄호를 친다 ()
  - 괄호 안에 struct의 name첫글자를 lowercase로 적어준다.
  - 그다음 띄우고 쓰려는 struct를 적어준다.

### 2

#### 왜 1에서 balance가 변경되지 않았는가.

- struct의 복사본을 이용하였기 때문.
- 해결법
  - Deposit method에서 receiver에 Account앞에 `*`을 붙여준다.
  - 누군가 Deposit method를 호출하면 account를 복사하지말고 Deposit method를 호출한 account를 사용하라는 뜻

#### go error handling

- try, catch 같은게 없으므로 if 로 직접 해주어야 한다.
- errors.New, error.Error 등을 리턴한다. 리턴 타입은 error
- return nil 은 null과 같은 개념.

- 에러를 리턴해도 출력이 따로 되지는 않는다. 에러 출력도 관리 해줘야함.
- return 된 값이 nil이 아니면 log.Fatalln(err)같이 써준다.
- 혹은 그냥 fmt.Println도 된다.

### 3 

#### account 자체를 fmt.println 해보면 어떨까

- 내부 method String이 호출된다. (파이썬의 `__str__ `같은 개념)
- 수정도 가능하다.