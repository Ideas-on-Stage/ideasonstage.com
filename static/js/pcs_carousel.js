/* global $this: true */
/* eslint no-unused-vars: ["error", { "varsIgnorePattern": "animationsSlider" }] */


$(function () {
  sliders()
})


/* sliders */
function sliders () {
  if ($('.owl-carousel').length) {
    $('.gallery').owlCarousel({
      items: 5,
      itemsDesktopSmall: [990, 5],
      itemsTablet: [768, 5],
      itemsMobile: [480, 3]
    })
  }
}


