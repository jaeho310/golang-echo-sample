# golang web sample
## architecture
- 3tier architecture(clean architecture)를 사용
- 

## testing
- 사용 라이브러리
```
golang 내장 testing 라이브러리
echo프레임워크에서 제공하는 test 관련 라이브러리
mockery(각 계층별 유닛테스트를 위해)
```

- 명령어
```bash
echo framework는 통합테스트만 지원, mockery 라이브러리를 사용한다. 
홈 디렉터리에서 아래 두 명령어를 입력하면 mock 구현체가 모두 구현된다.
$ go get github.com/vektra/mockery/v2/.../
$ mockery --all --keeptree
```
- test code 작성법
```
고랭 testing과 echo를 사용하면 통합테스트만 가능하다.
mock(가짜)을 사용한 unit 테스트를 위해 mockery를 사용
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