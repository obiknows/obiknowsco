require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.min.js");
$(() => {
  // Setup tooltips everywhere
  $(function () {
    $('[data-toggle="tooltip"]').tooltip()
  })
});
