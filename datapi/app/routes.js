var cacheList = "/app/cachelist"
var appRouter = function(app) {

  app.get("/", function(req, res) {
    res.contentType('application/json');
    res.send(JSON.stringify({"message":"hello"}));
  });

  app.get("/commands", function(req, res) {
    var commands = [];
    commands.push({"name":"create-cache","parameters": ["name,engine"]})
    commands.push({"name":"delete-cache","parameters": ["name"]})
    commands.push({"name":"show-cache","parameters": ["name"]})
    commands.push({"name":"list-cache","parameters": []})

    res.contentType('application/json');
    res.send(JSON.stringify(commands));
  });

  app.post("/caches", function(req, res) {
    if(!req.body.name) {
      return res.send({"status": "error", "message": "missing name"});
    }
    if(!req.body.engine) {
      return res.send({"status": "error", "message": "missing engine"});
    }
    var fs = require('fs');
    var obj = JSON.parse(fs.readFileSync(cacheList, 'utf8'));
    obj["caches"].push({"name":req.body.name,"engine": req.body.engine})
    data = JSON.stringify(obj);

    fs.writeFile(cacheList, data, function(err) {
        res.contentType('application/json');
        if(err) {
            res.send({"message":"error saving cache list"});
            return;
        }
        res.send(data);
    });

  });

  app.get("/caches", function(req, res) {
    var fs = require('fs');
    fs.readFile(cacheList, 'utf8', function (err, data) {
      res.contentType('application/json');
      if(err) {
          res.send({"message":"error reading cache list"});
          return;
      }

      res.send(data);
    });
  });

  app.get('/caches/:name', function (req, res) {

    var fs = require('fs');
    fs.readFile(cacheList, 'utf8', function (err, data) {
      res.contentType('application/json');
      if(err) {
          res.send({"message":"error reading cache list"});
          return;
      }

      var obj = JSON.parse(data);
      var caches = obj["caches"];
      for (var i = 0, len = caches.length; i < len; i++) {
        if (caches[i]["name"] == req.params.name) {
          res.send(JSON.stringify(caches[i]));
          return
        }
      }

      res.send({"message":"cache not found"});
    });
  })

}

module.exports = appRouter;
