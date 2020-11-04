# KHUMU Comment API Server

khumu의 comment 관련 API를 제공하는 서버. khumu API 서버 중 khumu-command-center 가 article을 비롯한 대부분의 API를 제공하는데 이때, article을 제공하는 서버와 comment를 제공하는 서버를 분리시켜 마이크로서비스 아키텍쳐를 구성하고자 khumu-comment 서버를 분리시킴.

## 개발 방향성 및 원칙

* Golang의 Echo라는 Web framework를 이용해 REST API 구축

* clean architecture를 적절하고 간단하게 Go의 문법으로 적용시켜봄.

    * 인증은 `http/middleware` 단에서 담당
    * 인가(권한)에 대한 것은 `http`의 `router` 혹은 `middleware` 에서 담당
    * `http` 는 주로 http나 권한에 대한 로직, `usecase` 는 도메인 로직(? 정확히 도메인이 뭔지 잘 모르겠음)
    * `repository` 는 단순히 정의된 Operation을 DB를 통해 수행
    * `model` 은 `gorm` 에서 사용할 model들을 정의
    * `container` 는 의존성 주입을 담당. Contianer에 Provide하여 Container가 어떠한 의존성을 가질 지 명시하고 Build 함으로써 해당 컨테이너를 생성.
    * **의존 순서**
        * Http
        * Usecase
        * Repository
            * Repository의 경우 어떤 DB(orm)에 대해서도 동작할 수 있도록 인터페이스로 설정 
        * Model
        * container에서 이러한 의존성을 주입한 뒤 앱을 실행할 수 있도록해준다.
        
* Test Driven까지는 아니어도 Test를 꽤 이용해보는 방향으로 개발하고 있음.
    * Test 할 때 의존성 주입을 편하게 하기 위해 concrete struct가 아닌 abstract한 **interface**를 사용함.
    * 그러면 Test 할 때 Mock interface를 이용할 수 있기 때문에 의존성 주입이 편하고, 추후 확장성이 용이하다.
    * Test 주도를 적극 수용하면서 DI(의존성 주입)을 위한 IoC Container를 이용하게 되었고, dig 패키지를 이용중이다.
    * IoC container를 이용하면 test를 짤 때에도 main과 동일하게 container만 만들고 뽑아쓰면 되기 때문에 편리하다.
      Mock은 줄줄이 dependency를 주입하지 않고도 가볍게 짤 수 있었지만 method를 다 mock해야하기도 하고, 정확도가 떨어진다.  

## API Examples

### List comments

_author 쪽은 아직 미정_
```json
{
  "statusCode": 200,
  "comments": [
    {
      "id": 1,
      "kind": "anonymous",
      "author": {
         "username": "jinsu",
         "type": ""
      },
      "article": 1,
      "content": "Lorem Ipsum passages, and ",
      "parent": null,
      "children": [
        {
          "id": 2,
          "kind": "named",
          "author": {
            "username": "jinsu",
            "type": ""
          },
          "article": 1,
          "content": "more recently with desktop ",
          "parent": null,
          "children": [],
          "created_at": "2020-11-01T14:10:40.016958Z"
        }
      ]
    }
  ]
}
```
## 설정

`config/default.yaml` 을 통해 필요한 설정을 작성한다.

`config/test.yaml` 을 통해 테스트할 때 사용할 설정을 작성한다.

`KHUMU_HOME` 환경변수를 통해 루트 경로를 설정한다. 예를 들어 config는 `$KHUMU_HOME/config/local.yaml` 과 같이 작동한다.

`KHUMU_SECRET` 환경변수를 통해 jwt를 verify할 secret을 설정한다.



## How to test

### 전체 프로젝트에 대한 유닛 테스트
```bash
# 프로젝트의 루트 경로에서
$ go test ./...
# 혹은 자세한 로그를 보고싶다면
$ go test ./... -v
```

### TDD를 위한 유닛 테스트 예시

* 개발할 기능에 대한 최소한의 기능과 사용하고자하는 타입, 네이밍등을 미리 `xxx_test.go` 파일에 작성한다.
* 해당 기능을 `xxx.go` 에서 구현한다.
* 구현해나가면서 아래의 테스트 코드를 통해 새로 구현하는 내용에 대해서만 간단히 테스트 해본다.
* 구현이 끝나면 전체 유닛 테스트를 실행해본다.

이렇게 새로 구현할 내용에 대한 유닛테스트를 통해 개발을 진행하면 전체 서버를 재실행하면서
같은 작업을 반복적으로 수행하는 불편을 없앨 수 있고, 한 기능에 대한 전체 계층을 구현하지 않아도
각 계층별로 미리 테스트가 가능하다. 

e.g. _TDD를 이용하지 않고 개발하면 우선 `repository` 계층만 미리 구현하고 테스트하고 싶은데, http 계층
까지 다 구현한 뒤 서버를 띄우고 그 엔드포인트에 요청을 보내고 로그를 관찰하며 개발을 해야하지만, TDD와 Unit test를
이용해 개발하면 unit test의 결과만 보고도 개발이 가능하다._

```bash
# 개발 중인 파일이 속한 패키지의 경로에 대해 실행하고자하는 함수명을 전달한다.
# 이때 TestInit에서 Initialize 관련한 내용도 테스트하도록 설계했기때문에 TestInit도 같이 전달한다.
$ go test ./repository/ -run TestInit TestLikeCommentRepositoryGorm_Create -v
```

## 개발 팁

* go test 대신 [gotest](https://github.com/rakyll/gotest)를 이용하면 좀 더 가시성 좋게 test를 진행할 수 있다.

* `{"message": "Not Found"}` 응답의 경우 해당 경로에 대해 route 할 수 없을 때 발생하는데, 주로 `/` 의 차이인듯한데, 이 레포의 컨벤션은 맨 뒤에 `/` 를 제거하는 것을 기본으로한다.