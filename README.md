# golang web sample
## 사용법
- go run main.go
- 웹브라우저에서 localhost:8080접근 or postman, curl 등 rest client를 사용

## architecture
- 3tier architecture(clean architecture)를 사용
  - ddd에서 파생된 국내에서 가장 많이 쓰는 구조
  - presentation layer, business layer, data access layer 로 나뉜다.
- presentation 계층(interfaces/controller)
  - 외부와 인터페이스 역할을 한다.
- business 계층(service)
  - 비즈니스 로직을 작성한다.
- data access 계층(repository)
  - 영속성 계층으로 쿼리문 등 영속성에 관련된 내용을 작성한다.
- infrastructure계층
  - 애플리케이션의 기반이 되는 패키지들을 모아놓는다.(server, db...etc)

## testing
- 사용 라이브러리
```
통합테스트는 golang에 내장된 test 라이브러리를 사용한다.
mockup 같은경우 두 진영이 존재한다.

go get -u github.com/golang/mock/gomock
                  vs
go get github.com/stretchr/testify/mock

start 수는 testify가 golang 기본 네임스페이스보다 많다(13.8k, 5.9k)

given when then을 사용하여 테스트 하는 방식은 거의 같으므로 

사용하고 싶은걸 사용하면 된다.

mock 구현체는  
golang/mock 진영은 mockgen을 사용
stretchr/testify 진영은 vektra/mockery를 사용하여 구현한다.

양 진영 모두 장단점이 존재하지만,
편의성때문에 testify 를 더 많은 사람들이 사용하고 있다.

ex)
mockgen -destination=$PWD/mocks -package mocks github.com/sgreben/gomock-vs-testify/comparison/gogenerate Interface1,Interface2
mockery --all

예시에서도 testfiy와 mockery 조합을 채택하였다.

두 진영에 대한 비교는 아래 블로그에 상세히 작성되어있다.
https://blog.codecentric.de/2019/07/gomock-vs-testify
```

- 명령어
```bash
echo framework는 통합테스트만 지원, mockery 라이브러리를 사용한다. 
홈 디렉터리에서 아래 명령어를 입력하면 mock 구현체가 모두 구현된다.
$ go get github.com/stretchr/testify/mock
$ go get github.com/vektra/mockery/v2/.../
$ mockery --all --keeptree
```
- test code 작성법
```
고랭 testing과 echo를 사용하면 통합테스트만 가능(영속성 계층을 갈아끼워서 테스트)
mock(가짜)을 사용한 unit 테스트를 위해 testify와 mockery를 사용한다.
각 계층간 interface역할을 하는 type ~~ interface의 mock을 mockery가 구현해준다. 
```



## sonarqube
- sonarqube config 파일
```
#Configure here general information about the environment, such as SonarQube server connection details for example
#No information about specific project should appear here

#----- Default SonarQube server
sonar.host.url=http://192.168.102.127:9000/
sonar.login=2a09a7d77a7c54c94b33b5a7270b30907d023127

# #----- Default source code encoding
sonar.sourceEncoding=UTF-8

sonar.projectKey=golang-echo-sample
sonar.projectName=golang-echo-sample
# sonar.projectVersion=1.0
sonar.language=go
sonar.sources=.
sonar.exclusions=**/mock/**,**/secret/**,**/docs/**,**/data/**,.idea/**,**/vendor/**
sonar.sourceEncoding=UTF-8
sonar.tests=.
sonar.test.inclusions=**/*_test.go
sonar.test.exclusions=**/vendor/**
sonar.go.coverage.reportPaths=**/coverage.out
```
- 명령어
```bash
$ go test -v ./... -coverprofile=coverage.out
$ go test -v ./... -json > report.json
$ sonar-scanner
```