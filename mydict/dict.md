# Banking

### 4

#### map

- map은 key를 주면 그 key가 존재하는지도 알려준다.
- item, ok := m["keyThing"] 
- return type을 두개 적어줬으면 두개 리턴해주어야 한다.
- 슬라이스와 맵은 참조 타입이므로 포인터를 지정하지 않아도 리시버 값을 수정할 수 있다.

### 6

#### delete, search

- search의 경우에는 `map[something]`  하면 value(찾았으면 그 값), bool(존재 하는지 안하는지) 을 리턴한다.
- delete는 `delete(map, something)` 하면 아무것도 리턴하지않고 something이라는 key가 없으면 아무것도 하지않는다.