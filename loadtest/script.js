import {randomIntBetween, randomString} from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';
import ws from 'k6/ws';
import {check} from 'k6';

const sessionDuration = 10000; // user session 10000ms

// export const options = {
//     stages: [
//         {duration: '10s', target: 2000},
//         {duration: '30s', target: 5000},
//         {duration: '10s', target: 1000},
//     ],
//     minIterationDuration: '15s',
// };

export const options = {
    stages: [
        // {duration: '10s', target: 2000},
        {duration: '30s', target: 10000},
        // {duration: '10s', target: 1000},
    ],
    minIterationDuration: '15s',
};

export default function () {
    const url = 'ws://localhost:3000/ws';
    const params = {tags: {my_tag: 'hello'}};


    const res = ws.connect(url, params, function (socket) {
        socket.on('open', function open() {
            console.log(`VU ${__VU}: connected`);

            socket.send(JSON.stringify({event: 'SET_NAME', new_name: `Croc ${__VU}`}));

            socket.setInterval(function timeout() {
                socket.send(JSON.stringify({event: 'SAY', message: `I'm saying ${randomString(5)}`}));
            }, randomIntBetween(1000, 3000)); // say something every 1-3seconds
        });

        socket.on('ping', function () {
            console.log('PING!');
        });

        socket.on('pong', function () {
            console.log('PONG!');
        });

        socket.on('close', function () {
            console.log(`VU ${__VU}: disconnected`);
        });

        socket.on('message', function (message) {
            const msg = JSON.parse(message);
            if (msg.event === 'CHAT_MSG') {
                console.log(`VU ${__VU} received: ${msg.user} says: ${msg.message}`);
            } else if (msg.event === 'ERROR') {
                console.error(`VU ${__VU} received:: ${msg.message}`);
            } else {
                console.log(`VU ${__VU} received unhandled message: ${msg.message}`);
            }
        });

        socket.setTimeout(function () {
            console.log(`VU ${__VU}: ${sessionDuration}ms passed, leaving the chat`);
            socket.send(JSON.stringify({event: 'LEAVE'}));
        }, sessionDuration);

        socket.setTimeout(function () {
            console.log(`Closing the socket forcefully 3s after graceful LEAVE`);
            socket.close();
        }, sessionDuration + 3000);
    });

    check(res, {'Connected successfully': (r) => r && r.status === 101});
}
