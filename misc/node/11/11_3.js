//p440

const SocketIO = require('socket.io');
module.exports = (server) => {
    // socket.io 패키지를 불러와서 익스프레스 서버와 연결. 두번째 인자로 옵션 객체를 넣어 서버에 관한 여러 설정 가능
    const io = SocketIO(server, {path: '/socket.io'});
  
    io.on('connection', (socket) => {
        //connection 이벤트는 클라이언트가 접속했을 때 발생하고, 콜백으로 socket 객체를 제공
        const req = socket.request;
        const ip = req.header['x-forwarded-for'] || req.connection.remoteAddress;
        console.log('새로운 클라이언트 접속!', ip, socket.id, req.ip); //socket.id로 소켓 고유의 아이디, 또 소켓 주인 알 수 있음
        
        socket.on('disconnect', () => {
            console.log('클라이언트 접속 해제', ip, socket.id);
            clearInterval(socket.interval)
        });
        socket.on('error', (error) => {
            console.error(error);
        });
        socket.on('reply', (data) => {
            console.log(data);
        });
        socket.interval = setInterval(() => {
            // news라는 이벤트 이름으로 Hello Socket.IO 라는 데이터를 클라이언트에 보냄.
            // 클라이언트가 이 메시지를 받기 위해선 new이벤트 리스너를 만들어두어야함
            socket.emit('news', 'Hello Socket.IO');
        }, 3000);
    });
};

/*

*/