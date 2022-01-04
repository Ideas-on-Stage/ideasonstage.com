(function($) {
  $.fn.autosubmit = function() {
      this.submit(function(event) {
        event.preventDefault();
        var form = $(this);
        var formtype = form.prop('name');
        var submitButton = $('input[type=submit]', form);
        if (formtype == "newsForm") {
          $.ajax({
            type: 'POST',
                url: form.prop('action'),
                accept: {
                  javascript: 'application/javascript'
                },
                data: form.serialize(),
                beforeSend: function() {
            var waitMessage = submitButton.attr("wait");
            submitButton.prop('value', waitMessage);
            submitButton.prop('disabled', true);
                }
          }).done(function(data) {
              var successMessage = submitButton.attr("success");
              submitButton.prop('value', successMessage);
              alert(submitButton.attr('success'));
    
          
              submitButton.prop('disabled', true);
          }).fail(function(data) {
            // Optionally alert the user of an error here...
          });
        } 
        else {
          if(checkform(form) == true){
            $.ajax({
              type: 'POST',
                  url: form.prop('action'),
                  accept: {
                    javascript: 'application/javascript'
                  },
                  data: form.serialize(),
                  beforeSend: function() {
              var waitMessage = submitButton.attr("wait");
              submitButton.prop('value', waitMessage);
              submitButton.prop('disabled', true);
                  }
            }).done(function(data) {
                var successMessage = submitButton.attr("success");
                submitButton.prop('value', successMessage);
                alert(submitButton.attr('success'));
      
            
                submitButton.prop('disabled', true);
            }).fail(function(data) {
              // Optionally alert the user of an error here...
            });
          }  
        }
          
        
      });
      return this;  
  }
})(jQuery)



