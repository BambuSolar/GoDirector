// server.js

// BASE SETUP
// =============================================================================

// call the packages we need
var express    = require('express');        // call express
var app        = express();                 // define our app using express
var app2        = express();                 // define our app using express
var bodyParser = require('body-parser');

// configure app to use bodyParser()
// this will let us get the data from a POST
app.use(bodyParser.urlencoded({ extended: true }));
app2.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());
app2.use(bodyParser.json());

var port = process.env.PORT || 5000;        // set our port

// ROUTES FOR OUR API
// =============================================================================
var router = express.Router();              // get an instance of the express Router
var router2 = express.Router();              // get an instance of the express Router


router.get('/builds', function(req, res) {
    res.json(
        {
            "beta": [
                "v0.0.2-beta_1",
                "v0.0.2-beta_2",
                "v1.0.2-beta_1",
                "v1.0.2-beta_2",
                "v1.0.3-beta_1",
                "v1.0.3-beta_2",
                "v1.0.3-beta_3",
                "v1.0.3-beta_4",
                "v1.0.3-beta_5",
                "v1.1.1-beta_1",
                "v1.1.1-beta_2",
                "v1.1.1-beta_3",
                "v1.1.1-beta_4",
                "v1.1.2-beta_1",
                "v1.1.3-beta_1"
            ],
            "prod": [
                "v0.0.1",
                "v1.0.1",
                "v1.0.2",
                "v1.1.1",
                "v1.1.2"
            ]
        }
    );
});

router.post('/builds', function(req, res) {
    //console.log(req.body);
    res.status(201).json({"data": "v0.0.2-beta_1"});
});

router.post('/deploys', function(req, res) {
    //console.log(req.body);
    res.status(201).json({"message": "Deploy success"});
});


// more routes for our API will happen here

// REGISTER OUR ROUTES -------------------------------
// all of our routes will be prefixed with /api
app.use('/api', router);
app2.use('', router2);

// START THE SERVER
// =============================================================================
app.listen(port);																							
app2.listen(3000);							
console.log('Magic happens on port ' + port);
