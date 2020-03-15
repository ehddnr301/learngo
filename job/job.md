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

### 5

- goroutine 플랜!
  - getPages 를 실행!
  - 그밑에 getPage가 동시에 실행되길 원함 (각 페이지별로 goroutine을 생성)
  - getPage밑에 extractJob을 동시에 실행되길 원함. (각 card별로 gorutine을 생성)

- gorutine 플랜!
  - extractJob 함수는 getPage함수와 channel을 이용해서 정보를 주고받음.
  - getPage함수는 main함수와 channel을 이용해서 정보를 주고받음.

- 작업순서
  - 제일 밑에있는 extractJob부터 goroutine을 생성.

- 상세작업
  - getPage에 channel을 만든다. `c := make(chan extractedJob)`
  - extractJob에 인자로 c라는 채널을 준다.
  - extractJob 함수에 `c chan<- extractedJob`을 적어준다.
  - 이제 extractJob 함수는 return 대신 채널에 값을 전송한다.
  - extractJob의 결과를 job에 저장하는 대신 `go extractJob(card, c)` 로 고루틴으로 만들어준다.
  - ```
      for i:=0; i<searchCards.Length();i++{
    		job := <-c
    		jobs = append(jobs, job)
    	}
  - 로 채널에서 값을받아준다.