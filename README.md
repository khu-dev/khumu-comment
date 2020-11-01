# KHUMU Comment API Server

khumu API 서버 중 article을 제공하는 서버와 comment를 제공하는 서버를 분리시켜 마이크로서비스 아키텍쳐를 구성.

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

## How to test

```bash
# 프로젝트의 루트 경로에서
$ go test ./...
# 혹은 자세한 로그를 보고싶다면
$ go test ./... -v
```