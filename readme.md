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

### method 만드는 방법
- function이랑 똑같이 적는다.
- receiver 추가
  - `func`와 `func의name` 사이에 괄호를 친다 ()
  - 괄호 안에 struct의 name첫글자를 lowercase로 적어준다.
  - 그다음 띄우고 쓰려는 struct를 적어준다.
- 