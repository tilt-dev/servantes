<!doctype html>
<html>
<head>
  <title>Doggos</title>
</head>
<body style="font-family: system, -apple-system, 'Roboto', sans-serif; color: rgba(0,0,0,0.8);overflow:hidden;">
<div id="container"></div>
</body>

<script>
window.twttr = (function(d, s, id) {
  var js, fjs = d.getElementsByTagName(s)[0],
    t = window.twttr || {};
  if (d.getElementById(id)) return t;
  js = d.createElement(s);
  js.id = id;
  js.src = "https://platform.twitter.com/widgets.js";
  fjs.parentNode.insertBefore(js, fjs);

  t._e = [];
  t.ready = function(f) {
    t._e.push(f);
  };

  return t;
}(document, "script", "twitter-wjs"));

// bishon: 1033600515811115008
twttr.ready(() => {
  twttr.widgets.createTweet(
    '1033600515811115008',
    document.getElementById('container'))
});
</script>

</html>
