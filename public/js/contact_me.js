/*
    This script was originally part of David Millers (https://github.com/davidtmiller)
    Bootstrap theme "Freelancer" add some functionality to the contact form.

    Licensed under the Apache License 2.0. For more information take a look at the orignal
    repository unde https://github.com/IronSummitMedia/startbootstrap-freelancer.
*/

$(function() {

    $("input,textarea").jqBootstrapValidation({
        preventSubmit: true,
        submitError: function($form, event, errors) {
            // additional error messages or events
        },
        submitSuccess: function($form, e) {
    	    e.preventDefault();
    	
    	    var submitButton = $('input[type=submit]', $form);
    		var success = $('#success', $form)
    	    $.ajax({
    	      type: 'POST',
    	      url: $form.prop('action'),
    	      accept: {
    	        javascript: 'application/javascript'
    	      },
    	      data: $form.serialize(),
    	      beforeSend: function() {
				submitButton.prop('value', '{{ with .Site.Params.contact.wait }}{{ . }}{{ end }}');
				submitButton.prop('disabled', 'disabled');
				setTimeout(function(){
				    submitButton.prop('value', '{{ with .Site.Params.contact.success }}{{ . }}{{ end }}');
				}, 3000);
    	      }
    	    }).done(function(data) {
    			hj('formSubmitSuccessful');
    			submitButton.prop('disabled', false);
    	    });
    	    
        },

        filter: function() {
            return $(this).is(":visible");
        },
    });

    $("a[data-toggle=\"tab\"]").click(function(e) {
        e.preventDefault();
        $(this).tab("show");
    });
});

/*When clicking on Full hide fail/success boxes */
$('#name').focus(function() {
    $('#success').html('');
});

