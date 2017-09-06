var module_path ='';
var express = require(module_path+'express');
var bodyParser = require(module_path+'body-parser');
var mongoose = require(module_path+'mongoose');

mongoose.connect('mongodb://localhost/myproject');

var Schema = mongoose.Schema;

var MhswSchema = new Schema({
	name: String,
	address: String,
	old: String
});

mongoose.model('Mhsw', MhswSchema);

var Mhsw = mongoose.model('Mhsw');

/*var mhsw = new Mhsw({
	name: 'Haidar',
	address: 'malang',
	old: '14'
});

mhsw.save();
*/
var app = express();

app.use(bodyParser.json());
app.use(express.static(__dirname + '/public'));

// ROUTES

app.get('/api/mhsws', function(req, res) {
	Mhsw.find(function(err, docs) {
		docs.forEach(function(item) {
			console.log("Received a GET request for _id: " + item._id);
		})
		res.send(docs);
	});
});

app.post('/api/mhsws', function(req, res) {
	console.log('Received a POST request:')
	for (var key in req.body) {
		console.log(key + ': ' + req.body[key]);
	}
	var mhsw = new Mhsw(req.body);
	mhsw.save(function(err, doc) {
		res.send(doc);
	});
});

app.delete('/api/mhsws/:id', function(req, res) {
	console.log('Received a DELETE request for _id: ' + req.params.id);
	Mhsw.remove({_id: req.params.id}, function(err, doc) {
		res.send({_id: req.params.id});
	});
});

app.put('/api/mhsws/:id', function(req, res) {
	console.log('Received an UPDATE request for _id: ' + req.params.id);
	Mhsw.update({_id: req.params.id}, req.body, function(err) {
		res.send({_id: req.params.id});
	});
}); 

var port = 3003;

app.listen(port);
console.log('server on ' + port);
