(function($) {
  $.fn.autosubmit = function() {
    this.submit(function(event) {
      event.preventDefault();
      var form = $(this);
      var formtype = form.prop('id');
      var submitButton = $('input[type=submit]', form);

      $.post($form.attr("action"), $form.serialize()).done(function(data) {
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