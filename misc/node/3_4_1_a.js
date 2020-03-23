//p083

/*
global 객체는 브라우저의 window와 같은 객체
전역 객체이므로 모든 파일에서 접근 가능
*/
module.exports = () => global.message;

//3_4_b.js참고

/*
global 객체 남용은 no.
프로그램의 규모가 커질수록 유지보수에 어려움을 겪게 되기 때문
다른 파일의 값을 사용하고 싶다면 모듈 형식으로 만들어서 명시적으로 값을 불러와 사용하는 것이 좋음.
*/