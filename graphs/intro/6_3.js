//p192
/*
요청과 응답의 중간에 위치. 
라우터와 에러 핸들러도 미들웨어의 일종.
미들웨어는 주로 app.use와 함께 사용됨

app.use 메서드의 인자로 들어 있는 함수가 미들웨어. 미들웨어는 use 메서드로 app에 장착
*/
app.use(logger('dev')); //미들웨어들을 순차적으로 거친 후 라우터에서 클라이언트로 응답을 보냄 
app.use(express.json());
app.use(express.urlencoded({extended: false}));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, 'public')));

app.use('/', indexRouter);
app.use('/users', usersRouter);

app.use(function(req, res, next){
    next(createError(404));
});


// 에러 핸들러
app.use(function(err, req, res, next){
    res.locals.message = err.message;
    res.locals.error = req.app.get('env') === 'development' ? err :{};

    res.status(err.status || 500);
    res.render('error');
});
module.exports = app;