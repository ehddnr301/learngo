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

### 3

- strings 라는 패키지를 사용합니다.
- trim space 라는 기능으로 space 를 지웁니다.
- Fields로 text부분을 다 조각내서 모든 space를 찾아내 지우려 합니다.
- `jobs = append(jobs, extractedJobs...)` jobs에 append 해줄때 `[[], [], []]` 이런식으로 되지않으려고 다 unpack되게함.

### 4

- csv 온라인 에디터로 csv 체크!
- ``` 
  w := csv.NewWriter(file)
  defer w.Flush()
- 로 생성 그리고 `wErr := w.Write()` 로 작성 Write는 error를 return