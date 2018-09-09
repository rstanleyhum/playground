const functions = require('firebase-functions');
const {Storage} = require('@google-cloud/storage');

const gcs = new Storage({
  projectId: 'logindemo-5f7b4',
});

const storageBucketName = "logindemo-5f7b4.appspot.com";

exports.assets = functions.https.onRequest((request, response) => {
  const bucket = gcs.bucket(storageBucketName);
  var query = {
    prefix: 'assets/',
//    delimiter: '/', limits recursive
  };
  var checkDate = _getCheckDate(request.query.time);

  bucket.getFiles(query)
    .then(data => {
      var files = data[0];
      var names = files.map(f => _returnFilenameIfUpdated(f.metadata.name, new Date(f.metadata.updated), checkDate));
      var json = names.filter(n => n != null);
      response.json(json);
    })
    .catch( err => {
      response.status(500).json({"error": err});
    });
});


const _returnFilenameIfUpdated = (name, lastupdated, checkdate) => {
  //return lastupdated > checkdate ? name : null;
  return name;
}


const _getCheckDate = (timestring) => {
  var checkTime = parseInt(timestring);
  if (!checkTime) {
    checkTime = Date.now();
  }
  return new Date(checkTime);
}