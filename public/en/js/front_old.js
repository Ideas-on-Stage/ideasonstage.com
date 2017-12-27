/* global $this: true */
/* eslint no-unused-vars: ["error", { "varsIgnorePattern": "animationsSlider" }] */


$(function () {
  sliders()
  utils()
})


/* sliders */
function sliders () {
  if ($('.owl-carousel').length) {
    $('.testimonials').owlCarousel({
      items: 4,
      itemsDesktopSmall: [990, 3],
      itemsTablet: [768, 2],
      itemsMobile: [480, 1]
    })

  }
  if ($('.owl-carousel').length) {
    $('.references').owlCarousel({
      items: 6,
      itemsDesktopSmall: [990, 6],
      itemsTablet: [768, 6],
      itemsMobile: [480, 3]
    })
  }
  
  if ($('.owl-carousel').length) {
    $('.gallery').owlCarousel({
      items: 5,
      itemsDesktopSmall: [990, 5],
      itemsTablet: [768, 5],
      itemsMobile: [480, 3]
    })
  }
}

function utils () {
  /* tooltips */
  $('[data-toggle="tooltip"]').tooltip()

  /* click on the box activates the radio */
  $('#checkout').on('click', '.box.shipping-method, .box.payment-method', function () {
    var radio = $(this).find(':radio')
    radio.prop('checked', true)
  })

  /* click on the box activates the link in it */
  $('.box.clickable').on('click', function () {
    window.location = $(this).find('a').attr('href')
  })

  /* external links in new window */
  $('.external').on('click', function (e) {
    e.preventDefault()
    window.open($(this).attr('href'))
  })

  /* animated scrolling */
  $('.scroll-to, .scroll-to-top').click(function (event) {
    var fullUrl = this.href
    var parts = fullUrl.split('#')

    if (parts.length > 1) {
      scrollTo(fullUrl)
      event.preventDefault()
    }
  })

  function scrollTo (fullUrl) {
    var parts = fullUrl.split('#')
    var trgt = parts[1]
    var targetOffset = $('#' + trgt).offset()
    var targetTop = targetOffset.top - 100

    if (targetTop < 0) {
      targetTop = 0
    }

    $('html, body').animate({
      scrollTop: targetTop
    }, 1000)
  }
}


