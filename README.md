# KHUMU Comment API Server

**khumu-comment**는 MSA로 개발중인 khumu의 comment 관련 API를 제공하는 서버이고, `Echo` 라는 Golang의 웹프렘워크를 바탕으로 개발 중입니다.
khumu API 서버 중 khumu-command-center 가 article을 비롯한 대부분의 API를 제공하고있으며, article을 제공하는 서버와 comment를 제공하는 서버를 분리시켜 마이크로서비스 형태로 개발하고자 
khumu-comment 서버를 이렇게 따로 개발하고 있습니다.

## API Documentation: 

https://documenter.getpostman.com/view/13384984/TVsvfkxs

## ⚙️설정

`config/default.yaml` 을 통해 필요한 설정을 작성한다.

`config/test.yaml` 을 통해 테스트할 때 사용할 설정을 작성한다.

`KHUMU_HOME` 환경변수를 통해 루트 경로를 설정한다. 예를 들어 config는 `$KHUMU_HOME/config/local.yaml` 과 같이 작동한다.

`KHUMU_SECRET` 환경변수를 통해 jwt를 verify할 secret을 설정한다.


## 💯 테스트를 진행하는 방법

### 프로젝트 내의 모든 유닛 테스트

현재는 Github Action에서 매 푸시마다 자동으로 전체에 대한 unit test를 진행하고,
이것이 모두 통과하면 docker image 빌드 후 private devops 레포지토리에 새 빌드를 적용시키고,
`ArgoCD` 를 통해 자동 배포가 된다.

```bash
# 프로젝트의 루트 경로에서
$ go test ./...
# 혹은 자세한 로그를 보고싶다면
$ go test ./... -v
```

### TDD 식의 개발 방법 - 선 유닛 테스트 작성. 후 개발/리팩토링 

> test는 MySQL이 아닌 SQLite3를 메모리를 바탕으로 간단하게 이용한다.

* 개발할 기능에 대한 최소한의 기능과 사용하고자하는 타입, 네이밍등을 미리 `xxx_test.go` 파일에 작성한다.
* 해당 기능을 `xxx.go` 에서 구현한다.
* 구현해나가면서 아래의 커맨드 코드를 통해 새로 구현하는 내용에 대해서만 간단히 테스트 해본다.
* 구현이 끝나면 전체 유닛 테스트를 실행해보고, 좀 더 나은 방향으로 리팩토링한다.
* 리팩토링이 완료되면 다시 새로운 기능에 대해 유닛테스트를 작성한다.

이렇게 새로 구현할 내용에 대한 유닛테스트를 통해 개발을 진행하면 전체 서버를 재실행하면서
같은 작업을 반복적으로 수행하는 불편을 없앨 수 있고, 한 기능에 대한 전체 계층을 구현하지 않아도
개발하는 동안 각 계층별로 미리 테스트가 가능하다.

**e.g.** _TDD 방식을 이용하지 않고 개발하면 우선 `repository` 계층만 미리 구현하고 테스트하고 싶은데, http 계층
까지 다 구현한 뒤 서버를 띄우고 그 엔드포인트에 요청을 보내고 로그를 관찰하며 개발을 해야하지만, TDD와 Unit test를
이용해 개발하면 unit test의 결과만 보고도 개발이 가능하다._

```bash
# 개발 중인 파일이 속한 패키지의 경로에 대해 실행하고자하는 함수명을 전달한다.
# 이때 TestSetUp에서 Initialize 관련한 내용도 테스트하도록 설계했기때문에 TestSetUp도 같이 전달한다.
$ go test ./repository/ -run TestSetUp TestLikeCommentRepositoryGorm_Create -v
```


## 📚 개발 팁 및 메모

* go test 대신 [gotest](https://github.com/rakyll/gotest)를 이용하면 좀 더 가시성 좋게 test를 진행할 수 있다.

* `{"message": "Not Found"}` 응답을 받는 경우 `echo` 가 해당 경로에 대해 route 할 수 없을 때 발생하는데, 주로 주소의 맨 끝 `/` 의 차이인듯한데
  이 레포의 컨벤션은 맨 뒤에 `/` 를 제거하는 것을 기본으로한다.


## 🚀 개발 방향성 및 원칙

### 1. clean architecture를 적절히 적용하자.

> 의존 순서 및 상위 계층 정도를 기준으로 각 계층의 역할에 대해 나열해보았습니다.

* **http**
  * 주로 http 통신 자체에 대한 로직을 담고있습니다. 가장 상위 계층입니다.
  * Router, Middleware, Authentication, Authorization 와 같은 작업을 다룹니다.
  * `struct` => `json` 으로 `marshal` 한 뒤 그 정보를 바탕으로 Response를 구성하는 로직을 담기도 합니다. 
    (e.g. `Comment` struct를 받아서 json으로 변환한 뒤 Response의 body를 작성하는 작업을 진행합니다.   

* **usecase**
  * repository 계층으로부터 단순히 model을 얻어온 뒤 json으로 직렬화되기 전까지의 로직, 즉 대부분의 비즈니스 로직(_? 맞는진 잘 모르겠음. 애플리케이션 로직?_)을 담당합니다.
  * 예를 들어 repository에서 얻어온 `Comment` 배열 중 자신이 작성자가 아닌 익명 댓글은 author의 username 필드의 값을 `익명` 으로 변경합니다.

* **repository**
  * Database와 직접 작업을 하는 계층입니다. 
  * 어떤 DB(orm)에 대해서도 동작할 수 있도록 인터페이스로 설정되어있습니다.(그렇다고 이 계층에서만 인터페이스를 사용하는 것은 아닙니다.)
  * 예를 들어 DB에서 Select 혹은 Find하는 명령과 직결된 작업을 수행합니다.
  * 예를 들어 연결하고자하는 DB가 SQLite3이든, MySQL이든 상위 계층을 repository interface를 인자로 받기 때문에
    어떤 database를 이용하든 상위계층은 동일하게 동작할 수 있습니다.
  
* **model**
  * Database에서 사용하는 Table에 대한 정의입니다.
  * 클린 아키텍쳐와 관련된 얘기는 아니지만, 한 Table도 여러 type에서 이용할 수 있습니다. 
    모든 type이 Table로 migrate 되는 것은 아니고, migrate할 type만 골라서 migrate합니다.
  * 위의 다른 계층들 중 아무 계층도 참조하지 않는 최하위 계층입니다.

* **container**
  * 컨테이너는 위의 모든 계층들과 달리 상위 계층, 하위 계층의 개념을 갖지 않고, IoC Container와 관련된 작업을 수행하는 패키지입니다.
  * IoC 컨테이너에 대한 정의와 의존성 주입을 담당합니다.
  * 개발자는 수작업으로 의존성을 주입해주거나, struct를 생성할 필요 없이 container가 type을 기반으로 자신(container)에 해당 타입의 변수가 존재하면
    그것을 이용할 수 있게해주고, 없다면 생성한 뒤 이용할 수 있게 해줍니다.
  * 현재는 uber의 `dig` 패키지를 의존성 주입 패키지로 사용 중입니다. 구글의 `wire` 가 꽤 유명한 것 같지만, 가독성을 해칠 것 같고, 유연하지 않은 듯하여 배제했습니다.
    uber의 `fx` 는 `dig` 를 한 단계 더 감싼 패키지인듯한데, 마찬가지로 유연성이 떨어지는 느낌을 받았습니다. 

* **config** : 프로그램에 대한 설정 정보나 그 정보를 불러오는 작업을 담당합니다.

### 2. TDD(Test Driven Development)를 통해 개발하자.

* 큰 장점들
  * 지속적인 개발에 대한 신뢰와 안정성이 상승하고, 이는 생산성으로도 연결된다.
  * 또한 당장의 개발에서도 unit test를 통해 계층을 나누어 개발하기 편리하게 때문에 생산성이 증가된다.

* 원래는 의존성 주입 패키지를 사용하지 않았는데, test code를 짜게 되면서 수동으로 의존성을 넣는 것이 번거롭기도 하고 가독성도 안 좋은 것 같아
  의존성 주입 패키지를 사용하기 시작했다.
  
* Mocking 하는 경우
  * struct 형 인자가 아닌 interface 형 인자를 이용하면 의존성을 주입할 때 mock type을 주입할 수 있다.
  * mock type을 이용하면 하위 계층의 내용과 독립되게 해당 계층만 테스트 할 수 있다.
  * mock type을 이용하면 하위 계층의 하위 계층에 대한 의존성, 그 하위 계층의 더 하위 계층에 대한 의존성을 모두 주입해 줄 필요 없이
    내가 직접 필요한 계층만 주입하면 된다는 점이 편리하다.
  * 다만 하위 계층을 흉내냈다는 점에서 실제 하위 계층의 동작과 다르게 동작할 수 있다는 면이 해당 계층의 테스트의 정확성을
    낮출 수 있다.
  * 의존성을 주입하는 것이 오히려 mock methods를 정의하는 것보다 편리한 경우도 많다.  
  
* 각 Test 별 독립성이 테스트가 가능하도록 하자.

  * Java Spring의 BeforeEach와 AfterEach에서 아이디어를 얻어 `B`와 `A`라는 함수를 정의하기로 했다. 기본적으로 초기 데이터가 필요한 테스트들은 모두 아래와 같이 B와 A를 이용해 Set up과 clean up을 진행한다.

    ```go
    func TestFoo(t *testing.T) {
        B()
        defer A()
        // some test scenarios.
    }
    ```

  * `github.com/khu-dev/khumu-comment/test`  패키지에 초기 데이터 형식과 필요한 몇 가지 함수를 정의해놓았다.

  * 독립적인 유닛 테스트가 가능해져 GoLand에서 지원하는 test를 바로 실행시켜주는 기능도 편리하게 이용 가능해졌다.


## 📚 Golang 개발 이야기

### `embedding` 을 통한 의존성 주입할 타입 정의하기

type을 기반으로 의존성 주입을 자동화하는 의존성 주입 패키지를 사용하는 경우
동일 타입이지만 다른 객체를 주입하고 싶은 경우 난감한 경우가 있다. 이런 경우에는 `embedding` 을 통해 원래 타입의
메소드와 필드를 모두 사용하면서 개별적인 type으로 이용할 수 있다. 예를 들어 router에서는 자식 router group은 parent router group을
인자로 받고싶은데, 그냥 `*echo.Group` 을 주입받겠다고 정의하면, 어떤 `*echo.Group` 을 주입받게 될 지 모른다. 따라서 아래와 같이
`embedding` 을 통해 원래의 메소드와 필드를 모두 사용하면서 주입받을 새로운 타입을 정의할 수 있다.

```go
// embedding을 통해 *echo.Group의 메소드, 필드를 이용할 수 있는 타입 정의
// 이 타입을 인자로 받는 메소드는 일반적인 *echo.Group 타입과 구별된 RootRouter Type을 이용할 수 있다.
type RootRouter struct{*echo.Group}

func NewRootRouter(echoServer *echo.Echo, ... 인자 생략) *RootRouter{
    g := RootRouter{Group: echoServer.Group("/api")}
    //... 작업 생략
    return &g
}

func NewCommentRouter(root *RootRouter, ... 인자 생략) *CommentRouter {
    // 특이하게 Type명과 이용하고자하는 메소드 명이 같아서 이렇게 사용할 뿐 원래는 embed 시 root.Group("/comments")로 사용 가능 
    group := root.Group.Group("/comments") 
    commentRouter := &CommentRouter{group, ... 인자 생략}
    return commentRouter
}
```

### `interface` 를 통해 mock type을 이용하여 의존성 주입이 필요 없는 테스트하기

`http` 계층에 대한 테스트 코드를 짠다고 가정하자. `http`는 `usecase` 계층에 의존적이다.
그럼 테스트 코드를 짤 때 `http` 생성 시 `usecase` 를 생성하여 주입시켜주어야한다. 근데 반복적으로
`usecase` 는 `repository` 에 의존적이므로 `repository` 를 생성하여 주입받아야한다. 따라서 이러한
의존 파이프라인을 만족시켜주기 번거롭기때문에 mock type을 이용해 테스트를 하고싶다.

만약 아래와 같은 코드에서 CommentUseCase에 관한 mock type을 정의한다면 사용이 가능할까?

```go
func NewCommentRouter(root *RootRouter, uc CommentUseCaseStruct) *CommentRouter {
    ... 작업 생략
    return commentRouter
}
```

mock type을 이용하기 불가능하다. 이유는 CommentUseCase 역할을 하는 인자가 struct 타입으로 정의되어있기 때문이다.
mock type을 정의한다고 해도 그 type은 위에 정의된 `CommentUseCaseStruct` type이 될 수 없다.

따라서 주입받는 인자의 type을 concrete한 struct가 아닌 abstract한 interface로 정의해주면 된다.
아래처럼 mock type이 해당 인터페이스가 되기 위한 메소드들만 필요한 만큼만 원래 타입을 흉내내어 구현해주면된다.

```go
type CommentUseCaseInterface interface{
    List() []*model.Comment
}
type CommentUseCaseMock struct{}

// 간단하게 필요한 만큼만 원래의 기능을 흉내낸다.
func (uc *CommentUseCaseMock) List []*model.Comment{
    return []*model.Comment{
        &model.Comment{...생략}, &model.Comment{...생략}, &model.Comment{...생략}
    }
}

// 주입받는 인자의 타입을 interface형으로 정의했기때문에 실제적인 CommentUseCase이든 가짜의 CommentUseCaseMock 타입이든
// 상관 없이 주입받을 수 있다.
// 서버를 돌릴 때에는 CommentUseCase를, 테스트 할 때는 의존성 주입이 편리한 CommentUseCaseMock을 사용하면 된다.
func NewCommentRouter(root *RootRouter, uc CommentUseCaseInterface) *CommentRouter {
    ... 작업 생략
    return commentRouter
}
```
