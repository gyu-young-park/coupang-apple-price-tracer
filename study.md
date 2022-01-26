# colly
colly는 golang에서 가장 괜찮은 크롤링 라이브러리로 평가받고 있다. 

user-agent를 먼저 설정해주어야 크롤링이 가능하다.
collector를 주 객체로 사용하는데, 이것으로 네트워크 통신과 콜백의 실행을 담당한다. e.ChildText를 통하여 텍스트 값을 넣을 수 있다. 