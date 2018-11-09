<!doctype html>
<html>
<head>
  <title>Servantes</title>
  <link href="https://fonts.googleapis.com/css?family=Lora:700|Varela+Round" rel="stylesheet">
</head>
<style>
  /* CSS Variables */
  :root {
    --white:        #FFFFFF;
    --white-var1:   #FFFDF6;
    --white-var2:   #FAF6E9;
    --tan:          #ECE8D9;
    --tan-dark:     #B3A16B;
    --yellow:       #C4A827;
    --blue:         #88ABAD;
    --dark-gray:    #19282B;
    /* --blue-green: #678072; */
  }
  
  html {
    margin: 0;
  }

  body {
    margin: 0;
    font-family: 'Varela Round', sans-serif;
    color: var(--dark-gray);
    background-color: var(--white-var1);
  }

  header {
    font-family: 'Lora', serif;
    color: var(--tan-dark);
    margin-top: 0;
    margin-left: auto;
    margin-right: auto;
    margin-bottom: 0;
    max-width: 100em;
    padding-top: 1em;
    padding-left: 2em;
    padding-right: 2em;
    font-size: 2em;
    text-align: right;
  }

  ul {
    margin: 0;
    padding: 0;
  }

  li {
    list-style-type: none;
    padding: 0;
  }

  ul.services {
    display: grid; 
    width: 100vw; 
    max-width: 100em;
    margin-left: auto;
    margin-right: auto;
    padding: 2em;
    grid-gap: 2em; 
    grid-template-columns: repeat(3, 1fr); 
    box-sizing: border-box;
  }

  li.service-item {
    background-color: var(--white);
    padding: 0;
    position: relative;
    min-height: 18em;
    border-right: 10px solid var(--tan);
    border-bottom: 10px solid var(--tan);
    /* Make room for K8s Data */
    margin-top: 2.5em; 
  }

  ul.k8s-data {
    color: var(--white-var1);
    background-color: var(--dark-gray);
    position: absolute;
    font-size: 0.8em; 
    top: -2.5em;
    height: 2.5em;
    z-index: 1; 
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding-left: 1em;
    padding-right: 1em;
    border-top-left-radius: 0.5em;
    border-top-right-radius: 0.5em;
    /* Align with border below: */
    left: -1px;
    right: -1px;
  }

  ul.k8s-data span + span {
    padding-left: 0.5em;
  }

  iframe {
    display: block;
    overflow: hidden;
    position: absolute;
    box-sizing: border-box;
    border: 1px solid var(--dark-gray);
    width: 100%;
    height: 100%;
  }
</style>
<body>
  <main>
    <header>Servantes ✴︎</header>
    <ul class="services">
      {{range $i, $service := .Services}}
        <li class="service-item">
          <ul class="k8s-data">
            <li>{{$service.Name}} — {{$service.Status}}</li>
            <li>
              <span>restarts: {{$service.RestartCount}}</span>
              <span>•</span>
              <span>age: {{$service.HumanAge}}</span>
            </li>
          </ul>
          <iframe frameborder="0" src="http://localhost:{{$service.Port}}/"></iframe>
        </li>
      {{end}}
    </ul>
  </main>
</body>
</html>
