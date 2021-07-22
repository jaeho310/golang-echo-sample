# golang web sample
## 요약
- golang echo framework를 사용한 api server sample
- 동작확인을 위해 간단하게 프론트단을 개발(웹브라우저에서 접근 가능)
- postman, curl 등 rest client를 사용하여도 가능


## profile파일 추가
- profile을 파일로 관리하여 원하는 자원을 편하게 쓸수 있게 한다.
    - (ex 매번 코브라 아규먼트나, 환경변수 등으로 여러개의 변수를 관리하는 방식은 생산성이 떨어진다.)
- 현재는 런타임시에만 일시적으로 환경변수를 적용시켜 관리하지만 yaml파일로 변경 예정

## Architecture
- 3tier architecture(clean architecture)를 사용
  - layered architecture중 국내에서 가장 많이 사용하는 구조(전자정부 프레임워크의 기본 아키텍쳐)
  - presentation layer, business layer, data access layer 로 나뉜다.
- presentation 계층(interfaces/controller)
  - 외부와 인터페이스 역할을 한다.
- business 계층(service)
  - 비즈니스 로직을 작성한다.
- data access 계층(repository)
  - 영속성 계층으로 쿼리문 등 영속성에 관련된 내용을 작성한다.
- infrastructure계층
  - 애플리케이션의 기반이 되는 패키지들을 모아놓는다.(server, db...etc)

## Inversion of Control과 Dependency Injection
- ioc와 di를 이용하여 계층구조의 의존성을 관리
    - 각 레이어는 내가 어떤 의존성을 갖는지 직접적으로 알 필요가 없으며(추상화된 주입이 이루어진다.)
    - 초기화시 ioc를 통해 계층관계의 의존성을 넣어준다.
- ioc
``` go
func (server Server) InjectDb() *gorm.DB {
	return server.MainDb
}

func (server Server) InjectUserRepository() *repository.UserRepositoryImpl {
	return repository.UserRepositoryImpl{}.NewUserRepositoryImpl(server.InjectDb())
}

func (server Server) InjectUserService() *service.UserServiceImpl {
	return service.UserServiceImpl{}.NewUserServiceImpl(server.InjectUserRepository())
}

func (server Server) InjectUserController() *api.UserController {
	return api.UserController{}.NewUserController(server.InjectUserService())
}

```
- di
```go
type UserServiceImpl struct {
	repository.UserRepository
}

func (UserServiceImpl) NewUserServiceImpl(repository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repository}
}
```

## testing
- 사용 라이브러리
```
통합테스트는 golang에 내장된 test 라이브러리를 사용한다.
mockup의 경우 두 진영이 존재한다.

github.com/golang/mock/gomock
            vs
github.com/stretchr/testify/mock

start 수는 testify가 golang 기본 네임스페이스보다 많다(13.8k, 5.9k)

given when then을 사용하여 테스트 하는 방식은 거의 같으므로 

사용하고 싶은걸 사용하면 된다.

mock 구현체는  
golang/mock 진영은 mockgen을 사용
stretchr/testify 진영은 vektra/mockery를 사용하여 구현한다.

양 진영에 장단점이 존재하지만,
편의성때문에 testify 를 더 많은 사람들이 사용하고 있다.(테스트의 기본적인 기능은 양진영 모두 포함하고 있어서 큰 차이가 없다.)

ex) mock 구현명령어만 봐도 mockery가 쉽게 해놨다.
mockgen -destination=$PWD/mocks -package mocks github.com/sgreben/gomock-vs-testify/comparison/gogenerate Interface1,Interface2
mockery --all

이번 예시에서도 testfiy와 mockery 조합을 채택하였다.

두 진영에 대한 비교는 아래 블로그에 상세히 작성되어 있다.
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
## test코드 작성법
- 통합테스트
```
고랭 testing과 echo(controller 사용을 위해)를 사용하여 테스트하며
영속성 계층이 의존하는 db만(직접 mock을 생성해 변경), 주입하여 테스트한다.

모든계층을 테스트 할 수 있으며, db만 같다면 실제 운영환경과 같은 환경으로 테스트가 가능하다.
(운영환경과 db까지 같은환경을 원한다면 db도 변경없이 테스트한다.)

원하는 계층의 mock을 직접 생성하면 단위테스트도 할 수있지만
 
mock객체 생성은 mockery 라이브러리를 이용한다. 
```

- 단위테스트
```
testify와 mockery를 사용한다.
각 계층간 interface역할을 하는 type ~~ interface의 mock을 mockery가 구현해주므로
나는 원하는 계층이 원하는 로직을 실행하는지 손쉽게 테스트 할 수 있다.
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
