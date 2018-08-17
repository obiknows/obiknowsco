require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.min.js");
$(() => {
    // Setup tooltips everywhere
    $(function() {
      $('[data-toggle="tooltip"]').tooltip();
    });

    // Entry Carousel JS
    $('.carousel').carousel({
      ride: true,
      interval: 2500,
    })


    // Listener 
    function swapCarousels(x) {
      if (x.matches) {
        // If media query matches
        document.body.style.backgroundColor = "yellow";
      } else {
        document.body.style.backgroundColor = "pink";
      }
    }

    var x = window.matchMedia("(max-width: 450px)");
    myFunction(x); // Call listener function at run time
    x.addListener(swapCarousels); // Attach listener function on state changes
  });
