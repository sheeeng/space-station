var http = require('http');

var url = 'http://api.open-notify.org/iss-now.json';

http.get(url, function(res){
    var body = '';

    res.on('data', function(chunk){
        body += chunk;
    });

    res.on('end', function(){
        var response = JSON.parse(body);
        console.log("Got response: ", response);

        console.log("Got response: ", response.message);
        console.log("Got response: ", response.timestamp);
        console.log("Got response: ", response.iss_position);
        console.log("Got response: ", response.iss_position['latitude']);
        console.log("Got response: ", response.iss_position['longitude']);

    });
}).on('error', function(e){
      console.log("Got an error: ", e);
});
