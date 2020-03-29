/*
p160
cluster 모듈은 싱글 스레드인 노드가 cpu코어를 모두 사용할 수 있게 해주는 모듈
포트를 공유하는 노드 프로세스를 여러 개 둘 수 있어서 병렬로 요청 분산 가능
*/
const cluster = require('cluster');
const http = require('http');
const numCPUs = require('os').cpus().length;

if (cluster.isMaster){
  console.log(`마스터 프로세스 아이디: ${process.pid}`);
  for (let i =0; i<numCPUs; i+=1){
    cluster.fork(); //CPU갯수만큼 워커를 생산
  }
  //워커가 종료되었을 때
  cluster.on('exit', (worker, code, signal) => {
    console.log(`${worker.process.pid}번 워커가 종료되었습니다`);
  });
} else{
  //워커들이 포트에서 대기
  http.createServer((req, res) => {
    res.write('<h1>Hello Node1</h1>');
    res.end('<p>Hello Cluster!</p>');
    setTimeout(() => {
      process.exit(1);
    }, 1000);
  }).listen(8085);

console.log(`${process.pid}번 워커 실행`);
}
/*
12번까지는 오류가 발생해도 서버가 정상 작동할 수 있다는 의미.
워커 죽을 때 종료된 워커를 다시 커면 오류가 발생해도 계속 버틸 수 있다.
cluster.fork();
이런 방식으로 오류를 막으려는 것은 좋지 않음. 오류의 원인을 찾아 해결해야 한다.

실무에서는 pm2 등의 모듈로 cluster 기능을 사용.

*/
