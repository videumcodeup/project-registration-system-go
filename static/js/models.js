
var CodeUpsModel = function () {
  this.codeUps = [];
  this.user = "Sebbe"
};

CodeUpsModel.prototype = {
    setData: function (codeUpsData) {
        console.log(codeUpsData);
        this.codeUps = codeUpsData;
        this.user = "Senästien"
    }
}

/*
{
  "Id": 1,
  "Title": "CodeUp-träff 1",
  "Description": "Vi träffas, kodar, äter pizza och dricker öl. Mys ^^",
  "StartTime": "2015-03-21T20:16:17+01:00",
  "EndTime": "2015-03-21T23:16:17+01:00",
  "Registrations": null
}
*/
