# Job Scrapper


### 0

- goquery를 이용합니다.
- `go get github.com/PuerkitoBio/goquery` 를 터미널에 입력.
- defer는 함수가 끝난뒤에 실행된다.
- goquery를 이용해서 pagination 클래스를 찾고 Each로 각각 함수를 실행시킨다.

### 1

- `strconv.Itoa` 를 사용하여 string으로 바꾸어 준다.


### 2

- `Each()` 는 `func(i int, s *goquery.Selection){}` 가 필요하다.