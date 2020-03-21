//p202
var logger = require('morgan');
var session = require('express-session');

var indexRouter = require('./routes/index');
var usersRouter = require('./routes/users');

app.use(cookieParser('secret code'));
app.use(sessioin({
    resave: false,
    saveUninitialized: false,
    secret: 'secret code',
    cookie: {
        httpOnly: true,
        secure: false,
    },
}));

/*
express-session은 인자로 세션에 대한 설정을 받음.
resave는 요청이 왔을 때 세션에 수정사항이 생기지 않더라도 세션을 다시 저장할지
saveUninitialized는 세션에 저장할 내역이 없더라도 세션을 저장할지. 보통 방문자를 추적할 때 사용.

express-session은 세션 관리 시 클라이언트에 쿠키를 보냄. 이를 세션 쿠키라고 부름
안전하게 쿠키를 전송하려면 쿠키에 서명을 추가해야하고, 쿠키를 서명하는 데 secret의 값이 필요.

maxAge, domain, path, expires, smaeSite, httpOnly, secure 등 일반적인 쿠키 옵션이 모두 제공됨
현재는 httpOnly:true 로 클라이언트에서 쿠키를 확인하지 못하도록
secure:false 로 https가 아닌 환경에서도 사용할 수 있게.

현재는 메모리에 세션을 저장. 서버를 재시작하면 메모리가 초기화되어 세션이 사라짐. 
배포 시에는 store에 데이터베이스를 연결하여 세션을 유지하는 것이 좋음 

express-session은 req 객체 안에 req.session객체를 만듦. 이 객체에 값을 대입하거나 삭제하여 세션 변경
나중에 세션을 한번에 삭제하려면 req.session.destroy() 호출
현재 세션의 아이디는 req.sessionID로 확인가능. 
*/