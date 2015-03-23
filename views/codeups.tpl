<header>

</header>
  <div class="container" id="codeUps">
    <!-- Example row of columns -->
    <div class="row">
      <div class="col-md-4" rv-each-codeup="scope.codeUps">
        <h2>{ codeup.Title }</h2>
        <p>{ codeup.Description }</p>
        <p><a class="btn btn-default" href="#" role="button">Anmäl dig här »</a></p>
      </div>
    </div>

    <hr>

    <footer>
      <p>{ scope.user }</p>
      <p>Videum CodeUp registreringssystem</p>
    </footer>
  </div>

  {{ define "js" }}
  <script>
  $( document ).ready(function() {
    app.codeUps.init();
  });
  </script>
  {{ end}}
