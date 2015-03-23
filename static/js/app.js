var app = {
  codeUps : {
    model : new CodeUpsModel(),
    init : function(){
      rivets.bind($('#codeUps'), {scope: app.codeUps.model})
      codeUpService.getAll(app.codeUps.gui.updateCodeUpList);
    },
    gui : {
        updateCodeUpList : function(codeUpsJsonData) {
            app.codeUps.model.setData(codeUpsJsonData);
        }
    }
  }
};

var codeUpService = {
  getAll : function(successCallback){
    $.get( "CodeUps/GetAll", function( data ) {
        successCallback(data);
    });
  }
};
