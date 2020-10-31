# KHUMU Comment API Server

khumu API 서버 중 article을 제공하는 서버와 comment를 제공하는 서버를 분리시켜 마이크로서비스 아키텍쳐를 구성.

* Golang의 Echo라는 Web framework를 이용해 REST API 구축

* clean architecture를 적절하고 간단하게 Go의 문법으로 적용시켜봄.

    * 인증은 middleware 단에서 담당
    * 인가(권한)에 대한 것은 http의 router 혹은 middleware에서 담당
    * http는 주로 http나 권한에 대한 로직, usecase는 도메인 로직(? 정확히 도메인이 뭔지 잘 모르겠음)
    * repository는 단순히 정의된 Operation을 DB를 통해 수행
    * model은 gorm에서 사용할 model들을 정의
    * container는 의존성 주입을 담당. 의존성 주입 패키지를 쓰는 것은 너무 과할 것 같아 일단 수작업으로 진행
    * **의존 순서**
        * Http
        * Usecase
        * Repository
            * Repository의 경우 어떤 DB(orm)에 대해서도 동작할 수 있도록 인터페이스로 설정 
        * Model
        * container에서 이러한 의존성을 주입한 뒤 앱을 실행할 수 있도록해준다.
        
* Test Driven까지는 아니어도 Test를 꽤 이용해보는 방향으로 개발하고 있음.
    * Test 할 때 의존성 주입을 편하게 하기 위해 concrete struct가 아닌 abstract한 interface를 사용함.
    * 그러면 Test 할 때 Mock interface를 이용할 수 있기 때문에 의존성 주입이 편하고, 테스트 하고자하는 계층이 독립된다. 
## 설정

`config/default.yaml` 을 통해 필요한 설정을 수행한다.