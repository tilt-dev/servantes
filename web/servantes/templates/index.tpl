<!doctype html>
<html>
<head>
  <title>Servantes</title>
</head>
<body style="font-family: system, -apple-system, 'Roboto', sans-serif; color: rgba(0,0,0,0.8); background-color: #fcfcfc;">

  <div style="margin: 2em; text-align: center; font-size: 3em;">
    Welcome to Servantes
  </div>
  
  <div style="display: grid; grid-template-columns: repeat(3, 1fr); grid-gap: 0.5em;">
    {{range $i, $service := .Services}}
      <div style="position:relative; box-shadow: inset 0 0 4em #00bfff;">
        <div style="position:absolute; font-size:0.8em; bottom:0.5em; right:0.5em; z-index:1; text-align:right;">
          service: {{$service.Name}}<br>
          phase: {{$service.Phase}}<br>
          age: {{$service.HumanAge}}
        </div>
        <iframe frameborder="none"
                style="padding:3em;width:100%;height:100%;" src="/s/{{$service.Name}}"></iframe>
      </div>
    {{end}}
      
  </div>

</body>
</html>
