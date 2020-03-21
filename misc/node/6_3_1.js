//p194
//요청이 들어왔을 때 콘솔에 메시지를 찍는 단순한 미들웨어
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'pug');

app.use(function(req, res, next){
    console.log(req.url, 'p195');
    next();
});
app.use(logger('dev'));

// GET / 와 GET /stylesheets/style.css 가 서버로 전달됨. 각각의 요청이 커스템 미들웨어를 작동시킴
// 주의점: 반드시 미들웨어 안에서 next()를 호출해야 다음 미들웨어로 넘어감.
// logger, express.json, express.urlencoded, cookieParser, express.static 모두 내부적으로는 next()를 호출

// next를 넣지 않으면 미들웨어에서 요청의 흐름이 끊겨버림
// next는 인자의 종류로 기능이 구분. 아무것도 넣지 않으면 단순하게 다음 미들웨어로 넘어감. 인자로 route를 넣으면 다음 라우터로
// route외의 다른 값을 넣으면 다른 미들웨어나 라우터를 건너 뛰고 바로 에러 핸들러로 이동.

//404처리 미들웨어
app.use(function(req, res, next){
    next(createError(404))
});
// 라우터에서 요청이 처리되지 않으면(일치하는 주소가 없다면) 요청은 라우터 다음에 위치한 이 미들웨어로 오게됨.
// next에 담아 에러 핸들러로 보냄.

//에러 핸들러
app.use(function(err, req, res, next){
    res.locals.message = err.message;
    res.locals.error = req.app.get('env') === 'development' ? error : {};

    res.status(err.status || 500);
    res.render('error');
}); //에러 핸들러 미들웨어는 일반적으로 미들웨어 중 제일 아래에 위치하여 위에 있는 미들웨어에서 발생하는 에러를 받아서 처리함 

// 하나의 use에 미들웨어를 여러 개 장착할 수 있음
app.use('/', function(req, res, next){
    console.log('첫 번째 미들웨어');
    next();

}, function(req, res, next){
    console.log('두 번째 미들웨어');
    next();
}, function(req, res, next) {
    console.log('세 번째 미들웨어');
    next();
});

//이 성질을 활용하여 다음과 같이 줄일 수도 있음
app.use(logger('dev'), express.json(), express.urlencoded({extended: false}),
cookieParser(), express.static(path.join(__dirname, 'public')));
