<html>
<head>
    <title>Emoji</title>
    <style>
        html {
            margin: 0;
            height: 100%;
        }

        body {
            margin: 0;
            font-family: 'BioRhyme', serif;
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
            margin: 0;
        }

        .time {
            font-size: 2em;
        }
    </style>
    <link href="https://fonts.googleapis.com/css?family=BioRhyme:400,800" rel="stylesheet">
</head>

<body>
    <main>
      {{range $row := .EmojiRows}}
        <p class="emoji">{{$row}}<p>
      {{end}}
    </main>
</body>

</html>
