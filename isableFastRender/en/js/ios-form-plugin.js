(function($) {
  $.fn.autosubmit = function() {
    this.submit(function(event) {
      event.preventDefault();
      var form = $(this);
      var formtype = form.prop('id');
      //this is a control point to see if the js script is getting to this point
      console.log('form:',form);
      console.log('formtype:', formtype);
      var submitButton = $('input[type=submit]', form);
      
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
    });
    return this;
  }
})(jQuery)



