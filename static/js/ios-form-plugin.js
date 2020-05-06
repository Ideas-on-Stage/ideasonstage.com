(function($) {
  $.fn.autosubmit = function() {
    this.submit(function(event) {
      event.preventDefault();
      var form = $(this);
      //this is a control point to see if the js script is getting to this point
      console.log('form:',form);
      $.ajax({
        type: form.attr('method'),
        url: form.attr('action'),
        data: form.serialize()
      }).done(function(data) {
        // Optionally alert the user of success here...
        alert("Hello! I am an alert box!");
      }).fail(function(data) {
        // Optionally alert the user of an error here...
      });
    });
    return this;
  }
})(jQuery)