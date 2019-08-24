// Shared Web Worker js for chixm servertemplate

onmessage = function (e) {
    console.log('Worker: Message received from main script');
    // multiply two numbers
    let result = e.data[0] * e.data[1];
    if (isNaN(result)) {
        postMessage('Please write two numbers');
    } else {
        let workerResult = 'Result: ' + result;
        console.log('Worker: Posting message back to main script');
        postMessage(workerResult);
    }
}