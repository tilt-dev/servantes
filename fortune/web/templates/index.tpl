<!doctype html>
  <head>
    <title>Fortune</title>
    <style>
      /* CSS Variables */
      :root {
        --white:        #FFFFFF;
        --blue:         #88ABAD;
        --dark-gray:    #19282B;
      }

      html {
        margin: 0;
        height: 100%;
      }

      body {
        margin: 0;
        font-family: 'Anonymous Pro', monospace;
        color: var(--blue);
        font-size: 1.5em;
        height: 100%;
      }

      main {
        text-align: center;
        height: 100%;
        display: flex;
        flex-direction: column;
        text-align: center;
        justify-content: center;
      }

      p {
        margin-top: 0;
        margin-bottom: 1em;
      }
    </style>
    <link href="https://fonts.googleapis.com/css?family=Anonymous+Pro:400,700" rel="stylesheet">
  </head>

  <body>
    <main>
      <p>Your fortune is:</p>
      <p><strong>{{.}}</strong></p>
    </main>
  </body>
</html>